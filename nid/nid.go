package nid

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Nid struct {
	Type        Type
	RoutingCode RoutingCode
	WeakID      WeakID
	Reserved    Reserved
}

// Deprecated: Use NewWithRoutingCode instead
func New(idType Type, weakID WeakID) (Nid, error) {
	return NewWithRoutingCode(idType, defaultRegion, weakID)
}

func NewWithRoutingCode(idType Type, idRoutingCode RoutingCode, weakID WeakID) (Nid, error) {
	nid := Nid{
		Type:        idType,
		RoutingCode: idRoutingCode,
		WeakID:      weakID,
		Reserved:    NoReserved,
	}

	err := nid.Validate()
	return nid, err
}

func Parse(s string) (Nid, error) {
	t, r, wid, rsrv, err := matchNid(s)
	if err != nil {
		return Nid{}, err
	}

	idRoutingCode, err := NewRoutingCode(r)
	if err != nil {
		return Nid{}, err
	}

	idType, err := NewType(t)
	if err != nil {
		return Nid{}, err
	}

	weakID, err := NewWeakID(wid)
	if err != nil {
		return Nid{}, err
	}

	reserved, err := NewReserved(rsrv)
	if err != nil {
		return Nid{}, err
	}

	return Nid{
		Type:        idType,
		RoutingCode: idRoutingCode,
		WeakID:      weakID,
		Reserved:    reserved,
	}, nil
}

func (n Nid) Serialize() string {
	// string builder is faster than concatenation
	var sb strings.Builder
	estimatedSize := len(n.Type) + 1 + len(n.RoutingCode) + len(n.WeakID)
	if n.Reserved != NoReserved {
		estimatedSize += 2 + len(n.Reserved)
	}
	sb.Grow(estimatedSize)

	sb.WriteString(string(n.Type))
	sb.WriteString("-")
	sb.WriteString(string(n.RoutingCode))
	sb.WriteString(string(n.WeakID))

	if n.Reserved != NoReserved {
		sb.WriteString("--")
		sb.WriteString(string(n.Reserved))
	}

	return sb.String()
}

func (n Nid) String() string {
	return n.Serialize()
}

func (n Nid) Validate() error {
	if !n.Type.IsValid() {
		return errors.New("type is invalid")
	}

	if !n.RoutingCode.IsValid() {
		return errors.New("routingCode is invalid")
	}

	if !n.WeakID.IsValid() {
		return errors.New("weak id is invalid")
	}

	if !n.Reserved.IsValid() {
		return errors.New("reserved is invalid")
	}

	return nil
}

// Ensure Nid implements the "database/sql".Scanner and "database/sql/driver".Valuer interfaces
// So it is (de)serializable by most of sql framework (gorm, xorm, sqlx, etc.) out of the box.
var (
	_ sql.Scanner   = (*Nid)(nil)
	_ driver.Valuer = Nid{} // method Value with **value receiver** is required for ydb
)

func (n Nid) Value() (driver.Value, error) {
	if n.Validate() != nil {
		return nil, errors.New("invalid NID")
	}
	return n.Serialize(), nil
}

func (n *Nid) Scan(src any) error {
	switch s := src.(type) {
	case string:
		parsed, err := Parse(s)
		if err != nil {
			return fmt.Errorf("parsing NID string: %w", err)
		}
		*n = parsed
	case []byte:
		parsed, err := Parse(string(s))
		if err != nil {
			return fmt.Errorf("parsing NID []byte: %w", err)
		}
		*n = parsed
	default:
		return fmt.Errorf("invalid NID source type: %T", src)
	}
	return nil
}

func (n Nid) MarshalYAML() (interface{}, error) {
	return n.Serialize(), nil
}

func (n *Nid) UnmarshalYAML(value *yaml.Node) error {
	var s string
	err := value.Decode(&s)
	if err != nil {
		return fmt.Errorf("decoding YAML node: %w", err)
	}

	parsed, err := Parse(s)
	if err != nil {
		return fmt.Errorf("parsing NID from YAML: %w", err)
	}

	*n = parsed
	return nil
}

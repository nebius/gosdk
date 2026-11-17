package mask

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"
)

// TokenType constants.
const (
	tokenComma = iota + 1
	tokenDot
	tokenLBrace
	tokenRBrace
	tokenWildCard
	tokenPlainKey
	tokenQuotedKey
	tokenEOL
)

// token represents a token type and its corresponding value.
type token struct {
	Type  int    // Type of the token (one of the constants above)
	Value string // Value associated with the token
	Pos   int    // Position of the token in the input
}

func (t token) TypeStr() string {
	switch t.Type {
	case tokenComma:
		return "Comma"
	case tokenDot:
		return "Dot"
	case tokenLBrace:
		return "LBrace"
	case tokenRBrace:
		return "RBrace"
	case tokenWildCard:
		return "WildCard"
	case tokenPlainKey:
		return "PlainKey"
	case tokenQuotedKey:
		return "QuotedKey"
	case tokenEOL:
		return "EOL"
	default:
		return fmt.Sprintf("Undefined(%d)", t.Type)
	}
}

// Unquote checks if the token is of type tokenQuotedKey and unquotes its value.
// It returns an error if unquoting fails.
func (t token) Unquote() (token, error) {
	if t.Type != tokenQuotedKey {
		return t, nil
	}
	var str string
	err := json.Unmarshal([]byte(t.Value), &str)
	if err != nil {
		return t, fmt.Errorf("failed to parse quoted string: %w", err)
	}
	t.Value = str
	t.Type = tokenPlainKey
	return t, nil
}

func (t token) String() string {
	if t.Type == tokenQuotedKey {
		return fmt.Sprintf("Token%s(%s pos %d)", t.TypeStr(), t.Value, t.Pos)
	}
	return fmt.Sprintf("Token%s(%q pos %d)", t.TypeStr(), t.Value, t.Pos)
}

// regular expression to match simple string tokens (alphanumeric and
// underscores).
const simpleStringTokenSet = "[a-zA-Z0-9_]" //nolint:gosec // false positive
var simpleStringTokenStart = regexp.MustCompile("^" + simpleStringTokenSet)
var simpleStringToken = regexp.MustCompile("^(" + simpleStringTokenSet + "+)")

const (
	spaceSet = " \r\n\t"
	// maximum length of context to show in error messages.
	maxContext = 30
	// number of characters to show before the error position in error messages.
	contextBack = 12
	errorMarker = "\u20DE" // combining enclosing square
)

type UnmarshalError struct {
	Input    string
	Position int
	Summary  string
	Cause    error // Cause is an optional original error.
}

// possibleEllipsis truncates the string if it exceeds maxContext and adds
// ellipsis.
func possibleEllipsis(str string, maxContext int) string {
	if len(str) > maxContext {
		return str[0:maxContext-3] + "..."
	}
	return str
}

// near returns a contextual error message indicating the position of an error
// in the input string.
func near(str string, pos int) string {
	ctxStart := pos
	if ctxStart >= contextBack {
		ctxStart -= contextBack
	} else {
		ctxStart = 0
	}
	delta := pos - ctxStart
	contextStr := possibleEllipsis(str[ctxStart:], maxContext)
	withMark := contextStr[0:delta] + errorMarker + contextStr[delta:]

	return fmt.Sprintf("at position %d near %q", pos, withMark)
}

func (e *UnmarshalError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf(
			"%s %s: %s",
			e.Summary, near(e.Input, e.Position), e.Cause.Error(),
		)
	}
	return fmt.Sprintf("%s %s", e.Summary, near(e.Input, e.Position))
}

func (e *UnmarshalError) Unwrap() error {
	return e.Cause
}

// lexer represents a tokenizer that breaks an input string into tokens.
type lexer struct {
	input string // input is the string to tokenize.
	pos   int    // pos is the current position in the input string.
}

// newLexer initializes a new Lexer with the input string.
func newLexer(input string) *lexer {
	return &lexer{input: input}
}

// nextToken returns the next token from the input string.
func (l *lexer) nextToken() (token, error) {
	// Skip whitespace characters
	for l.pos < len(l.input) && isWhitespace(l.input[l.pos]) {
		l.pos++
	}

	if l.pos >= len(l.input) {
		return token{Type: tokenEOL, Value: "", Pos: l.pos}, nil
	}
	start := l.pos

	// Check for specific token patterns
	switch {
	case l.startsWith(","):
		l.consume(",")
		return token{Type: tokenComma, Value: ",", Pos: start}, nil
	case l.startsWith("."):
		l.consume(".")
		return token{Type: tokenDot, Value: ".", Pos: start}, nil
	case l.startsWith("("):
		l.consume("(")
		return token{Type: tokenLBrace, Value: "(", Pos: start}, nil
	case l.startsWith(")"):
		l.consume(")")
		return token{Type: tokenRBrace, Value: ")", Pos: start}, nil
	case l.startsWith("*"):
		l.consume("*")
		return token{Type: tokenWildCard, Value: "*", Pos: start}, nil
	case l.startsWith("\""):
		return l.scanQuotedKey()
	case simpleStringTokenStart.MatchString(l.input[l.pos:]):
		return l.scanPlainKey()
	default:
		return token{}, &UnmarshalError{
			Input:    l.input,
			Position: l.pos,
			Summary:  "unexpected symbol",
		}
	}
}

// scanPlainKey scans a PlainKey token.
func (l *lexer) scanPlainKey() (token, error) {
	start := l.pos
	match := simpleStringToken.FindString(l.input[start:])
	l.consume(match)
	return token{Type: tokenPlainKey, Value: match, Pos: start}, nil
}

// scanQuotedKey scans a QuotedKey token.
func (l *lexer) scanQuotedKey() (token, error) {
	start := l.pos
	l.consume("\"")
	for l.pos < len(l.input) && !l.startsWith("\"") {
		if l.startsWith("\\") {
			l.pos += 2 // skip the escaped character
		} else {
			l.pos++
		}
	}
	if l.startsWith("\"") {
		l.consume("\"")
		return token{
			Type:  tokenQuotedKey,
			Value: l.input[start:l.pos],
			Pos:   start,
		}, nil
	}
	return token{}, &UnmarshalError{
		Input:    l.input,
		Position: start,
		Summary:  "unterminated quoted string",
	}
}

// consume moves the lexer position forward by the given length.
func (l *lexer) consume(s string) {
	l.pos += len(s)
}

// startsWith checks if the input string starts with a given prefix at the
// current position.
func (l *lexer) startsWith(prefix string) bool {
	return l.pos+len(prefix) <= len(l.input) &&
		l.input[l.pos:l.pos+len(prefix)] == prefix
}

// isWhitespace checks if a given byte is a whitespace character.
func isWhitespace(b byte) bool {
	return strings.ContainsRune(spaceSet, rune(b))
}

// tokenize runs the lexer to generate a list of tokens from the input string.
func (l *lexer) tokenize() ([]token, error) {
	ret := []token{}
	for {
		tok, err := l.nextToken()
		if err != nil {
			return nil, fmt.Errorf("failed to tokenize: %w", err)
		}
		tok, err = tok.Unquote()
		if err != nil {
			return nil, &UnmarshalError{
				Input:    l.input,
				Position: tok.Pos,
				Summary:  "failed to unquote token",
				Cause:    err,
			}
		}
		if tok.Type == tokenEOL {
			break
		}
		ret = append(ret, tok)
	}
	return ret, nil
}

// level represents a part of mask encapsed in braces.
//
// Starts, Ends and Active are sets: one mask occurs one time at most. Two
// different masks never share a child. Thus, the size of each set is not more
// than the number of masks in the result.
type level struct {
	Starts []*Mask // Starts is a list of mask starting nodes of this level.
	Ends   []*Mask // Ends is a list of finished masks for this level.
	Active []*Mask // Active is the current unfinished masks for this level.
	// The depths are the maximum numbers of edges from the root to the
	// masks in the corresponding sets. The root has depth zero.
	startsDepth int
	endsDepth   int
	activeDepth int
	// ended contains the masks of Ends.
	ended map[*Mask]struct{}
	// collectEnds is true when a key or a wildcard can extend the masks that
	// this level finishes. Only then the level fills Ends.
	collectEnds bool
	// Prev is a reference to the previous level in the hierarchy.
	Prev *level
	// Pos is the position in the input where this level was created.
	Pos int
}

// newLevel initializes a new mask level with a root mask.
func newLevel() *level {
	root := New()
	return &level{
		Starts: []*Mask{root},
		Ends:   []*Mask{},
		Active: []*Mask{root},
	}
}

// AddKey adds a new field key to the unfinished masks in the current level.
// Each unfinished mask gets its own child.
func (l *level) AddKey(k FieldKey) error {
	if err := l.descend(); err != nil {
		return err
	}
	newActive := make([]*Mask, 0, len(l.Active))
	for _, a := range l.Active {
		m, ok := a.FieldParts[k]
		if !ok {
			m = New()
			a.FieldParts[k] = m
		}
		newActive = append(newActive, m)
	}
	l.Active = newActive
	return nil
}

// AddAny adds a new wildcard to the unfinished masks in the current level.
// Each unfinished mask gets its own child.
func (l *level) AddAny() error {
	if err := l.descend(); err != nil {
		return err
	}
	newActive := make([]*Mask, 0, len(l.Active))
	for _, a := range l.Active {
		if a.Any == nil {
			a.Any = New()
		}
		newActive = append(newActive, a.Any)
	}
	l.Active = newActive
	return nil
}

// descend checks the depth before a key or wildcard extends Active. Each
// active mask gets a child one edge deeper, so only the maximum is needed.
func (l *level) descend() error {
	if len(l.Active) == 0 {
		return nil
	}
	if l.activeDepth+1 >= recursionTooDeep {
		return ErrRecursionTooDeep
	}
	l.activeDepth++
	return nil
}

// finish moves the unfinished masks to Ends. Branches that end at the same
// mask add it one time.
func (l *level) finish() {
	if !l.collectEnds {
		return
	}
	if l.ended == nil {
		l.ended = map[*Mask]struct{}{}
	}
	for _, a := range l.Active {
		if _, ok := l.ended[a]; !ok {
			l.ended[a] = struct{}{}
			l.Ends = append(l.Ends, a)
		}
	}
	l.endsDepth = max(l.endsDepth, l.activeDepth)
}

// NewMask finalizes current mask and creates a new one in the current level.
func (l *level) NewMask() {
	l.finish()
	l.Active = l.Starts
	l.activeDepth = l.startsDepth
}

// PushLevel creates a new level as a child to a current, and returns it.
// pos saves the position at which level was created. beforeDot tells that a
// dot follows the closing brace of the new level.
func (l *level) PushLevel(pos int, beforeDot bool) *level {
	nl := newLevel()
	nl.Pos = pos
	nl.Prev = l
	// The parent level gets the finished masks of the new level. Thus, the
	// new level also collects them when the parent level does.
	nl.collectEnds = beforeDot || l.collectEnds
	nl.Starts = l.Active
	nl.Active = l.Active
	nl.startsDepth = l.activeDepth
	nl.activeDepth = l.activeDepth
	return nl
}

// PopLevel finalizes the level, updates the parent level with the new
// unfinished masks and returns it. The parent level gets no masks when the
// level does not collect them, because no token can use them.
func (l *level) PopLevel() *level {
	l.finish()
	p := l.Prev
	p.Active = l.Ends
	p.activeDepth = l.endsDepth
	return p
}

// groupsBeforeDot finds the groups that a dot follows. It returns one value for
// each token. The value is true for the opening brace of such a group.
func groupsBeforeDot(tokens []token) []bool {
	ret := make([]bool, len(tokens))
	opened := []int{} // opened contains the indexes of the open braces.
	for i, tok := range tokens {
		switch tok.Type {
		case tokenLBrace:
			opened = append(opened, i)
		case tokenRBrace:
			if len(opened) == 0 {
				// The parser reports the unmatched brace.
				return ret
			}
			last := len(opened) - 1
			ret[opened[last]] = i+1 < len(tokens) &&
				tokens[i+1].Type == tokenDot
			opened = opened[:last]
		}
	}
	return ret
}

const (
	// stateKey represents the state of expecting a field key.
	stateKey = iota + 1
	// stateSeparator represents the state of expecting a separator (dot or
	// comma).
	stateSeparator
	// stateLevelStart represents the state of starting a new mask level.
	stateLevelStart
)

// parseStepsPerByte is the number of steps that [Parse] permits for each byte
// of the input.
const parseStepsPerByte = 8

// ErrTooComplex is the cause of the error that [Parse] returns when a mask
// needs more steps than [Parse] permits.
var ErrTooComplex = errors.New("mask is too complex")

// Parse parses the input mask string into a [Mask] object. It tokenizes the
// input string, processes the tokens based on the [Mask Syntax], and constructs
// the corresponding [Mask] structure representing the parsed mask.
//
// Parse limits its work, because a sequence of groups multiplies the number of
// fields. For example, "(a,b).(c,d).(e,f)" contains 14 fields. Parse does one
// step for each unfinished mask when it applies a key, a wildcard, a comma or a
// closing brace. Parse permits 8 steps for each byte of the input. Above this
// limit, Parse returns an [UnmarshalError] with the cause [ErrTooComplex]. The
// number of fields in the result is thus not more than 8 times the length of
// the input.
//
// The limit has no effect on a mask in which no dot follows a group. Such a
// mask needs one step for each token at most. [Mask.Marshal] writes only such
// masks.
//
// Parameters:
//
//	mask string - The input mask string to be parsed.
//
// Returns:
//
//   - *Mask - The parsed [Mask] object representing the structure defined by
//     the mask string.
//   - error - An error if there was any issue during parsing or tokenization.
//
// [Mask Syntax]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func Parse(mask string) (*Mask, error) {
	return parse(mask, parseStepsPerByte*int64(len(mask)))
}

// parse is [Parse] with an explicit limit for the number of steps.
func parse(mask string, maxSteps int64) (*Mask, error) { //nolint:funlen,gocognit
	if len(strings.TrimLeft(mask, spaceSet)) == 0 {
		// Return an empty mask if the input string is empty.
		return New(), nil
	}
	tokens, err := newLexer(mask).tokenize() // Tokenize the input mask string.
	if err != nil {
		return nil, fmt.Errorf("tokenizer failure: %w", err)
	}
	beforeDot := groupsBeforeDot(tokens)
	lvl := newLevel()     // Create a new mask level.
	root := lvl.Starts[0] // Get the root mask.
	pos := 0              // Initialize position tracker.
	state := stateKey     // Initialize state to expecting a field key.
	steps := int64(0)     // Count the visits of the unfinished masks.

	for pos < len(tokens) { // End the loop if the end of tokens is reached.
		tok := tokens[pos] // Get the current token.
		if state != stateLevelStart &&
			tok.Type != tokenDot && tok.Type != tokenLBrace {
			// All other tokens visit each unfinished mask. Count the visits
			// before the level creates or collects masks.
			steps += int64(len(lvl.Active))
			if steps > maxSteps {
				return nil, &UnmarshalError{
					Input:    mask,
					Position: tok.Pos,
					Summary:  "cannot expand mask",
					Cause:    ErrTooComplex,
				}
			}
		}
		switch state {
		case stateLevelStart:
			if tok.Type == tokenRBrace {
				// If it's an empty level, treat the next token as a separator.
				state = stateSeparator
			} else {
				state = stateKey // Expect a field key after opening brace.
			}
			continue // Don't move to the next token, just switch state.
		case stateKey:
			switch tok.Type {
			case tokenPlainKey:
				err = lvl.AddKey(FieldKey(tok.Value))
				state = stateSeparator
			case tokenWildCard:
				err = lvl.AddAny()
				state = stateSeparator
			case tokenLBrace:
				lvl = lvl.PushLevel(tok.Pos, beforeDot[pos])
				state = stateLevelStart
			default:
				return nil, &UnmarshalError{
					Input:    mask,
					Position: tok.Pos,
					Summary: fmt.Sprintf(
						"unexpected token %s, expecting field or submask", tok,
					),
				}
			}
			if err != nil {
				return nil, &UnmarshalError{
					Input:    mask,
					Position: tok.Pos,
					Summary:  "cannot expand mask",
					Cause:    err,
				}
			}
		case stateSeparator:
			switch tok.Type {
			case tokenDot:
				state = stateKey
			case tokenComma:
				lvl.NewMask()
				state = stateKey
			case tokenRBrace:
				if lvl.Prev == nil {
					return nil, &UnmarshalError{
						Input:    mask,
						Position: tok.Pos,
						Summary:  "unmatched right brace",
					}
				}
				lvl = lvl.PopLevel()
				state = stateSeparator
			default:
				return nil, &UnmarshalError{
					Input:    mask,
					Position: tok.Pos,
					Summary: fmt.Sprintf(
						"unexpected token %s, expecting separator or closing "+
							"brace",
						tok,
					),
				}
			}
		default:
			return nil, errors.New("state machine corruption")
		}
		pos++
	}
	if lvl.Prev != nil {
		return nil, &UnmarshalError{
			Input:    mask,
			Position: lvl.Pos,
			Summary:  "unclosed left brace",
		}
	}
	if state != stateSeparator {
		return nil, errors.New("unexpected end of mask")
	}

	return root, nil
}

// ParseMust is a wrapper around [Parse] that panics if [Parse] fails.
func ParseMust(mask string) *Mask {
	ret, err := Parse(mask)
	if err != nil {
		panic(err)
	}
	return ret
}

type format string

const (
	jsonFormat format = "json"
	yamlFormat format = "yaml"
)

// ParseJSON parses the input json into a [Mask] object. It finds the
// leafs, build their [FieldPath] and merge them to one [Mask].
//
// Parameters:
//
//	input []byte - The input json to be parsed.
//
// Returns:
//
//   - *Mask - The parsed [Mask] object representing the structure defined by
//     the json input.
//   - error - An error if there was any issue during parsing or tokenization.
//
// [Mask Syntax]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func ParseJSON(input []byte) (*Mask, error) {
	return parseFormat(input, jsonFormat)
}

// ParseYAML parses the input yaml into a [Mask] object. It finds the
// leafs, build their [FieldPath] and merge them to one [Mask].
//
// Parameters:
//
//	input []byte - The input yaml to be parsed.
//
// Returns:
//
//   - *Mask - The parsed [Mask] object representing the structure defined by
//     the yaml input.
//   - error - An error if there was any issue during parsing or tokenization.
//
// [Mask Syntax]: https://nebius.atlassian.net/wiki/spaces/NEWBIUS/pages/131367768/ResetMask+tech+design
func ParseYAML(input []byte) (*Mask, error) {
	return parseFormat(input, yamlFormat)
}

func parseFormat(input []byte, format format) (*Mask, error) {
	var data any

	switch format {
	case jsonFormat:
		if err := json.Unmarshal(input, &data); err != nil {
			return nil, fmt.Errorf("json unmarshal: %w", err)
		}
	case yamlFormat:
		if err := yaml.Unmarshal(input, &data); err != nil {
			return nil, fmt.Errorf("yaml unmarshal: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}

	ret := New()
	for _, fp := range extractLeafFP(data, nil) {
		if err := ret.Merge(fp.ToMask()); err != nil {
			return nil, fmt.Errorf("merge mask and field path: %w", err)
		}
	}
	return ret, nil
}

func extractLeafFP(data any, prefix FieldPath) []FieldPath {
	var fieldPaths []FieldPath

	switch v := data.(type) {
	case map[string]any:
		for key, value := range v {
			newPrefix := prefix.Join(FieldKey(key))
			fieldPaths = append(fieldPaths, extractLeafFP(value, newPrefix)...)
		}
	case []any:
		for i, value := range v {
			newPrefix := prefix.Join(FieldKey(strconv.Itoa(i)))
			childPaths := extractLeafFP(value, newPrefix)
			if len(childPaths) == 0 {
				fieldPaths = append(fieldPaths, newPrefix)
			} else {
				fieldPaths = append(fieldPaths, childPaths...)
			}
		}
	default:
		fieldPaths = append(fieldPaths, prefix)
	}

	return fieldPaths
}

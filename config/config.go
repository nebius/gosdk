package config

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const DefaultConfigDir = ".nebius/"
const defaultConfigName = "config.yaml"

const DefaultAPIEndpoint = "api.nebius.cloud"

type AuthType string

const (
	AuthTypeFederation     AuthType = "federation"
	AuthTypeServiceAccount AuthType = "service account"
)

type Config struct {
	Default  string         `yaml:"default,omitempty"`
	Profiles ProfilesConfig `yaml:"profiles"`
}

type ProfilesConfig map[string]*Profile

type Profile struct {
	Name string `yaml:"-"`

	Endpoint           string   `yaml:"endpoint,omitempty"`
	TokenFile          string   `yaml:"token-file,omitempty"`
	AuthType           AuthType `yaml:"auth-type,omitempty"`
	FederationEndpoint string   `yaml:"federation-endpoint,omitempty"`
	FederationID       string   `yaml:"federation-id,omitempty"`
	// TODO: move the following 4 values to a nested struct (config command doesn't support that)
	ServiceAccountID string `yaml:"service-account-id,omitempty"`
	PublicKeyID      string `yaml:"public-key-id,omitempty"`
	PrivateKey       string `yaml:"private-key,omitempty"`

	ParentID string `yaml:"parent-id,omitempty"`
}

func NewConfig() *Config {
	cfg := &Config{
		Profiles: make(ProfilesConfig),
	}
	return cfg
}

func (c *Config) LoadFromFile(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return NewMissingConfigError(err)
		}
		return NewError(err)
	}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return NewError(err)
	}
	for name, profile := range c.Profiles {
		profile.Name = name
	}
	return nil
}

func (c *Config) GetDefaultProfile() (*Profile, error) {
	if c.Default == "" {
		switch len(c.Profiles) {
		case 0:
			return nil, NewError(NewGetProfileError(errors.New("no profile configured"), c.Profiles))
		case 1:
			for _, profile := range c.Profiles {
				return profile, nil
			}
		default:
			return nil, NewError(NewGetProfileError(
				errors.New("malformed configuration: default profile is empty"),
				c.Profiles,
			))
		}
	}
	profile, err := c.GetProfile(c.Default)
	if err != nil {
		return nil, NewError(err)
	}
	return profile, nil
}

func (c *Config) HasProfile(name string) bool {
	_, exists := c.Profiles[name]
	return exists
}

func (c *Config) GetProfile(name string) (*Profile, error) {
	profile, exists := c.Profiles[name]
	if !exists {
		return nil, NewGetProfileError(fmt.Errorf("unknown profile %q", name), c.Profiles)
	}
	return profile, nil
}

func GetDefaultConfigPath() (string, error) {
	home, err := UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, DefaultConfigDir, defaultConfigName), nil
}

func UserHomeDir() (string, error) {
	home, envErr := os.UserHomeDir()
	if envErr == nil {
		return home, nil
	}
	usr, uidErr := user.Current()
	if uidErr != nil {
		return "", errors.Join(envErr, uidErr)
	}
	return usr.HomeDir, nil
}

func GetAbsoluteDefaultConfigDir() (string, error) {
	home, err := UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, DefaultConfigDir), nil
}

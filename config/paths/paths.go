package paths

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
)

const DefaultConfigDir = ".nebius/"
const defaultConfigName = "config.yaml"

const DefaultAPIEndpoint = "api.nebius.cloud"

func GetDefaultConfigPath() (string, error) {
	home, err := UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, DefaultConfigDir, defaultConfigName), nil
}

func ExpandHomeDir(path string) (string, error) {
	if len(path) == 0 {
		return path, nil
	}

	if path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != os.PathSeparator && path[1] != '/' {
		return "", errors.New("cannot expand user-specific home dir")
	}

	dir, err := UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, path[1:]), nil
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

/*
Copyright 2022-2023 Rishvic Pushpakaran

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package xdg

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/shlex"
)

var XDGUnexpectedError = errors.New("Unexpected error while resolving XDG DIR")

func XDGCacheHome() (string, error)  { return envThenDef("XDG_CACHE_HOME", ".cache") }
func XDGConfigHome() (string, error) { return envThenDef("XDG_CONFIG_HOME", ".config") }
func XDGDataHome() (string, error)   { return envThenDef("XDG_DATA_HOME", ".local", "share") }
func XDGRuntimeDir() (string, error) { return envOnly("XDG_RUNTIME_DIR") }
func XDGStateHome() (string, error)  { return envThenDef("XDG_STATE_HOME", ".local", "state") }
func XDGBinHome() (string, error)    { return envThenDef("XDG_BIN_HOME", ".local", "bin") }

func XDGUserDir(dirType string) (string, error) {
	if dir, err := userDirLookup(dirType); err == nil {
		return dir, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "/tmp", nil
	}
	if dirType == "DESKTOP" {
		return filepath.Join(homeDir, "Desktop"), nil
	}
	return homeDir, nil
}

var errNoValidFound = errors.New("No valid set value found")

func envOnly(env string) (string, error) {
	if runtime.GOOS == "plan9" {
		env = strings.ToLower(env)
	}
	// Check environment variable only
	// According to XDG spec, any specified path needs to be absolute
	if path, ok := os.LookupEnv(env); ok && filepath.IsAbs(path) {
		return path, nil
	}
	return "", errNoValidFound
}

func envThenDef(env string, path ...string) (string, error) {
	if dir, err := envOnly(env); err == nil {
		return dir, nil
	}

	// If did not find appropriate env var, go for default
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Join(errors.New("Could not resolve user home directory"), err)
	}
	elem := append([]string{homeDir}, path...)
	return filepath.Join(elem...), nil
}

func userDirLookup(dirType string) (string, error) {
	// Fetch user home directory for relative paths
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", errors.Join(errors.New("Could not resolve user home directory"), err)
	}

	// Fetch XDG_CONFIG_DIR for finding user-dirs.dir
	path, err := XDGConfigHome()
	if err != nil {
		return "", errors.Join(errors.New("Could not resolve XDG_CONFIG_DIR"), err)
	}

	path = filepath.Join(path, "user-dirs.dirs")
	file, err := os.Open(path)
	if err != nil {
		return "", errors.Join(errors.New("Could not open user-dirs.dir"), err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		subStrs, err := shlex.Split(line)
		if err != nil || len(subStrs) == 0 {
			continue
		}

		key, val, found := strings.Cut(subStrs[0], "=")
		if !found {
			continue
		}
		if key != "XDG_"+dirType+"_DIR" {
			continue
		}

		var relative bool
		if strings.HasPrefix(val, "$HOME/") {
			relative = true
		} else if !filepath.IsAbs(val) {
			continue
		}

		if relative {
			val = filepath.Join(homeDir, val[len("$HOME/"):])
		}

		return val, nil
	}

	return "", errNoValidFound
}

//go:build !(darwin || windows)

/*
Copyright 2022 Rishvic Pushpakaran

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
package dirs

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/google/shlex"
	"github.com/mitchellh/go-homedir"
)

func HomeDir() (string, error) { return homedir.Dir() }

func CacheDir() (string, error)      { return envThenDef("XDG_CACHE_HOME", ".cache") }
func ConfigDir() (string, error)     { return envThenDef("XDG_CONFIG_HOME", ".config") }
func DataDir() (string, error)       { return envThenDef("XDG_DATA_HOME", ".local", "share") }
func DataLocalDir() (string, error)  { return DataDir() }
func PreferenceDir() (string, error) { return ConfigDir() }
func RuntimeDir() (string, error)    { return envOnly("XDG_RUNTIME_DIR") }
func StateDir() (string, error)      { return envThenDef("XDG_STATE_HOME", ".local", "state") }
func ExecutableDir() (string, error) { return envThenDef("XDG_BIN_HOME", ".local", "bin") }

func AudioDir() (string, error)    { return userDir("MUSIC") }
func DesktopDir() (string, error)  { return userDir("DESKTOP") }
func DocumentDir() (string, error) { return userDir("DOCUMENTS") }
func DownloadDir() (string, error) { return userDir("DOWNLOAD") }
func FontDir() (string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataDir, "fonts"), nil
}
func PictureDir() (string, error)  { return userDir("PICTURES") }
func PublicDir() (string, error)   { return userDir("PUBLICSHARE") }
func TemplateDir() (string, error) { return userDir("TEMPLATES") }
func VideoDir() (string, error)    { return userDir("VIDEOS") }

func envOnly(env string) (string, error) {
	if runtime.GOOS == "plan9" {
		env = strings.ToLower(env)
	}
	// Check environment variable only
	if path, ok := os.LookupEnv(env); ok && filepath.IsAbs(path) {
		return path, nil
	}
	return "", ErrNotFound
}

func envThenDef(env string, path ...string) (string, error) {
	if dir, err := envOnly(env); err == nil {
		return dir, nil
	}

	// If did not find appropriate env var, go for default
	homeDir, err := HomeDir()
	if err != nil {
		return "", err
	}
	elem := make([]string, 1, len(path)+1)
	elem[0] = homeDir
	elem = append(elem, path...)
	return filepath.Join(elem...), nil
}

func userDirLookup(dirType string) (val string, ok bool) {
	path, err := ConfigDir()
	if err != nil {
		return "", false
	}

	path = filepath.Join(path, "user-dirs.dirs")
	file, err := os.Open(path)
	if err != nil {
		return "", false
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
			homeDir, err := HomeDir()
			if err != nil {
				return "", false
			}
			val = filepath.Join(homeDir, val[len("$HOME/"):])
		}

		return val, true
	}

	return "", false
}

func userDir(dirType string) (string, error) {
	if dir, ok := userDirLookup(dirType); ok {
		return dir, nil
	}

	homeDir, err := HomeDir()
	if err != nil {
		return "", err
	}
	if dirType == "DESKTOP" {
		return filepath.Join(homeDir, "Desktop"), nil
	}
	return homeDir, nil
}

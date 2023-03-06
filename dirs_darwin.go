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
package dirs

import (
	"os"
	"path/filepath"
)

func HomeDir() (string, error) { return os.UserHomeDir() }

func CacheDir() (string, error)      { return joinHome("Library", "Caches") }
func ConfigDir() (string, error)     { return joinHome("Library", "Application Support") }
func DataDir() (string, error)       { return joinHome("Library", "Application Support") }
func DataLocalDir() (string, error)  { return DataDir() }
func PreferenceDir() (string, error) { return joinHome("Library", "Preferences") }
func ExecutableDir() (string, error) { return "", ErrNotFound }
func RuntimeDir() (string, error)    { return "", ErrNotFound }
func StateDir() (string, error)      { return "", ErrNotFound }

func AudioDir() (string, error)    { return joinHome("Music") }
func DesktopDir() (string, error)  { return joinHome("Desktop") }
func DocumentDir() (string, error) { return joinHome("Documents") }
func DownloadDir() (string, error) { return joinHome("Downloads") }
func FontDir() (string, error)     { return joinHome("Library", "Fonts") }
func PictureDir() (string, error)  { return joinHome("Pictures") }
func PublicDir() (string, error)   { return joinHome("Public") }
func TemplateDir() (string, error) { return "", ErrNotFound }
func VideoDir() (string, error)    { return joinHome("Movies") }

func joinHome(path ...string) (string, error) {
	homeDir, err := HomeDir()
	if err != nil {
		return "", err
	}

	elem := append([]string{homeDir}, path...)
	return filepath.Join(elem...), nil
}

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
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func HomeDir() (string, error) { return homedir.Dir() }

func CacheDir() (string, error)      { return cacheCache.cur(joinHome("Library", "Caches")) }
func ConfigDir() (string, error)     { return configCache.cur(joinHome("Library", "Application Support")) }
func DataDir() (string, error)       { return dataCache.cur(joinHome("Library", "Application Support")) }
func DataLocalDir() (string, error)  { return DataDir() }
func PreferenceDir() (string, error) { return preferenceCache.cur(joinHome("Library", "Preferences")) }
func ExecutableDir() (string, error) { return "", ErrNotFound }
func RuntimeDir() (string, error)    { return "", ErrNotFound }
func StateDir() (string, error)      { return "", ErrNotFound }

func AudioDir() (string, error)    { return audioCache.cur(joinHome("Music")) }
func DesktopDir() (string, error)  { return desktopCache.cur(joinHome("Desktop")) }
func DocumentDir() (string, error) { return documentCache.cur(joinHome("Documents")) }
func DownloadDir() (string, error) { return downloadCache.cur(joinHome("Downloads")) }
func FontDir() (string, error)     { return fontCache.cur(joinHome("Library", "Fonts")) }
func PictureDir() (string, error)  { return pictureCache.cur(joinHome("Pictures")) }
func PublicDir() (string, error)   { return publicCache.cur(joinHome("Public")) }
func TemplateDir() (string, error) { return "", ErrNotFound }
func VideoDir() (string, error)    { return videoCache.cur(joinHome("Movies")) }

func joinHome(path ...string) func() (string, error) {
	return func() (string, error) {
		homeDir, err := HomeDir()
		if err != nil {
			return "", err
		}

		elem := make([]string, 1, len(path)+1)
		elem[0] = homeDir
		elem = append(elem, path...)
		return filepath.Join(elem...), nil
	}
}

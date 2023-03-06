//go:build !(darwin || windows)

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

	"github.com/Colocasian/dirs-go/pkg/xdg"
)

func HomeDir() (string, error) { return os.UserHomeDir() }

func CacheDir() (string, error)      { return xdg.XDGCacheHome() }
func ConfigDir() (string, error)     { return xdg.XDGConfigHome() }
func DataDir() (string, error)       { return xdg.XDGDataHome() }
func DataLocalDir() (string, error)  { return xdg.XDGDataHome() }
func PreferenceDir() (string, error) { return xdg.XDGConfigHome() }
func RuntimeDir() (string, error)    { return xdg.XDGRuntimeDir() }
func StateDir() (string, error)      { return xdg.XDGStateHome() }
func ExecutableDir() (string, error) { return xdg.XDGBinHome() }

func AudioDir() (string, error)    { return xdg.XDGUserDir("MUSIC") }
func DesktopDir() (string, error)  { return xdg.XDGUserDir("DESKTOP") }
func DocumentDir() (string, error) { return xdg.XDGUserDir("DOCUMENTS") }
func DownloadDir() (string, error) { return xdg.XDGUserDir("DOWNLOAD") }
func FontDir() (string, error) {
	dataDir, err := DataDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dataDir, "fonts"), nil
}
func PictureDir() (string, error)  { return xdg.XDGUserDir("PICTURES") }
func PublicDir() (string, error)   { return xdg.XDGUserDir("PUBLICSHARE") }
func TemplateDir() (string, error) { return xdg.XDGUserDir("TEMPLATES") }
func VideoDir() (string, error)    { return xdg.XDGUserDir("VIDEOS") }

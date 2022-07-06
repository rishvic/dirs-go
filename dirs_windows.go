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

import "golang.org/x/sys/windows"

func HomeDir() (string, error) { return homeCache.cur(kf(windows.FOLDERID_Profile)) }

func DataDir() (string, error)       { return dataCache.cur(kf(windows.FOLDERID_RoamingAppData)) }
func DataLocalDir() (string, error)  { return dataLocalCache.cur(kf(windows.FOLDERID_LocalAppData)) }
func CacheDir() (string, error)      { return DataLocalDir() }
func ConfigDir() (string, error)     { return DataDir() }
func ExecutableDir() (string, error) { return "", ErrNotFound }
func PreferenceDir() (string, error) { return DataDir() }
func RuntimeDir() (string, error)    { return "", ErrNotFound }
func StateDir() (string, error)      { return "", ErrNotFound }

func AudioDir() (string, error)    { return audioCache.cur(kf(windows.FOLDERID_Music)) }
func DesktopDir() (string, error)  { return desktopCache.cur(kf(windows.FOLDERID_Desktop)) }
func DocumentDir() (string, error) { return documentCache.cur(kf(windows.FOLDERID_Documents)) }
func DownloadDir() (string, error) { return downloadCache.cur(kf(windows.FOLDERID_Downloads)) }
func FontDir() (string, error)     { return "", ErrNotFound }
func PictureDir() (string, error)  { return pictureCache.cur(kf(windows.FOLDERID_Pictures)) }
func PublicDir() (string, error)   { return publicCache.cur(kf(windows.FOLDERID_Public)) }
func TemplateDir() (string, error) { return templateCache.cur(kf(windows.FOLDERID_Templates)) }
func VideoDir() (string, error)    { return videoCache.cur(kf(windows.FOLDERID_Videos)) }

var (
	homeCache,
	dataCache,
	dataLocalCache,

	audioCache,
	desktopCache,
	documentCache,
	downloadCache,
	fontCache,
	pictureCache,
	publicCache,
	templateCache,
	videoCache cache
)

func kf(folderID *windows.KNOWNFOLDERID) func() (string, error) {
	return func() (string, error) {
		return windows.KnownFolderPath(folderID, 0)
	}
}

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

import "testing"

func TestDirs(t *testing.T) {
	fs := []func() (string, error){
		HomeDir,

		CacheDir,
		ConfigDir,
		DataDir,
		DataLocalDir,
		PreferenceDir,
		RuntimeDir,
		StateDir,
		ExecutableDir,

		AudioDir,
		DesktopDir,
		DocumentDir,
		DownloadDir,
		FontDir,
		PictureDir,
		PublicDir,
		TemplateDir,
		VideoDir,
	}

	names := []string{
		"HomeDir",

		"CacheDir",
		"ConfigDir",
		"DataDir",
		"DataLocalDir",
		"PreferenceDir",
		"RuntimeDir",
		"StateDir",
		"ExecutableDir",

		"AudioDir",
		"DesktopDir",
		"DocumentDir",
		"DownloadDir",
		"FontDir",
		"PictureDir",
		"PublicDir",
		"TemplateDir",
		"VideoDir",
	}

	for i := range fs {
		dir, err := fs[i]()
		t.Logf("%v:%*s%q %v", names[i], 15-len(names[i]), " ", dir, err)
		if i == 0 || i == 8 {
			t.Logf("")
		}
	}
}

func TestCaching(t *testing.T) {
	testcases := [][2]string{
		{"1", "2"},
		{"Lorem ipsum", "dolor sit amet"},
		{"Is this the real life?", "Is this just fantasy?"},
	}

	for _, tc := range testcases {
		m := newMockF(tc[:])
		c := NewCache(m.str)
		// Load s1 into cache
		s1, err := c.Cur()
		if err != nil || s1 != tc[0] {
			t.Errorf("Dir: (%q, %v), want (%q, nil)", s1, err, tc[0])
		}

		// If s2 is different from s1, check if cache uses previous data
		if tc[1] == tc[0] {
			t.Log("testcase's s1 == s2, continuing...")
			continue
		}
		s2, err := c.Cur()
		if err != nil || s2 != tc[0] {
			t.Errorf("Cached Dir: (%q, %v), want (%q, nil)", s2, err, tc[0])
		}
	}
}

func TestCacheReset(t *testing.T) {
	testcases := [][2]string{
		{"1", "2"},
		{"Lorem ipsum", "dolor sit amet"},
		{"Is this the real life?", "Is this just fantasy?"},
	}

	for _, tc := range testcases {
		m := newMockF(tc[:])
		c := NewCache(m.str)
		// Load s1 into cache
		s1, err := c.Cur()
		if err != nil || s1 != tc[0] {
			t.Errorf("Dir: (%q, %v), want (%q, nil)", s1, err, tc[0])
		}

		// Reset cache, and test with s2
		if tc[1] == tc[0] {
			t.Log("testcase's s1 == s2, continuing...")
			continue
		}
		c.Reset()
		s2, err := c.Cur()
		if err != nil || s2 != tc[1] {
			t.Errorf("Reset Dir: (%q, %v), want (%q, nil)", s2, err, tc[1])
		}
	}
}

type mockF struct {
	i    int
	strs []string
}

func newMockF(strs []string) *mockF {
	return &mockF{strs: strs}
}

func (m *mockF) str() (string, error) {
	defer func() { m.i = (m.i + 1) % len(m.strs) }()
	return m.strs[m.i], nil
}

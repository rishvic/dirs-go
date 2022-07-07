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
	"errors"
	"sync"
)

var ErrNotFound = errors.New("directory not found")

type Cache struct {
	f    func() (string, error)
	once sync.Once
	str  string
	err  error
}

func NewCache(f func() (string, error)) *Cache {
	return &Cache{f: f}
}

func (c *Cache) Cur() (string, error) {
	c.once.Do(func() { c.str, c.err = c.f() })
	if c.err != nil {
		return "", c.err
	}
	return c.str, nil
}

func (c *Cache) Reset() {
	var again sync.Once
	c.once = again
}

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

type cache struct {
	once sync.Once
	dir  string
	err  error
}

func (c *cache) Dir(f func() (string, error)) (string, error) {
	c.once.Do(func() { c.dir, c.err = f() })
	if c.err != nil {
		return "", c.err
	}
	return c.dir, nil
}

// Copyright 2015 The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package source

import (
	"bytes"
	"fmt"
)

type ByteSource struct {
	Name    string
	Content []byte
}

func (b *ByteSource) String() string {
	return fmt.Sprintf("%s %s", b.Name, string(b.Content))
}

type InMemorySource struct {
	ByteSource []ByteSource
}

func (i *InMemorySource) Files() (files []*File) {
	files = make([]*File, len(i.ByteSource))
	for i, fake := range i.ByteSource {
		files[i] = NewFileWithContents(fake.Name, bytes.NewReader(fake.Content))
	}
	return
}

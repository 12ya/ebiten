// Copyright 2025 The Ebitengine Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package textutil

import (
	"iter"
	"strings"
	"unicode/utf8"

	"github.com/rivo/uniseg"
)

func Lines(str string) iter.Seq[string] {
	return func(yield func(s string) bool) {
		var line string
		state := -1
		for len(str) > 0 {
			segment, nextStr, mustBreak, nextState := uniseg.FirstLineSegmentInString(str, state)
			line += segment
			if mustBreak {
				if !yield(line) {
					return
				}
				line = ""
			}
			state = nextState
			str = nextStr
		}
		if len(line) > 0 {
			if !yield(line) {
				return
			}
		}
	}
}

func TrimTailingLineBreak(str string) string {
	if !uniseg.HasTrailingLineBreakInString(str) {
		return str
	}

	// https://en.wikipedia.org/wiki/Newline#Unicode
	if strings.HasSuffix(str, "\r\n") {
		return str[:len(str)-2]
	}

	_, s := utf8.DecodeLastRuneInString(str)
	return str[:len(str)-s]
}

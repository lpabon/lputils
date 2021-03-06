//
// Copyright (c) 2015 Luis Pabón <lpabon@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package lputils

import (
	"testing"

	"github.com/lpabon/lputils/tests"
)

func TestNewStringSet(t *testing.T) {
	s := NewStringSet()
	tests.Assert(t, s.Set != nil)
	tests.Assert(t, len(s.Set) == 0)
}

func TestStringSet(t *testing.T) {
	s := NewStringSet()

	s.Add("one")
	s.Add("two")
	s.Add("three")
	tests.Assert(t, s.Len() == 3)
	tests.Assert(t, SortedStringHas(s.Set, "one"))
	tests.Assert(t, SortedStringHas(s.Set, "two"))
	tests.Assert(t, SortedStringHas(s.Set, "three"))

	s.Add("one")
	tests.Assert(t, s.Len() == 3)
	tests.Assert(t, SortedStringHas(s.Set, "one"))
	tests.Assert(t, SortedStringHas(s.Set, "two"))
	tests.Assert(t, SortedStringHas(s.Set, "three"))

	s.Add("three")
	tests.Assert(t, s.Len() == 3)
	tests.Assert(t, SortedStringHas(s.Set, "one"))
	tests.Assert(t, SortedStringHas(s.Set, "two"))
	tests.Assert(t, SortedStringHas(s.Set, "three"))

	s.Add("four")
	tests.Assert(t, s.Len() == 4)
	tests.Assert(t, SortedStringHas(s.Set, "one"))
	tests.Assert(t, SortedStringHas(s.Set, "two"))
	tests.Assert(t, SortedStringHas(s.Set, "three"))
	tests.Assert(t, SortedStringHas(s.Set, "four"))
}

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
	"sort"
	"testing"

	"github.com/lpabon/lputils/tests"
)

func TestSortedStringsHas(t *testing.T) {
	s := sort.StringSlice{"z", "b", "a"}
	s.Sort()
	tests.Assert(t, len(s) == 3)
	tests.Assert(t, s[0] == "a")
	tests.Assert(t, s[1] == "b")
	tests.Assert(t, s[2] == "z")

	tests.Assert(t, SortedStringHas(s, "a"))
	tests.Assert(t, SortedStringHas(s, "b"))
	tests.Assert(t, SortedStringHas(s, "z"))
	tests.Assert(t, !SortedStringHas(s, "c"))
	tests.Assert(t, !SortedStringHas(s, "zz"))
}

func TestSortedStringsDelete(t *testing.T) {
	s := sort.StringSlice{"z", "b", "a"}
	s.Sort()
	tests.Assert(t, len(s) == 3)
	tests.Assert(t, s[0] == "a")
	tests.Assert(t, s[1] == "b")
	tests.Assert(t, s[2] == "z")

	tests.Assert(t, SortedStringHas(s, "a"))
	tests.Assert(t, SortedStringHas(s, "b"))
	tests.Assert(t, SortedStringHas(s, "z"))
	tests.Assert(t, !SortedStringHas(s, "c"))
	tests.Assert(t, !SortedStringHas(s, "zz"))

	s = SortedStringsDelete(s, "notthere")
	tests.Assert(t, len(s) == 3)
	s = SortedStringsDelete(s, "zzzznotthere")
	tests.Assert(t, len(s) == 3)
	s = SortedStringsDelete(s, "1azzzznotthere")
	tests.Assert(t, len(s) == 3)
	tests.Assert(t, SortedStringHas(s, "a"))
	tests.Assert(t, SortedStringHas(s, "b"))
	tests.Assert(t, SortedStringHas(s, "z"))
	tests.Assert(t, !SortedStringHas(s, "c"))
	tests.Assert(t, !SortedStringHas(s, "zz"))

	s = SortedStringsDelete(s, "z")
	tests.Assert(t, len(s) == 2)
	tests.Assert(t, SortedStringHas(s, "a"))
	tests.Assert(t, SortedStringHas(s, "b"))
	tests.Assert(t, !SortedStringHas(s, "z"))
	tests.Assert(t, !SortedStringHas(s, "c"))
	tests.Assert(t, !SortedStringHas(s, "zz"))

	s = SortedStringsDelete(s, "a")
	tests.Assert(t, len(s) == 1)
	tests.Assert(t, !SortedStringHas(s, "a"))
	tests.Assert(t, SortedStringHas(s, "b"))
	tests.Assert(t, !SortedStringHas(s, "z"))
	tests.Assert(t, !SortedStringHas(s, "c"))
	tests.Assert(t, !SortedStringHas(s, "zz"))

	s = SortedStringsDelete(s, "b")
	tests.Assert(t, len(s) == 0)
	tests.Assert(t, !SortedStringHas(s, "a"))
	tests.Assert(t, !SortedStringHas(s, "b"))
	tests.Assert(t, !SortedStringHas(s, "z"))
	tests.Assert(t, !SortedStringHas(s, "c"))
	tests.Assert(t, !SortedStringHas(s, "zz"))

}

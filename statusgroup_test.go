//
// Copyright (c) 2015 The heketi Authors
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

package utils

import (
	"errors"
	"fmt"
	"github.com/heketi/tests"
	"testing"
	"time"
)

func TestNewStatusGroup(t *testing.T) {
	s := NewStatusGroup()
	tests.Assert(t, s != nil)
	tests.Assert(t, s.results != nil)
	tests.Assert(t, len(s.results) == 0)
	tests.Assert(t, s.err == nil)
}

func TestStatusGroupSuccess(t *testing.T) {

	s := NewStatusGroup()

	max := 100
	s.Add(max)

	for i := 0; i < max; i++ {
		go func(value int) {
			defer s.Done()
			time.Sleep(time.Millisecond * 1 * time.Duration(value))
		}(i)
	}

	err := s.Result()
	tests.Assert(t, err == nil)

}

func TestStatusGroupFailure(t *testing.T) {
	s := NewStatusGroup()

	for i := 0; i < 100; i++ {

		s.Add(1)
		go func(value int) {
			defer s.Done()
			time.Sleep(time.Millisecond * 1 * time.Duration(value))
			if value%10 == 0 {
				s.Err(errors.New(fmt.Sprintf("Err: %v", value)))
			}

		}(i)

	}

	err := s.Result()

	tests.Assert(t, err != nil)
	tests.Assert(t, err.Error() == "Err: 90", err)

}

func TestStatusGroupAbort(t *testing.T) {
	s := NewStatusGroup()
	var last int

	for i := 1; i < 100; i++ {

		s.Add(1)
		go func(value int) {
			defer s.Done()
			time.Sleep(time.Millisecond * 1 * time.Duration(value))
			for !s.Abort() {
				if value%10 == 0 {
					s.Err(errors.New(fmt.Sprintf("Err: %v", value)))
					last = value
				}
			}

		}(i)

	}

	err := s.Result()

	tests.Assert(t, err != nil)
	tests.Assert(t, err.Error() == "Err: 10", err)
	tests.Assert(t, last == 10)
}

func TestResultFailFast(t *testing.T) {
	s := NewStatusGroup()
	var last int

	for i := 1; i < 100; i++ {

		s.Add(1)
		go func(value int) {
			defer s.Done()
			time.Sleep(time.Millisecond * 1 * time.Duration(value))
			if value%10 == 0 {
				s.Err(errors.New(fmt.Sprintf("Err: %v", value)))
				last = value
			}

		}(i)

	}

	err := s.ResultFailFast()

	tests.Assert(t, err != nil)
	tests.Assert(t, err.Error() == "Err: 10", err)
	tests.Assert(t, last == 10)
}

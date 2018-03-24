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
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/lpabon/lputils/tests"
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

func TestStatusGroupAbortNoError(t *testing.T) {
	s := NewStatusGroup()

	for i := 0; i < 100; i++ {

		s.Add(1)
		go func(value int) {
			defer s.Done()
			for j := 0; j < 10 && !s.Aborted(); j++ {
				time.Sleep(time.Millisecond * 1 * time.Duration(value))
				s.Err(nil)
			}

		}(i)

	}

	err := s.Result()
	tests.Assert(t, err == nil)
}

func TestStatusGroupAbort(t *testing.T) {
	s := NewStatusGroup()

	wait := make(chan bool)
	go func() {
		wait <- true
		s.Err(fmt.Errorf("ERROR"))
	}()
	<-wait

	for i := 0; i < 100; i++ {
		s.Add(1)
		go func() {
			defer s.Done()
			if !s.Aborted() {
				panic("Should not be here")
			} else {
				s.Err(nil)
			}
		}()
	}

	err := s.Result()

	tests.Assert(t, err != nil)
	tests.Assert(t, err.Error() == "ERROR")
}

func TestResultFailFast(t *testing.T) {
	s := NewStatusGroup()

	wait := make(chan bool)
	go func() {
		wait <- true
		s.Err(fmt.Errorf("ERROR"))
	}()
	<-wait

	for i := 1; i < 100; i++ {

		s.Add(1)
		go func(value int) {
			defer s.Done()
			time.Sleep(time.Millisecond * 1 * time.Duration(value))
			if value%10 == 0 {
				s.Err(errors.New(fmt.Sprintf("Err: %v", value)))
			}
		}(i)

	}

	err := s.ResultFailFast()
	tests.Assert(t, err != nil)
	tests.Assert(t, err.Error() == "ERROR")
}

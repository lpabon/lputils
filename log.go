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
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/lpabon/godbc"
)

type LogLevel int

// Log levels
const (
	LEVEL_NOLOG LogLevel = iota
	LEVEL_CRITICAL
	LEVEL_ERROR
	LEVEL_WARNING
	LEVEL_INFO
	LEVEL_DEBUG
)

var (
	stderr io.Writer = os.Stderr
	stdout io.Writer = os.Stdout
)

type Logger struct {
	critlog, errorlog, infolog *log.Logger
	debuglog, warninglog       *log.Logger

	level LogLevel
}

func logWithLongFile(l *log.Logger, format string, v ...interface{}) {
	_, file, line, _ := runtime.Caller(2)
	l.Print(fmt.Sprintf("%v:%v: ", path.Base(file), line) +
		fmt.Sprintf(format, v...))
}

// Create a new logger
func NewLogger(prefix string, level LogLevel) *Logger {
	godbc.Require(level >= 0, level)
	godbc.Require(level <= LEVEL_DEBUG, level)

	l := &Logger{}

	if level == LEVEL_NOLOG {
		l.level = LEVEL_DEBUG
	} else {
		l.level = level
	}

	l.critlog = log.New(stderr, prefix+" CRITICAL ", log.LstdFlags)
	l.errorlog = log.New(stderr, prefix+" ERROR ", log.LstdFlags)
	l.warninglog = log.New(stdout, prefix+" WARNING ", log.LstdFlags)
	l.infolog = log.New(stdout, prefix+" INFO ", log.LstdFlags)
	l.debuglog = log.New(stdout, prefix+" DEBUG ", log.LstdFlags)

	godbc.Ensure(l.critlog != nil)
	godbc.Ensure(l.errorlog != nil)
	godbc.Ensure(l.warninglog != nil)
	godbc.Ensure(l.infolog != nil)
	godbc.Ensure(l.debuglog != nil)

	return l
}

// Return current level
func (l *Logger) Level() LogLevel {
	return l.level
}

// Set level
func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

// Log critical information
func (l *Logger) Critical(format string, v ...interface{}) error {
	if l.level >= LEVEL_CRITICAL {
		logWithLongFile(l.critlog, format, v...)
	}

	return fmt.Errorf(format, v...)
}

// Log error string
func (l *Logger) LogError(format string, v ...interface{}) error {
	if l.level >= LEVEL_ERROR {
		logWithLongFile(l.errorlog, format, v...)
	}

	return fmt.Errorf(format, v...)
}

// Log error variable
func (l *Logger) Err(err error) error {
	if l.level >= LEVEL_ERROR {
		logWithLongFile(l.errorlog, "%v", err)
	}

	return err
}

// Log warning information
func (l *Logger) Warning(format string, v ...interface{}) error {
	if l.level >= LEVEL_WARNING {
		logWithLongFile(l.warninglog, format, v...)
	}

	return fmt.Errorf(format, v...)
}

// Log error variable as a warning
func (l *Logger) WarnErr(err error) error {
	if l.level >= LEVEL_WARNING {
		logWithLongFile(l.warninglog, "%v", err)
	}

	return err
}

// Log string
func (l *Logger) Info(format string, v ...interface{}) {
	if l.level >= LEVEL_INFO {
		logWithLongFile(l.infolog, format, v...)
	}
}

// Log string as debug
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.level >= LEVEL_DEBUG {
		logWithLongFile(l.debuglog, format, v...)
	}
}

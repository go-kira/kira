// Copyright 2018 Rachid Lafriakh
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

package log

import (
	"fmt"
	"sync"
	"time"
)

// Strings mapping.
var Strings = [...]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	FatalLevel: "FATAL",
	PanicLevel: "PANIC",
}

// DefaultLogFormatter - default log formatter
type DefaultLogFormatter struct {
	mu sync.Mutex
}

// NewDefaultFormatter ...
func NewDefaultFormatter() *DefaultLogFormatter {
	return &DefaultLogFormatter{}
}

// Format - it's format the output log
func (d *DefaultLogFormatter) Format(log *Logger, l Level, msg interface{}, t time.Time) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	level := Strings[l]

	// if the writer not terminal
	_, err := fmt.Fprintf(log.Writer, "%s %s %s", time.Now().Format(time.RFC3339Nano), level, msg)
	if err != nil {
		return err
	}

	if len(log.fields) > 0 {
		fmt.Fprint(log.Writer, " | ")
		for name, field := range log.fields {
			_, err := fmt.Fprintf(log.Writer, "%s=%s ", name, field)
			if err != nil {
				return err
			}
		}
		fmt.Fprint(log.Writer, "\n")
	}

	return nil
}

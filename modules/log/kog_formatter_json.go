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
	"encoding/json"
	"sync"
	"time"
)

// JSONFormatter to format the log into json format.
type JSONFormatter struct {
	Level   string      `json:"level,omitempty"`
	Time    string      `json:"time,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Fields  Fields      `json:"fields,omitempty"`
}

// JSONLogFormatter - default log formatter
type JSONLogFormatter struct {
	mu sync.Mutex
}

// NewJSONFormatter ...
func NewJSONFormatter() *JSONLogFormatter {
	return &JSONLogFormatter{}
}

// Format - it's format the output log
func (d *JSONLogFormatter) Format(log *Logger, l Level, msg interface{}, t time.Time) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	level := Strings[l]

	// Struct
	st := JSONFormatter{
		Time:    time.Now().Format(time.RFC3339Nano),
		Level:   level,
		Message: msg,
		Fields:  log.fields,
	}

	// if the writer not terminal
	err := json.NewEncoder(log.Writer).Encode(st)
	if err != nil {
		return err
	}

	return nil
}

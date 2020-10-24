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

// Level of severity.
type Level int

const (
	// DebugLevel - Detailed debug information.
	DebugLevel Level = iota
	// InfoLevel - Interesting events. Examples: User logs in, SQL logs.
	InfoLevel
	// WarnLevel - Exceptional occurrences that are not errors. Examples: Use of deprecated APIs, poor use of an API, undesirable things that are not necessarily wrong.
	WarnLevel
	// ErrorLevel - Runtime errors that do not require immediate action but should typically be logged and monitored.
	ErrorLevel
	// FatalLevel - Logs and then calls `os.Exit(1)`.
	FatalLevel
	// PanicLevel - highest level of severity
	PanicLevel
)

var levelNames = [...]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warn",
	ErrorLevel: "error",
	FatalLevel: "fatal",
	PanicLevel: "panic",
}

// LevelStrings - to return level type by level name
var LevelStrings = map[string]Level{
	"debug": DebugLevel,
	"info":  InfoLevel,
	"warn":  WarnLevel,
	"error": ErrorLevel,
	"fatal": FatalLevel,
	"panic": PanicLevel,
}

func (l Level) String() string {
	return levelNames[l]
}

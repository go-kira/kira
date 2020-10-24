package config

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"sync"
)

// Parse the file and return a map of configs.
func Parse(src io.Reader) (map[string]interface{}, error) {
	var data = make(map[string]interface{})
	var lock = sync.RWMutex{}

	lock.Lock()
	defer lock.Unlock()

	// Start reading from the reader using a scanner.
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		// the line.
		line := scanner.Text()

		// Skip the line if we start with hashtag.
		if len(line) > 0 && string(line[0]) == "#" {
			continue
		}

		before := beforeDelimiter(line, delimiter)
		after := afterDelimiter(line, delimiter)

		// if one of the above is empty ignore that line.
		if before == "" || after == "" {
			continue
		}

		// log.Debug(after)
		afterString := after.(string)
		// log.Debug(afterString)
		// log.Debug("first:", afterString[:1], " | last:", afterString[len(afterString)-1:])

		if afterString[:1] == `"` && afterString[len(afterString)-1:] == `"` { // String
			data[before[:len(before)]] = afterString[1 : len(afterString)-1]
		} else if toInt, err := strconv.ParseInt(afterString, 10, 64); err == nil { // Number
			data[before[:len(before)]] = toInt
		} else if toFloat, err := strconv.ParseFloat(afterString, 64); err == nil { // float
			data[before[:len(before)]] = toFloat
		} else if toBool, err := strconv.ParseBool(afterString); err == nil { // bool
			data[before[:len(before)]] = toBool
		} else {
			data[before[:len(before)]] = after
		}
	}

	if scanner.Err() != nil {
		return data, scanner.Err()
	}

	return data, nil
}

func beforeDelimiter(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return strings.TrimSpace(value[0:pos])
}

func afterDelimiter(value string, a string) interface{} {
	// Get substring after a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return nil
	}
	return strings.TrimSpace(value[pos+1 : len(value)])
}

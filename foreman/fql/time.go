// Copyright (C) 2017 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fql

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	queryRelativeTimeSecondsRegex = "-[0-9]+s"
	queryRelativeTimeMinutesRegex = "-[0-9]+min"
	queryRelativeTimeHoursRegex   = "-[0-9]+h"
	queryRelativeTimeDaysRegex    = "-[0-9]+d"
	queryRelativeTimeWeeksRegex   = "-[0-9]+w"
	queryRelativeTimeMonthsRegex  = "-[0-9]+mon"
	queryRelativeTimeYearsRegex   = "-[0-9]+y"

	queryRelativeTimeFormat = "-%d%s"

	queryAbsoluteTimeNowRegex                 = "now"
	queryAbsoluteTimeSQLNowRegex              = "now()"
	queryAbsoluteTimeSQLCurrentTimestampRegex = "current_timestamp"
	queryAbsoluteTimeRegexRFC                 = "[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}"
	queryAbsoluteTimeRegexSQL                 = "[0-9]{4}/[0-9]{2}/[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}"
	queryAbsoluteTimeUnixRegex                = "[0-9]*"
	queryAbsoluteTimeFormatRFC                = "2006-01-02 15:04:05" // RFC3339
	queryAbsoluteTimeFormatSQL                = "2006/01/02 15:04:05" // SQL

)

// TimeStringToTime returns a time of the specfied time string.
func TimeStringToTime(timeStr string) (*time.Time, error) {
	if IsRelativeTimeString(timeStr) {
		return RelativeTimeStringToTime(timeStr)
	}

	if IsAbsoluteTimeString(timeStr) {
		return AbsoluteTimeStringToTime(timeStr)
	}

	return nil, fmt.Errorf(errorInvalidTimeFormat, timeStr)
}

var queryAbsoluteTimeRegexes []*regexp.Regexp

// getAbsoluteTimeRegexes returns a Regexp slice of the absolute time formats.
func getAbsoluteTimeRegexes() []*regexp.Regexp {
	if queryAbsoluteTimeRegexes == nil {
		queryAbsoluteTimeRegexes = []*regexp.Regexp{
			regexp.MustCompile(queryAbsoluteTimeNowRegex),
			regexp.MustCompile(queryAbsoluteTimeSQLNowRegex),
			regexp.MustCompile(queryAbsoluteTimeSQLCurrentTimestampRegex),
			regexp.MustCompile(queryAbsoluteTimeRegexRFC),
			regexp.MustCompile(queryAbsoluteTimeRegexSQL),
			regexp.MustCompile(queryAbsoluteTimeUnixRegex),
		}
	}
	return queryAbsoluteTimeRegexes
}

// IsAbsoluteTimeString returns the specified string whether it is based on the absolute time formats.
func IsAbsoluteTimeString(timeStr string) bool {
	for _, regex := range getAbsoluteTimeRegexes() {
		matched := regex.MatchString(timeStr)
		if matched {
			return true
		}
	}

	return false
}

// AbsoluteTimeStringToTime returns a time based on the specified relative time string.
func AbsoluteTimeStringToTime(timeStr string) (*time.Time, error) {
	for n, regex := range getAbsoluteTimeRegexes() {
		matched := regex.MatchString(strings.ToLower(timeStr))
		if !matched {
			continue
		}

		switch n {
		case 0: // queryAbsoluteTimeNowRegex
			now := time.Now()
			return &now, nil
		case 1: // queryAbsoluteTimeSQLNowRegex
			now := time.Now()
			return &now, nil
		case 2: // queryAbsoluteTimeSQLCurrentTimestampRegex
			now := time.Now()
			return &now, nil
		case 3: // queryAbsoluteTimeRegexRFC
			time, err := time.ParseInLocation(queryAbsoluteTimeFormatRFC, timeStr, time.Local)
			if err != nil {
				return nil, err
			}
			return &time, nil
		case 4: // queryAbsoluteTimeRegexSQL
			time, err := time.ParseInLocation(queryAbsoluteTimeFormatSQL, timeStr, time.Local)
			if err != nil {
				return nil, err
			}
			return &time, nil
		case 5: // queryAbsoluteTimeUnixRegex
			unixTime, err := strconv.ParseInt(timeStr, 10, 64)
			if err != nil {
				break
			}
			time := time.Unix(unixTime, 0)
			return &time, nil
		}

		return nil, fmt.Errorf(errorInvalidTimeFormat, timeStr)
	}

	return nil, fmt.Errorf(errorInvalidTimeFormat, timeStr)
}

var queryRelativeTimeRegexes []*regexp.Regexp

// getRelativeTimeRegexes returns a Regexp slice of the relative time formats.
func getRelativeTimeRegexes() []*regexp.Regexp {
	if queryRelativeTimeRegexes == nil {
		queryRelativeTimeRegexes = []*regexp.Regexp{
			regexp.MustCompile(queryRelativeTimeSecondsRegex),
			regexp.MustCompile(queryRelativeTimeMinutesRegex),
			regexp.MustCompile(queryRelativeTimeHoursRegex),
			regexp.MustCompile(queryRelativeTimeDaysRegex),
			regexp.MustCompile(queryRelativeTimeWeeksRegex),
			regexp.MustCompile(queryRelativeTimeMonthsRegex),
			regexp.MustCompile(queryRelativeTimeYearsRegex),
		}
	}
	return queryRelativeTimeRegexes
}

// IsRelativeTimeString returns the specified string whether it is based on the releative time formats.
func IsRelativeTimeString(timeStr string) bool {
	for _, regex := range getRelativeTimeRegexes() {
		matched := regex.MatchString(timeStr)
		if matched {
			return true
		}
	}

	return false
}

// RelativeTimeStringToTime returns a time based on the specified relative time string.
func RelativeTimeStringToTime(timeStr string) (*time.Time, error) {
	now := time.Now()

	for n, regex := range getRelativeTimeRegexes() {
		matched := regex.MatchString(timeStr)
		if !matched {
			continue
		}

		var timeNum int
		var timeUnit string
		_, err := fmt.Sscanf(timeStr, queryRelativeTimeFormat, &timeNum, &timeUnit)
		if err != nil {
			break
		}

		now = time.Now()
		switch n {
		case 0: // queryRelativeTimeSecondsRegex
			time := now.Add(-(time.Duration(timeNum) * time.Second))
			return &time, nil
		case 1: // queryRelativeTimeMinutesRegex
			time := now.Add(-(time.Duration(timeNum) * time.Minute))
			return &time, nil
		case 2: // queryRelativeTimeHoursRegex
			time := now.Add(-(time.Duration(timeNum) * time.Hour))
			return &time, nil
		case 3: // queryRelativeTimeDaysRegex
			time := now.Add(-(time.Duration(timeNum) * time.Hour * 24))
			return &time, nil
		case 4: // queryRelativeTimeWeeksRegex
			time := now.Add(-(time.Duration(timeNum) * time.Hour * 24 * 7))
			return &time, nil
		case 5: // queryRelativeTimeMonthsRegex
			time := now.Add(-(time.Duration(timeNum) * time.Hour * 24 * 30))
			return &time, nil
		case 6: // queryRelativeTimeYearsRegex
			time := now.Add(-(time.Duration(timeNum) * time.Hour * 24 * 365))
			return &time, nil
		}

		return nil, fmt.Errorf(errorInvalidTimeFormat, timeStr)
	}

	return nil, fmt.Errorf(errorInvalidTimeFormat, timeStr)
}

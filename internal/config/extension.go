package config

import (
	"fmt"
	"strconv"
	"time"
)

type (
	layoutTZ struct {
		l string // Layout
		n bool   // Is named timezone or no timezone
	}
)

var (
	// map[format-name]map[format-layout]has-timezone
	// NOTE: Implement this using a proper struct.
	timeFormats = map[string]layoutTZ{
		"Layout":      {l: "01/02 03:04:05PM '06 -0700", n: false},
		"ANSIC":       {l: "Mon Jan _2 15:04:05 2006", n: true},
		"UnixDate":    {l: "Mon Jan _2 15:04:05 MST 2006", n: true},
		"RubyDate":    {l: "Mon Jan 02 15:04:05 -0700 2006", n: false},
		"RFC822":      {l: "02 Jan 06 15:04 MST", n: true},
		"RFC822Z":     {l: "02 Jan 06 15:04 -0700", n: false},
		"RFC850":      {l: "Monday, 02-Jan-06 15:04:05 MST", n: true},
		"RFC1123":     {l: "Mon, 02 Jan 2006 15:04:05 MST", n: true},
		"RFC1123Z":    {l: "Mon, 02 Jan 2006 15:04:05 -0700", n: false},
		"RFC3339":     {l: "2006-01-02T15:04:05Z07:00", n: false},
		"RFC3339Nano": {l: "2006-01-02T15:04:05.999999999Z07:00", n: false},
		"Kitchen":     {l: "3:04PM", n: true},
		"Stamp":       {l: "Jan _2 15:04:05", n: true},
		"StampMilli":  {l: "Jan _2 15:04:05.000", n: true},
		"StampMicro":  {l: "Jan _2 15:04:05.000000", n: true},
		"StampNano":   {l: "Jan _2 15:04:05.000000000", n: true},
		"DateTime":    {l: "2006-01-02 15:04:05", n: true},
		"DateOnly":    {l: "2006-01-02", n: true},
		"TimeOnly":    {l: "15:04:05", n: true},
	}
)

// GetString returns the value associated with the key as a string.
func (cfg *Config) GetString(key string) string {
	vals := cfg.get(true)
	val, ok := vals[key]
	if !ok {
		val = cfg.defaults[key]
	}

	return val
}

// GetBool returns the value associated with the key as a boolean.
func (cfg *Config) GetBool(key string) bool {
	val := cfg.GetString(key)
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return false
	}

	return boolVal
}

// GetInt returns the value associated with the key as an integer.
func (cfg *Config) GetInt(key string) int {
	val := cfg.GetString(key)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return int(intVal)
}

// GetInt32 returns the value associated with the key as an integer.
func (cfg *Config) GetInt32(key string) int32 {
	val := cfg.GetString(key)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return int32(intVal)
}

// GetInt64 returns the value associated with the key as an integer.
func (cfg *Config) GetInt64(key string) int64 {
	val := cfg.GetString(key)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return intVal
}

// GetUint returns the value associated with the key as an unsigned integer.
func (cfg *Config) GetUint(key string) uint {
	val := cfg.GetString(key)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return uint(intVal)
}

// GetUint16 returns the value associated with the key as an unsigned integer.
func (cfg *Config) GetUint16(key string) uint16 {
	val := cfg.GetString(key)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return uint16(intVal)
}

// GetUint32 returns the value associated with the key as an unsigned integer.
func (cfg *Config) GetUint32(key string) uint32 {
	val := cfg.GetString(key)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return uint32(intVal)
}

// GetUint64 returns the value associated with the key as an unsigned integer.
func (cfg *Config) GetUint64(key string) uint64 {
	val := cfg.GetString(key)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}
	return uint64(intVal)
}

// GetFloat64 returns the value associated with the key as a float64.
func (cfg *Config) GetFloat64(key string) float64 {
	val := cfg.GetString(key)
	floatVal, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0
	}

	return floatVal
}

// GetTime returns the value associated with the key as time.
func (cfg *Config) GetTime(key string) time.Time {
	val := cfg.GetString(key)
	timeVal, err := cfg.toTime(val)
	if err != nil {
		return time.Time{}
	}

	return timeVal
}

// GetDuration returns the value associated with the key as a duration.
func (cfg *Config) GetDuration(key string) time.Duration {
	panic("not implemented yet")
}

// GetIntSlice returns the value associated with the key as a slice of int values.
func (cfg *Config) GetIntSlice(key string) []int {
	panic("not implemented yet")
}

// GetStringSlice returns the value associated with the key as a slice of strings.
func (cfg *Config) GetStringSlice(key string) []string {
	panic("not implemented yet")
}

// GetStringMap returns the value associated with the key as a map of interfaces.
func (cfg *Config) GetStringMap(key string) map[string]interface{} {
	panic("not implemented yet")
}

// GetStringMapString returns the value associated with the key as a map of strings.
func (cfg *Config) GetStringMapString(key string) map[string]string {
	panic("not implemented yet")
}

// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
func (cfg *Config) GetStringMapStringSlice(key string) map[string][]string {
	panic("not implemented yet")
}

// GetSizeInBytes returns the size of the value associated with the given key
func (cfg *Config) GetSizeInBytes(key string) uint {
	panic("not implemented yet")
}

// Helpers
func (cfg *Config) toTime(timeString string) (timeVal time.Time, err error) {
	for _, format := range timeFormats {
		timeVal, err = time.Parse(format.l, timeString)
		if err != nil {
			return time.Time{}, fmt.Errorf("time cannot be parses from: %s", timeString)
		}
		if format.n {
			// Named timezone or not timezone
			year, month, day := timeVal.Date()
			hour, min, sec := timeVal.Clock()
			timeVal = time.Date(year, month, day, hour, min, sec, timeVal.Nanosecond(), cfg.location)
		}

		return timeVal, nil
	}

	return time.Time{}, fmt.Errorf("time cannot be parses from: %s", timeString)
}

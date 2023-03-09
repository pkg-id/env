// Package env provides a straightforward way to retrieve environment variables and offers a default value if the
// specified key is not present.
package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// String returns the string value if the key exists; otherwise, it returns the fallback value.
func String(key, fallback string) string { return Parse(key, Parsers.Identity(), fallback) }

// Int returns the int value if the key exists; otherwise, it returns the fallback value.
func Int(key string, fallback int) int { return Parse(key, Parsers.Int(), fallback) }

// Int64 returns the int64 value if the key exists; otherwise, it returns the fallback value.
func Int64(key string, fallback int64) int64 { return Parse(key, Parsers.Int64(), fallback) }

// Float64 returns the float64 value if the key exists; otherwise, it returns the fallback value.
func Float64(key string, fallback float64) float64 { return Parse(key, Parsers.Float64(), fallback) }

// Duration returns the time.Duration value if the key exists; otherwise, it returns the fallback value.
func Duration(key string, fallback time.Duration) time.Duration {
	return Parse(key, Parsers.Duration(), fallback)
}

// Bool returns the bool value if the key exists; otherwise, it returns the fallback value.
func Bool(key string, fallback bool) bool { return Parse(key, Parsers.Bool(), fallback) }

// List retrieves the environment variable with the specified key and attempts to parse it into a slice of type T.
// The Parser function is used to parse each value in the comma-separated string obtained from the environment variable.
// If the environment variable is not set or parsing fails for any of the values,
// the function returns the fallback value.
func List[T any](key string, parser Parser[T], fallback []T) []T {
	comaSeperated, exists := getEnv(key)
	if !exists {
		return fallback
	}

	values := make([]T, 0)
	words := strings.Split(comaSeperated, ",")
	for _, word := range words {
		word = strings.TrimSpace(word)
		t, err := parser(word)
		if err != nil {
			return fallback
		}
		values = append(values, t)
	}

	return values
}

// getEnv returns the value of the environment variable and a flag indicating whether the variable exists.
func getEnv(key string) (string, bool) {
	v, exists := os.LookupEnv(key)
	return v, exists
}

// must panics if the error parameter is not nil.
func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}

	return v
}

// Parser is a generic function type that takes a string input and returns a value of type T along with an error.
type Parser[T any] func(v string) (T, error)

// Parse the env value to the given type.
// Parse is a generic function that takes the key, parser, and a fallback value as arguments.
// It returns the parsed value of the environment variable if the key exists, or the fallback value if it does not.
func Parse[T any](key string, parser Parser[T], fallback T) T {
	v, exists := getEnv(key)
	if !exists {
		return fallback
	}

	return must(parser(v))
}

// parsers is a type for Parsers namespace.
type parsers int

// Parsers is a namespace that can be used to access all available parsers.
const Parsers = parsers(1)

// Identity returns the given input as the output.
// This Identity is used to parse a string into a string.
func (parsers) Identity() Parser[string] {
	return func(v string) (string, error) { return v, nil }
}

// Duration parses a string into a time.Duration.
func (parsers) Duration() Parser[time.Duration] { return time.ParseDuration }

// Bool parses a string into a boolean.
func (parsers) Bool() Parser[bool] { return strconv.ParseBool }

// Int parses a string into an int.
func (parsers) Int() Parser[int] { return strconv.Atoi }

// Int64 parses a string into an int64.
func (parsers) Int64() Parser[int64] {
	return func(v string) (int64, error) { return strconv.ParseInt(v, 10, 64) }
}

// Float64 parses a string into a float64.
func (parsers) Float64() Parser[float64] {
	return func(v string) (float64, error) { return strconv.ParseFloat(v, 64) }
}

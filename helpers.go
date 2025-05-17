package gohelpers

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// Check if provided int slice includes specified int value
func IncludesInt(slice []int, value int) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}

// Check if provided string slice includes specified string value
func IncludesString(slice []string, value string) bool {
	for _, element := range slice {
		if element == value {
			return true
		}
	}
	return false
}

// Create a UNIX timestamp in milliseconds (13 digits)
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Create a UNIX timestamp in seconds (10 digits)
func MakeTimestampSeconds() int64 {
	return time.Now().Unix()
}

// This is an alias for StructFields function.
// Get a slice of struct field names (for both public and private fields).
// This function is similar to Object.keys() in Javascript.
// Method names are not included.
// Field names of nested structs are not included.
// Works only for structs.
func ObjectKeys(value interface{}) ([]string, error) {
	return StructFields(value)
}

// This is an alias for StructFieldsJson function.
// Get a slice of struct JSON field tags (for both public and private fields).
// This function is similar to Object.keys() in Javascript.
// Method names are not included.
// Field names of nested structs are not included.
// Works only for structs.
func ObjectKeysJson(value interface{}, params StructKeysJsonParams) ([]string, error) {
	return StructFieldsJson(value, params)
}

// This is an alias for StructValues function.
// Get a slice of string values from struct fields (similar to Object.values() in JS).
// Nested struct values are returned as a single string.
// Works only for structs.
func ObjectValues(value interface{}) []string {
	return StructValues(value)
}

// Create a random alphanumeric string with specified length
func RandomString(length int) string {
	if length <= 0 {
		return ""
	}

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	accumulator := make([]byte, length)
	for i := range accumulator {
		accumulator[i] = randomStringCharset[seededRand.Intn(len(randomStringCharset))]
	}

	return string(accumulator)
}

// Get a slice of struct field names (for both public and private fields).
// This function is similar to Object.keys() in Javascript.
// Method names are not included.
// Field names of nested structs are not included.
// Works only for structs.
func StructFields(value interface{}) ([]string, error) {
	reflected := reflect.TypeOf(value)
	if reflected.Kind() != reflect.Struct {
		return nil, errors.New(structFieldsTypeError)
	}
	keys := make([]string, reflected.NumField())
	for i := 0; i < reflected.NumField(); i += 1 {
		keys[i] = reflected.Field(i).Name
	}
	return keys, nil
}

// Get a slice of struct JSON field tags (for both public and private fields).
// This function is similar to Object.keys() in Javascript.
// Method names are not included.
// Field names of nested structs are not included.
// Works only for structs.
func StructFieldsJson(value interface{}, params StructKeysJsonParams) ([]string, error) {
	reflected := reflect.TypeOf(value)
	if reflected.Kind() != reflect.Struct {
		return nil, errors.New(structFieldsTypeError)
	}

	var keys []string
	for i := 0; i < reflected.NumField(); i += 1 {
		jsonTag := reflected.Field(i).Tag.Get("json")

		if jsonTag == "" || jsonTag == "," || jsonTag == ",omitempty" {
			if !params.SkipMissingFields {
				jsonTag = reflected.Field(i).Name
				keys = append(keys, jsonTag)
				continue
			} else {
				continue
			}
		}

		if jsonTag == "-" {
			if !params.SkipIgnoredFields {
				jsonTag = reflected.Field(i).Name
				keys = append(keys, jsonTag)
				continue
			} else {
				continue
			}
		}

		tagPartials := strings.Split(jsonTag, ",")

		if len(tagPartials) > 1 {
			name := tagPartials[0]
			if name != "-" {
				keys = append(keys, name)
				continue
			}
			if !params.SkipIgnoredFields {
				keys = append(keys, reflected.Field(i).Name)
			} else {
				continue
			}
		}

		if len(tagPartials) == 1 {
			keys = append(keys, jsonTag)
		}
	}

	return keys, nil
}

// Get a slice of string values from struct fields (similar to Object.values() in JS).
// Nested struct values are returned as a single string.
// Works only for structs.
func StructValues(value interface{}) []string {
	var list []string
	elements := reflect.ValueOf(value)

	// if its a pointer, resolve its value
	if elements.Kind() == reflect.Ptr {
		elements = reflect.Indirect(elements)
	}

	for i := 0; i < elements.NumField(); i += 1 {
		list = append(list, fmt.Sprintf("%v", elements.Field(i).Interface()))
	}
	return list
}

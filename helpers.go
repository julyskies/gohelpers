package gohelpers

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

// Check if provided int array includes specified int value
func IncludesInt(array []int, value int) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}

// Check if provided string array includes specified string value
func IncludesString(array []string, value string) bool {
	for _, element := range array {
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

// Get an array (slice) of struct key names (similar to Object.keys() in JS).
// Works only for structs.
func ObjectKeys(object interface{}) ([]string, error) {
	reflected := reflect.TypeOf(object)
	if reflected.Kind() != reflect.Struct {
		return nil, errors.New(OBJECT_KEYS_TYPE_ERROR)
	}
	keys := make([]string, reflected.NumField())
	for i := 0; i < reflected.NumField(); i += 1 {
		keys[i] = reflected.Field(i).Name
	}
	return keys, nil
}

// Get an array (slice) of string values from struct keys (similar to Object.values() in JS).
// Works only for structs.
func ObjectValues(object interface{}) []string {
	var list []string
	elements := reflect.ValueOf(object)

	// if its a pointer, resolve its value
	if elements.Kind() == reflect.Ptr {
		elements = reflect.Indirect(elements) // TODO: coverage
	}

	for i := 0; i < elements.NumField(); i++ {
		list = append(list, fmt.Sprintf("%v", elements.Field(i).Interface()))
	}
	return list
}

// Create a random alphanumeric string with specified length
func RandomString(length int) string {
	if length <= 0 {
		return ""
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	accumulator := make([]byte, length)
	for i := range accumulator {
		accumulator[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(accumulator)
}

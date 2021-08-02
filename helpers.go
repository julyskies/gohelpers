package gohelpers

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func IncludesInt(array []int, value int) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}

func IncludesString(array []string, value string) bool {
	for _, element := range array {
		if element == value {
			return true
		}
	}
	return false
}

func ObjectValues(object interface{}) []string {
	var list []string
	elements := reflect.ValueOf(object)

	// if its a pointer, resolve its value
	if elements.Kind() == reflect.Ptr {
		elements = reflect.Indirect(elements)
	}

	for i := 0; i < elements.NumField(); i++ {
		list = append(list, fmt.Sprintf("%v", elements.Field(i).Interface()))
	}
	return list
}

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

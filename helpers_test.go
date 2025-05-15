package gohelpers

import (
	"fmt"
	"testing"
	"time"
)

func TestIncludesInt(t *testing.T) {
	arrayOfInts := []int{1, 2, 3, 4, 9}

	if ok := IncludesInt(arrayOfInts, 9); !ok {
		t.Error("value is included")
	}

	if ok := IncludesInt(arrayOfInts, 5); ok {
		t.Error("value is not included")
	}
}

func TestIncludesString(t *testing.T) {
	arrayOfStrings := []string{"a", "b", "c"}

	if ok := IncludesString(arrayOfStrings, "a"); !ok {
		t.Error("value is included")
	}

	if ok := IncludesString(arrayOfStrings, "x"); ok {
		t.Error("value is not included")
	}
}

func TestMakeTimestamp(t *testing.T) {
	timestampControl := time.Now().UnixNano() / int64(time.Millisecond)

	timestamp := MakeTimestamp()

	if timestamp < timestampControl {
		t.Error("timestamp is invalid")
	}

	timestampString := fmt.Sprint(timestamp)
	if len(timestampString) != 13 {
		t.Error("timestamp is invalid")
	}
}

func TestMakeTimestampSeconds(t *testing.T) {
	timestampControl := time.Now().Unix()

	timestamp := MakeTimestampSeconds()
	if timestamp < timestampControl {
		t.Error("timestamp is invalid")
	}

	timestampString := fmt.Sprint(timestamp)
	if len(timestampString) != 10 {
		t.Error("timestamp is invalid")
	}
}

func TestObjectKeys(t *testing.T) {
	keys, keysError := ObjectKeys("invalid argument")
	if keysError == nil {
		t.Error("invalid handling of the wrong argument type")
	}
	if keysError.Error() != objectKeysTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if keys != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	var testStruct objectKeysTestStruct

	keys, keysError = ObjectKeys(testStruct)
	if keysError != nil {
		t.Error("invalid error when providing correct argument")
	}
	if len(keys) != 2 {
		t.Error("invalid resulting slice length")
	}
	if IncludesString(keys, "method") {
		t.Error("should not include method names")
	}
}

func TestObjectKeysJson(t *testing.T) {
	keys, keysError := ObjectKeysJson("invalid argument", DefaultStructKeysJsonParams)
	if keysError == nil {
		t.Error("invalid handling of the wrong argument type")
	}
	if keysError.Error() != objectKeysTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if keys != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	var testStruct objectKeysJsonTestStruct

	keys, keysError = ObjectKeysJson(testStruct, DefaultStructKeysJsonParams)
}

func TestObjectValues(t *testing.T) {
	animals := animals{
		Elephant: "elephant",
		Hippo:    "hippo",
		Lion:     "lion",
	}

	values := ObjectValues(animals)
	if len(values) != 3 {
		t.Error("invalid resulting slice length")
	}

	if !IncludesString(values, "lion") {
		t.Error("invalid values are included in resulting slice")
	}
}

func TestRandomString(t *testing.T) {
	emptyString1 := RandomString(-5)
	if len(emptyString1) > 0 {
		t.Error("negative string length should result in empty string")
	}

	emptyString2 := RandomString(0)
	if len(emptyString2) > 0 {
		t.Error("zero string length should result in empty string")
	}

	randomString := RandomString(32)
	if len(randomString) != 32 {
		t.Error("invalid string length")
	}
}

func TestStructFields(t *testing.T) {
	fields, fieldsError := StructFields("invalid argument")
	if fieldsError == nil {
		t.Error("invalid handling of the wrong argument type")
	}
	if fieldsError.Error() != objectKeysTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if fields != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	var testStruct objectKeysTestStruct

	fields, fieldsError = StructFields(testStruct)
	if fieldsError != nil {
		t.Error("invalid error when providing correct argument")
	}
	if len(fields) != 2 {
		t.Error("invalid resulting slice length")
	}
	if IncludesString(fields, "method") {
		t.Error("should not include method names")
	}
}

func TestStructFieldsJson(t *testing.T) {
	fields, fieldsError := StructFieldsJson("invalid argument", DefaultStructKeysJsonParams)
	if fieldsError == nil {
		t.Error("invalid handling of the wrong argument type")
	}
	if fieldsError.Error() != objectKeysTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if fields != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	var testStruct objectKeysJsonTestStruct

	fields, fieldsError = StructFieldsJson(testStruct, DefaultStructKeysJsonParams)
	if fieldsError != nil {
		t.Error("invalid error when providing correct arguments")
	}
	if len(fields) == 0 {
		t.Error("resulting slice should not be empty when using default params")
	}

	params := StructKeysJsonParams{
		ReplaceIgnoredFieldsWithFieldNames: false,
		ReplaceMissingTagsWithFieldNames:   true,
	}
	fields, fieldsError = StructFieldsJson(testStruct, params)
	if fieldsError != nil {
		t.Error("invalid error when providing correct arguments")
	}
	if IncludesString(fields, "Ignored") {
		t.Error("resulting slice contains field name that should be ignored according to params")
	}
	if IncludesString(fields, "IgnoredWithOmitempty") {
		t.Error("resulting slice contains field name that should be ignored according to params")
	}
	if len(fields) != 7 {
		t.Error("resulting slice has invalid length")
	}

	params = StructKeysJsonParams{false, false}
	fields, fieldsError = StructFieldsJson(testStruct, params)
	if fieldsError != nil {
		t.Error("invalid error when providing correct argument type")
	}
	if IncludesString(fields, "Ignored") {
		t.Error("resulting slice contains field name that should be ignored according to params")
	}
	if IncludesString(fields, "IgnoredWithOmitempty") {
		t.Error("resulting slice contains field name that should be ignored according to params")
	}

}

func TestStructValues(t *testing.T) {
	animals := animals{
		Elephant: "elephant",
		Hippo:    "hippo",
		Lion:     "lion",
	}

	values := StructValues(animals)
	if len(values) != 3 {
		t.Error("invalid resulting slice length")
	}

	if !IncludesString(values, "lion") {
		t.Error("invalid values are included in resulting slice")
	}
}

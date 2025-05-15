package gohelpers

import (
	"fmt"
	"testing"
	"time"
)

type animalsStruct struct {
	Elephant string
	Hippo    string
	Lion     string
}

type testStructFields struct {
	Public  string
	private string
}

func (obj testStructFields) testMethod() string {
	return obj.Public + obj.private
}

type testStructFieldsEmbedding struct {
	testStructFields
	A int
	B string
}

type testStructFieldsJson struct {
	Comma                int `json:","`
	Empty                int `json:""`
	Ignored              int `json:"-"`
	Missing              int
	EmptyWithOmitempty   int `json:",omitempty"`
	IgonredWithOmitempty int `json:"-,omitempty"`
	NormalField          int `json:"normalField"`
	private              int `json:"private"`
	privateWithOmitempty int `json:"privateWithOmitempty,omitempty"`
}

func (obj testStructFieldsJson) testMethod() int {
	return obj.NormalField + obj.private
}

type testStructFieldsJsonEmbedding struct {
	testStructFields
	A int
	B string
}

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
	if keysError.Error() != structFieldsTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if keys != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	keys, keysError = ObjectKeys(testStructFields{})
	if keysError != nil {
		t.Error("invalid error when providing correct argument")
	}
	if len(keys) != 2 {
		t.Error("invalid resulting slice length")
	}
	if IncludesString(keys, "testMethod") {
		t.Error("should not include method names")
	}

	keys, keysError = ObjectKeys(testStructFieldsEmbedding{})
	if keysError != nil {
		t.Error("invalid error when providing correct argument")
	}
	if len(keys) != 3 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(keys, "testStructFields") {
		t.Error("resulting slice should contain the name of the embedded struct as an entry")
	}
	if IncludesString(keys, "Public") || IncludesString(keys, "private") {
		t.Error("resulting slice should not contain any embedded struct field names")
	}
}

// TODO: finish this after completing TestStructFieldsJson
func TestObjectKeysJson(t *testing.T) {
	keys, keysError := ObjectKeysJson("invalid argument", DefaultStructKeysJsonParams)
	if keysError == nil {
		t.Error("invalid handling of the wrong argument type")
	}
	if keysError.Error() != structFieldsTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if keys != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	var testStruct testStructFieldsJson

	keys, keysError = ObjectKeysJson(testStruct, DefaultStructKeysJsonParams)
}

func TestObjectValues(t *testing.T) {
	animals := animalsStruct{
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
	if fieldsError.Error() != structFieldsTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if fields != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	fields, fieldsError = StructFields(testStructFields{})
	if fieldsError != nil {
		t.Error("invalid error when providing correct argument")
	}
	if len(fields) != 2 {
		t.Error("invalid resulting slice length")
	}
	if IncludesString(fields, "testMethod") {
		t.Error("should not include method names")
	}

	fields, fieldsError = StructFields(testStructFieldsEmbedding{})
	if fieldsError != nil {
		t.Error("invalid error when providing correct argument")
	}
	if len(fields) != 3 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(fields, "testStructFields") {
		t.Error("resulting slice should contain the name of the embedded struct as an entry")
	}
	if IncludesString(fields, "Public") || IncludesString(fields, "private") {
		t.Error("resulting slice should not contain any embedded struct field names")
	}
}

func TestStructFieldsJson(t *testing.T) {
	// testing with invalid argument type
	fields, fieldsError := StructFieldsJson("invalid argument", DefaultStructKeysJsonParams)
	if fieldsError == nil {
		t.Error("invalid handling of the wrong argument type")
	}
	if fieldsError.Error() != structFieldsTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if fields != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	// testing with default params: do not skip ignored tags & replace missing tags with field names
	params := StructKeysJsonParams{true, true}
	fields, fieldsError = StructFieldsJson(testStructFieldsJson{}, params)
	if fieldsError != nil {
		t.Error("invalid error when providing correct arguments")
	}
	if len(fields) == 0 {
		t.Error("resulting slice should not be empty when using default params")
	}

	// testing with non-default params: skip ignored tags & replace missing tags with field names
	params = StructKeysJsonParams{
		ReplaceIgnoredFieldsWithFieldNames: false,
		ReplaceMissingTagsWithFieldNames:   true,
	}
	fields, fieldsError = StructFieldsJson(testStructFieldsJson{}, params)
	if fieldsError != nil {
		t.Error("invalid error when providing correct arguments")
	}
	if IncludesString(fields, "Ignored") {
		t.Error("resulting slice contains field name that should be skipped according to params")
	}
	if IncludesString(fields, "IgnoredWithOmitempty") {
		t.Error("resulting slice contains field name that should be ignored according to params")
	}
	if len(fields) != 7 {
		t.Error("resulting slice has invalid length")
	}

	params = StructKeysJsonParams{false, false}
	fields, fieldsError = StructFieldsJson(testStructFieldsJson{}, params)
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
	animals := animalsStruct{
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

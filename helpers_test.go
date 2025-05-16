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

type nestedAnimals struct {
	Cat           string
	Dog           string
	AnimalsStruct animalsStruct
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
	Comma                    int `json:","`
	Empty                    int `json:""`
	Ignored                  int `json:"-"`
	Missing                  int
	EmptyWithOmitempty       int `json:",omitempty"`
	IgonredWithOmitempty     int `json:"-,omitempty"`
	NormalField              int `json:"normalField"`
	NormalFieldWithOmitempty int `json:"normalFieldWithOmitempty,omitempty"`
	private                  int `json:"private"`
	privateWithOmitempty     int `json:"privateWithOmitempty,omitempty"`
}

func (obj testStructFieldsJson) testMethod() int {
	return obj.NormalField + obj.private
}

type testStructFieldsJsonEmbedding struct {
	testStructFields
	A int    `json:"a"`
	B string `json:"b"`
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

func TestObjectKeysJson(t *testing.T) {
	// testing with invalid argument type
	fields, fieldsError := ObjectKeysJson("invalid argument", DefaultStructKeysJsonParams)
	if fieldsError == nil {
		t.Error("invalid handling of the wrong argument type")
	}
	if fieldsError.Error() != structFieldsTypeError {
		t.Error("invalid error message when providing the wrong argument type")
	}
	if fields != nil {
		t.Error("invalid returned value when providing the wrong argument type")
	}

	// testing embedding
	fields, fieldsError = ObjectKeysJson(
		testStructFieldsJsonEmbedding{},
		DefaultStructKeysJsonParams,
	)
	if fieldsError != nil {
		t.Error("invalid error when providing correct arguments")
	}
	if len(fields) != 3 {
		t.Error("invalid resulting slice length when using struct with embedding")
	}
	if IncludesString(fields, "normalField") {
		t.Error("embedded fields should not be present in the resulting slice")
	}

	// testing basic functionality regardless of the used params
	basicFunctionalityTests := func(fields []string, fieldsError error) {
		if fieldsError != nil {
			t.Error("invalid error when providing correct arguments")
		}
		if len(fields) == 0 {
			t.Error("resulting slice should not be empty")
		}
		if IncludesString(fields, "testMethod") {
			t.Error("resulting slice should not include method names")
		}
		if !IncludesString(fields, "normalField") {
			t.Error("resulting slice should contain non-private field tags without any modifiers")
		}
		if !IncludesString(fields, "normalFieldWithOmitempty") {
			t.Error(
				"resulting slice should contain non-private field tags with 'omitempty' modifier",
			)
		}
		if IncludesString(fields, "normalFieldWithOmitempty,omitempty") {
			t.Error(
				"the 'omitempty' substring should be removed from the tag string for non-private fields",
			)
		}
		if !IncludesString(fields, "private") {
			t.Error("resulting slice should contain private field tags or field names")
		}
		if !IncludesString(fields, "privateWithOmitempty") {
			t.Error("resulting slice should contain private field tags with 'omitempty' modifier")
		}
		if IncludesString(fields, "privateWithOmitempty,omitempty") {
			t.Error(
				"the 'omitempty' substring should be removed from the tag string for private fields",
			)
		}
		if IncludesString(fields, "-") {
			t.Error("resulting slice should not contain '-' symbol")
		}
		if IncludesString(fields, "-,omitempty") {
			t.Error("resulting slice should not contain '-,omitempty'")
		}
		if IncludesString(fields, ",") {
			t.Error("resulting slice should not contain ',' symbol")
		}
		if IncludesString(fields, ",omitempty") {
			t.Error("resulting slice should not contain ',omitempty'")
		}
	}

	// testing with default params: do not skip ignored tags & replace missing tags with field names
	params := StructKeysJsonParams{true, true}
	fields, fieldsError = ObjectKeysJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 10 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(fields, "Ignored") {
		t.Error("resulting slice should contain ignored field name")
	}
	if !IncludesString(fields, "IgonredWithOmitempty") {
		t.Error("resulting slice should contain ignored field name even if it has 'omitempty'")
	}
	if !IncludesString(fields, "Comma") {
		t.Error("default field name should be used if JSON tag string is ','")
	}
	if !IncludesString(fields, "Empty") {
		t.Error("default field name should be used if JSON tag is an empty string")
	}
	if !IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("default field name should be used if JSON tag string is ',omitempty'")
	}
	if !IncludesString(fields, "Missing") {
		t.Error("default field name should be used if JSON tag is not set")
	}

	// testing with non-default params: skip ignored tags & replace missing tags with field names
	params = StructKeysJsonParams{
		ReplaceIgnoredFieldsWithFieldNames: false,
		ReplaceMissingTagsWithFieldNames:   true,
	}
	fields, fieldsError = ObjectKeysJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 8 {
		t.Error("resulting slice has invalid length")
	}
	if IncludesString(fields, "Ignored") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "IgnoredWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if !IncludesString(fields, "Comma") {
		t.Error("default field name should be used if JSON tag string is ','")
	}
	if !IncludesString(fields, "Empty") {
		t.Error("default field name should be used if JSON tag is an empty string")
	}
	if !IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("default field name should be used if JSON tag string is ',omitempty'")
	}
	if !IncludesString(fields, "Missing") {
		t.Error("default field name should be used if JSON tag is not set")
	}

	// testing with non-default params: skip ignored tags & skip missing tags
	params = StructKeysJsonParams{false, false}
	fields, fieldsError = ObjectKeysJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 4 {
		t.Error("resulting slice has invalid length")
	}
	if IncludesString(fields, "Ignored") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "IgnoredWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Comma") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Empty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Missing") {
		t.Error("resulting slice contains field name that should be skipped")
	}

	// testing with non-default params: do not skip ignored tags & skip missing tags
	params = StructKeysJsonParams{
		ReplaceIgnoredFieldsWithFieldNames: true,
		ReplaceMissingTagsWithFieldNames:   false,
	}
	fields, fieldsError = ObjectKeysJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 6 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(fields, "Ignored") {
		t.Error("resulting slice should contain ignored field name")
	}
	if !IncludesString(fields, "IgonredWithOmitempty") {
		t.Error("resulting slice should contain ignored field name even if it has 'omitempty'")
	}
	if IncludesString(fields, "Comma") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Empty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Missing") {
		t.Error("resulting slice contains field name that should be skipped")
	}
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

	// pass a pointer
	values = ObjectValues(&animals)
	if len(values) != 3 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(values, "lion") {
		t.Error("invalid values are included in resulting slice")
	}

	moreAnimals := nestedAnimals{
		AnimalsStruct: animalsStruct{
			Elephant: "elephant",
			Hippo:    "hippo",
			Lion:     "lion",
		},
		Cat: "cat",
		Dog: "dog",
	}
	values = ObjectValues(&moreAnimals)
	if len(values) != 3 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(values, "cat") {
		t.Error("missing values from resulting slice")
	}
	if IncludesString(values, "hippo") {
		t.Error("nested field values should be returned as a single string")
	}
	if !IncludesString(values, "{elephant hippo lion}") {
		t.Error("nested field values should be returned as a single string")
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

	// testing embedding
	fields, fieldsError = StructFieldsJson(
		testStructFieldsJsonEmbedding{},
		DefaultStructKeysJsonParams,
	)
	if fieldsError != nil {
		t.Error("invalid error when providing correct arguments")
	}
	if len(fields) != 3 {
		t.Error("invalid resulting slice length when using struct with embedding")
	}
	if IncludesString(fields, "normalField") {
		t.Error("embedded fields should not be present in the resulting slice")
	}

	// testing basic functionality regardless of the used params
	basicFunctionalityTests := func(fields []string, fieldsError error) {
		if fieldsError != nil {
			t.Error("invalid error when providing correct arguments")
		}
		if len(fields) == 0 {
			t.Error("resulting slice should not be empty")
		}
		if IncludesString(fields, "testMethod") {
			t.Error("resulting slice should not include method names")
		}
		if !IncludesString(fields, "normalField") {
			t.Error("resulting slice should contain non-private field tags without any modifiers")
		}
		if !IncludesString(fields, "normalFieldWithOmitempty") {
			t.Error(
				"resulting slice should contain non-private field tags with 'omitempty' modifier",
			)
		}
		if IncludesString(fields, "normalFieldWithOmitempty,omitempty") {
			t.Error(
				"the 'omitempty' substring should be removed from the tag string for non-private fields",
			)
		}
		if !IncludesString(fields, "private") {
			t.Error("resulting slice should contain private field tags or field names")
		}
		if !IncludesString(fields, "privateWithOmitempty") {
			t.Error("resulting slice should contain private field tags with 'omitempty' modifier")
		}
		if IncludesString(fields, "privateWithOmitempty,omitempty") {
			t.Error(
				"the 'omitempty' substring should be removed from the tag string for private fields",
			)
		}
		if IncludesString(fields, "-") {
			t.Error("resulting slice should not contain '-' symbol")
		}
		if IncludesString(fields, "-,omitempty") {
			t.Error("resulting slice should not contain '-,omitempty'")
		}
		if IncludesString(fields, ",") {
			t.Error("resulting slice should not contain ',' symbol")
		}
		if IncludesString(fields, ",omitempty") {
			t.Error("resulting slice should not contain ',omitempty'")
		}
	}

	// testing with default params: do not skip ignored tags & replace missing tags with field names
	params := StructKeysJsonParams{true, true}
	fields, fieldsError = StructFieldsJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 10 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(fields, "Ignored") {
		t.Error("resulting slice should contain ignored field name")
	}
	if !IncludesString(fields, "IgonredWithOmitempty") {
		t.Error("resulting slice should contain ignored field name even if it has 'omitempty'")
	}
	if !IncludesString(fields, "Comma") {
		t.Error("default field name should be used if JSON tag string is ','")
	}
	if !IncludesString(fields, "Empty") {
		t.Error("default field name should be used if JSON tag is an empty string")
	}
	if !IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("default field name should be used if JSON tag string is ',omitempty'")
	}
	if !IncludesString(fields, "Missing") {
		t.Error("default field name should be used if JSON tag is not set")
	}

	// testing with non-default params: skip ignored tags & replace missing tags with field names
	params = StructKeysJsonParams{
		ReplaceIgnoredFieldsWithFieldNames: false,
		ReplaceMissingTagsWithFieldNames:   true,
	}
	fields, fieldsError = StructFieldsJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 8 {
		t.Error("resulting slice has invalid length")
	}
	if IncludesString(fields, "Ignored") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "IgnoredWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if !IncludesString(fields, "Comma") {
		t.Error("default field name should be used if JSON tag string is ','")
	}
	if !IncludesString(fields, "Empty") {
		t.Error("default field name should be used if JSON tag is an empty string")
	}
	if !IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("default field name should be used if JSON tag string is ',omitempty'")
	}
	if !IncludesString(fields, "Missing") {
		t.Error("default field name should be used if JSON tag is not set")
	}

	// testing with non-default params: skip ignored tags & skip missing tags
	params = StructKeysJsonParams{false, false}
	fields, fieldsError = StructFieldsJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 4 {
		t.Error("resulting slice has invalid length")
	}
	if IncludesString(fields, "Ignored") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "IgnoredWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Comma") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Empty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Missing") {
		t.Error("resulting slice contains field name that should be skipped")
	}

	// testing with non-default params: do not skip ignored tags & skip missing tags
	params = StructKeysJsonParams{
		ReplaceIgnoredFieldsWithFieldNames: true,
		ReplaceMissingTagsWithFieldNames:   false,
	}
	fields, fieldsError = StructFieldsJson(testStructFieldsJson{}, params)
	basicFunctionalityTests(fields, fieldsError)
	if len(fields) != 6 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(fields, "Ignored") {
		t.Error("resulting slice should contain ignored field name")
	}
	if !IncludesString(fields, "IgonredWithOmitempty") {
		t.Error("resulting slice should contain ignored field name even if it has 'omitempty'")
	}
	if IncludesString(fields, "Comma") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Empty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "EmptyWithOmitempty") {
		t.Error("resulting slice contains field name that should be skipped")
	}
	if IncludesString(fields, "Missing") {
		t.Error("resulting slice contains field name that should be skipped")
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

	// pass a pointer
	values = StructValues(&animals)
	if len(values) != 3 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(values, "lion") {
		t.Error("invalid values are included in resulting slice")
	}

	moreAnimals := nestedAnimals{
		AnimalsStruct: animalsStruct{
			Elephant: "elephant",
			Hippo:    "hippo",
			Lion:     "lion",
		},
		Cat: "cat",
		Dog: "dog",
	}
	values = StructValues(&moreAnimals)
	if len(values) != 3 {
		t.Error("invalid resulting slice length")
	}
	if !IncludesString(values, "cat") {
		t.Error("missing values from resulting slice")
	}
	if IncludesString(values, "hippo") {
		t.Error("nested field values should be returned as a single string")
	}
	if !IncludesString(values, "{elephant hippo lion}") {
		t.Error("nested field values should be returned as a single string")
	}
}

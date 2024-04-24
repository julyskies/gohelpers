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

func TestObjectValues(t *testing.T) {
	type Animals struct {
		Elephant string
		Hippo    string
		Lion     string
	}

	animals := Animals{
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

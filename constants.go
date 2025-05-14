package gohelpers

import "fmt"

const OBJECT_KEYS_TYPE_ERROR string = "provided argument type is not a struct"

type objectKeysTestStruct struct {
	Public  string
	private string
}

func (obj objectKeysTestStruct) testMethod() string {
	return fmt.Sprintf("%s %s", obj.Public, obj.private)
}

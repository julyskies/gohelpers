package gohelpers

import "fmt"

const objectKeysTypeError string = "provided argument type is not a struct"

const randomStringCharset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type animals struct {
	Elephant string
	Hippo    string
	Lion     string
}

type objectKeysTestStruct struct {
	Public  string
	private string
}

func (obj objectKeysTestStruct) testMethod() string {
	return fmt.Sprintf("%s %s", obj.Public, obj.private)
}

type objectKeysJsonTestStruct struct {
	Comma                int `json:","`
	Empty                int `json:""`
	Ignored              int `json:"-"`
	Missing              int
	EmptyWithOmitempty   int `json:",omitempty"`
	IgonredWithOmitempty int `json:"-,omitempty"`
	NormalField          int `json:"normalField"`
	private              int `json:"private"`
	privateWithOmitempty int `json:"private,omitempty"`
}

func (obj objectKeysJsonTestStruct) testMethod() string {
	return fmt.Sprintf("%d %d", obj.NormalField, obj.private)
}

type StructKeysJsonParams struct {
	ReplaceIgnoredFieldsWithFieldNames bool // use struct field name if `json:"-"` or if `json:"-,omitempty"`
	ReplaceMissingTagsWithFieldNames   bool // use struct field name if there's no JSON tag, if `json:""`, if `json:",omitempty"`, if `json:","`
}

// Replace ignored fields with field names: true
// Replace missing tags with field names: true
var DefaultStructKeysJsonParams = StructKeysJsonParams{
	ReplaceIgnoredFieldsWithFieldNames: true,
	ReplaceMissingTagsWithFieldNames:   true,
}

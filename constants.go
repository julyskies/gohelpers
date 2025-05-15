package gohelpers

const structFieldsTypeError string = "provided argument type is not a struct"

const randomStringCharset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

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

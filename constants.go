package gohelpers

const structFieldsTypeError string = "provided argument type is not a struct"

const randomStringCharset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type StructKeysJsonParams struct {
	SkipIgnoredFields bool // skip field if `json:"-"` / `json:"-,omitempty"`
	SkipMissingFields bool // skip field if there is no JSON tag, or tag is empty, i. e. `json:""` / `json:",omitempty"` / `json:","`
}

// Skip ignored fields: false -> ignored JSON tags will be replaced with default field names.
// Skip missing fields: false -> default field names will be used instead of missing JSON tags.
var DefaultStructKeysJsonParams = StructKeysJsonParams{
	SkipIgnoredFields: false,
	SkipMissingFields: false,
}

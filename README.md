## gohelpers

This package contains helper functions for Golang applications.

Minimal required Golang version: **`v1.16`**.

### Install

```shell script
go get github.com/julyskies/gohelpers
```

### Available functions

- **`IncludesInt(slice []int, value int) bool`**

  This helper function returns a boolean value if a slice of `int` values contains a specified `int` value.

  **Example:**

  ```go
  slice := []int{1, 2, 4, 9}
  result := gohelpers.IncludesInt(slice, 8)
  fmt.Println(result) // false
  ```

- **`IncludesString(slice []string, value string) bool`**

  This helper function returns a boolean value if slice of `string` values contains a specified `string` value.

  **Example:**

  ```go
  slice := []string{"a", "b", "c"}
  result := gohelpers.IncludesString(slice, "a")
  fmt.Println(result) // true
  ```

- **`MakeTimestamp() int64`**

  This helper function returns a UNIX timestamp in milliseconds.

  **Example:**

  ```go
  timestamp := gohelpers.MakeTimestamp()
  fmt.Println(timestamp) // 1627987461201
  ```

- **`MakeTimestampSeconds() int64`**

  This helper function returns a UNIX timestamp in seconds.

    **Example:**

  ```go
  timestamp := gohelpers.MakeTimestampSeconds()
  fmt.Println(timestamp) // 1713957122
  ```

- **`RandomString(length int) string`**

  This helper function returns a random alphanumeric string of the provided length.

  **Example:**

  ```go
  randomString := gohelpers.RandomString(8)
  fmt.Println(randomString) // A9is5Try
  ```

- **`StructFields(value interface{}) ([]string, error)`**

  This helper function returns a slice of struct field names (similar to `Object.keys()` in Javascript). Method names are not included. The `value` argument should be a struct. Both public and private struct field names are returned.
  
  `ObjectKeys` is an alias for this function.

  Please notice: this function will not return field names for the nested structs.

  **Example:**

  ```go
  type someStruct struct {
    A       int
    B       string
    private bool
    Public  int
  }
  fields, _ := gohelpers.StructFields(someStruct{})
  fmt.Println(fields) // [A B private Public]

  // method names are not returned
  func (s someStruct) myMethod() string {
    return s.B
  }
  fields, _ = gohelpers.StructFields(someStruct{})
  fmt.Println(fields) // [A B private Public]

  // working with nesting is not supported
  type withNesting struct {
    SomeStruct someStruct
    J int
    K string
  }
  fields, _ = gohelpers.StructFields(withNesting{})
  fmt.Println(fields) // [SomeStruct J K]

  // handling an error
  result, err := gohelpers.StructFields("not a struct")
  if err != nil {
    fmt.Println(result)      // []
    fmt.Println(err.Error()) // provided argument type is not a struct
  }
  ```

- **`StructFieldsJson(value interface{}, params gohelpers.StructKeysJsonParams) ([]string, error)`**

  This helper function returns a slice of struct field JSON tags (similar to `Object.keys()` in Javascript). Method names are not included. The `value` argument should be a struct. Both public and private struct field JSON tags are returned.

  `ObjectKeysJson` is an alias for this function.

  Please notice: this function will not return field names for the nested structs.

  This function requires a second argument - `gohelpers.StructKeysJsonParams` struct, where you can specify the following options:

  ```go
  type StructKeysJsonParams struct {
	  SkipIgnoredFields bool
	  SkipMissingFields bool
  }
  ```

  - `SkipIgnoredFields` field determines if ignored JSON tags should be skipped (i. e. `json:"-"` and `json:"-,omitempty"`). If ignored fields are skipped, resulting slice will not include these struct fields. If they are not skipped, default field name will be included in resulting slice.

  - `SkipMissingFields` field determines if missing or empty JSON tags should be skipped (i. e. if there is no JSON tag for a field, or if JSON tag equals to `json:""`, `json:",omitempty"` or `json:","`). If missing fields are skipped, resulting slice will not include these struct fields. If they are not skipped, default field name will be included in resulting slice.

  Default options for this function available as `gohelpers.DefaultStructKeysJsonParams`:

  ```go
  // Skip ignored fields: false -> ignored JSON tags will be replaced with default field names.
  // Skip missing fields: false -> default field names will be used instead of missing JSON tags.
  var DefaultStructKeysJsonParams = StructKeysJsonParams{
	  SkipIgnoredFields: false,
	  SkipMissingFields: false,
  }
  ```

  **Example:**

  ```go
  package main

  import (
    "fmt"

    "github.com/julyskies/gohelpers"
  )

  type User struct {
    Age       uint8  `json:""`
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Password  string `json:"-"`
    role      string
    Status    string `json:"status,omitempty"`
  }

  func main() {
    // get all fields and available JSON tags
    fields, _ := gohelpers.StructFieldsJson(
      User{},
      gohelpers.DefaultStructKeysJsonParams,
    )
    fmt.Println(fields) // [Age firstName lastName Password role status]

    // get all non-ignored fields and JSON tags
    params := gohelpers.StructKeysJsonParams{
      SkipIgnoredFields: true,
      SkipMissingFields: false,
    }
    fields, _ = gohelpers.StructFieldsJson(User{}, params)
    fmt.Println(fields) // [Age firstName lastName Role status]

    // get all non-missing fields and JSON tags
    params = gohelpers.StructKeysJsonParams{
      SkipIgnoredFields: false,
      SkipMissingFields: true,
    }
    fields, _ = gohelpers.StructFieldsJson(User{}, params)
    fmt.Println(fields) // [firstName lastName Password status]

    // skip ignored and missing tags / fields
    params = gohelpers.StructKeysJsonParams{
      SkipIgnoredFields: true,
      SkipMissingFields: true,
    }
    fields, _ = gohelpers.StructFieldsJson(User{}, params)
    fmt.Println(fields) // [firstName lastName status]

    // handling an error
    result, err := gohelpers.StructFieldsJson(
      "not a struct",
      gohelpers.StructKeysJsonParams{
        SkipIgnoredFields: true,
        SkipMissingFields: true,
      },
    )
    if err != nil {
      fmt.Println(result)      // []
      fmt.Println(err.Error()) // provided argument type is not a struct
    }
  }
  ```

- **`StructValues(value interface{}) []string`**

  This helper function returns a slice of values as strings. These values are taken from the provided  `struct`. This function is similar to `Object.values()` in Javascript. Methods are not returned.
  
  `ObjectValues` is an alias for this function.

  Please notice: nested struct values are returned as a single string.

  **Example:**

  ```go
  type animalsStruct struct {
    Elephant string
    Hippo    string
    Lion     string
  }
  animals := animalsStruct{
    Elephant: "elephant",
    Hippo:    "hippo",
    Lion:     "lion",
  }
  values := gohelpers.StructValues(animals)
  fmt.Println(values) // [elephant, hippo, lion]

  // nested struct
  type nestedAnimals struct {
    AnimalsStruct animalsStruct
    Cat           string
    Dog           string
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
  values := gohelpers.StructValues(moreAnimals)
  fmt.Println(values) // [{elephant hippo lion} cat dog]
  ```

### Aliases

Aliases are added for compatibility.

- **`ObjectKeys(value interface{}) ([]string, error)`**

  This is an alias for `StructFields` function.

- **`ObjectKeysJson(value interface{}, params gohelpers.StructKeysJsonParams) ([]string, error)`**

  This is an alias for `StructFieldsJson` function.

- **`ObjectValues(value interface{}) []string`**

  This is an alias for `StructValues` function.


### Testing

Tests are located in [`helpers_test.go`](./helpers_test.go)

Run tests:

```shell script
go test
```

### License

[MIT](./LICENSE.md)

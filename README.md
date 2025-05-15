## gohelpers

This package contains helper functions for Golang applications.

Minimal required Golang version: **`v1.16`**.

### Install

```shell script
go get github.com/julyskies/gohelpers
```

### Available helper functions

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

  This helper function returns a slice of struct field names (similar to `Object.keys()` in Javascript). Method names are not included. The `value` argument should be a struct. Both public and private struct field names are returned. `ObjectKeys` is an alias for this function.

  Please notice: this function will not return embedded field names (i. e. field names from the embedded structs).

  **Example:**

  ```go
  type SomeStruct struct {
    A       int
    B       string
    private bool
    Public  int
  }
  fields, _ := gohelpers.StructFields(SomeStruct{})
  fmt.Println(fields) // [A B private Public]

  // working with embedding is not supported
  type Embedding struct {
    SomeStruct
    J int
    K string
  }
  fields, _ = gohelpers.StructFields(Embedding{})
  fmt.Println(fields) // [SomeStruct J K]

  // handling an error
  result, err := gohelpers.StructFields("invalid argument type")
  if err != nil {
    fmt.Println(err.Error()) // provided argument type is not a struct
    fmt.Println(result) // nil
  }
  ```

- **`StructFieldsJson(value interface{}, params gohelpers.StructKeysJsonParams) ([]string, error)`**

  This helper function returns a slice of struct JSON field tags (similar to `Object.keys()` in Javascript). Method names are not included. The `value` argument should be a struct. Both public and private struct JSON field tags are returned.

  A second argument is required for this function, it should be `gohelpers.StructKeysJsonParams` struct, where you can specify the following:

  ```go
  type StructKeysJsonParams struct {
	  ReplaceIgnoredFieldsWithFieldNames bool
	  ReplaceMissingTagsWithFieldNames   bool
  }
  ```

- **`StructValues(value interface{}) []string`**

  This helper function returns an array of values as strings. These values are taken from the provided  `struct`. Behaviour is similar to the `Object.values()` from JS. `ObjectValues` is an alias for this function.

  **Example:**

  ```go
  type animalsStruct struct {
    Elephant string
    Hippo    string
    Lion     string
  }

  animals := animalsStruct{
    Elephant: "elephant",
    Hippo: "hippo",
    Lion: "lion",
  }

  values := gohelpers.StructValues(animals)
  fmt.Println(values) // ["elephant", "hippo", "lion"]
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

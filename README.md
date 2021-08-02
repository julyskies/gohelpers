## gohelpers

This repository contains helper functions for Golang applications.

### Install

```shell script
go get github.com/julyskies/gohelpers
```

### Available helper functions

- `IncludesInt(array []int, value int) bool`

  This helper function returns a boolean value if array of `int` values contains a specified `int` value.

- `IncludesString(array []string, value string) bool`

  This helper function returns a boolean value if array of `string` values contains a specified `string` value.

- `ObjectValues(object interface{}) []string`

  This helper function returns an array of values as strings. These values are taken from the provided  `struct`. Behaviour is similar to the `Object.values()` from JS.

- `RandomString(length int) string`

  This helper function returns a random alphanumeric string of the provided length.

### Example

```go
import "github.com/julyskies/gohelpers"

type Animals struct {
  Elephant string
  Hippo    string
  Lion     string
}

func example() {
  arrayOfInts := []int{1, 2, 3, 4, 9}
  arrayOfStrings := []string{"a", "b", "c"}
  animals := Animals{
    Elephant: "elephant",
    Hippo: "hippo",
    Lion: "lion",
  }

  includesInt := gohelpers.IncludesInt(arrayOfInts, 8) // false
  includesString := gohelpers.IncludesString(arrayOfStrings, "a") // true
  values := gohelpers.ObjectValues(animals) // ["elephant", "hippo", "lion"]

  randomString := gohelpers.RandomString(8) // A9is5Try
}
```

### License

[MIT](./LICENSE.md)

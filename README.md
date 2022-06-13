# Utils

A Set of useful functions for working with files, errors and strings.

## Functions

### If

Generate quick if result.

```go
// Signature:
func If[T any](cond bool, yes T, no T) T

// Example:
res := If[string](name == "john", "it's john", "anonymous")
```

### Contains

Check if slice contains item.

```go
// Signature:
func Contains[T comparable](items []T, item T) bool

// Example:
res := Contains[string](items, "john")
```

## Error

### TaggedError

Generate a tagged error.

```go
// Signature:
TaggedError(tags []string, format string, args ...interface{}) error

// Example:
TaggedError([]string{"MyLib","MyMethod"}, "failed on %s file!", "main.json")
// [MyLib] [MyMethod] failed on main.json file!
```

### IsErrorOf

Check if error has tag.

```go
// Signature:
IsErrorOf(tag string, err error) bool

// Example:
IsErrorOf("MyLib", err) // true
```

### HasError

Check if error is nil or not.

```go
// Signature:
HasError(err error) bool
```

### PanicOnError

Generate panic from error if error not nil.

```go
// Signature:
PanicOnError(err error)
```

### Handle Functions Return Value

This set of functions help you to handle output of function that have a return type with error value. this functions generate panic on error or return value. You can use this functions instead of if else block for checking error.

```go
// Signature:
UIntOrPanic(res uint, err error) uint
UInt8OrPanic(res uint8, err error) uint8
UInt16OrPanic(res uint16, err error) uint16
UInt32OrPanic(res uint32, err error) uint32
UInt64OrPanic(res uint64, err error) uint64
IntOrPanic(res int, err error) int
Int8OrPanic(res int8, err error) int8
Int16OrPanic(res int16, err error) int16
Int32OrPanic(res int32, err error) int32
Int64OrPanic(res int64, err error) int64
Float32OrPanic(res float32, err error) float32
Float64OrPanic(res float64, err error) float64
StringOrPanic(res string, err error) string
BoolOrPanic(res bool, err error) bool
InterfaceOrPanic(res interface{}, err error) interface{}

// Example:
import "github.com/bopher/utils"
func TestIt(a int, b int) (int, error) { ... }
res = utils.IntOrPanic(TestIt(1, 3))
```

## File

### FileExists

Check if file exists or not.

```go
// Signature:
FileExists(path string) (bool, error)

// Example:
import "github.com/bopher/utils"
exists, err := utils.FileExists("path/to/file")
```

### IsDirectory

Check if path is a directory.

```go
// Signature:
IsDirectory(path string) (bool, error)

// Example
import "github.com/bopher/utils"
ok, err := utils.IsDirectory("path/to/dir")
```

### FindFile

Search for files in directory by a regex pattern.

```go
// Signature:
FindFile(dir string, pattern string) []string

// Example:
import "github.com/bopher/utils"
files := utils.FindFile("path/to/dir", ".+\.sql") // => Get All file with sql extension
```

### ClearDirectory

Delete all files and sub-directory in directory.

```go
// Signature:
ClearDirectory(dir string) error
```

### GetSubDirectory

Get list of sub directories.

```go
// Signature:
GetSubDirectory(dir string) ([]string, error)
```

### CreateDirectory

Create nested directory.

```go
// Signature:
CreateDirectory(path string) error

// Example:
import "github.com/bopher/utils"
err := utils.CreateDirectory("a/b/c/d") // => Create all a, b, c and d directory
```

### DetectMime

Detect file mime info from content

```go
// Signature:
DetectMime(data []byte) *mimetype.MIME

// Example:
if mime := DetectMime(myFileData); mime != nil {
    // do something
}
```

### Extension

Get file extension.

```go
// Signature:
Extension(file string) string

// Example
Extension("file") // ""
Extension("file.JPG") // ".jpg"
Extension("file.png") // ".png"
Extension("file.") // "."
```

## String

### ExtractNumbers

Extract numbers from string.

```go
import "github.com/bopher/utils"
numbers := utils.ExtractNumbers("(+1) 234-56789") // => 123456789
```

### ExtractAlphaNum

Extract alpha and numbers from string `[a-zA-Z0-9]`. You can add extra character to add in extraction.

```go
import "github.com/bopher/utils"
numbers := utils.ExtractAlphaNum("this is a: 123", ":") // => "thisisa:123"
```

### ExtractAlphaNumPersian

Extract persian alpha, alpha and numbers from string `[ا-یa-zA-Z0-9]`. You can add extra character to add in extraction.

```go
import "github.com/bopher/utils"
numbers := utils.ExtractAlphaNumPersian("My name is: مجتبی", " ") // => "My name is مجتبی"
```

### RandomStringFromCharset

Generate random string from character list.

```go
import "github.com/bopher/utils"
str, err := utils.RandomStringFromCharset(5, "1234567890") // => "59102"
str2, err2 := utils.RandomStringFromCharset(3, "ABCDEFGH") // => "DFC"
```

### RandomString

Generate random string from Alpha-Num Chars

```go
import "github.com/bopher/utils"
str, err := utils.RandomString(5) // => "AB5S2"
```

### Slugify

Generate dash separated string.

```go
import "github.com/bopher/utils"
str := utils.Slugify("welcome to", "my site") // => "welcome-to-my-site"
```

### SlugifyPersian

Make slugify string for persian string. this function only keep `persian alphabet`, `a-z`, `A-Z` and `0-9` characters.

```go
import "github.com/bopher/utils"
str := utils.SlugifyPersian("خوش آمدید \n \r \t - to گچپژ") // => "خوش-آمدید-to-گچپژ"
```

### ConcatStr

Join strings with separator.

```go
// Signature:
ConcatStr(sep string, str ...string) string

// Example:
import "github.com/bopher/utils"
str := utils.ConcatStr(" ", "John", "", "Doe") // => "John Doe"
```

### FormatNumber

Format number with comma separator.

```go
import "github.com/bopher/utils"
func FormatNumber(format string, v ...interface{}) string {
str := utils.FormatNumber("$ %d [total] $ %d [remain]", 10000, 2500) // => "$ 10,000 [total] $ 2,500 [remain]"
```

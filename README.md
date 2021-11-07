# Utils

A Set of useful functions for working with files, date, errors and strings.

## Date

### ParseJalali

Parse jalali date from string. if you pass nil location it use **Asia/Tehran** by default.

this function return nil when failed to parse jalali date.

```go
// Signature:
ParseJalali(str string, loc *time.Location) *ptime.Time

// Example:
import "github.com/bopher/utils"
jDate := utils.ParseJalali("1370-09-30 14:30", nil)
```

### TimeToJalali

Convert standard go time to jalali time.

```go
// Signature:
TimeToJalali(t time.Time) (ptime.Time, error)

// Example:
import "github.com/bopher/utils"
jDate, err := utils.TimeToJalali(myTime)
```

## Error

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

## String

### ExtractNumbers

Extract numbers from string.

```go
import "github.com/bopher/utils"
numbers := utils.ExtractNumbers("(+1) 234-56789") // => 123456789
```

### RandomStringFromCharset

Generate random string from character list.

```go
import "github.com/bopher/utils"
str, err := RandomStringFromCharset(5, "1234567890") // => "59102"
str2, err2 := RandomStringFromCharset(3, "ABCDEFGH") // => "DFC"
```

### RandomString

Generate random string from Alpha-Num Chars

```go
import "github.com/bopher/utils"
str, err := RandomString(5) // => "AB5S2"
```

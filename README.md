# Utils

A Set of useful functions for working with files, date, errors and strings.

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
str := utils.FormatNumber("total: $ %d", 12500000) // => "total: $ 12,500,000"
```

## Mongo DB

### MongoArray

Generate `primitive.A` array from args.

```go
import "github.com/bopher/utils"
m := utils.MongoArray("John", "Kim", "Jimmy") // => ["John", "Kim", "Jimmy"]
```

### MongoMap

Generate `primitive.M` structure from args. Args count must be even!

```go
import "github.com/bopher/utils"
m := utils.MongoMap("_id", 1, "name", "John") // => { _id: 1, name: "John"}
```

### MongoMaps

generate `[]primitive.M` from args. Args count must be even!

```go
import "github.com/bopher/utils"
m := utils.MongoMaps("_id", 1, "name", "John") // => [{ _id: 1}, {name: "John"}]
```

### MongoDoc

Generate `primitive.D` from args. Args count must be even!

```go
import "github.com/bopher/utils"
m := utils.MongoDoc("_id", 1, "name", "John") // => [{Key: "_id", Value: 1}, {Key: "name", Value: "John"}]
```

### MongoOr

Generate mongo $or.

```go
import "github.com/bopher/utils"
orCond := utils.MongoOr(utils.MongoMaps("_id", 1, "name", "John")) // => {$or: [{ _id: 1}, {name: "John"}]}
```

### MongoAnd

Generate mongo $and.

```go
import "github.com/bopher/utils"
andCond := utils.MongoAnd(utils.MongoMaps("_id", 1, "name", "John")) // => {$and: [{ _id: 1}, {name: "John"}]}
```

### MongoIn

Generate mongo $in.

```go
import "github.com/bopher/utils"
inCond := utils.MongoIn("name", []int{1,2,3}) // => {name: {$in: [1,2,3]}}
```

### MongoSet

Generate mongo $set.

```go
import "github.com/bopher/utils"
setVal := utils.MongoSet(utils.MongoMap("name", "John")) // => {$set: {name: "John"}}
```

### MongoNestedSet

Generate nested mongo $set with key/value.

```go
import "github.com/bopher/utils"
setVal := utils.MongoNestedSet("name", "John") // => {$set: {name: "John"}}
```

### MongoMatch

Generate mongo $match.

```go
import "github.com/bopher/utils"
match := utils.MongoMatch(utils.MongoMap("name", "John")) // => {$match: {name: "John"}}
```

### MongoLimit

Generate mongo $limit.

```go
import "github.com/bopher/utils"
limit := utils.MongoLimit(2) // => {$limit: 2}
```

### MongoSort

Generate mongo $sort.

```go
import "github.com/bopher/utils"
sorts := utils.MongoSort(utils.MongoMap({"_id", 1})) // => {$sort: {_id:1}}
```

### MongoSkip

Generate mongo $skip.

```go
import "github.com/bopher/utils"
sorts := utils.MongoSkip(30)) // => {$skip: 30}
```

### MongoLookup

Generate mongo $lookup.

```go
import "github.com/bopher/utils"
lookup := utils.MongoLookup("users", "user_id", "_id", "user"))
// => {
//     $lookup: {
//         from: "users",
//         localField: "user_id",
//         foreignField: "_id",
//         as: "user"
//     }
// }
```

### MongoUnwind

Generate mongo $unwind (with preserveNull).

```go
import "github.com/bopher/utils"
unwind := utils.MongoUnwind("invoices"))
// => {
//     $unwind: {
//         path: "invoices",
//         preserveNullAndEmptyArrays: true,
//     }
// }
```

### MongoUnwrap

Get first item of array and insert to doc using $addFields.

```go
import "github.com/bopher/utils"
unwrapped := utils.MongoUnwrap("logins", "first_login")) // => {$addFields: {first_login: {$first: logins}}}
```

### MongoGroup

Generate mongo $group.

```go
import "github.com/bopher/utils"
group := utils.MongoGroup(utils.MongoMap("_id", "$_id", "last_login", utils.MongoMap("$last", "$logins"))))
// => {
//     $group:{
//         _id: $_id,
//         last_login: {$last: $logins}
//     }
// }
```

### MongoSetRoot

Generate mongo $replaceRoot.

```go
import "github.com/bopher/utils"
setRoot := utils.MongoSetRoot("$_record")) // => {$replaceRoot: {newRoot: $_record}}
```

### MongoMergeRoot

Generate $replaceRoot with $mergeObject.

```go
import "github.com/bopher/utils"
mergeRoot := utils.MongoMergeRoot("$_record", "$$ROOT")) // => {$replaceRoot: {newRoot: {$mergeObjects: ["$_record", "$$ROOT"]}}}
```

### MongoUnProject

Generate $project to remove fields from result.

```go
import "github.com/bopher/utils"
unProject := utils.MongoUnProject("_record", "referral")) // => {$project: {_record:0, referral: 0 }}
```

### ParseObjectID

Parse object id from string. This function return nil if object id is invalid!

```go
import "github.com/bopher/utils"
oId := utils.ParseObjectID("6184011af9530d2ec143ae38")
```

### IsValidOId

Check if object id is valid and not zero.

```go
import "github.com/bopher/utils"
valid := utils.IsValidOId(oIdObject)
```

### Behaviour 
If a field is included in the fieldmask and the receiver is a pointer, the field will be initialized to the zero value of the type pointed to.
```go
/*
proto: message s {
    string field_a = 1;
}
field_mask: {"field_a"}
message: {}
*/
type s struct {
    FieldA *string `field_mask:"field_a"`
}
var ss s
_ = field_mask.Apply(field_mask, message, &ss)

// FieldA is now a pointer to an empty string to indicate that the value should be set to zero value as opposed to not set at all
// *ss.FieldA == ""
```

<hr>
If a struct is embedded and lacking a struct tag; the type of the struct is not used as it is in a golang struct literal.

```go
/*
proto: message s {
    string name = 1;
    int32 count = 2;
}
field_mask: {"name", "count"}
message: {"name": "test", "count": 1} <- note that it's just "count", not "ChildStruct.count"
*/

type s struct {
    ChildStruct
    Name string `field_mask:"name"`
}

type ChildStruct {
    Name string `field_mask:"name"`
    Count int32 `field_mask:"count"`
}

var ss s
_ = field_mask.Apply(field_mask, message, &ss)

fmt.Println("%#v", ss)
// Output: s{ChildStruct: ChildStruct{Name:"", Count:1}, Name:"test"}
```

If the embedded struct is tagged it will behave as if it were a normal field with type struct

```go
/*
proto: message s {
    string name = 1;
    embedded e = 2;
}
message embedded {
    string name = 1;
    int32 count = 2;
}

field_mask.Paths: ["name", "e.count", "e.name"]

message: {"name": "test", e: {"name: "embedded name", "count": 1}}
*/

type s struct {
    ChildStruct `field_mask:"e"`
    Name string `field_mask:"name"`
}

type ChildStruct {
    Name string `field_mask:"name"`
    Count int32 `field_mask:"count"`
}

var ss s
_ = field_mask.Apply(field_mask, message, &ss)

fmt.Println("%#v", ss)
// Output: s{ChildStruct: ChildStruct{Name:"embedded name", Count:1}, Name:"test"}
```
### Testing

#### Benchmarks

Benchmarks are run with 
```shell
go test -bench=. -benchtime=100000x -benchmem -count=10 ./... | tee testResults.txt && benchstat benchmarks/testResults.txt testResults.txt
```

# Field Mask

`field_mask` is a Go library designed to help manage field masks in protocol buffers.

This library is meant to complement the [google.golang.org/protobuf/types/known/fieldmaskpb](https://google.golang.org/protobuf/types/known/fieldmaskpb) package with helpers for working with field masks and mapping between protobuf messages and custom struct types.

See [the protobuf documentation](https://protobuf.dev/reference/protobuf/google.protobuf/#field-mask), [the go documentation for fieldmaskpb.FieldMask](https://pkg.go.dev/google.golang.org/protobuf/types/known/fieldmaskpb#FieldMask), [134. Standard methods: Update](https://google.aip.dev/134), [161. Field masks](https://google.aip.dev/161) for more information on field masks

### Features

* Use struct tags to map protobuf fields to custom structs
  * Flatten message structures by hoisting child-fields
  * Supports missing field_mask wildcard case
  * Supports nested structures, lists and maps
* Update mask: Map a protobuf message to a custom type filtered by a field mask
  * Use pointers to differentiate between unset/filtered values and values to be set to their zero-value as described [here](https://protobuf.dev/reference/protobuf/google.protobuf/#field-masks-updates)

### What
This library is meant to be used to filter protobuf messages based on a field mask and then populate a user defined `struct` based on struct tags.

### Why
Doing full updates can become problematic when a client and the backend has different versions of the protobuf messages of the resource being updated; the client will set any fields that it doesn't know about to zero. This is described in more detail [here](https://protobuf.dev/programming-guides/api/#support-partial-updates).  
The [api best practices](https://protobuf.dev/programming-guides/api/#support-partial-updates) suggest either defining smaller endpoints updating only a single field and batching the requests **or** 
 using a field mask to specify what fields are meant to be set and ignore the rest.   

## Quickstart

### Installation
```bash
go get -u github.com/captainlettuce/field_mask
```

### Usage
```go
type MyUpdateRequest struct {
    FieldA *string `field_mask:"field_a"`
    FieldB *string `field_mask:"field_b"`
}

var req MyUpdateReqest
fieldMask := fieldmaskpb.FieldMask{Paths: []string{"field_a"}}

// pb.UpdateRequest is a generated protobuf message
_ = fieldmask.Apply(&pb.UpdateRequest{FieldA: "test", FieldB: "other field"}, fieldMask, &req)

fmt.Printf("%v\n", req)
// output: {FieldA: <pointer to "test">, FieldB: nil}
```

## Behaviour

### Struct mapping

#### Handling field resetting (deleting current value)

If a field is included in the field mask, the corresponding field in the message is empty, and the receiver is a pointer, the field will be initialized to the zero value of the type pointed to, indicating that the field should be deleted.

```go
/*
proto: message s {
    string field_a = 1;
}
field_mask: {"field_a"}
message: {"field_a": ""}
*/
type s struct {
    FieldA *string `field_mask:"field_a"`
}
var ss s
_ = field_mask.Apply(field_mask, message, &ss)

// FieldA is now a pointer to an empty string to indicate that the value should be set to zero value as opposed to not touched at all
// *ss.FieldA == ""
```

<hr>

#### Go struct embedding
If a struct is embedded and lacking a struct tag; the embedded structs fields are treated as direct children.  
Name conflicts are resolved by default go precedence rules with the altercation that if two fields are at the same level and one is tagged and the other is not: the tagged field will take precedence (like the encoding/json library).

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

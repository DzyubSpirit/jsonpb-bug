package main

import (
	fmt "fmt"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	st1 := &Structure1{}
	jsonpb.UnmarshalString(`{"field1":"hello","field2":{},"field3":"world"}`, st1)
	fmt.Printf("%v\n", st1)

	st2 := &Structure2{}
	jsonpb.UnmarshalString(`{"field1":"hello","field3":{},"field2":"world"}`, st2)
	fmt.Printf("%v\n", st2)
}

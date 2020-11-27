package main

import (
	"fmt"
	"encoding/json"
	"time"
)

type testStruct struct {
	A int
	B int
	C string
	D string
	E []string
	F []string
	G []string
	H string
	I string
}

func main() {

	encoded, _ := json.Marshal( testStruct{ 1, 2, "one", "two", []string{ "cow", "chicken", "pig", "horse" },
	[]string{ "cow", "chicken", "pig", "horse" }, []string{ "cow", "chicken", "pig", "horse" }, "foo", "buz" } );
	fmt.Println(string(encoded))
	

	{
	    start := time.Now()
		var in testStruct
		for i := 0; i < 10000; i++ {
			json.Unmarshal( encoded, &in )
		}
	    fmt.Println(time.Since(start).Milliseconds() / 1000.0 )
	    fmt.Println(in)
	}

	{
	    start := time.Now()
		var in struct { I string }
		for i := 0; i < 10000; i++ {
			json.Unmarshal( encoded, &in )
		}
	    fmt.Println(time.Since(start).Milliseconds() / 1000.0 )
	    fmt.Println(in)
	}

	//fmt.Println(in)
}

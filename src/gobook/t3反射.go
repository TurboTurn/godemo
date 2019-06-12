package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name string
	Life int
}

func (b *Bird) Fly() {
	fmt.Println("I am flying...")
}

func main() {
	sparrow := &Bird{"sparrow", 3}
	s := reflect.ValueOf(sparrow).Elem()

	//fmt.Println(reflect.TypeOf(sparrow),reflect.ValueOf(sparrow))

	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("\n%d: %s %s = %v", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

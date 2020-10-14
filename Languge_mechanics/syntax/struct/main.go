package main

import (
	"fmt"
	"reflect"
)

type Parrents struct {
	MotherName string
	FatherName string
}

type Person struct {
	Name   string
	Age    int
	Family Parrents
}

func main() {
	p1 := &Person{
		Name: "genson",
		Age:  20,
		Family: Parrents{
			FatherName: "my father :))",
			MotherName: "my mother :))",
		},
	}

	fmt.Printf("%v\n", p1)

	v := reflect.ValueOf(p1).Elem()

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)

		fmt.Printf("- index: %d - name: %s - type: %s - value: %v\n",
			i, f.Type().Name(), f.Type(), f.Interface())

	}

	fmt.Println(v)

	/*
		&{genson 20 {my mother :)) my father :))}}
		- index: 0 - name: string - type: string - value: genson
		- index: 1 - name: int - type: int - value: 20
		- index: 2 - name: Parrents - type: main.Parrents - value: {my mother :)) my father :))}
		{genson 20 {my mother :)) my father :))}}
	*/
}

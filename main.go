package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Age    int
	Nested []Nested
}

type Nested struct {
	Field1 string
	Field2 int
}

func main() {
	user1 := User{
		Name: "hiro",
		Age:  30,
		Nested: []Nested{
			{
				Field1: "test",
				Field2: 42,
			},
		},
	}
	user2 := User{
		Name: "john",
		Age:  30,
		Nested: []Nested{
			{
				Field1: "different",
				Field2: 42,
			},
		},
	}

	NotEqualStruct(user1, user2)
}

func NotEqualStruct(s1, s2 any) {
	v1 := reflect.ValueOf(s1)
	v2 := reflect.ValueOf(s2)

	// If s1 or s2 is not a struct, return early
	if v1.Kind() != reflect.Struct || v2.Kind() != reflect.Struct {
		fmt.Println("Both argument should be struct or pointers to a struct")
	}

	// If s1 or s2 is a pointer to a struct, get the actual value
	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}
	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	for i := 0; i < v1.NumField(); i++ {
		f1 := v1.Field(i)
		f2 := v2.Field(i)
		typeField := v1.Type().Field(i)

		if !reflect.DeepEqual(f1.Interface(), f2.Interface()) {
			fmt.Printf("Not equal:field %s\n", typeField.Name)
		}
	}
}

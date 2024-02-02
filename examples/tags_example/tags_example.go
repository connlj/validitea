package main

import (
	"fmt"
	"github.com/connerlj/validitea/validation"
)

type User struct {
	Id    int    `validate:"required"`
	Name  string `validate:"min=2,max=10"`
	Email string `validate:"required"`
}

func main() {
	//name := "e"

	user := User{
		Id:    1,
		Name:  "1",
		Email: "john@example",
	}

	v := validation.FromStructTags(user)
	err := v.Validate()
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println()
	//userNoEmail := User{
	//	Id:    1,
	//	Name:  "",
	//	Email: "",
	//}
	//
	//v = validation.FromStructTags(userNoEmail)
	//err = v.Validate()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//s := "test"
	//fmt.Println(reflect.ValueOf(s), reflect.ValueOf(s).Kind())
	//
	//u := User{Id: 1}
	//fmt.Println(reflect.ValueOf(u), reflect.ValueOf(u).Kind())
	//fmt.Println(reflect.ValueOf(u.Id), reflect.ValueOf(u.Id).Kind())
	//
	//func(a any) {
	//	fmt.Println(reflect.ValueOf(a), reflect.ValueOf(a).Kind())
	//}(&User{})

	//p := "pretty please"
	//val := reflect.ValueOf(p)
	//fmt.Println(val.Kind())
	//
	//switch val.Kind() {
	//case reflect.String:
	//	fmt.Println("here")
	//}

	v = validation.New()
	v.Add("test", "normal add", validation.ValidateMinLength(5))
	//v.Add("test", "no", validation.ValidateMinLength(5))
	err = v.Validate()
	if err != nil {
		fmt.Println(err)
	}

	//v := reflect.ValueOf("test")
	//i := reflect.Indirect(reflect.ValueOf(v))
	//fmt.Println(v, i)

	//reflectValue := reflect.Indirect(reflect.ValueOf(value))
	//for reflectValue.Kind() == reflect.Ptr || reflectValue.Kind() == reflect.Interface {
	//	reflectValue = reflect.Indirect(reflectValue)
	//}
}

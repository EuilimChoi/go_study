package study

import "fmt"

//일반적인 const,var 선언은 타입을 붙혀줘야함
const name string = "lim"
const theName string = "Euilim"

//짧은 방식으로 선언하면 자동으로 타입을 찾아주지만 함수내에서만 사용할 수 있다!
func showname() {
	fullname := "Choieuilim" //string
	boo := false             //bool
	fmt.Println(fullname, boo)
}

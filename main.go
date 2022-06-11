package main

import "fmt"

func main() { //대문자로 작성한 함수, 클래스만 export 할수 있다
	// const name string = "nico"
	name := "nico" //이렇게 정의하면 자동으로 타입을 정해주고 var로 정의한것과 같다.
	fmt.Printf((name))
}

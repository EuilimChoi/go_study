package study

import "fmt"

func canIDrink(age int) bool { //자바스크립트의 if만 비슷하네

	if koreanAge := age + 2; koreanAge < 18 { //if문 앞에 생성한 변수는 해당 if문 블럭에서만 사용된다. 이건 코드 설명해좀 꿀이겠네.
		return false
	} else {
		return true
	}
}

func getcanIDrink() {
	fmt.Println(canIDrink(16))
} // false

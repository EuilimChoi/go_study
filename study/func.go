package study

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int { //파라미터를 저렇게 선언하면 둘다 숫자라는 뜻이다, 그게 아니면 각각 선언해 주면된다. 타입스크립트랑 살짝 비슷한 느낌
	return a * b
}

//go는 함수에서 두개의 리턴을 가질 수 있다, 이왜됨;;

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
} // 이 함수는 이름의 길이와 대문자로 만든 이름 두개를 뱉는다. 따라서 함수를 사용할때도 두개의 변수로 받아야 한다.

func anotherLenAndUpper(name string) (length int, uppername string) { //이렇게 하면 naked return 이라고 부르는데 특별히 리턴으로 변수를 설정하지 않아도 함수 선언 부분에서 리턴 변수를 선언함으로서 자동적으로 변수에 할당되서 리턴된다.
	length = len(name)
	uppername = strings.ToUpper(name)

	return
}

func dolenAndUpper() {
	Namelength, uppername := lenAndUpper("Euilimchoi")
	fmt.Print(Namelength, uppername)
} // 결과는 10 EUILIMCHOI
// 만약 둘중 하나만 받고 싶다면 받는 변수에 _로 작성해 무시해주면된다.
// _, uppername := lenAndUpper("Euilimchoi") => EUILIMCHOI 가된다.

func repeatMe(word ...string) { //많은 단어를 한번에 뱉어주는? 함수다 ...을 붙히면 무제한 개의 변수를 같은 타입으로 받을 수 있다.
	fmt.Println((word))
}

func dorepeatMe() {
	repeatMe("nico", "lynn", "dal", "marl", "flynn")
}

//defer!!
//defer는 함수가 리턴 된고 나서 실행되는 부분을 정의한다. 이건 좀 개쩐다.

func withDeferLenAndUpper(name string) (length int, uppername string) {
	//defer를 작성하면 이 함수의 결과가 리턴되고나면 defer의 함수가 실행된다 이경우 i'm done 이 출력된다!!
	defer fmt.Println(("I'm done"))
	uppername = strings.ToUpper(name)

	return
}

package study

import "fmt"

func array() { //아래의 경우 length가 5로 고정되기 때문에 이 안에서 자유롭게 추가, 제거가 가능하다.
	names := [5]string{"sldknflk", "sdfsdf", "basdfsss"}
	names[3] = "ssss"
	names[4] = "bbbb"

	fmt.Println(names)
}

func slice() { //아래의 경우 length가 없는 slice가 된다, 즉 array와 같은 방식으로는 추가가 안된다 추가하려면 append를 사용해야한다.
	names := []string{"sldknflk", "sdfsdf", "basdfsss"}
	names[3] = "ssss"             // 이건 작동안함
	names = append(names, "bbbb") // append는 요소가 추가된 새로운 slice를 반환한다.

	fmt.Println(names)
}

func mapMaking() { //map은 자바스크립트의 객체와 비슷한 개념이다
	lim := map[string]string{"name": "lim", "age": "32"}

	for key, value := range lim {
		fmt.Println(key, value) // name lim age 32 가 출력됨
	}
}

type person struct { //struct 는 타입스트립트에서 인터페이스 같은 느낌?
	name    string
	age     int
	favFood []string
}

func structMaking() {
	favFood := []string{"apple", "banana"}
	//lim := person{"lim", 18, favFood} 이런식으로 써도 되지만 코드 가독성이 안좋다, 어떤 필드가 어떤 내용인지 한눈에 보기 힘들고 위에서 person을 다시 한번 봐야하니까...
	lim := person{name: "lim", age: 18, favFood: favFood}
	fmt.Println(lim)
}

package study

// go는 for 밖에 없다 다른건 읎어~~~~

func superAdd(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total = total + number // 또는 total += number <= 자바스크립트에서 자주봤죠?

		//range로 그냥 돌리면 index를 뱉는다. 0,1,2,3,4,5 가 나온다는 것이지. 우리가 의도한대로 하려면 for index, number := range numbers{...} 형태로 해야한다.
	}

	//  for i:=0; i < len(numbers); i++ {} 형태 처럼 그냥 자바스크립트 for 처럼 사용해도 된다.
	return total
}

func getSuperAdd() {
	superAdd(1, 2, 3, 4) // 10
}

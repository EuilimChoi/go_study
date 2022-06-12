package main

import (
	"fmt"

	"github.com/usrname/learngo/mydict"
)

// func main() { //private한 상태로 import해와서 사용
// 	account := accounts.NewAccount("lim")
// 	account.Deposit(1000) // 이런식으로 account 타입을 가지고 있는 리시버를 사용할 수 있다.
// 	fmt.Println(account.Balance())
// 	err := account.Withdraw(2000)
// 	if err != nil {
// 		// log.Fatalln(err) 이렇게 하면 프로그램이 죽어버린다. 죽이면서 에러 체크하기
// 		fmt.Println(err)
// 	}
// 	fmt.Println(account.Balance(), account.ReturnOwner())
// }

// func main() { //단어를 찾고 없으면 추가하기 테스트
// 	dictionary := mydict.Dictionary{"first": "First word"}
// 	word := "hello"
// 	definition := "Greeting"
// 	err := dictionary.Add(word, definition)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	hello, _ := dictionary.Search(word)
// 	fmt.Println(hello)

// 	err2 := dictionary.Add(word, definition)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}
// }

// func main() { //업데이트 예시
// 	dictionary := mydict.Dictionary{}
// 	baseWord := "hello"
// 	dictionary.Add(baseWord, "First")
// 	err := dictionary.Update("sbasdfasdfs", "Second")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	word, _ := dictionary.Search(baseWord)
// 	fmt.Println(word)
// }

func main() { // 삭제 예시 만들고 찾고 삭제하고 다시 찾으면 Notfound!
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}

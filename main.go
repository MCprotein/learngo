package main

import (
	"fmt"
	abc "fmt"

	"github.com/MCprotein/learngo/something"

	"rsc.io/quote/v4"
)

/*
*

  - go에서는 function을 export 하려면 첫글자를 대문자로 해야한다.

    축약형 := 는 변수를 선언하고 값을 할당하는 것을 한번에 해준다.
    func 안에서만 가능하고, 변수에만 적용할 수 있다.
    := 를 사용하여 지정된 타입은 개발자가 변경할 수 없다.
*/
func main() {
	// const name string = "nico"

	var name2 string = "nico"
	name2 = "lynn"

	name3 := "nico"

	abc.Println("Hello, World!")
	abc.Println(name2)
	abc.Println(name3)

	abc.Println(quote.Go())

	sth_case := something.CaseStruct{}
	sth_case.UpperValue = "Upper value."
	fmt.Println("sth_case's UpperValue:", sth_case.UpperValue)
	something.UpperSth()
	something.ExternalPackage()
	fmt.Println("final Cnt :", something.Cnt)
}

func init() {
	fmt.Println("init called in main package")
	something.Cnt = 1
}

package main

import (
	"fmt"
	abc "fmt"
	"strings"

	"github.com/MCprotein/learngo/accounts"
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

	/**
	* &는 메모리 주소를 보여준다.
	* *는 메모리 주소에 있는 값을 보여준다.
	* *를 사용하여 해당 메모리 주소에 있는 값을 변경할 수 있다.

	 */
	name3 := &name2

	abc.Println("Hello, World!")
	abc.Println(&name2, "name2의 메모리 주소")
	abc.Println(name3, "name3")
	abc.Println(*name3, "name3의 메모리 주소에 있는 값")

	*name3 = "dal"
	abc.Println(name2, "*name3에 의해 변경된 name2")

	abc.Println(quote.Go())

	sth_case := something.CaseStruct{}
	sth_case.UpperValue = "Upper value."
	fmt.Println("sth_case's UpperValue:", sth_case.UpperValue)
	something.UpperSth()
	something.ExternalPackage()
	fmt.Println("final Cnt :", something.Cnt)

	fmt.Println(multiply(2, 2))
	/**
	* _ 는 무시하고 싶을 때 사용한다.
	* totalLength, _ := lenAndUpper("nico")
	* 객체가 아니기때문에 값을 하나만 받을 수는 없다.
	 */
	totalLength, _ := lenAndUpper("nico")
	fmt.Println(totalLength)

	totalLength2, upper := lenAndUpperNakedReturn("nico")
	fmt.Println(totalLength2, upper)

	repeatMe("nico", "lynn", "dal", "marl", "flynn")

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)

	fmt.Println(canIDrink(16))

	array()
	mapFunc()
	structFunc()

	account := accounts.NewAccount("nico")
	fmt.Println(account)
}

func init() {
	fmt.Println("init called in main package")
	something.Cnt = 1
}

/*
*
func multiply(a int, b int) int {
인자의 타입이 같으면 아래와 같이 표현할 수 있다.
*/
func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

/*
*
naked return
*/
func lenAndUpperNakedReturn(name string) (length int, uppercase string) {
	/**
	* defer는 함수가 끝나고 실행된다.
	 */
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

/*
*
javascript의 spread operator와 비슷하다.
*/
func repeatMe(words ...string) {
	fmt.Println(words)

}

func superAdd(numbers ...int) int {
	// for index, number := range numbers {
	// 	fmt.Println(index, number)
	// }

	total := 0

	for _, number := range numbers {
		total += number
	}

	return total

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
}

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}
	return true

}

func canIDrinkSwitch(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}

	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	return false
}

func array() {
	names := [5]string{"nico", "lynn", "dal"}
	names[3] = "alala"
	names[4] = "alala"
	fmt.Println(names)

	/**
	slice도 array긴한데 length가 없다.
	append는 javascript의 array.prototype.push 와는 다르게 새로운 array를 리턴한다.
	*/
	namess := []string{"nico", "lynn", "dal"}
	namess = append(namess, "flynn", "oooo")
	fmt.Println(namess)
}

func mapFunc() {
	nico := map[string]string{"name": "nico", "age": "12"}
	fmt.Println(nico)

	for key, value := range nico {
		fmt.Println(key, value)
	}
}

func structFunc() {
	/**
	* struct는 객체를 만들어준다.
	* 객체를 만들 때는 new를 사용하지 않는다.
	 */
	type person struct {
		name    string
		age     int
		favFood []string
	}
	favFood := []string{"kimchi", "ramen"}
	// nico := person{"nico", 18, favFood}
	nico := person{name: "nico", age: 18, favFood: favFood}
	fmt.Println(nico.name)
}

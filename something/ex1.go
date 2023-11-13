package something

import "fmt"

func init() {
	fmt.Println("ex1.go init 1")
	Cnt = 10
}

type CaseStruct struct {
	UpperValue string
	lowerValue string
}

func UpperSth() {
	fmt.Println("UpperSth() Called")
	lowerSth()
	ExportSth()
	cannotExportSth()
}

func lowerSth() {
	fmt.Println("lowerSth() Called")
}

func init() {
	fmt.Println("ex1.go init 2")
}

package something

import (
	"fmt"

	"github.com/google/uuid"
)

var Cnt int8

func ExportSth() {
	fmt.Println("exportSth() Called")
}

func cannotExportSth() {
	fmt.Println("cannotExportSth() Called")
}

func init() {
	fmt.Println("ex2.go init 1")
	Cnt = 20
}

func init() {
	fmt.Println("ex2.go init 2")
}

func ExternalPackage() {
	fmt.Println(uuid.New(), "uuid.New()")
}

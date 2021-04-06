package zlog

import (
	"fmt"
)

func Print(str string, data ...interface{}) {
	fmt.Print(Format(str, data...))
}

func Println(str string, data ...interface{}) {
	fmt.Println(Format(str, data...))
}

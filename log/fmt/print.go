package fmt

import (
	"fmt"

	"github.com/liuxiaobopro/qsgo/global"
)

func Println(a ...any) (n int, err error) {
	if global.Debug {
		return fmt.Println(a...)
	} else {
		return 0, nil
	}
}

func Printf(format string, a ...any) (n int, err error) {
	if global.Debug {
		return fmt.Printf(format, a...)
	} else {
		return 0, nil
	}
}
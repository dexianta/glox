package errorhandle

import (
	"fmt"
)

var HadError = false

func Report(line int, where, msg string) {
	fmt.Printf("[line \"%d\"] Error %s \": \" %s\n", line, where, msg)
	HadError = true
}
package akcl_mod_cal

import "fmt"

var Version = "0.0.0"
var LogMessage = "cal"

func Sum(x, y int) int {
	info()
	return x + y
}

func info() {
	fmt.Println(LogMessage)
}

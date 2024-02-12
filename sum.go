package calculator

// private
var logMessage = "[LOG]"

// public
var Version = "1.0"

// private
func internalSum(num int) int {
	return num - 1
}

// public
func Sum(num1, num2 int) int {
	return num1 + num2
}

package cfmt

import "fmt"

const clrReset = "\033[0m"

const clrRed = "\033[31m"
const clrGreen = "\033[32m"
const clrYellow = "\033[33m"
const clrBlue = "\033[34m"
const clrMagenta = "\033[35m"
const clrCyan = "\033[36m"
const clrGray = "\033[37m"
const clrWhite = "\033[97m"

const clrLRed = "\033[31;1m"
const clrLGreen = "\033[32;1m"
const clrLYellow = "\033[33;1m"
const clrLBlue = "\033[34;1m"
const clrLMagenta = "\033[35;1m"
const clrLCyan = "\033[36;1m"
const clrLGray = "\033[37;1m"
const clrLWhite = "\033[97;1m"

//todo: add log level

func PrintlnClr(clr string, args ...interface{}) {
	fmt.Print(clr)
	for _, arg := range args {
		fmt.Print(arg)
		fmt.Print(" ")
	}
	fmt.Println(clrReset)
}

func PrintlnInfo(args ...interface{}) {
	PrintlnClr(clrLCyan, args...)
}

func PrintlnUser(args ...interface{}) {
	PrintlnClr(clrLWhite, args...)
}

func PrintlnErr(args ...interface{}) {
	PrintlnClr(clrRed, args...)
}

func PrintlnWarn(args ...interface{}) {
	PrintlnClr(clrLRed, args...)
}

func PrintlnOk(args ...interface{}) {
	PrintlnClr(clrGreen, args...)
}

func PrintlnLine(args ...interface{}) {
	PrintlnClr(clrCyan, args...)
}

func PrintlnFunc(args ...interface{}) {
	PrintlnClr(clrLWhite, args...)
}

func PrintlnDbg(args ...interface{}) {
	PrintlnClr(clrLMagenta, args...)
}

func PrintlnImp(args ...interface{}) {
	PrintlnClr(clrLYellow, args...)
}

func PrintlnArg(args ...interface{}) {
	PrintlnClr(clrLBlue, args...)
}

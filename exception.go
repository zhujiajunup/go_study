package main

import (
	"fmt"
)

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dvidee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func main() {
	if result, msg := Divide(100, 10); msg == "" {
		fmt.Println("100/10 = ", result)
	}

	if _, msg := Divide(19, 0); msg != "" {
		fmt.Println("error msg: ", msg)
	}
}

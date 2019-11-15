package util

import (
	"errors"
	"fmt"
)

func OutputStep(msg string) {
	fmt.Printf("%s💡%s %s\n", Blue, Reset, msg )
}

func OutputOk(msg string) {
	output(Green, "[OK]", msg)
}

func OutputInfo(title, msg string) {
	output(Blue, "[INFO]", title+":"+msg)
}

func OutputWarn(msg string) {
	output(Yellow, "[WARNING]️", msg)
}

func OutputError(msg string) error {
	return errors.New(soutput(Red, "[ERROR]", msg))
}

func OutputFail(msg string) error {
	return errors.New(soutput(Red, "[FAIL]", msg))
}

func output(color, header, msg string) {
	fmt.Printf("%s %s %s %s\n", color, header, Reset, msg)
}
func soutput(color, header, msg string) string {
	return fmt.Sprintf("%s %s %s %s\n", color, header, Reset, msg)
}

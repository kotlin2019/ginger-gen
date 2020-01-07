package xprint

import (
	"errors"
	"fmt"
)

func Step(msg string) {
	fmt.Printf("%s💡%s %s\n", Blue, Reset, msg )
}

func Ok(msg string) {
	output(Green, "[OK]", msg)
}

func Info(title, msg string) {
	output(Blue, "[INFO]", title+":"+msg)
}

func Warn(msg string) {
	output(Yellow, "[WARNING]️", msg)
}

func Error(msg string) error {
	return errors.New(soutput(Red, "[ERROR]", msg))
}

func Fail(msg string) error {
	return errors.New(soutput(Red, "[FAIL]", msg))
}

func output(color, header, msg string) {
	fmt.Printf("%s %s %s %s\n", color, header, Reset, msg)
}
func soutput(color, header, msg string) string {
	return fmt.Sprintf("%s %s %s %s\n", color, header, Reset, msg)
}

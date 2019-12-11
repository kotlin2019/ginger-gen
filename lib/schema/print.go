package schema

import (
	"errors"
	"fmt"
	"github.com/gofuncchan/ginger-cli/util"
)

func OutputStep(msg string) {
	fmt.Printf("%s💡%s %s\n", util.Blue, util.Reset, msg )
}

func OutputOk(msg string) {
	output(util.Green, "[OK]", msg)
}

func OutputInfo(title, msg string) {
	output(util.Blue, "[INFO]", title+":"+msg)
}

func OutputWarn(msg string) {
	output(util.Yellow, "[WARNING]️", msg)
}

func OutputError(msg string) error {
	return errors.New(soutput(util.Red, "[ERROR]", msg))
}

func OutputFail(msg string) error {
	return errors.New(soutput(util.Red, "[FAIL]", msg))
}

func output(color, header, msg string) {
	fmt.Printf("%s %s %s %s\n", color, header, util.Reset, msg)
}
func soutput(color, header, msg string) string {
	return fmt.Sprintf("%s %s %s %s\n", color, header, util.Reset, msg)
}

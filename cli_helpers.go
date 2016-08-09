package trousseau

import (
	"fmt"

	"github.com/howeyc/gopass"
)

func PromptForPassword() string {
	fmt.Printf("Password: ")
	pass, _ := gopass.GetPasswd()

	return string(pass)
}

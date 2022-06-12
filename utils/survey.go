package utils

import (
	"bufio"
	"fmt"
	"os"
)

type Question struct {
	Validate func(value string) error
	Default  string
	Message  string
}

func Ask(ques Question, ans *string) {

	a := ques.Default

	if ques.Default != "" {
		fmt.Print(Colorize("green", "? ") + ques.Message + " (" + ques.Default + ")" + colorGray + " ")
	} else {
		fmt.Print(Colorize("green", "? ") + ques.Message + colorGray + " ")
	}

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		if scanner.Text() != "" {
			a = scanner.Text()
		}
	}

	fmt.Print(colorReset)

	ok := ques.Validate(a)

	if ok != nil {
		fmt.Println("\033[A                                                                  \033[A")
		fmt.Println(Colorize("red", "X ") + Colorize("red", ok.Error()))
		Ask(ques, ans)
		return
	}

	fmt.Println("\033[A                                                                      \033[A")
	fmt.Println(Colorize("green", "? ") + ques.Message + " " + Colorize("cyan", a))

	*ans = a

}

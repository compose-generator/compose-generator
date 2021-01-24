package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func Heading(text string) {
	green := color.New(color.FgGreen).Add(color.Bold)
	green.Println(text)
}

func TextQuestion(question string) string {
	reader := bufio.NewReader(os.Stdin)
	cyan := color.New(color.FgCyan)

	cyan.Print(question)

	result_string, _ := reader.ReadString('\n')
	result_string = strings.TrimRight(result_string, "\r\n")
	if result_string == "" {
		fmt.Println("Error. This value is required!")
		os.Exit(1)
	}
	return result_string
}

func TextQuestionWithDefault(question string, default_value string) string {
	reader := bufio.NewReader(os.Stdin)
	cyan := color.New(color.FgCyan)

	cyan.Print(question)

	result_string, _ := reader.ReadString('\n')
	result_string = strings.TrimRight(result_string, "\r\n")
	if result_string != "" {
		return result_string
	} else {
		return default_value
	}
}

func YesNoQuestion(question string, default_value bool) bool {
	reader := bufio.NewReader(os.Stdin)
	cyan := color.New(color.FgCyan)

	if default_value {
		cyan.Print(question + " [Y/n]: ")
	} else {
		cyan.Print(question + " [y/N]: ")
	}

	result_string, _ := reader.ReadString('\n')
	result_string = strings.TrimRight(result_string, "\r\n")
	result := default_value
	if result_string != "" {
		result = strings.ToLower(result_string) == "y"
	}
	return result
}

func MenuQuestion(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
		//Stdout: &bellSkipper{},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}

// -------------------------- Skip bell sound output on select questions --------------------------

type bellSkipper struct{}

func (bs *bellSkipper) Write(b []byte) (int, error) {
	const charBell = 7 // c.f. readline.CharBell
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

func (bs *bellSkipper) Close() error {
	return os.Stderr.Close()
}

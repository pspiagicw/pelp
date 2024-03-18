package pelp

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Version(name string, version string) {
	fmt.Printf("%s version: '%s'\n", name, version)
}
func HeaderWithDescription(header string, description []string) {
	fmt.Println(strings.ToUpper(header))
	for _, desc := range description {
		fmt.Printf("  %s\n", desc)
	}
	fmt.Println()
}
func Aligned(header string, leftElement []string, rightElements []string) {
	left := strings.Join(leftElement, "\n")
	right := strings.Join(rightElements, "\n")
	fmt.Println(strings.ToUpper(header))

	commandCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(left).MarginLeft(2).String()
	messageCol := lipgloss.NewStyle().Align(lipgloss.Left).SetString(right).MarginLeft(5).String()

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Bottom, commandCol, messageCol))
	fmt.Println()
}
func Examples(header string, examples []string) {
	fmt.Println(strings.ToUpper(header))
	for _, example := range examples {
		fmt.Printf("  $ %s\n", example)
	}
	fmt.Println()
}
func Flags(header string, flags []string, description []string) {

	for i, flag := range flags {
		flags[i] = "--" + flag
	}

	Aligned(header, flags, description)
}
func Print(message string) {
	fmt.Println(message)
	fmt.Println()
}

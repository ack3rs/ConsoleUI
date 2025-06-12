package main

import (
	"fmt"

	v1 "github.com/ack3rs/ConsoleUI/v1"
)

func main() {

	// Clear the Screen.
	// v1.Clear()

	// Draw a Up/Down Prompt called Title in Style Blue at Position X, Y. for the Items in the List. (c)
	prompt := v1.DrawPrompt(
		[]string{"1. test 1",
			"2. test 2",
			"3. test 3",
			"4. What time is it",
			"5. Exit"},
		v1.Option{
			Title:    "Title",
			Style:    v1.PresetStyle("blank"),
			Position: v1.Relative,
		})

	fmt.Printf("\nThe Answer was %v", prompt)

}

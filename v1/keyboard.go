package v1

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
)

// Get retrieves a single character and a key press event from the keyboard. It initializes and closes the keyboard session.
func Get() (rune, keyboard.Key) {

	err := keyboard.Open()
	if err != nil {
		fmt.Printf("Unable to open  Keyboard: %v", err)
		os.Exit(0)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	char, key, err := keyboard.GetKey()
	if err != nil {
		fmt.Printf("Unable to get Keyboard: %v", err)
	}

	return char, key
}

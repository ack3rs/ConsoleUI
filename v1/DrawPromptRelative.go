package v1

import (
	"strings"

	"github.com/eiannone/keyboard"
)

func drawPromptRelativeInital(opt Option, items []string) {

	SelectedItem := 0

	// Go through all the Items and find the max length
	maxLength := 1

	for _, item := range items {
		if len(item) > maxLength {
			maxLength = len(item)
		}
	}

	Out(opt.Style.TopLeft)
	Out(opt.Title)
	Out(strings.Repeat(opt.Style.TopCenter, (len(opt.Style.Gap)*2)+maxLength-(len(opt.Style.HeaderText))))
	Out(opt.Style.TopRight)
	MoveDownAndBeginning(1)

	for i := 1; i <= opt.Style.SpaceHeader; i++ {
		MoveDownAndBeginning(1)
		Out(opt.Style.MiddleLeft + opt.Style.Gap + strings.Repeat(" ", maxLength) + opt.Style.Gap + opt.Style.MiddleRight)
	}

	for i, item := range items {

		if len(item) < maxLength {
			item = item + strings.Repeat(" ", maxLength-len(item))
		}

		if SelectedItem == i {
			item = opt.Style.TextHighLight + item + opt.Style.TextNormal
		}

		Out(opt.Style.MiddleLeft + opt.Style.Gap + item + opt.Style.Gap + opt.Style.MiddleRight)
		MoveDownAndBeginning(1)
	}

	Out(opt.Style.BottomLeft)
	Out(strings.Repeat(opt.Style.TopCenter, (len(opt.Style.Gap)*2)+maxLength-6))
	Out(opt.Style.Legend + "[ ↑↓ ]")

	Out(opt.Style.BottomRight)

}

func drawPromptRelativeUpdate(opt Option, SelectedItem int, items []string) {

	// Go back to the beginning
	maxLength := 1

	for _, item := range items {
		if len(item) > maxLength {
			maxLength = len(item)
		}
	}

	MoveY(len(items) + 1)
	MoveDownAndBeginning(1)
	for i, item := range items {

		if len(item) < maxLength {
			item = item + strings.Repeat(" ", maxLength-len(item))
		}

		if SelectedItem == i {
			item = opt.Style.TextHighLight + item + opt.Style.TextNormal
		}

		Out(opt.Style.MiddleLeft + opt.Style.Gap + item + opt.Style.Gap + opt.Style.MiddleRight)
		MoveDownAndBeginning(1)
	}

}

func drawPromptRelative(items []string, opt Option) string {

	drawPromptRelativeInital(opt, items)
	SelectedItem := 0

	HideCursor()
	defer Reset()
	defer ShowCursor()

	for {
		r, key := Get()

		// If the first character of the item is the same as the key pressed
		// TODO: Ignore if this is a special key

		for i, item := range items {
			rs := []rune(item)
			if rs[0] == r {
				return items[i]
			}
		}

		if key == keyboard.KeyArrowDown {
			SelectedItem++
			if SelectedItem > len(items)-1 {
				SelectedItem = 0
			}
		}

		if key == keyboard.KeyArrowUp {
			SelectedItem--
			if SelectedItem < 0 {
				SelectedItem = len(items) - 1
			}
		}

		if key == keyboard.KeyEnter {
			return items[SelectedItem]
		}

		// Nothing was pressed
		if key == keyboard.KeyEsc {
			return ""
		}
		drawPromptRelativeUpdate(opt, SelectedItem, items)
	}

}

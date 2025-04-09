package v1

import (
	"strings"

	"github.com/eiannone/keyboard"
)

func drawPrompAbsolutetIntial(myStyle Style, x int, y int, items []string) {

	MoveTo(x, y)
	SelectedItem := 0

	// Go through all the Items and find the max length
	maxLength := 1

	for _, item := range items {
		if len(item) > maxLength {
			maxLength = len(item)
		}
	}

	// Draw a Box Starting at X, Y with the MaxLength
	Out(myStyle.TopLeft)
	Out(myStyle.HeaderText)
	Out(strings.Repeat(myStyle.TopCenter, (len(myStyle.Gap)*2)+maxLength-(len(myStyle.HeaderText))))
	Out(myStyle.TopRight)

	for i := 1; i <= myStyle.SpaceHeader; i++ {
		MoveTo(x, y+i)
		Out(myStyle.MiddleLeft + myStyle.Gap + strings.Repeat(" ", maxLength) + myStyle.Gap + myStyle.MiddleRight)
	}

	newY := y + 1 + myStyle.SpaceHeader

	MoveTo(x, newY)

	for i, item := range items {

		if len(item) < maxLength {
			item = item + strings.Repeat(" ", maxLength-len(item))
		}

		if SelectedItem == i {
			item = myStyle.TextHighLight + item
		} else {
			item = myStyle.TextNormal + item
		}

		Out(myStyle.MiddleLeft + myStyle.Gap + item + myStyle.Gap + myStyle.MiddleRight)
		newY++
		MoveTo(x, newY)
	}

	for i := 1; i <= myStyle.SpaceFooter; i++ {
		Out(myStyle.MiddleLeft + myStyle.Gap + strings.Repeat(" ", maxLength) + myStyle.Gap + myStyle.MiddleRight)
		newY++
		MoveTo(x, newY)
	}

	Out(myStyle.BottomLeft)
	Out(strings.Repeat(myStyle.TopCenter, (len(myStyle.Gap)*2)+maxLength-6))
	Out(myStyle.Legend + "[ ↑↓ ]")

	Out(myStyle.BottomRight)
}

func drawPromptAbsoluteUpdate(myStyle Style, SelectedItem int, x int, y int, items []string) {

	maxLength := 1
	newY := y + 1 + myStyle.SpaceHeader

	for _, item := range items {
		if len(item) > maxLength {
			maxLength = len(item)
		}
	}

	MoveTo(x, newY)

	for i, item := range items {

		if len(item) < maxLength {
			item = item + strings.Repeat(" ", maxLength-len(item))
		}

		if SelectedItem == i {
			item = myStyle.TextHighLight + item
		} else {
			item = myStyle.TextNormal + item
		}

		Out(myStyle.MiddleLeft + myStyle.Gap + item + myStyle.Gap + myStyle.MiddleRight)
		newY++
		MoveTo(x, newY)
	}
}

func drawPromptAbsolute(items []string, opt Option) string {

	myStyle := opt.Style
	myStyle.HeaderText = opt.Title

	drawPrompAbsolutetIntial(myStyle, opt.Xpos, opt.Ypos, items)
	SelectedItem := 0

	HideCursor()
	defer Reset()
	defer ShowCursor()

	for {
		r, key := Get()

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

		if key == keyboard.KeyEsc {
			return ""
		}

		drawPromptAbsoluteUpdate(myStyle, SelectedItem, opt.Xpos, opt.Ypos, items)

	}
}

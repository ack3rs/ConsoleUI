package v1

// DrawPrompt renders a selectable prompt with items using the specified Option configuration and position mode.
func DrawPrompt(items []string, opt Option) string {

	if opt.Position == Relative {
		return drawPromptRelative(items, opt)
	} else {
		return drawPromptAbsolute(items, opt)
	}
}

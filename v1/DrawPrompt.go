package v1

func DrawPrompt(items []string, opt Option) string {

	if opt.Position == Relative {
		return drawPromptRelative(items, opt)
	} else {
		return drawPromptAbsolute(items, opt)
	}
}

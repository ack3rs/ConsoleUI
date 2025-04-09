package v1

type Style struct {
	TopLeft       string
	TopCenter     string
	TopRight      string
	MiddleLeft    string
	MiddleRight   string
	BottomLeft    string
	BottomCenter  string
	BottomRight   string
	SpaceAround   int
	SpaceHeader   int
	SpaceFooter   int
	HeaderText    string
	Gap           string
	TextNormal    string
	TextHighLight string
	Legend        string
}

func PresetStyle(PresetStyle string) Style {

	switch PresetStyle {
	case "blank":
		return Style{
			TopLeft:       " ",
			TopCenter:     " ",
			TopRight:      " ",
			MiddleLeft:    " ",
			MiddleRight:   " ",
			BottomLeft:    " ",
			BottomCenter:  " ",
			BottomRight:   " ",
			Gap:           " ",
			SpaceHeader:   0,
			SpaceFooter:   0,
			HeaderText:    "",
			TextNormal:    "[F-WHITE][B-BLACK]",
			TextHighLight: "[B-WHITE][F-BLACK]",
		}
	case "simple":
		return Style{
			TopLeft: "[F-WHITE][B-BLACK]┌", TopCenter: "─", TopRight: "┐",
			MiddleLeft:  "[F-WHITE][B-BLACK]│",
			MiddleRight: "[F-WHITE][B-BLACK]│",
			BottomLeft:  "[F-WHITE][B-BLACK]└", BottomCenter: "─", BottomRight: "[F-WHITE][B-BLACK]┘",
			Gap:           "    ",
			SpaceHeader:   0,
			SpaceFooter:   0,
			HeaderText:    "",
			TextNormal:    "[F-WHITE][B-BLACK]",
			TextHighLight: "[B-WHITE][F-BLACK]",
			Legend:        "[B-BLACK][F-GREEN]",
		}
	case "blue":
		return Style{
			TopLeft: "[B-BLUE][F-WHITE]╔", TopCenter: "═", TopRight: "╗",
			MiddleLeft:  "[B-BLUE][F-WHITE]║",
			MiddleRight: "[B-BLUE][F-WHITE]║",
			BottomLeft:  "[B-BLUE][F-WHITE]╚", BottomCenter: "═", BottomRight: "[B-BLUE][F-WHITE]╝",

			Gap:           "  ",
			SpaceHeader:   0,
			SpaceFooter:   0,
			HeaderText:    "",
			TextNormal:    "[B-BLUE][F-WHITE]",
			TextHighLight: "[B-RED][F-WHITE]",
			Legend:        "[B-BLUE][F-GREEN]",
		}
	default:
		return Style{
			TopLeft: "[F-WHITE][B-BLACK]┌", TopCenter: "─", TopRight: "┐",
			MiddleLeft:  "[F-WHITE][B-BLACK]│",
			MiddleRight: "[F-WHITE][B-BLACK]│",
			BottomLeft:  "[F-WHITE][B-BLACK]└", BottomCenter: "─", BottomRight: "[F-WHITE][B-BLACK]┘",

			Gap:           "  ",
			SpaceHeader:   2,
			SpaceFooter:   2,
			HeaderText:    "",
			TextNormal:    "[F-WHITE][B-BLACK]",
			TextHighLight: "[B-WHITE][F-BLACK]",
			Legend:        "[B-BLACK][F-GREEN]",
		}
	}
}

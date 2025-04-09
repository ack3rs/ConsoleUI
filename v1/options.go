package v1

const (
	Relative = "relative"
	Absolute = "absolute"
)

type Option struct {
	Title    string
	Style    Style
	Xpos     int
	Ypos     int
	Position string
}

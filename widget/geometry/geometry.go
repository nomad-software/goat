package geometry

var (
	// Side specifies the side to place the widget inside its parent.
	Side = side{
		Left:   "left",
		Right:  "right",
		Top:    "top",
		Bottom: "bottom",
	}

	// Fill specifies the space to fill inside its parent.
	Fill = fill{
		None:       "none",
		Horizontal: "x",
		Vertical:   "y",
		Both:       "both",
	}

	// Anchor specifies the anchor position of the widget inside its parent.
	Anchor = anchor{
		North:     "n",
		NorthEast: "ne",
		East:      "e",
		SouthEast: "se",
		South:     "s",
		SouthWest: "sw",
		West:      "w",
		NorthWest: "nw",
		Center:    "center",
	}

	// BorderMode specifies interaction modes for parent borders.
	BorderMode = borderMode{
		Inside:  "inside",
		Outside: "outside",
		Ignore:  "ignore",
	}
)

type side struct {
	Left   string
	Right  string
	Top    string
	Bottom string
}

type fill struct {
	None       string
	Horizontal string
	Vertical   string
	Both       string
}

type borderMode struct {
	Inside  string
	Outside string
	Ignore  string
}

type anchor struct {
	North     string
	NorthEast string
	East      string
	SouthEast string
	South     string
	SouthWest string
	West      string
	NorthWest string
	Center    string
}

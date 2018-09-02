package fullcoverage

type Shape struct {
	Sides int
	Size  Size
	Color Color
}

type Size int

const (
	SizeLarge  Size = 100
	SizeMedium      = 10
	SizeSmall  Size = 1
)

type Color int

const (
	ColorYellow Color = iota + 1
	ColorGreen
	ColorBlue
	ColorPurple
	ColorRed
	ColorOrange
	ColorBrown
)

func BucketName(s Shape) string {
	if s.Color == ColorBlue {
		return "Blue bucket"
	}
	if s.Size == SizeSmall {
		return "Bucket o' small"
	}
	if s.Size == SizeMedium && s.Sides == 4 {
		return "Medium quadrangles"
	}
	if s.Color == ColorYellow {
		return "The yellow yins"
	}
	return ""
}

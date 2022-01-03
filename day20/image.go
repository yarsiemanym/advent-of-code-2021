package day20

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
	"github.com/yarsiemanym/advent-of-code-2021/vt100"
)

type Image struct {
	plane *common.BoundedPlane
}

func NewImage(pixels [][]uint64) *Image {
	plane := common.NewBoundedPlane(len(pixels), len(pixels[0]))

	for row := range pixels {
		for col := range pixels[row] {
			plane.SetValueAt(common.New2DPoint(col, row), pixels[row][col])
		}
	}

	return &Image{
		plane: plane,
	}
}

func NewImageFromString(text string) *Image {
	text = strings.Trim(text, " \n")
	lines := common.Split(text, "\n")
	pixels := make([][]uint64, len(lines))

	for row, line := range lines {
		pixels[row] = make([]uint64, len(line))

		for col, character := range line {
			switch character {
			case '#':
				pixels[row][col] = 1
			case '.':
				pixels[row][col] = 0
			default:
				log.Fatalf("'%c' is not a valid pixel value.", character)
			}
		}
	}

	return NewImage(pixels)
}

func (image *Image) Height() int {
	return image.plane.Span().End().Y() + 1
}

func (image *Image) Width() int {
	return image.plane.Span().End().X() + 1
}

func (image *Image) GetPixelAt(point *common.Point, fill uint64) uint64 {
	if point.X() < 0 || point.X() >= image.Width() || point.Y() < 0 || point.Y() >= image.Height() {
		return fill
	}

	value := image.plane.GetValueAt(point)

	if value == nil {
		return 0
	}

	return value.(uint64)
}

func (image *Image) CountIlluminatedPixels() uint64 {
	illuminatedPixels := uint64(0)

	for _, point := range image.plane.GetAllPoints() {
		illuminatedPixels += image.GetPixelAt(point, 0)
	}

	return illuminatedPixels
}

func (image *Image) String() string {
	output := ""

	for y := 0; y < image.Height(); y++ {
		for x := 0; x < image.Width(); x++ {
			point := common.New2DPoint(x, y)
			pixel := image.GetPixelAt(point, 0)
			switch pixel {
			case 1:
				output += vt100.Sprint(" ", vt100.YellowBackgroundAttribute)
			case 0:
				output += vt100.Sprint(" ")
			default:
				log.Fatalf("%d is not a valid pixel value.", pixel)
			}
		}

		output += "\n"
	}

	return output
}

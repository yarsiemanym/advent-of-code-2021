package day20

import (
	"strings"

	"github.com/Workiva/go-datastructures/bitarray"
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type ImageEnhancer struct {
	algorithm bitarray.BitArray
	fill      uint64
}

func NewImageEnhancer(text string) *ImageEnhancer {
	text = strings.Trim(text, " \n")

	if len(text) != 512 {
		log.Fatal("Image enhancement algorithm must be 512 bits long.")
	}

	bits := bitarray.NewBitArray(uint64(len(text)))

	for index, character := range text {
		switch character {
		case '#':
			bits.SetBit(uint64(index))
		case '.':
			bits.ClearBit(uint64(index))
		default:
			log.Fatalf("'%c' is not a valid character.", character)
		}
	}

	return &ImageEnhancer{
		algorithm: bits,
		fill:      0,
	}
}

func (enhancer *ImageEnhancer) GetValueAt(index uint64) uint64 {
	bit, err := enhancer.algorithm.GetBit(index)
	common.Check(err)

	if bit {
		return 1
	} else {
		return 0
	}
}

func (enhancer *ImageEnhancer) EnhanceImage(originalImage *Image) *Image {
	growthFactor := common.New2DPoint(1, 1)
	newPixels := make([][]uint64, originalImage.Height()+(2*growthFactor.Y()))

	for y := 0; y < len(newPixels); y++ {
		newPixels[y] = make([]uint64, originalImage.Width()+(2*growthFactor.X()))

		for x := 0; x < len(newPixels[y]); x++ {
			newImagePoint := common.New2DPoint(x, y)
			originalImagePoint := newImagePoint.Subtract(growthFactor)
			newPixels[y][x] = enhancer.EnhancePixel(originalImage, originalImagePoint)
		}
	}

	bit0, err := enhancer.algorithm.GetBit(0)
	common.Check(err)

	if bit0 {
		if enhancer.fill == 0 {
			enhancer.fill = 1
		} else {
			enhancer.fill = 0
		}
	}

	enhancedImage := NewImage(newPixels)
	return enhancedImage
}

func (enhancer *ImageEnhancer) EnhancePixel(image *Image, point *common.Point) uint64 {
	neighbors := point.GetMooreNeighborhood()
	index := uint64(0)

	for _, neighbor := range neighbors {
		pixelValue := image.GetPixelAt(neighbor, enhancer.fill)
		index = (index << 1) | pixelValue
	}

	bit, err := enhancer.algorithm.GetBit(index)
	common.Check(err)

	if bit {
		return uint64(1)
	} else {
		return uint64(0)
	}
}

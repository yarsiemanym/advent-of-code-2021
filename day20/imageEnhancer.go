package day20

import (
	"strings"

	"github.com/Workiva/go-datastructures/bitarray"
	log "github.com/sirupsen/logrus"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type ImageEnhancer struct {
	algorithm bitarray.BitArray
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

func (enhancer *ImageEnhancer) Enhance(orginalImage *Image) *Image {
	growthFactor := common.New2DPoint(2, 2)
	newPixels := make([][]uint64, orginalImage.Height()+(2*growthFactor.Y()))

	for y := 0; y < len(newPixels); y++ {
		newPixels[y] = make([]uint64, orginalImage.Width()+(2*growthFactor.X()))

		for x := 0; x < len(newPixels[y]); x++ {
			newImagePoint := common.New2DPoint(x, y)
			originalImagePoint := newImagePoint.Subtract(growthFactor)
			originalImagePointNeighbors := originalImagePoint.GetMooreNeighborhood()
			index := uint64(0)

			for _, originalImagePointNeighbor := range originalImagePointNeighbors {
				index = (index << 1) | orginalImage.GetPixelAt(originalImagePointNeighbor)
			}

			bit, err := enhancer.algorithm.GetBit(index)
			common.Check(err)

			if bit {
				newPixels[y][x] = 1
			} else {
				newPixels[y][x] = 0
			}
		}
	}

	return NewImage(newPixels)
}

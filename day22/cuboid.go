package day22

import (
	"github.com/Workiva/go-datastructures/set"
	"github.com/yarsiemanym/advent-of-code-2021/common"
)

type Cuboid struct {
	span    *common.LineSegment
	onCubes *set.Set
}

func NewCuboid(start *common.Point, end *common.Point) *Cuboid {
	return &Cuboid{
		span:    common.NewLineSegment(start, end),
		onCubes: set.New(),
	}
}

func NewCuboidFromPoints(points []*common.Point) *Cuboid {
	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	minZ := 0
	maxZ := 0

	for _, point := range points {
		minX = common.MinInt(minX, point.X())
		maxX = common.MaxInt(maxX, point.X())
		minY = common.MinInt(minY, point.Y())
		maxY = common.MaxInt(maxY, point.Y())
		minZ = common.MinInt(minZ, point.Z())
		maxZ = common.MaxInt(maxZ, point.Z())
	}

	return NewCuboid(common.New3DPoint(minX, minY, minZ), common.New3DPoint(maxX, maxY, maxZ))
}

func (cuboid *Cuboid) GetCube(point *common.Point) bool {
	exists := cuboid.onCubes.Exists(*point)
	return exists
}

func (cuboid *Cuboid) SetCube(point *common.Point, on bool) {
	if on {
		cuboid.onCubes.Add(*point)
	} else {
		cuboid.onCubes.Remove(*point)
	}
}

func (cuboid *Cuboid) CountOnCubes() int {
	return int(cuboid.onCubes.Len())
}

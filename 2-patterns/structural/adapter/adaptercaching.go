package adapter

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

// When working with adapters, it may be wise to watch out so we don't create too many temporary objects

var pointCache = map[[16]byte][]Point{}

func (a *vectorToRasterAdapter) addLineCached(line Line) {
	// Generate the hash for a point
	hash := func(obj interface{}) [16]byte {
		bytes, _ := json.Marshal(obj)
		return md5.Sum(bytes)
	}

	h := hash(line)
	if pts, ok := pointCache[h]; ok {
		// Instead of calculating these points again, we simply add them to the adapter
		a.points = append(a.points, pts...)

		return
	}

	left, right := minmax(line.X1, line.X2)
	top, bottom := minmax(line.Y1, line.Y2)
	dx := right - left
	dy := line.Y2 - line.Y1

	if dx == 0 {
		for y := top; y <= bottom; y++ {
			a.points = append(a.points, Point{left, y})
		}
	} else if dy == 0 {
		for x := left; x <= right; x++ {
			a.points = append(a.points, Point{x, top})
		}
	}

	pointCache[h] = a.points
	fmt.Println("we have", len(a.points), "points")
}

func VectorToRasterCached(vi *VectorImage) RasterImage {
	adapter := vectorToRasterAdapter{}
	for _, line := range vi.Lines {
		adapter.addLineCached(line)
	}

	return adapter
}

func AdapterCaching() {
	rc2 := NewRectangle(6, 4)
	cachedAdapter := VectorToRasterCached(rc2)
	_ = VectorToRasterCached(rc2) // unnecessary 2nd adapter call
	fmt.Println(DrawPoints(cachedAdapter))
}

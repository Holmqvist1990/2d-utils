package utils

/*
	1D index (sprite) to tileset rectangle.
	Assumes a square tileset.

		sprite = 3, tileSize = N, tilesetWidth = 4
		. . x x .
		. . x x .
		. . . . .
		. . . . .
*/
func TilesetRect[T Integer](sprite, tileSize, tilesetWidth T) (x1, y1, x2, y2 T) {
	x1 = (sprite % tilesetWidth) * tileSize
	y1 = (sprite / tilesetWidth) * tileSize
	x2 = x1 + tileSize
	y2 = y2 + tileSize
	return x1, y1, x2, y2
}

// 1D to 2D coordinate conversion.
func IndexToXY[T Integer](idx, mapWidth, mapHeight T) (x, y T) {
	x = idx % mapWidth
	y = idx / mapHeight
	return x, y
}

// 2D to 1D coordinate conversion.
func XYToIndex[T Integer](x, y, mapHeight T) T {
	return x + y*mapHeight
}

// Removes under/overflow.
func ClampXYToMap[T Integer](x, y, mapWidth, mapHeight T) (T, T) {
	return Max(0, Min(x, mapWidth-1)), Max(0, Min(y, mapHeight-1))
}

// Removes under/overflow.
func ClampIndexToMap[T Integer](index, mapWidth, mapHeight T) T {
	return Clamp(index, 0, (mapWidth*mapHeight)-1)
}

func IsXYInRange[T Integer](x, y, mapWidth, mapHeight T) bool {
	return x >= 0 && x < mapWidth && y >= 0 && y < mapHeight
}

func IsIndexInRange[T Integer](index, mapWidth, mapHeight T) bool {
	return index >= 0 && index < mapWidth*mapHeight
}

// Returns range through min/max for XY.
func XYRange[T Number](x1, y1, x2, y2 T) (T, T, T, T) {
	return Min(x1, x2), Min(y1, y2), Max(x1, x2), Max(y1, y2)
}

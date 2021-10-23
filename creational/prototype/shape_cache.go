// @file: shape_cache.go
// @date: 2021/10/28

package main

type ShapeCache struct {
	shapeMap map[string]Shape
}

func NewShapeCache() *ShapeCache {
	return &ShapeCache{shapeMap: make(map[string]Shape)}
}

func (c ShapeCache) getShape(shapeId string) Shape {
	cachedShape, ok := c.shapeMap[shapeId]
	if !ok {
		return nil
	}
	return cachedShape.clone()
}

func (c *ShapeCache) loadCache() {
	circle := NewCircle()
	circle.setId("1")
	c.shapeMap[circle.getId()] = circle

	rectangle := NewRectangle()
	rectangle.setId("2")
	c.shapeMap[rectangle.getId()] = rectangle

	square := NewSquare()
	square.setId("3")
	c.shapeMap[square.getId()] = square
}

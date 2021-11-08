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
	return cachedShape.Clone()
}

func (c *ShapeCache) loadCache() {
	circle := NewCircle()
	circle.SetId("1")
	c.shapeMap[circle.GetId()] = circle

	rectangle := NewRectangle()
	rectangle.SetId("2")
	c.shapeMap[rectangle.GetId()] = rectangle

	square := NewSquare()
	square.SetId("3")
	c.shapeMap[square.GetId()] = square
}

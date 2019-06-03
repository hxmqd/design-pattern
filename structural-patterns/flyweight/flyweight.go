// 享元模式（Flyweight Pattern）主要用于减少创建对象的数量，以减少内存占用和提高性能。
// 这种类型的设计模式属于结构型模式，它提供了减少对象数量从而改善应用所需的对象结构的方式。
// 享元模式尝试重用现有的同类对象，如果未找到匹配的对象，则创建新对象。
//
// 创建了20个不同效果的圆，但相同颜色的圆只需要创建一次便可，相同颜色的只需要引用原有对象，改变其坐标值便可。
// 此种模式下，同一颜色的圆虽然位置不同，但其地址都是同一个，所以说此模式适用于注重单一结果的情况。

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Shape interface {
	draw()
}

type Circle struct {
	color  string
	x      int
	y      int
	radius int
}

func (c *Circle) SetX(x int) {
	c.x = x
}

func (c *Circle) SetY(y int) {
	c.y = y
}

func (c *Circle) setRadius(radius int) {
	c.radius = radius
}

func (c *Circle) draw() {
	fmt.Println("Circle: Draw() [Color:" + c.color + ", x:" + strconv.Itoa(c.x) + ", y:" +
		strconv.Itoa(c.y) + ", radius:" + strconv.Itoa(c.radius))
}

type ShapeFactory struct {
	circleMap map[string]Shape
}

func (s *ShapeFactory) GetCircle(color string) (circle *Circle) {
	value, ok := s.circleMap[color]
	if ok {
		circle, _ = value.(*Circle)
	} else {
		circle = &Circle{
			color: color,
		}
		s.circleMap[color] = circle
		fmt.Println("Creating circle of color : " + color)
	}
	return circle
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]Shape),
	}
}

func main() {
	colors := [...]string{"Red", "Green", "Blue", "White", "Black"}
	shapeFactory := NewShapeFactory()
	for i := 0; i < 20; i++ {
		circle := shapeFactory.GetCircle(colors[rand.Intn(5)])
		circle.SetX(rand.Intn(100))
		circle.SetY(rand.Intn(100))
		circle.setRadius(100)
		circle.draw()
	}
}

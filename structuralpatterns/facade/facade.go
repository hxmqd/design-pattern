// 外观系统（Facade Pattern）隐藏系统的复杂性，并向客户端提供了一个可以访问系统的接口，来隐藏系统的复杂性。
// 1、为复杂的模块或子系统提供外界访问的模块。 2、子系统相对独立。 3、预防低水平人员带来的风险
package main

import "fmt"

type Shape interface {
	draw()
}

type Rectangle struct {
}

func (r *Rectangle)draw(){
	fmt.Println("Rectangle::draw()")
}

type Square struct {
}

func (s *Square)draw(){
	fmt.Println("Square::draw()")
}

type Circle struct {
}

func (c *Circle)draw(){
	fmt.Println("Circle::draw()")
}

// 外观类
type ShapeMaker struct {
	circle Shape
	rectangle Shape
	square Shape
}

func NewShapeMaker() *ShapeMaker {
	sm := &ShapeMaker{}
	sm.circle = &Circle{}
	sm.rectangle = &Rectangle{}
	sm.square = &Square{}
	return sm
}

func (s *ShapeMaker) drawCircle(){
	s.circle.draw()
}

func (s *ShapeMaker) drawRectangle(){
	s.rectangle.draw()
}

func (s *ShapeMaker) drawSquare(){
	s.square.draw()
}

func main(){
	shapeMaker := NewShapeMaker()
	shapeMaker.drawCircle()
	shapeMaker.drawRectangle()
	shapeMaker.drawSquare()
}


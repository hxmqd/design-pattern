// 装饰器模式允许向一个现有的对象添加新的功能，同时又不改变其结构。
// 这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。
// 意图：动态地给一个对象添加一些额外的职责。就增加功能来说，装饰器模式相比生成子类更为灵活。
// 主要解决：一般的，我们为了扩展一个类经常使用继承方式实现，由于继承为类引入静态特征，并且随着扩展功能的增多，子类会很膨胀。
// 何时使用：在不想增加很多子类的情况下扩展类。
package main

import "fmt"

type Shape interface {
	draw()
}

type Rectangle struct {
}

func (r *Rectangle) draw(){
	fmt.Println("Shape: Rectangle")
}

type Circle struct {
}

func (r *Circle) draw(){
	fmt.Println("Shape: Circle")
}

// 抽象装饰类
type ShapeDecorator struct {
	decoratedShape Shape
}

func (s *ShapeDecorator) draw(){
	s.decoratedShape.draw()
}

// 实体装饰类
type RedShapeDecorator struct {
	ShapeDecorator
}

func (s *RedShapeDecorator) draw(){
	s.ShapeDecorator.draw()
	fmt.Println("Border Color: Red")
}

func NewRedShapeDecorator(decoratedShape Shape) *RedShapeDecorator {
	return &RedShapeDecorator{ShapeDecorator{decoratedShape}}
}

func main(){
	var circle Shape = &Circle{}

	var redCircle Shape = NewRedShapeDecorator(new(Circle))

	var redRectangle Shape = NewRedShapeDecorator(new (Rectangle))

	fmt.Println("circle with normal border")
	circle.draw()

	fmt.Printf("\n")
	fmt.Println("circle of red border")
	redCircle.draw()

	fmt.Printf("\n")
	fmt.Println("rectangle of red border")
	redRectangle.draw()
}









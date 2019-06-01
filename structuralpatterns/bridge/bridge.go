// 桥接模式(Bridge)是用于把抽象化与实现化解耦，使得两者可以独立变化
//
// 如果软件系统中某个类存在两个独立变化的维度，通过该模式可以将这两个维度分离出来，
// 使两者可以独立扩展，让系统更加符合单一职责原则。桥接模式主要使用抽象关联取代传统的多重继承，
// 将类之间的静态继承关系转换为动态的对象组合关系，使得系统更加灵活，并易于扩展，同时有效控制了系统中类的个数
//
// 桥接模式的作用
// 1.分离接口及其实现部分 2.提高可扩展性 3.实现细节对使用者透明
//
// 桥接模式很好的利用了合成-聚合复用原则，其实我们在设计软件结构的时候应该优先使用这一原则，而不是类继承
// 类继承一定要在‘is-a’的关系时考虑使用，而不是任何时候去使用
//
// 在毛笔中，颜色和型号实现了分离，增加新的颜色或者型号都对另外一个维度没有任何影响
package main

import "fmt"

// 颜色
type Color interface {
	Use()
}

type Red struct{}

func (r Red) Use() {
	fmt.Println("Use Red color")
}

type Green struct{}

func (g Green) Use() {
	fmt.Println("Use Green color")
}

type Yellow struct{}

func (y Yellow) Use() {
	fmt.Println("Use Yellow color")
}

// 毛笔 定义抽象
type BrushPen interface {
	DrawPicture()
}

// 组合关系，实现桥接
type BigBrushPen struct {
	Color
}

// 实现
func (b BigBrushPen) DrawPicture() {
	fmt.Println("Draw picture with big brush pen")
	b.Use()
}

// 组合关系，实现桥接
type SmallBrushPen struct {
	Color
}

// 实现
func (s SmallBrushPen) DrawPicture() {
	fmt.Println("Draw picture with small brush pen")
	s.Use()
}

// 工厂方法:生产具体的BrushPen
func NewBrushPen(t string, color Color) BrushPen {
	switch t {
	case "BIG":
		return BigBrushPen{
			Color: color,
		}
	case "SMALL":
		return SmallBrushPen{
			Color: color,
		}
	default:
		return nil
	}
}

func main() {
	var tColor Color
	tColor = Red{}
	tBrushPen := NewBrushPen("BIG", tColor)
	tBrushPen.DrawPicture()
	fmt.Println("=========================")
	tColor = Green{}
	tBrushPen = NewBrushPen("SMALL", tColor)
	tBrushPen.DrawPicture()
	fmt.Println("=========================")
	tColor = Yellow{}
	tBrushPen = NewBrushPen("BIG", tColor)
	tBrushPen.DrawPicture()
}





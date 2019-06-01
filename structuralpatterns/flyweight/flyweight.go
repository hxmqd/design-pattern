package flyweight

import (
	"fmt"
	"strconv"
)

type Shape interface {
	draw()
}

type Circle struct {
	x      int
	y      int
	radius int
}

func (c *Circle) SetX(x int) {
    c.x = x
}

func (c *Circle)SetY(y int){
	c.y = y
}

func (c *Circle)setRadius(radius int){
	c.radius = radius
}

func (c *Circle)draw(){
0








                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              	fmt.Println("Circle: Draw(), c.x :"+ strconv.Itoa(c.x) +", c.y :" + strconv.Itoa(c.y) +", radius :" + strconv.Itoa(c.radius))
}

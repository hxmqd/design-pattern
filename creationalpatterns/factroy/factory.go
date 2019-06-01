// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/14.
//
// 工厂模式属于创建型模式，它提供了一种创建对象的最佳方式
package factroy

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (s *Rectangle) Draw(){
	fmt.Println("inside Rectangle::draw() method.")
}

type Square struct {
}

func (s *Square) Draw(){
	fmt.Println("inside Square::draw() method.")
}

type Circle struct {
}

func (s *Circle) Draw(){
	fmt.Println("inside Circle::draw() method.")
}

type ShapeFactory struct {
}

func NewShapeFactory() *ShapeFactory{
	return &ShapeFactory{}
}

func (s *ShapeFactory)GetShape(shapeType string) Shape {
	switch shapeType {
	case "rectangle":
		return new(Rectangle)
	case "square":
		return new(Square)
	case "circle":
		return new(Circle)
	default:
		return nil
	}
	return nil
}
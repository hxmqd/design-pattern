// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/14.
//
// 抽象工厂模式是围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂。该模式提供了一个接口，用于创建相关或依赖对象的家族，而不需要明确指定具体类。
// 在抽象工厂模式中，接口是负责创建一个相关对象的工厂，不需要显式指定它们的类。每个生成的工厂都能按照工厂模式提供对象。

package abstractfactory

import "fmt"

// 形状接口
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

// 颜色接口
type Color interface {
	Fill()
}

type Green struct {
}

func (g *Green)Fill(){
	fmt.Println("Inside Green::fill() method.")
}

type Red struct {
}

func (r *Red)Fill(){
	fmt.Println("Inside Red::fill() method.")
}

type IFactory interface {
	GetShape(shape string) Shape
	GetColor(color string) Color
}

type ShapeFactory struct {
}

func (s *ShapeFactory)GetShape(shape string) Shape{
	switch shape {
	case "rectangle":
		return new(Rectangle)
	case "square":
		return new(Square)
	default:
		return nil
	}
	return nil
}

func (s *ShapeFactory)GetColor(color string) Color{
	return nil
}

type ColorFactory struct {
}

func (s *ColorFactory)GetColor(color string) Color{
	switch color {
	case "red":
		return new(Red)
	case "green":
		return new(Green)
	default:
		return nil
	}
	return nil
}

func (s *ColorFactory)GetShape(shape string) Shape{
	return nil
}

// 工厂创建类
type FactoryProducer struct {
}

func NewFactoryProducer() *FactoryProducer{
	return &FactoryProducer{}
}

func (f * FactoryProducer) GetFactory(factory string) IFactory{
	switch factory {
	case "shape":
		return new(ShapeFactory)
	case "color":
		return new(ColorFactory)
	}
	return nil
}

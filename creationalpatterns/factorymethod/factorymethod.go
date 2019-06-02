// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/14.
//
// 工厂方法模式定义了一个创建实例的接口，但由子类决定要实例化的类是哪一个。工厂方法让类把实例化推迟到子类

package factorymethod

type IOperation interface {
	SetA(float64)
	SetB(float64)
	GetResult() float64
}

type Operation struct {
	a float64
	b float64
}

func (op *Operation) SetA(a float64) {
	op.a = a
}

func (op *Operation) SetB(b float64) {
	op.b = b
}

// addition
type AddOperation struct {
	Operation
}

func (a *AddOperation) GetResult() float64 {
	return a.a + a.b
}

// subtraction
type SubOperation struct {
	Operation
}

func (s *SubOperation) GetResult() float64 {
	return s.a - s.b
}

type IFactory interface {
	CreateOperation() Operation
}

type AddFactory struct {
}

func (a *AddFactory) CreateOperation() IOperation {
	return &AddOperation{}
}

type SubFactory struct {
}

func (s *SubFactory) CreateOperation() IOperation {
	return &SubOperation{}
}




















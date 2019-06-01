// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/14.
//
//建造者模式（Builder Pattern）使用多个简单的对象一步一步构建成一个复杂的对象
package builder

import "fmt"

// 食物包装
type Packing interface {
	Pack() string
}

type Wrapper struct {
}

func (w *Wrapper) Pack() string{
	return "Wrapper"
}

type Bottle struct {
}

func (w *Bottle) Pack() string{
	return "Bottle"
}

// 食物
type Item interface {
	Name() string
	Packing() Packing
	Price() float64
}

//汉堡
type Burger struct {
}

func (b *Burger)Packing() Packing{
	return new (Wrapper)
}

// 冷饮
type ColorDrink struct {
}

func (c *ColorDrink)Packing() Packing{
	return new(Bottle)
}

// 蔬菜汉堡
type VegBurger struct {
	Burger
}

func (v *VegBurger)Name() string{
	return "vegBurger"
}

func (v *VegBurger)Price() float64{
	return 18.5
}

// 鸡肉汉堡
type ChickenBurger struct {
	Burger
}

func (v *ChickenBurger)Name() string{
	return "ChickenBurger"
}

func (v *ChickenBurger)Price() float64{
	return 22.5
}

// 可口可乐
type CocaCola struct {
	ColorDrink
}

func (c *CocaCola) Name() string{
	return "CocaCola"
}

func (c *CocaCola) Price() float64{
	return 3.0
}

// 冰激凌
type IceCream struct {
	ColorDrink
}

func (c *IceCream) Name() string{
	return "iceCream"
}

func (c *IceCream) Price() float64{
	return 8.0
}

// 快餐
type Meal struct {
	items []Item
}

func (m *Meal) addItem(item Item){
	m.items = append(m.items, item)
}

func (m *Meal) getCost() float64{
	cost := 0.0
	for _ , item := range m.items{
		cost += item.Price()
	}
	return cost
}

func (m *Meal) showItems(){
	for _, item := range m.items{
		fmt.Print("item:"+item.Name())
		fmt.Print(",Packing:"+item.Packing().Pack())
		fmt.Printf(",Price:%f\n",item.Price())
	}
}

// 快餐生产者
type MealBuilder struct {
}

func NewMealBuilder() *MealBuilder{
	return &MealBuilder{}
}

func (m *MealBuilder) prepareVegMeal(){
	var meal Meal
	meal.addItem(&VegBurger{})
	meal.addItem(&CocaCola{})
	meal.showItems()
	fmt.Printf("cost:%f\n",meal.getCost())
}

func (m *MealBuilder) prepareMeatMeal(){
	var meal Meal
	meal.addItem(&ChickenBurger{})
	meal.addItem(&IceCream{})
	meal.showItems()
	fmt.Printf("cost:%f\n",meal.getCost())
}
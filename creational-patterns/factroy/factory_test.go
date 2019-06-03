// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/14.

package factroy

import "testing"

func TestFactory(t *testing.T){

	shape1 := NewShapeFactory().GetShape("rectangle")
	shape1.Draw()

	shape2 :=NewShapeFactory(). GetShape("square")
	shape2.Draw()

	shape3 := NewShapeFactory().GetShape("circle")
	shape3.Draw()
}

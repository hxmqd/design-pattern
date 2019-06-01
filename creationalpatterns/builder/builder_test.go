// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/15.

package builder

import "testing"

func TestBuilder(T *testing.T){
	NewMealBuilder().prepareMeatMeal()
	NewMealBuilder().prepareVegMeal()
}
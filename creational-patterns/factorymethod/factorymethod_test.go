// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/14.

package factorymethod

import (
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T ){
	fac1 := new(AddFactory)
	oper1 := fac1.CreateOperation()
	oper1.SetA(1.0)
	oper1.SetB(2.0)
	fmt.Printf("%f + %f = %f \n",1.0, 2.0, oper1.GetResult())

	fac2 := new(SubFactory)
	oper2 := fac2.CreateOperation()
	oper2.SetA(1.0)
	oper2.SetB(2.0)
	fmt.Printf("%f - %f = %f \n",1.0, 2.0, oper2.GetResult())

}
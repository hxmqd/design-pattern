// Copyright 2019 didi. All rights reserved.
// Created by hexiaomin on 2019/4/14.

package abstractfactory

import "testing"

func TestAbstractFactory(t *testing.T ){
	NewFactoryProducer().GetFactory("color").GetColor("red").Fill()
	NewFactoryProducer().GetFactory("shape").GetShape("square").Draw()
}

/*
抽象工厂：工厂方法中再嵌套多次工厂方法
1. 一工厂要生产 摩托车 和 汽车 两种不同类型的交通工具
2. 汽车 可以分为两种：5门家庭用车、4门小轿车
3. 摩托车 也分为两种：2坐的巡逻摩托车、1坐的越野摩托车

首先我们需要 指令来区分生产 汽车还是摩托车 这是一个工厂方法
其次我们继续还需要 指令来生产 汽车或者摩托车 中的某一种具体类型 这又是一个工厂方法
*/
package abstract_factory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	GetVehicle(v int) (Vehicle, error)
}

const (
	CarFactoryType       = 1
	MotorbikeFactoryType = 2
)

func GetVehicleFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized\n", f))
	}
}

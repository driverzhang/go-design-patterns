package builder

// 抽象类定义
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// Director
type ManufacturingDirector struct {
	builder BuildProcess
}

// Director 负责安排已有模块的顺序，然后告诉Builder开始建造，
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

//  ConcreteBuilder 实现抽象类定义(BuildProcess)的所有方法，并且返回一个组建好的对象。
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Product 实现了模板方法模式，也就是有模板方法（Construct）和基本方法（BuildProcess内的四个方法）
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// A Builder of type car; Builder 规范产品的组建，一般是由子类实现
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// A Builder of type motorbike
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

// A Builder of type motorbike
type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}

func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}

func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}

func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}



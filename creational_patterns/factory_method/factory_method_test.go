package factory_method

import "testing"

func TestGet(t *testing.T) {
	factory := Get("logs")
	rsp := factory.GetData("123")
	if rsp != "this is logs id:123" {
		t.Error("error logs")
	}

	sim := Get("simBase")
	rspSim := sim.GetData("321")
	if rspSim != "this is simbase ids:321" {
		t.Error("error simbase")
	}
}

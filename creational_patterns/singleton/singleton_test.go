package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	singleton := GetInstance()
	t.Log(singleton)
	count := singleton.AddOne()
	newCount:=  singleton.GetCount()
	t.Logf("AddOne After is %d", count)
	t.Logf("AddOne After singleton.count %d", newCount)
}

func TestGetInstance2(t *testing.T) {
	counter1 := GetInstance()
	if counter1 == nil {
		//Test 1 failed
		t.Error("A new connection object must have been made")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}

func TestGetInstanceOnce(t *testing.T) {
	counter1 := GetInstanceOnce()
	if counter1 == nil {
		//Test 1 failed
		t.Error("A new connection object must have been made")
	}
	expectedCounter := counter1

	currentCount := counter1.AddOne()
	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstanceOnce()
	if counter2 != expectedCounter {
		//Test 2 failed
		t.Error("Singleton instances must be different")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}
}
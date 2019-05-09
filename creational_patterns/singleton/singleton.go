package singleton

import "sync"

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		//instance = new(singleton)
		instance = &singleton{}
	}

	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func (s *singleton) GetCount() int {
	return s.count
}

var once sync.Once
// go 特有包单例模式高并发终极版推荐
// 支持高并发单例模式
func GetInstanceOnce() *singleton  {
	once.Do(func() {
		if instance == nil {
			instance =  new(singleton)
		}
	})
	return instance
}
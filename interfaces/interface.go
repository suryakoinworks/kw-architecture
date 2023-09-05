package interfaces

import (
	"sort"
	"time"
)

type (
	Interface interface {
		Run()
		IsBackground() bool
		Priority() int
	}

	factory struct {
		interfaces []Interface
	}
)

func NewInterfaceFactory(interfaces ...Interface) *factory {
	return &factory{interfaces: interfaces}
}

func (f *factory) Run() {
	sort.Slice(f.interfaces, func(i int, j int) bool {
		return f.interfaces[i].Priority() > f.interfaces[j].Priority()
	})

	for _, application := range f.interfaces {
		if application.IsBackground() {
			go application.Run()
		} else {
			time.Sleep(100 * time.Millisecond)
			application.Run()
		}
	}
}

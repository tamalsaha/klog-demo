package main

import (
	"flag"

	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
)

type myError struct {
	str string
}

func (e myError) Error() string {
	return e.str
}

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "3")
	flag.Parse()
	logger := klogr.New().WithName("MyName")
	l1 := logger.WithValues("user", "you")
	l1.Info("hello", "val1", 1, "val2", map[string]int{"k": 1})
	l1.V(3).Info("nice to meet you")

	l2 := logger.WithValues("user", "they")
	l2.Info("hello", "val2", 1, "val2", map[string]int{"k": 1})
	l2.V(3).Info("nice to meet you")
	l2.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	l2.Error(myError{"an error occurred"}, "goodbye", "code", -1)
	l2.V(1).Error(myError{"an error occurred_v(1)"}, "goodbye", "code", -1)

	l1.Error(nil, "uh oh", "trouble", true, "reasons", []float64{0.1, 0.11, 3.14})
	l1.Error(myError{"an error occurred"}, "goodbye", "code", -1)
	l1.V(1).Error(myError{"an error occurred_v(1)"}, "goodbye", "code", -1)

	klog.Flush()
}

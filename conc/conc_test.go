package conc_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/srcio/go/conc"
)

func TestRun(t *testing.T) {
	conc.Run(func() {
		fmt.Println("Hello")
	}, func() {
		fmt.Println("World")
	})
	time.Sleep(1 * time.Second)
}

func TestWait(t *testing.T) {
	conc.Wait(func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Hello")
	}, func() {
		time.Sleep(4 * time.Second)
		fmt.Println("World")
	})
}

func TestWaitAndRecover(t *testing.T) {
	if p := conc.WaitAndRecover(func() {
		fmt.Println("Hello")
		panic("panic when print hello")
	}, func() {
		fmt.Println("World")
		time.Sleep(1 * time.Second)
		panic("panic when print world")
	}); p != nil {
		fmt.Printf("p.PanicInfo(): %v\n", p.PanicInfo())
		fmt.Printf("p.PanicVal(): %v\n", p.PanicVal())
	}
}

func TestForIndex(t *testing.T) {
	conc.ForIndex(10, func(i int) {
		fmt.Println(i)
	})
}

func TestForEach(t *testing.T) {
	var names = []string{"Alice", "Bob", "Cindy", "David"}

	conc.ForEach(names, func(name string) {
		fmt.Printf("hello: %s\n", name)
	})
}

func TestForEachWitLimit(t *testing.T) {
	var names = []string{"Alice", "Bob", "Cindy", "David", "Eve", "Frank", "Grace", "Helen", "Irene", "Jack", "Karl", "Lily", "Marry", "Nancy", "Oscar", "Peter", "Queen", "Rose", "Sara", "Tom", "Uma", "Vivian", "Wendy", "Xavier", "Yvonne", "Zoe"}

	conc.ForEachWithLimit(names, func(name string) {
		time.Sleep(4 * time.Second)
		fmt.Printf("hello: %s\n", name)
	}, 5)
}

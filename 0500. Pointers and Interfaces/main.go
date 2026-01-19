package main

import "fmt"

type Counter interface {
	Inc(delta int)
	Dec(delta int)
	Value() int
}

type SimpleCounter struct {
	value int
}

func (c *SimpleCounter) Inc(delta int) {
	c.value += delta
}

func (c *SimpleCounter) Dec(delta int) {
	c.value -= delta
}

func (c *SimpleCounter) Value() int {
	return c.value
}

type LimitedCounter struct {
	value int
}

func (c *LimitedCounter) Inc(delta int) {
	c.value += delta
}

func (c *LimitedCounter) Dec(delta int) {
	c.value -= delta
	if c.value < 0 {
		c.value = 0
	}
}

func (c *LimitedCounter) Value() int {
	return c.value
}

func UseCounter(c Counter) {
	// пример действий с счётчиком
	c.Inc(5)
	c.Dec(2)
	c.Inc(10)
	c.Dec(20)
}

func main() {
	fmt.Print("\033[H\033[2J")

	// обычный счётчик
	simple := &SimpleCounter{}
	UseCounter(simple)
	fmt.Printf("SimpleCounter: %d\n", simple.Value())

	// счётчик с ограничением
	limited := &LimitedCounter{}
	UseCounter(limited)
	fmt.Printf("LimitedCounter: %d\n", limited.Value())
}

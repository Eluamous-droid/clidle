package main

import "math"

type Crab struct {
	name       string
	cost       int
	production int
	count int
}

const incrementAmount = 1.25

func (c *Crab) buyCrab() {
	c.cost = int(math.Ceil(float64(c.cost) * incrementAmount))
	c.count++
	increaseIncome(c.production)

}


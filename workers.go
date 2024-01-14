package main

type Crab struct {
	name       string
	cost       int
	production int
	count int
}

const incrementAmount = 5

func (c *Crab) buyCrab() {
	c.cost = int(c.cost + incrementAmount)
	c.count++
	increaseIncome(c.production)

}


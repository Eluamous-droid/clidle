package main 

type crab struct{
	name string
	cost int
	production int
}

const incrementAmount = 5 

func (c crab) buyCrab(){
	c.cost = int(c.cost + incrementAmount)	
	increaseIncome(c.production)
}

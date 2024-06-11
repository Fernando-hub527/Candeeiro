package consumption

type IConsumptionCost interface {
	getConsumptionCost(kw uint32) float64
}

type ConsumptionCost struct {
	costKw uint32
}

func FactoryConsumptionCost() IConsumptionCost {
	return &ConsumptionCost{costKw: 1}
}

func (c *ConsumptionCost) getConsumptionCost(kw uint32) float64 {
	return float64(kw * c.costKw)
}

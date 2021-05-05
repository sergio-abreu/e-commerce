package discount

func NewDiscount(valueInCents int, maxPercentage Percentage) Discount {
	return Discount{ValueInCents: valueInCents, originalValueInCents: valueInCents, maxPercentage: maxPercentage}
}

type Percentage float64

const Percent Percentage = 0.01

type Discount struct {
	Percentage           Percentage
	ValueInCents         int
	originalValueInCents int
	maxPercentage        Percentage
}

func (d *Discount) AddByPercentage(percentage Percentage) {
	d.Percentage += percentage
	if d.Percentage > d.maxPercentage {
		d.Percentage = d.maxPercentage
	}
	d.ValueInCents = int(float64(d.originalValueInCents) * float64(1-d.Percentage))
}

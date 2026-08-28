package pricing

type Policy struct{ Standard, VIP int }

func Default() Policy { return Policy{Standard: 100, VIP: 220} }
func (p Policy) Price(kind string, quantity int) int {
	if quantity < 0 {
		quantity = 0
	}
	unit := p.Standard
	if kind == "vip" {
		unit = p.VIP
	}
	return unit * quantity
}
func Discount(total int, rate float64) int {
	if rate < 0 {
		rate = 0
	}
	if rate > 1 {
		rate = 1
	}
	return int(float64(total) * (1 - rate))
}
func Tax(total int, rate float64) int {
	if rate < 0 {
		rate = 0
	}
	return int(float64(total) * (1 + rate))
}

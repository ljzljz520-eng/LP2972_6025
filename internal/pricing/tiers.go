package pricing

func Tier(kind string) string {
	if kind == "vip" || kind == "premium" {
		return "vip"
	}
	return "standard"
}
func Bundle(p Policy, kind string, n int) int {
	if n <= 0 {
		return 0
	}
	total := p.Price(kind, n)
	if n >= 10 {
		return Discount(total, 0.1)
	}
	return total
}

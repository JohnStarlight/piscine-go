package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	switch order {
	case "burger":
		b := food{preptime: 15}
		return b.preptime

	case "nuggets":
		n := food{preptime: 12}
		return n.preptime

	case "chips":
		c := food{preptime: 10}
		return c.preptime

	default:
		return 404
	}
}

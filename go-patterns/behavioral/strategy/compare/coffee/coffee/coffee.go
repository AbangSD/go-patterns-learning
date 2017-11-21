package coffee

type Coffee struct {
	Name  string
	Price float32
}

type Seasoning func(Coffee) Coffee

func Make(cs, ce Seasoning) Seasoning {
	return func(c Coffee) Coffee {
		return ce(cs(c))
	}
}

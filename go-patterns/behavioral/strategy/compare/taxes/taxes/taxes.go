package taxes

type Person interface {
	Taxes() float32
}

type People struct {
	Person Person
}

func (o *People) Taxes() float32 {
	return o.Person.Taxes()
}

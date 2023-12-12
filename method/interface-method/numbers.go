package interface_method

type Integer int

type Math interface {
	Add(i Integer)
	Multiply(i Integer) Integer
}

func (i2 *Integer) Add(i Integer) {
	*i2 = *i2 + i
}

func (i2 *Integer) Multiply(i Integer) Integer {
	return *i2 * i
}

package interface_method

type Addr struct {
	start int
}

func (addr Addr) AddTo(i int) int {
	return addr.start + i
}

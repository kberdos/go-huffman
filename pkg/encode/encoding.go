package encode

type encoding struct {
	R uint64 `json:"R"`
	N uint8  `json:"N"`
	// number of bits in r
}

func (e encoding) push0() encoding { // pushes a 0 to the right
	return encoding{
		R: e.R * 2,
		N: e.N + 1,
	}
}

func (e encoding) push1() encoding { // pushes a 1 to the right
	return encoding{
		R: e.R*2 + 1,
		N: e.N + 1,
	}
}

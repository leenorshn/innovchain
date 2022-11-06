package main

type Trans struct {
	from   string
	to     string
	amount int64
}

func (t *Trans) SendMoney(from, to string, amount int64) *Trans {
	return &Trans{
		amount: amount,
		from:   from,
		to:     to,
	}
}

package scoreboard

import "io"

type Aggregator interface {
	Save(string, ...string)
	Dump()
}

type aggregate struct {
	_records []string
	iowriter io.Writer
}

func NewAggregate(io io.Writer) *aggregate {
	return &aggregate{
		iowriter: io,
	}
}

func (a *aggregate) Save(s string, d ...string) {
	if len(d) > 0 {
		a._records = append(d, a._records...)
	} else {
		a._records = append(a._records, s)
	}
}

func (a *aggregate) Dump() {
	for _, s := range a._records {
		a.iowriter.Write([]byte(s))
	}
}

func (a *aggregate) Adj() {

}

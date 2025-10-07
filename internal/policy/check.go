package policy

import (
	"context"
)

type Checker struct {
	fileName string
	input    interface{}
	ctx      context.Context
}

func Check(fileName string) *Checker {
	return &Checker{
		fileName: fileName,
		ctx:      context.Background(),
	}
}

func (c *Checker) Security() *SecurityService {
	return &SecurityService{
		checker: c,
	}
}

func (c *Checker) Replica() *ReplicaService {
	return &ReplicaService{
		checker: c,
	}
}

package policy

import (
	"fmt"
)

type ReplicaService struct {
	checker *Checker
}

func (r *ReplicaService) Run() (*Result, error) {
	if err := r.checker.yaml(); err != nil {
		return nil, fmt.Errorf("failed to load input: %w", err)
	}

	violations, err := evaluatePolicies(r.checker.ctx, r.checker.input, []string{"replica_policy.rego"})
	if err != nil {
		return nil, err
	}

	return &Result{
		Violations: violations,
		Passed:     len(violations) == 0,
	}, nil
}

package policy

import (
	"fmt"
)

type SecurityService struct {
	checker       *Checker
	imageChecks   bool
	contextChecks bool
}

func (s *SecurityService) Images() *SecurityService {
	s.imageChecks = true
	return s
}

func (s *SecurityService) Context() *SecurityService {
	s.contextChecks = true
	return s
}

func (s *SecurityService) Run() (*Result, error) {
	if err := s.checker.yaml(); err != nil {
		return nil, fmt.Errorf("failed to load input: %w", err)
	}

	var policies []string
	if s.imageChecks {
		policies = append(policies, "image_policy.rego")
	}
	if s.contextChecks {
		policies = append(policies, "security_context_policy.rego")
	}

	if len(policies) == 0 {
		return nil, fmt.Errorf("no policy checks selected")
	}

	violations, err := evaluatePolicies(s.checker.ctx, s.checker.input, policies)
	if err != nil {
		return nil, err
	}

	return &Result{
		Violations: violations,
		Passed:     len(violations) == 0,
	}, nil
}

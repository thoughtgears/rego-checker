package policy

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/open-policy-agent/opa/v1/rego"
)

const (
	policies = "policy"
)

type Result struct {
	Violations []string
	Passed     bool
}

func evaluatePolicies(ctx context.Context, input interface{}, policyFiles []string) ([]string, error) {
	if len(policyFiles) == 0 {
		return nil, fmt.Errorf("no policy files provided")
	}

	regoOpts := []func(*rego.Rego){
		rego.Query("data.main.deny"),
	}

	for _, fileName := range policyFiles {
		policyPath := filepath.Join(policies, fileName)
		content, err := os.ReadFile(policyPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read policy file %s: %w", policyPath, err)
		}
		regoOpts = append(regoOpts, rego.Module(fileName, string(content)))
	}

	query, err := rego.New(regoOpts...).PrepareForEval(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		return nil, fmt.Errorf("failed to evaluate policy: %w", err)
	}

	var violations []string
	if len(results) > 0 && len(results[0].Expressions) > 0 {
		if violationList, ok := results[0].Expressions[0].Value.([]interface{}); ok {
			for _, v := range violationList {
				violations = append(violations, fmt.Sprintf("%v", v))
			}
		}
	}

	return violations, nil
}

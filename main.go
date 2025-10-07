package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thoughtgears/rego-checker/internal/policy"
)

const (
	PolicyDir = "policy"
)

func main() {
	fileName := flag.String("file", "", "file name")
	flag.StringVar(fileName, "f", "", "file name")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Error: filename is required")
		flag.Usage()
		os.Exit(1)
	}

	if _, err := os.Stat(*fileName); os.IsNotExist(err) {
		fmt.Printf("Error: file '%s' does not exist\n", *fileName)
		os.Exit(1)
	}

	checks := []struct {
		name string
		run  func() (*policy.Result, error)
	}{
		{
			name: "Security",
			run:  func() (*policy.Result, error) { return policy.Check(*fileName).Security().Images().Context().Run() },
		},
		{
			name: "Replica",
			run:  func() (*policy.Result, error) { return policy.Check(*fileName).Replica().Run() },
		},
	}

	var allViolations []string
	for _, check := range checks {
		result, err := check.run()
		if err != nil {
			fmt.Printf("Error running %s checks: %s\n", check.name, err)
			os.Exit(1)
		}

		if !result.Passed {
			allViolations = append(allViolations, result.Violations...)
		}
	}

	if len(allViolations) > 0 {
		fmt.Println("❌ Policy violations found:")
		for _, v := range allViolations {
			fmt.Printf("- %v\n", v)
		}
		os.Exit(1)
	}

	fmt.Println("✅ All policy checks passed!")
}

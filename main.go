package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/open-policy-agent/opa/rego"
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

	content, err := os.ReadFile(*fileName)
	if err != nil {
		log.Fatalf("Failed to read input file '%s': %s", *fileName, err)
	}

	var input interface{}
	if err := yaml.Unmarshal(content, &input); err != nil {
		log.Fatalf("Failed to parse input: %s", err)
	}

	policies, err := os.ReadDir(PolicyDir)
	if err != nil {
		log.Fatalf("Failed to read policy directory: %s", err)
	}

	regoOpts := []func(*rego.Rego){
		rego.Query("data.main.deny"),
	}

	for _, file := range policies {
		if filepath.Ext(file.Name()) == ".rego" {
			policyPath := filepath.Join(PolicyDir, file.Name())
			content, err := os.ReadFile(policyPath)
			if err != nil {
				log.Fatalf("Failed to read policy file %s: %s", policyPath, err)
			}
			regoOpts = append(regoOpts, rego.Module(file.Name(), string(content)))
		}
	}

	ctx := context.Background()
	query, err := rego.New(regoOpts...).PrepareForEval(ctx)
	if err != nil {
		log.Fatalf("Failed to prepare query: %s", err)
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		log.Fatalf("Failed to evaluate policy: %s", err)
	}

	if len(results) > 0 && len(results[0].Expressions) > 0 {
		violations, ok := results[0].Expressions[0].Value.([]interface{})

		if ok && len(violations) > 0 {
			fmt.Println("❌ Policy violations found:")
			for _, v := range violations {
				fmt.Printf("- %v\n", v)
			}
			os.Exit(1)
		}
	}

	fmt.Println("✅ Policy check passed!")
}

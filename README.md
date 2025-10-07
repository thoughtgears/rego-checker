# Rego Checker

A POC to test out writing your own rego checker in golang to support the [OPA](https://www.openpolicyagent.org/) policy 
engine and be able to write your own checkers in a deployment application for example. This would be useful for having 
a deployment service that does OPA policies before allowing a deployment to proceed.

## Usage 

`go run main.go --file examples/{good,bad}-values.yaml`
package main

deny[msg] {
    not input.securityContext
    msg := "A securityContext must be defined for the deployment."
}
package main

deny contains msg if {
    not input.securityContext
    msg := "A securityContext must be defined for the deployment."
}
package main

deny contains msg if {
    input.replicaCount <= 1
    msg := "Replica count must be greater than 1."
}
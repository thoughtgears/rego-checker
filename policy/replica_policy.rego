package main

deny[msg] {
    input.replicaCount <= 1
    msg := "Replica count must be greater than 1."
}
package main

deny[msg] {
    input.image.tag == "latest"
    msg := "Using the 'latest' image tag is not recommended."
}
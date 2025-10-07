package main

deny contains msg if {
    input.image.tag == "latest"
    msg := "Using the 'latest' image tag is not recommended."
}
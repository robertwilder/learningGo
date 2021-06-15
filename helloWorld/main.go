package main

import "fmt"

func main() {
	fmt.Println("Hi there")
}

//1. How do we run the code inside our code 

// This is how you run - compile and immediately execute
// go run "file name"

// go build
// - This compiles but does not execute

// go fmt 
// - go format, formats the code 

// go install 

// go get

// go test






//2. What does 'package main' mean

// Package = Project or Workspace
// top of the file has to state which package it relates too 
// Two different types of Packages in Go 
// 1. Executable 
// when finished it spits out an executable file - used for actually doing something
// Main shows that it is an executable type package
// only use for runable files
// if package is main then func has to be main too 
// 2. Reusable
// We add reusable logic or helper functions.



//3. What does 'import "fmt' mean 

// give access to all the code and functionality form fmt - shortend version of format

// main, math, encoding, crypto, io, fmt, debug

// you can go beyond the standard Libaries

// golang.org/pkg - shows all packages in the standard library 



// 4. What is that func thing?

// short for function - 

// How is the main.go file organized?




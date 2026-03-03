package main

import "fmt"

const pi = 3.14
const appName = "Go Variables Example"
const Version = 1.0
const DebugMode = true
const MaxUsers = 100

// In go we need declare variables near to where they are used
// var can be declared in block
var (
	country  string = "USA"
	zipCode  int    = 94107
	isActive bool   = true
)

// const can be declared in block
const (
	language = "Go"
	platform = "Cross-Platform"
)

// when const is initialized with lower case it is not exported
// whrn const is initialized with upper case it is exported

func main() {
	var name string = "Cristian"
	// go use ligic assignment
	fmt.Println("Hello, " + name + "!")
	fmt.Printf("Name: %s %T\n", name, name)

	// Simplified declaration
	// When data type is obvious
	city := "San Francisco"
	fmt.Printf("City: %s %T\n", city, city)
	age := 30
	fmt.Printf("Age: %d %T\n", age, age)
	fmt.Printf("Pi: %f %T\n", pi, pi)
	fmt.Printf("App Name: %s %T\n", appName, appName)
	fmt.Printf("Version: %f %T\n", Version, Version)
	fmt.Printf("Debug Mode: %t %T\n", DebugMode, DebugMode)
	fmt.Printf("Max Users: %d %T\n", MaxUsers, MaxUsers)
	fmt.Printf("Country: %s %T\n", country, country)
	fmt.Printf("Zip Code: %d %T\n", zipCode, zipCode)
	fmt.Printf("Is Active: %t %T\n", isActive, isActive)
	fmt.Printf("Language: %s %T\n", language, language)
	fmt.Printf("Platform: %s %T\n", platform, platform)
}

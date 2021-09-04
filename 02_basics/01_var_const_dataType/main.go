package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float32
	type EuropeanVelocity float32

	const eur EuropeanVelocity = 120.4
	const us AmericanVelocity = 130

	var eurSpeed EuropeanVelocity = eur * 3.6
	var usSpeed AmericanVelocity = ((us * 1000) / 3600) / 1.609

	usSpeed1 := float64(usSpeed)
	usSpeed2 := math.Round((usSpeed1 * 100) / 100)

	fmt.Println(eurSpeed, usSpeed1)
	fmt.Printf("%.2f\n", usSpeed1)	
	fmt.Printf("%.2f\n", usSpeed2)	
}
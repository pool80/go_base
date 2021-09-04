package main

import (
	"fmt"
	"math"
)

func main() {
	type AmericanVelocity float64
	type EuropeanVelocity float64

	const eur EuropeanVelocity = 120.4
	const us AmericanVelocity = 130

	var eurSpeed EuropeanVelocity = eur * 3.6
	var usSpeed AmericanVelocity = ((us * 1000) / 3600) / 1.609

	amV := float64(usSpeed)

	fmt.Println(eurSpeed, usSpeed)
	fmt.Println(math.Round(amV * 100) / 100)
}
package main

import (
	"fmt"
	"math"
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinYearsToSurpass(aliceInitial, bobBonus, aliceRate, bobRate float64) int {
	// If already strictly greater at t=0
	if bobBonus > aliceInitial {
		return 0
	}
	// Equal rates
	if aliceRate == bobRate {
		// If equal growth and Bob <= Alice now, never surpasses
		return -1
	}
	// If Bob grows slower or equal but not already ahead → impossible
	if bobRate < aliceRate {
		return -1
	}

	// bobRate > aliceRate
	// Find smallest integer n such that: bobBonus*(1+b)^n > aliceInitial*(1+a)^n
	// => ((1+b)/(1+a))^n > aliceInitial/bobBonus
	r := (1.0 + bobRate) / (1.0 + aliceRate)
	if bobBonus <= 0 {
		// Bob starts at 0; still fine mathematically.
		if bobBonus == 0 && aliceInitial == 0 {
			return -1 // equal forever at 0
		}
	}
	// If bobBonus == 0, RHS is +Inf; use iteration fallback.
	if bobBonus == 0 {
		ai, bi := aliceInitial, 0.0
		years := 0
		for years <= 1_000_000 {
			years++
			ai *= 1.0 + aliceRate
			bi *= 1.0 + bobRate
			if bi > ai {
				return years
			}
		}
		return -1
	}

	target := aliceInitial / bobBonus
	// If target <= 1 and r > 1, n could be 0 (but we already checked t=0 case)
	n := int(math.Floor(math.Log(target) / math.Log(r)))
	n = maxInt(0, n) + 1

	// verify by simulation to avoid floating error
	ai := aliceInitial
	bi := bobBonus
	for year := 0; year <= n+2; year++ {
		if bi > ai {
			return year
		}
		ai *= 1.0 + aliceRate
		bi *= 1.0 + bobRate
	}
	// Shouldn't reach here unless extreme precision issues; fall back to bounded loop
	for year := n + 3; year <= n+1_000; year++ {
		if bi > ai {
			return year
		}
		ai *= 1.0 + aliceRate
		bi *= 1.0 + bobRate
	}
	return -1
}

func main() {
	// Example usages
	fmt.Println("Example 1:", MinYearsToSurpass(100.0, 50.0, 0.05, 0.10))  // Bob starts behind but grows faster
	fmt.Println("Example 2:", MinYearsToSurpass(100.0, 200.0, 0.05, 0.03)) // Bob already ahead
	fmt.Println("Example 3:", MinYearsToSurpass(100.0, 0.0, 0.05, 0.10))   // Bob starts at 0
	fmt.Println("Example 4:", MinYearsToSurpass(0.0, 0.0, 0.01, 0.02))     // both zero
}

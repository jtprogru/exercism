package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var rate float32
	if balance < 0 {
		rate = 3.213
	} else if balance >= 0 && balance < 1000 {
		rate = 0.5
	} else if balance >= 1000 && balance < 5000 {
		rate = 1.621
	} else if balance >= 5000 {
		rate = 2.475
	}

	return rate
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	var interest float64
	var rate float32
	rate = InterestRate(balance)

	interest = balance / 100.0 * float64(rate)

	return interest
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int
	var b float64 = balance
	for b < targetBalance {
		b = AnnualBalanceUpdate(b)
		years++
	}

	return years
}

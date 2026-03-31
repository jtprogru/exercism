package bookstore

func Cost(basket []int) int {
	if len(basket) == 0 {
		return 0
	}

	counts := make(map[int]int)
	for _, b := range basket {
		counts[b]++
	}

	groupCounts := make(map[int]int)

	for {
		groupSize := 0
		for k := range counts {
			if counts[k] > 0 {
				counts[k]--
				groupSize++
			}
		}

		if groupSize == 0 {
			break
		}

		groupCounts[groupSize]++
	}

	for groupCounts[5] > 0 && groupCounts[3] > 0 {
		groupCounts[5]--
		groupCounts[3]--
		groupCounts[4] += 2
	}

	total := 0
	for size, count := range groupCounts {
		total += price(size) * count
	}

	return total
}

func price(n int) int {
	switch n {
	case 1:
		return 800
	case 2:
		return 2 * 800 * 95 / 100
	case 3:
		return 3 * 800 * 90 / 100
	case 4:
		return 4 * 800 * 80 / 100
	case 5:
		return 5 * 800 * 75 / 100
	}
	return 0
}

package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var count int
    for _, v := range birdsPerDay {
        count += v
    }
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    weekLen := 7
    startIndex := (week * weekLen) - weekLen
    endIndex := week * weekLen

    return TotalBirdCount(birdsPerDay[startIndex:endIndex])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    var idx int

    for idx = range birdsPerDay {
        if idx % 2 == 0 {
            birdsPerDay[idx] += 1
        }
    }
    return birdsPerDay
}

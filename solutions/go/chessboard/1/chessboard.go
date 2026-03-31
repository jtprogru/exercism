package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
    var counter int
    r := cb[rank]
    for _, v := range r {
        if v {
            counter++
        }
    }
    return counter
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
    var counter int
	if file < 1 || file > 8 {
		return 0
	}

	for _, line := range cb {
		if line[file-1] {
			counter++
		}
	}
	return counter
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
    lenLine := len(cb["A"])
    countLine := len(cb)
    return lenLine * countLine
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
    var counter int
    for _, line := range cb {
        for _, square := range line {
            if square {
                counter++
            }
        }
    }
	return counter
}

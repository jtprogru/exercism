package strand

func ToRNA(dna string) string {
	var res string
	for _, v := range dna {
		switch string(v) {
		case "G":
			res += "C"
		case "C":
			res += "G"
		case "T":
			res += "A"
		case "A":
			res += "U"
		}
	}
	return res
}

package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var res Ints

	for _, elem := range i {
		if filter(elem) {
			res = append(res, elem)
		}
	}

	return res
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var res Ints

	for _, elem := range i {
		if !filter(elem) {
			res = append(res, elem)
		}
	}

	return res
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var res Lists

	for _, elem := range l {
		if filter(elem) {
			res = append(res, elem)
		}
	}

	return res
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var res Strings

	for _, elem := range s {
		if filter(elem) {
			res = append(res, elem)
		}
	}

	return res
}

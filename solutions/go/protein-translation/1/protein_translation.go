package protein

import (
	"errors"
)

var ErrStop = errors.New("stop")

var ErrInvalidBase = errors.New("invalid base")

func IsStopCodon(c string) bool {
	return c == "UAA" || c == "UAG" || c == "UGA"
}

func FromRNA(rna string) ([]string, error) {
	var proteins []string
	bases := []rune(rna)
	for i := 0; i < len(bases); i += 3 {
		p, err := FromCodon(string(bases[i : i+3]))
		switch err {
		case ErrInvalidBase:
			return nil, err
		case ErrStop:
			return proteins, nil
		default:
			proteins = append(proteins, p)
		}
	}
	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	if IsStopCodon(codon) {
		return "", ErrStop
	}
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	default:
		return "", ErrInvalidBase
	}
}

// Package protein implements a solution to the same-named problem on Go track of Exercism
package protein

import (
	"errors"
)

// ErrStop terminates the translation and the final methionine is not translated into the protein sequence
var ErrStop = errors.New("STOP")

// ErrInvalidBase indicates that a given codon was invalid
var ErrInvalidBase = errors.New("INVALID")

var proteins = map[string]string{"AUG": "Methionine", "UUU": "Phenylalanine", "UUC": "Phenylalanine",
	"UUA": "Leucine", "UUG": "Leucine", "UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
	"UAU": "Tyrosine", "UAC": "Tyrosine", "UGU": "Cysteine", "UGC": "Cysteine", "UGG": "Tryptophan",
	"UAA": "STOP", "UAG": "STOP", "UGA": "STOP"}

// FromCodon returns a protein given a valid codon string and an error otherwise
func FromCodon(input string) (string, error) {
	if protein, ok := proteins[input]; ok {
		if protein == "STOP" {
			return "", ErrStop
		}
		return protein, nil
	}
	return "", ErrInvalidBase
}

// FromRNA translates RNA sequences into proteins
func FromRNA(input string) (proteins []string, err error) {
	idx := 0
	for idx < len(input)/3 {
		protein, e := FromCodon(input[idx*3 : idx*3+3])
		if e != nil {
			if e == ErrStop {
				return proteins, nil
			}
			return proteins, e
		}
		proteins = append(proteins, protein)
		idx++
	}

	return proteins, nil
}

package main

import (
    "fmt"
    "math"
    "math/cmplx"
)
// Superposition of 2 (binary/bit) states
type Qubit struct {
    C0 complex128
    C1 complex128
}

func singletonQubit(c0,c1 complex128) (*Qubit, error) {
    norm := math.Pow(cmplx.Abs(c0), 2) + math.Pow(cmplx.Abs(c1), 2)
    if math.Abs(norm - 1.0) < TOLERANCE {
        qb := &Qubit{ C0: c0, C1: c1 }
        return qb, nil
    }
    return nil, fmt.Errorf("src/main.go : singletonQubit() :: ERROR ::: Invalid complex number arguments for a valid qubit.")
}

func (qb *Qubit) measure(observation float64) int {
    if real(qb.C0) == 1 && imag(qb.C0) == 0 && real(qb.C1) == 0 && imag(qb.C1) == 0 {
	    return 0
    } else if real(qb.C0) == 0 && imag(qb.C0) == 0 && real(qb.C1) == 1 && imag(qb.C1) == 0 {
	    return 1
    }
    prob0 := math.Pow(cmplx.Abs(qb.C0), 2)
    prob1 := math.Pow(cmplx.Abs(qb.C1), 2)
    if prob0 > prob1 && observation < prob0 {
        return 0
    } else if prob1 > prob0 && observation < prob1 {
	    return 1
    } else {
	    return int(math.Round(observation))
    }
}
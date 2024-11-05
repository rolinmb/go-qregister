package main

import (
    "fmt"
    "math"
    "math/cmplx"
)
// Superposition of n (positive integer >= 1) states
type Qudit struct {
    Amplitudes []complex128
}

func singletonQudit(cargs []complex128) (*Qudit, error) {
    norm := 0.0
    for _, c := range cargs {
        norm += math.Pow(cmplx.Abs(c), 2)
    }
    if math.Abs(norm - 1.0) < TOLERANCE {
        qd := &Qudit { Amplitudes: cargs }
        return qd, nil
    }
    return nil, fmt.Errorf("src/main.go : singletonQudit() :: ERROR ::: Invalid complex number arguments for a valid qudit.")
}

func (qd *Qudit) measure() int {
    observation := getObservation()
    cuumProb := 0.0
    for i, c := range qd.Amplitudes {
	    cuumProb += math.Pow(cmplx.Abs(c), 2)
	    if observation < cuumProb {
            return i
	    }
    }
    return len(qd.Amplitudes) - 1 
}

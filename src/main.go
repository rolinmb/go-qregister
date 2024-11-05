package main

import (
    "fmt"
    "math"
    "math/cmplx"
    "math/rand"
    "time"
)

const (
    N_ITERS = 1000
)

func getObservation() float64 {
    time.Sleep(1*time.Nanosecond)
    rand.Seed(time.Now().UnixNano())
    return rand.Float64()
}

type Qubit struct {
    C0 complex128
    C1 complex128
}

func singletonQubit(c0,c1 complex128) (*Qubit, error) {
    norm := math.Pow(cmplx.Abs(c0), 2) + math.Pow(cmplx.Abs(c1), 2)
    if math.Abs(norm - 1.0) < 1e-6 {
        qb := &Qubit{ C0: c0, C1: c1 }
        return qb, nil
    }
    return nil, fmt.Errorf("src/main.go : singletonQubit() :: ERROR ::: Invalid complex number arguments for a valid qubit.")
}

func (qb *Qubit) measure() int {
    if real(qb.C0) == 1 && imag(qb.C0) == 0 && real(qb.C1) == 0 && imag(qb.C1) == 0 {
	return 0
    } else if real(qb.C0) == 0 && imag(qb.C0) == 0 && real(qb.C1) == 1 && imag(qb.C1) == 0 {
	return 1
    }
    observation := getObservation()
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


func main() {
    c0 := complex(math.Sqrt(3.0)/2.0, 0.0) // 75% chance of '0'
    c1 := complex(0.0, 0.5) // 25% chance of '1'
    qba, err := singletonQubit(c0, c1)
    if err != nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Qubit A ", err)
        return
    }
    countZeros := 0
    countOnes := 0
    for i := 0; i < N_ITERS; i++ {
        result := qba.measure()
        if result == 0 {
            countZeros += 1
        } else {
            countOnes += 1
        }
        //fmt.Printf("src/main.go : main() :: Qubit Test A Measurement Result = %d\n", result)
    }
    fmt.Printf("src/main.go : main() :: Qubit Test A Results after %d iterations ::: \n", N_ITERS)
    fmt.Printf("src/main.go : main() ::\n\t-> Number of Zeros = %d\n\t-> Number of Ones = %d\n", countZeros, countOnes)
    c2 := complex(0.0, 1.0/math.Sqrt(6.0)) // 1/6 chance of '0'
    c3 := complex(0.0, -math.Sqrt(5.0)/math.Sqrt(6.0)) // 5/6 chance of '1'
    qbb, err := singletonQubit(c2, c3)
    if err != nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Qubit B ", err)
        return
    }
    countZeros = 0
    countOnes = 0
    for i := 0; i < N_ITERS; i++ {
        result := qbb.measure()
        if result == 0 {
            countZeros += 1
        } else {
            countOnes += 1
        }
        //fmt.Printf("src/main.go : main() :: Qubit Test B Measurement Result = %d\n", result)
    }
    fmt.Printf("src/main.go : main() :: Qubit Test B Results after %d iterations ::: \n", N_ITERS)
    fmt.Printf("src/main.go : main() ::\n\t-> Number of Zeros = %d\n\t-> Number of Ones = %d\n", countZeros, countOnes)
}


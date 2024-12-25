package main

import (
    "fmt"
    "math"
)

func main() {
    c0 := complex(math.Sqrt(0.75), 0.0) // 75% chance of '0'
    c1 := complex(0.0, -1.0*math.Sqrt(0.25)) // 25% chance of '1'
    qb, err := singletonQubit(c0, c1)
    if err != nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Qubit A ", err)
        return
    }
    countZeros := 0
    countOnes := 0
    for i := 0; i < N_ITERS; i++ {
        result := qb.measure(getObservation())
        if result == 0 {
            countZeros++
        } else {
            countOnes++
        }
        //fmt.Printf("src/main.go : main() :: Qubit Test measurement result = %d\n", result)
    }
    fmt.Printf("\nsrc/main.go : main() :: Qubit Test Results after %d iterations :::", N_ITERS)
    fmt.Printf("\n\t-> Qubit State 0 = %d occurrences\n\t-> Qubit State 1 = %d occurrences", countZeros, countOnes)
    cargs0 := []complex128{
        complex(math.Sqrt(0.1), 0),       // 10% chance of '0'
        complex(0, -1.0*math.Sqrt(0.2)),  // 20% chance of '1'
        complex(-1.0*math.Sqrt(0.3), 0),  // 30% chance of '2'
        complex(0, math.Sqrt(0.4)),       // 40% chance of '3'
    }
    qd0, err := singletonQudit(cargs0)
    if err != nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Qudit 0", err)
        return
    }
    counts0 := make([]int, len(cargs0))
    for i := 0; i < N_ITERS; i++ {
        result := qd0.measure(getObservation())
        counts0[result]++
        //fmt.Printf("src/main.go : main() :: Qudit 0 Test measurement result = %d\n", result)
    }
    fmt.Printf("\nsrc/main.go : main() :: Qudit 0 Test Results after %d iterations :::", N_ITERS)
    for i, count := range counts0 {
        fmt.Printf("\n\t-> Qudit 0 State %d: %d occurrences", i, count)
    }
    cargs1 := []complex128{
        complex(math.Sqrt(0.05), 0),
        complex(0, -1.0*math.Sqrt(0.15)),
        complex(-1.0*math.Sqrt(0.25), 0),
        complex(math.Sqrt(0.25), 0),
        complex(0, math.Sqrt(0.3)),
    }
    qd1, err := singletonQudit(cargs1)
    if err != nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Qudit1", err)
        return
    }
    counts1 := make([]int, len(cargs1))
    for i := 0; i < N_ITERS; i++ {
        result := qd1.measure(getObservation())
        counts1[result]++
        //fmt.Printf("src/main.go : main() :: Qudit 1 Test measurement result = %d\n", result)
    }
    fmt.Printf("\nsrc/main.go : main() :: Qudit 1 Test Results after %d iterations :::", N_ITERS)
    for i, count := range counts1 {
        fmt.Printf("\n\t-> Qudit 1 State %d: %d occurrences", i, count)
    }
    qr := QuantumRegister{
        Qudits: []*Qudit{qd0, qd1},
    }
    jointState := qr.TensorProduct()
    if jointState == nil {
        fmt.Println("src/main.go : main() :: ERROR ::: Tensor product failed. Quantum register is empty.")
        return
    }
    fmt.Println("\nsrc/main.go : main() :: Tensor Product of Quantum Register:", jointState.Amplitudes)
    qrCounts := make(map[string]int)
    for i := 0; i < N_ITERS; i++ {
        measurement := qr.measure(getObservation())
        stateKey := fmt.Sprint(measurement)
        qrCounts[stateKey]++
    }
    fmt.Printf("\nsrc/main.go : main() :: Quantum Register results after %d iterations :::\n", N_ITERS)
    for state, count := range qrCounts {
        fmt.Printf("\t -> Quantum Register State %s: %d occurrences\n", state, count)
    }
}

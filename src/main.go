package main

import (
    "fmt"
    "math"
)

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
    fmt.Printf("\n\t-> Number of Zeros = %d\n\t-> Number of Ones = %d\n", countZeros, countOnes)
    /*c2 := complex(0.0, 1.0/math.Sqrt(6.0)) // 1/6 chance of '0'
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
    
    fmt.Printf("src/main.go : main() :: Qubit Test B Results after %d iterations ::: \n", N_ITERS)
    fmt.Printf("\n\t-> Number of Zeros = %d\n\t-> Number of Ones = %d\n", countZeros, countOnes)*/
    cargs := []complex128{
        complex(math.Sqrt(0.1), 0),  // 10% chance of '0'
        complex(math.Sqrt(0.2), 0),  // 20% chance of '1'
        complex(math.Sqrt(0.3), 0),  // 30% chance of '2'
        complex(math.Sqrt(0.4), 0),  // 40% chance of '3'
    }
    qd, err := singletonQudit(cargs)
    if err != nil {
        fmt.Println("src/main.go () : main() :: ERROR ::: Qudit", err)
        return
    }
    counts := make([]int, len(cargs))
    for i := 0; i < N_ITERS; i++ {
        result := qd.measure()
        counts[result]++
    }
    fmt.Printf("src/main.go : main() :: Qudit Results after %d iterations:\n", N_ITERS)
    for i, count := range counts {
        fmt.Printf("\n\t-> Qudit State %d: %d occurrences", i, count)
    }
}

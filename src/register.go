package main

import (
    "fmt"
)

func tensorProduct(a0,a1 []complex128) []complex128 {
    result := make([]complex128, len(a0)*len(a1))
    for i, amp0 := range a0 {
        for j, amp1 := range a1 {
            result[i*len(a1)+j] = amp0 * amp1
        }
    }
    return result
}

type QuantumRegister struct {
    Qudits []*Qudit
}

func (qr *QuantumRegister) TensorProduct() *Qudit {
    if len(qr.Qudits) == 0 {
        return nil
    }
    jointAmplitudes := qr.Qudits[0].Amplitudes
    for i := 1; i < len(qr.Qudits); i++ {
        jointAmplitudes = tensorProduct(jointAmplitudes, qr.Qudits[i].Amplitudes)
    }
    return &Qudit{Amplitudes: jointAmplitudes}
}

func (qr *QuantumRegister) measure(observation float64) []int {
    jointState := qr.TensorProduct()
    if jointState == nil {
        fmt.Println("src/register.go : measure() :: ERROR ::: No Qudits in Quantum Register to take Tensor Product from.")
        return nil
    }
    result := jointState.measure(observation)
    // Decode the joint measurement result into individual qudit states
    numQudits := len(qr.Qudits)
    results := make([]int, numQudits)
    base := len(qr.Qudits[0].Amplitudes)
    for i := numQudits - 1; i >= 0; i-- {
        results[i] = result % base
        result /= base
    }
    return results
}
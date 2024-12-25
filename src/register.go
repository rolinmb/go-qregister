package main

type QuantumRegister struct {
    Qudits []*Qudit
}

func (qr *QuantumRegister) TensorProduct() []complex128 {
    if len(qr.Qudits) == 0 {
        return nil
    }
    result := qr.Qudits[0].Amplitudes
    for i := 1; i < len(qr.Qudits); i++ {
        result = tensorProduct(result, qr.Qudits[i].Amplitudes)
    }
    return result
}

func (qr *QuantumRegister) measure(observation float64) []int {
    results := make([]int, len(qr.Qudits))
    for i, qudit := range qr.Qudits {
        results[i] = qudit.measure(observation)
    }
    return results
}

func tensorProduct(a,b []complex128) []complex128 {
    product := make([]complex128, len(a)*len(b))
    for i, ai := range a {
        for j, bj := range b {
            product[i*len(b)+j] = ai * bj
        }
    }
    return product
}
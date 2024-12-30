THIS PROJECT IS DEPRECATED: view [my qudit-register-go repo](https://github.com/rolinmb/qudit-register-go) for an implementation that allows quantum logic gates to be used on individual qudits.

This repo references other repos of mine to implement a quantum register in golang. Entanglement is 'simulated' by sharing the same 'observation' across all quantum logic units when the register is measured.

My ['py-qregister' repo](https://github.com/rolinmb/py-qregister/) is just less verbose that golang, but isnt as thorough of an implementation as this repo.

Recreate this in rust/C/C++/zig if you wanna go fast. I also have a ['qubits' repo](https://github.com/rolinmb/qubits) for reference trying to implement these concepts in other programming languages.

Build and run from root: go build -C src -o main && ./src/main

TODO/WARNING: src/qubit.go code doesnt work as desired but src/qudit.go does work as expected; so just use the abstraction my friend.

Referencing other repos of mine to implement a quantum register in golang.

My 'py-qregister' repo is just less verbose that golang, but golang aint that bad and is just faster than python at executing.

Recreate this in rust/C/C++/zig if you wanna go fast. I have a 'qubits' repo for reference trying to implement these concepts in other programming languages.

Build and run from root: go build -C src -o main && ./src/main

TODO: Qubit type/struct doesn't work properly in this implementation, but qudit does, so just abstract brother

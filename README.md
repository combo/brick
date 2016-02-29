# brick: A software engineering clean room

_This is a pre-alpha concept, still in the design phase_

## Rationale

We want software to be reproducible and able to run anywhere.

Current technologies such as Docker and Nix are orthogonal solutions
to create and distribute atomic, isolated and reproducible applications that are
guaranteed to run anywhere (TM).

But, in their current state, their are not the panacea we are looking for:

* Docker leverages Linux namespacing to create isolated environments but requires complex wiring for any non-trivial application and its layering technology enables a lot of duplication and container bloat.

* Nix tries to solve the problem with a functional, immutable package manager but adoption is slow and not many energies are focused on the deployment and isolation problem.

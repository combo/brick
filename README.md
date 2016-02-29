# brick: A software engineering clean room

_This is a pre-alpha concept, still in the design phase_

## Rationale

We want software to be reproducible and able to run anywhere.

Current technologies such as Docker and Nix are orthogonal solutions
to create and distribute atomic, isolated and reproducible applications that are
guaranteed to run anywhere (TM).

But, in their current state, they are not the panacea we are looking for:

* Docker leverages Linux namespacing to create isolated environments but requires complex wiring for any non-trivial application and its layering technology enables a lot of duplication and container bloat.

* Nix tries to solve the problem with a functional, immutable package manager but adoption is slow and not many energies are focused on the deployment and isolation problem.

brick wants to merge the selling points of Docker (containerization) and Nix (declarative and immutable packaging) into a single product, that accompanies the developer from the initial development to packaging to release and deployment.

Applications should be developed into a **clean room** container that contains
all and uniquely its dependencies. Dependencies are both versioned and their contents hashed, and, if we can guarantee that the application works in the clean room, given that set of dependencies, we can guarantee it'll run anywhere.

A working application can then be distributed in a single file, and brick will make sure to download, verify all the dependencies required by the application, and run it in a container isolated from the rest of the host system.

Like with any other package manager, brick packages can be hosted anywhere, and built from source if no binary distribution is available. Again, the final package is built in the clean room which will contain everything needed to create the final product without any particular host system configuration.

The host system will be able to keep in cache multiple versions of the same package without conflict, as only the versions specified in the brick package will be loaded in the clean room.

## Project status

We are still in the design phase, things are moving and changing very fast.

Please refer to our [wiki](https://github.com/combo/brick/wiki) for the design documents.

## Contributing

Until the basic idea is fully fleshed out we are not accepting any contribution, but any kind of discussion is welcome.

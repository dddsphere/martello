# Description

Martello is a reference implementation of a micro monolith that serves as an example of Go microservices architecture that integrates Domain-Driven Design (DDD), Command and Query Responsibility Segregation (CQRS), and Event Sourcing principles. This repo should provide with a reliable reference for creating custom services based on these principles while minimizing external dependencies and leveraging the Go standard library.

A code generator to develop in another repository is envisioned to be developed based on Martello's architecture. This generator will utilize the reference implementation as a foundation for generating custom code that aligns with the established principles.

## Key Features

- **DDD-Driven Design**: Embrace modularity, separation of concerns, and a deep understanding of the domain model for well-organized and maintainable microservices.
- **CQRS Paradigm**: Promote a clear separation between write and read operations to design optimized read models and ensure a consistent and scalable write side.
- **Event Sourcing**: Capture every state change as an immutable event to enable precise auditing, temporal querying, and the ability to rebuild state at any given point.
- **Dependency Minimization**: Strive to minimize external dependencies by relying primarily on the functionality provided by the Go standard library for simplicity, efficiency, and ease of maintenance. The goal is to ensure that the generated code is self-contained. Any common structures or reusable components, such as struct embeddings, will be generated within the `internal` package of the project, avoiding the need to import additional packages from outside of the generated codebase.

[TopSpin](https://github.com/dddsphere/topspin) and its generator will eventually be modeled after this implementation. The goal is to leverage the principles and design patterns demonstrated in this reference implementation to shape the development of TopSpin and its associated code generator. By aligning with this implementation, we aim to ensure consistency, maintainability, and adherence to best practices.

## Usage

To use this reference implementation, follow the instructions below:

1. Clone the repository: `git clone github.com/dddsphere/martello.git`
2. Explore the different components and modules to understand the implementation.
3. Adapt and modify the codebase according to your specific project requirements.


## License

This project is licensed under the [MIT License](LICENSE).


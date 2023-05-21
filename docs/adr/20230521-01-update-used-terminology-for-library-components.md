# Name
Update used terminology for library components

## Status
Proposed

## Context
We are developing a library for creating microservices following Domain-Driven Design (DDD) patterns. Within the DDD vocabulary, the terms "service" is used to define specific structures and functionalities at the application and domain levels. However, our library also includes its own concept of a "service" to define some internal component. This results in the usage of the term "service" in three different contexts, which can lead to confusion and lack of clarity.

## Decision
To address the issue of ambiguity and improve the clarity of our library's terminology, we propose the following changes:

1. Rename the library's internal "service" object to "module" which more accurately represents their purpose and functionality.
2. Update the documentation, code comments, and examples throughout the library to reflect the new terminology consistently.

## Consequences
The decision to update the terminology for library components will have the following consequences:

1. Improved clarity: The revised terminology will differentiate between the DDD-related services and the library's internal objects, reducing confusion and enhancing understanding for developers using the library.

2. Enhanced maintainability: The new terminology will make the library's codebase more maintainable and readable, as the intent and purpose of each module will be explicitly conveyed through the updated naming conventions.

## Date
May 21, 2023

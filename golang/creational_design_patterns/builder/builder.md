# Builder Pattern

Builder pattern is used when an object has many optional fields.
Instead of using constructors with long parameter lists, we create a builder that:

- takes required fields upfront
- allows step-by-step configuration
- validates before building final object

**steps:**
1. Define final struct
2. Define builder struct
3. Add constructor with required fields
4. Add methods (return *Builder)
5. Add Build() with validation that returns final struct

**Use case:**
- building requests / configs / queries
- Many fields are optional
- Validation matters
- Readability is critical
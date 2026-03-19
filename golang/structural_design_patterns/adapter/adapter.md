# Adapter Pattern 

- Converts one interface into another that the client expects
- Acts as a wrapper/translator

**Problem it solves**

- Incompatible interfaces
- Reusing existing/third-party code without modification

**Structure:**

- Target → Product Interface expected by client
- Adaptee → Existing incompatible class
- Adapter → Concrete implemetaion of the product, Converts Target → Adaptee
- Client → Uses Target only

**Usecase**

- Integrating third-party APIs (payments, cloud, logging)
- Supporting multiple providers
- Working with legacy systems

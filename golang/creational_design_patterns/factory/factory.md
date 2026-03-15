# Factory Pattern:
- The Factory Pattern is a creational design pattern used to create objects without exposing the object creation logic to the client.
- Instead of the client directly creating objects using constructors, the client asks a factory to provide the required object. 
- The factory decides which concrete implementation to create based on input or configuration.
- The client only interacts with a common interface, not the actual implementation.
---------------
**note:**
Factory = a place that decides which object to create.

# Steps to Implement
When explaining or writing code in an interview, follow this order.
1. Define the Product Interface
- Create an interface that defines the behavior expected from all implementations.

**Example:**
- Notifier
- PaymentProcessor
- Storage

2. Implement Concrete Products

- Create multiple structs that implement the interface.
- Each struct represents a specific implementation.

**Example:**
- EmailNotifier
- SMSNotifier
- PushNotifier

All of them implement the same interface method.

3. Create the Factory

-Write a factory function that decides which implementation to create.
-The factory:
-Takes some input (type, mode, config)

-Returns the appropriate interface implementation

**Example:**
-GetNotifier("email") → EmailNotifier
-GetNotifier("sms") → SMSNotifier
-
-The client does not know which struct is created.

4. Client Code Uses the Interface

- The client requests an object from the factory and works only with the interface.
- The client should not create concrete implementations directly.

**Example:**
notifier := GetNotifier("email")
notifier.Send("Hello")

The client doesn't care whether it is Email, SMS, or Push.

----------------
**Use Cases**

1. Multiple Implementations Exist

When there are multiple implementations of the same behavior.

Examples:

- Notification services (Email, SMS, Push)
- Payment gateways (Stripe, PayPal, Razorpay)
- Storage systems (S3, GCS, Local)

2. Object Creation Depends on Runtime Input
config.notification = "sms"
3. Complex Object Creation
4. Decoupling Business Logic

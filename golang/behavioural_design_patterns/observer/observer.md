## Observer Pattern

- It defines a one-to-many relationship
- When subject changes → all observers get notified
- Achieved using:
  1. Interface (Subscriber)
  2. Concrete Subscriber implementations (Email, Mobile)
  3. Interface Subject - subject maintains a collection of observers
  4. Concrete Subject implementations (youtube channel)
- Promotes loose coupling & extensibility

**Use case:**

- News or Blog Subscriptions
- Social Media Notifications
- Event systems (logs, alerts)
- WebSocket listeners
- Cache invalidation listeners

**Situations Where the Observer Pattern May Fall Short**

- Excessive Observer Load
In high-scale systems with millions of observers (e.g., when a celebrity with 10M followers goes live), a direct notification loop becomes inefficient. Such cases are better handled using event queues, pub-sub architectures, or broadcast systems optimized for massive concurrency.

- In short, Observer Pattern works really well with a small number of observers **(sending snaps where friends count would be lesser)**, but to scale, it becomes essential to move toward an event-driven architecture.

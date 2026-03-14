# Singleton Pattern in Go
- Singleton belongs to Creational Design Patterns because it controls how objects are created.
- The Singleton Pattern ensures that a struct has only one instance and provides a global access point to that instance.
- In Go applications, this pattern is commonly used for components like:

1. Logger
2. Configuration manager
3. Database connection pool

-------------
Approaches to Implement Singleton in Go

The most common implementations are:

1. Eager Initialization
2. Lazy Initialization (Not Thread Safe)
3. Thread-Safe Lazy Initialization
4. sync.Once (Best Practice)

-----------------------
# 1. Eager singleton: 
- initialized when package is loaded

**Pros:**
- Very simple
- Thread-safe automatically
- No locking required

**Cons:**
- Instance created even if never used
- Not ideal for heavy objects

----------
# 2. Lazy Loading (Not Thread Safe)
- instance is created only when first requested

------------
# 3. Mutex Based Thread-Safe Singleton
- Every call locks the mutex.
- Even after initialization, this introduces performance overhead.

**Double Checked Locking**
- This improves performance by avoiding lock after initialization.

------------------
**sync.Once (Best Practice in Go)**
- Go provides a built-in mechanism for Singleton initialization.
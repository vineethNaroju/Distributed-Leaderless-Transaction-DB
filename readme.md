# Single Node Transaction DB
Node supports ACID transaction.

# Node
Node contains store of records. Each record has its own mutex for fine atomicity.

# Write Operations
Explicit new key creation using CreateRecord(key, val). 
Each transaction is atomic.

# Read Operations
Simple Get(key).

# Transaction
It contains list of `keys` that needs to be locked and `executeFn` to support business logic.

# Internals
A transaction can be owned by any node. Transaction will try to acquire lock over all keys, fetch
a local copy of records, execute logic, update changes to node's store and release locks.

Locks are acquired synchronosly in sequential order of keys. Acquiring locks async causes Dining
philosopher problem.

# Results
Check single-node.log

# Todo
1. Extend this to multi-node leaderless transaction with multiple co-ordinator.
2. Try to introduce idempotency and maintain transaction states.
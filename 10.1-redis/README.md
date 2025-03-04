# Redis (Remote Dictionary Server)

## What is Redis?

Redis is an open-source, in-memory data structure store that can be used as a:

- Database
- Cache
- Message broker
- Queue

## Key Features

### 1. In-Memory Storage

- Extremely fast performance (microsecond latency)
- Can persist data to disk for durability
- Supports data structures like strings, hashes, lists, sets, sorted sets

### 2. Data Structures

- Strings: Basic key-value pairs
- Lists: Ordered collections of strings
- Sets: Unordered collections of unique strings
- Sorted Sets: Sets ordered by a score
- Hashes: Collections of field-value pairs
- Bitmaps: Binary operations
- HyperLogLogs: Probabilistic data structure
- Streams: Append-only collections

### 3. Common Use Cases

- Caching: Frequently accessed data
- Session Management: Web application sessions
- Real-time Analytics: Counting and statistics
- Rate Limiting: API request throttling
- Queues: Message brokers
- Leaderboards: Using sorted sets
- Real-time Chat: Pub/sub functionality

### 4. Basic Commands

```
SET key value           # Set a key-value pair
GET key                 # Get value by key
DEL key                # Delete a key
INCR key               # Increment a number
EXPIRE key seconds     # Set key expiration
HSET hash field value  # Set hash field
LPUSH list value       # Push to list
SADD set value         # Add to set
```

### 5. Advanced Features

- Pub/Sub: Message broadcasting
- Transactions: Atomic operations
- Lua Scripting: Custom operations
- Replication: Master-slave setup
- Cluster Mode: Horizontal scaling
- Persistence: RDB and AOF options

### 6. Performance Characteristics

- Single-threaded (per core)
- Typical operations: < 1ms
- Can handle 100K+ operations per second
- Memory-bound (limited by RAM)

### 7. Best Practices

1. Use appropriate data structures
2. Set TTL (Time To Live) for cached data
3. Monitor memory usage
4. Plan for persistence if needed
5. Consider replication for high availability
6. Use pipelining for bulk operations

### 8. Common Patterns

```redis
# Caching
SET user:123 "user_data" EX 3600    # Cache for 1 hour

# Counter
INCR pageviews

# Rate Limiting
INCR rate:ip:123
EXPIRE rate:ip:123 60

# Session Store
HSET session:123 user_id 456
HSET session:123 login_time "2024-03-20"
```

## Getting Started

1. Install Redis
2. Start Redis server:

```bash
redis-server
```

3. Connect with Redis CLI:

```bash
redis-cli
```

## Security Considerations

- Set strong passwords
- Configure firewall rules
- Enable SSL/TLS for remote connections
- Regularly update Redis
- Limit exposed commands
- Monitor access logs

Redis is particularly powerful for:

- High-performance applications
- Real-time data processing
- Distributed systems
- Microservices architecture
- Scalable web applications

Understanding these fundamentals will help you effectively use Redis in your applications. Remember to choose the right data structures and patterns based on your specific use cases.

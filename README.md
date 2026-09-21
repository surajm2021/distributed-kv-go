# Distributed KV Store (Go)

A resilient, high-throughput key-value storage engine backed by the Raft consensus algorithm.

## Features
- **Raft Consensus**: Leader election, heartbeat heartbeats, and replicated log consistency.
- **Lock-Free Skiplist Storage**: In-memory sorted string tables for sub-millisecond lookups.
- **gRPC Cluster Interface**: Streaming client RPCs with automated leader redirection.

## Benchmark Results
- Write Throughput: ~85,000 ops/sec (3-node quorum)
- Read Latency (p99): 0.42 ms

## Running Tests
```bash
go test -v -race ./...
```

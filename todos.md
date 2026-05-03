# Portfolio Enhancement Plan

## 1. Articles (2)

### Article A: Designing Sub-Second Event-Driven Architectures in Go
**Focus**: Comparing channels vs. lock-free ring buffers for autonomous systems and websocket routing.
- [ ] **Research**: Draft benchmarks comparing unbuffered, buffered channels, and lock-free ring buffers.
- [ ] **Implementation**: Code minimal event loops demonstrating the differences.
- [ ] **Content Outline**:
  - Introduction to real-time agent/trading constraints.
  - The anatomy of a sub-second event loop.
  - Memory analysis and garbage collection pressure of channels vs buffers.
  - Conclusion/Takeaway.
- [ ] **Write Article**: Use `tech_article_prompt.md` to optimize for systems engineers.
- [ ] **Publish & Promote**: Post to LinkedIn and X highlighting trade-off matrix.

### Article B: Rust vs Go Threading Models: A Performance Deep Dive
**Focus**: Goroutines vs async tasks, memory analysis, and performance tables.
- [ ] **Research**: Set up identical concurrent workloads (e.g. million message passing) in Go and Rust.
- [ ] **Implementation**: Collect CPU/Memory profiles for both.
- [ ] **Content Outline**:
  - M:N Scheduling (Go) vs OS Thread / Async Runtime (Rust/Tokio).
  - Benchmark visualizations and tables.
  - Context switching overhead.
- [ ] **Write Article**: Finalize graphs and text.
- [ ] **Publish & Promote**: Post to LinkedIn and X aiming at the systems engineering community.

---

## 2. rteval Library
**Repository**: `github.com/omar391/rteval`
**Goal**: High-throughput Go library for sub-second rule evaluation and event routing.

### Optimised Task List
- [ ] **Phase 1: Core Primitives**
  - [ ] Implement `Event` interface with generic payload encapsulation.
  - [ ] Implement lock-free `RingBuffer` for ultra-low latency event queueing.
  - [ ] Setup core `EventLoop` loop runner with graceful shutdown and panic recovery.
- [ ] **Phase 2: Rule Evaluator**
  - [ ] Create `Rule` interface mapping `Event` -> `Action`.
  - [ ] Implement `Evaluator` struct that maps multiple rules to an incoming event stream.
  - [ ] Implement fast path pattern matching (exact match, regex, prefix) for routing.
- [ ] **Phase 3: Network / Agent Tooling**
  - [ ] Build a generic `WebsocketRouter` optimized for sub-second delivery.
  - [ ] Provide integration hooks for external Agent execution (e.g., executing tool calls).
- [ ] **Phase 4: Benchmarks & CI**
  - [ ] Add `go test -bench` covering ring buffer vs go-channels.
  - [ ] Add end-to-end load tests mimicking multi-agent event spikes.
  - [ ] Setup GitHub Actions for testing and linting (golangci-lint).
- [ ] **Phase 5: Release & Documentation**
  - [ ] Write comprehensive `README.md` with usage examples and architecture diagram.
  - [ ] Publish `v0.1.0` module.
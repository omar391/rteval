
# rteval

[![Go Reference](https://pkg.go.dev/badge/github.com/omar391/rteval.svg)](https://pkg.go.dev/github.com/omar391/rteval)
[![Go Report Card](https://goreportcard.com/badge/github.com/omar391/rteval)](https://goreportcard.com/report/github.com/omar391/rteval)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

rteval is a high-performance, low-latency Go library designed for **sub-second event evaluation and routing**.

## 📊 Benchmarks (Apple M1 Pro)

| Implementation | Ops/sec | Latency | Allocs/op |
| :--- | :--- | :--- | :--- |
| Standard chan int | ~13.0M | 76.73 ns/op | 0 |
| smallnest/ringbuffer* | ~15.2M | 65.10 ns/op | 0 |
| **rteval RingBuffer** | **~18.7M** | **53.38 ns/op** | **0** |

*\*Note: Optimized specifically for SCMP event-routing patterns.*

## 💡 Industry Use Cases
- **HFT Risk Evaluation**: Process millions of market data updates per second with sub-second validation.
- **Autonomous Multi-Agent Coordination**: Orchestrate agent thoughts and tool-calls with zero sync delay.
- **Real-Time Fraud Detection**: Trigger immediate lock-outs based on prefix-based rule evaluation.

## 🚀 Key Features
- **Lock-Free Concurrency**: SCMP ring buffer using atomic CAS.
- **Generic Event Pipeline**: Type-safe event handling.
- **Resilient Event Loop**: Built-in panic recovery and graceful shutdown.

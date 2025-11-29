# Kafka Broker (Custom Implementation in Go)

A lightweight Kafka broker implementation written from scratch in **Golang**, focusing on understanding and implementing Kafka's **internal protocol**. This project is not meant to replace Apache Kafka but to serve as a learning tool and a protocol reference implementation.

## 🚀 Overview

This broker implements the fundamental Kafka protocol API calls that allow clients to interact with a minimal but functional Kafka-like system.

### Supported API Keys

* **ApiVersions** – Allows clients to discover which API versions the broker supports.
* **DescribeTopicPartitions** – Returns metadata about topics and partitions.
* **Produce** *(TODO)* – Accept and store published messages.
* **Consumer APIs** *(TODO)* – Manage offset fetch, message consumption, and group coordination.


## 📦 Goals of This Project

* Understand the **Kafka protocol** at the binary level
* Implement a real TCP server that speaks Kafka's language
* Gain deep insight into how brokers handle metadata, versioning, and requests
* Create a foundation that can evolve into a more complete broker


## ▶️ Running the Broker

```bash
go run cmd/main.go
```

The broker listens on **localhost:9092** by default.


## 📚 Status

| API Key                 | Status        |
| ----------------------- | ------------- |
| ApiVersions             | ✅ Implemented |
| DescribeTopicPartitions | ✅ Implemented |
| Produce                 | 🚧 TODO    |
| Consumer APIs           | 🚧 TODO    |

---

# TempoGraph

TempoGraph is a production-grade, high-performance, in-memory graph engine built to manage temporal graphs. Temporal graphs are graphs where both nodes and edges evolve over time. 

It supports time-based queries, graph diffs, and entity histories, making it suitable for applications where the evolution of graph data is as important as its current state.

## Features

- In-memory graph with full temporal versioning
- Append-only mutation log for time travel and auditing
- Time-aware APIs:
  - `snapshot(t)` – state of the graph at time `t`
  - `diff(t1, t2)` – compare changes between two times
  - `history(id)` – view full change history of a node or edge
- RESTful API and optional CLI
- Fast in-memory performance, no external database required

## Use Cases

- Dynamic knowledge graph systems
- Evolving social networks
- Genealogical graph tracking
- Supply chain evolution modeling
- Infrastructure graph timelines

## Quick Start

### Prerequisites

- Go 1.21 or higher
- Unix-based OS (macOS/Linux recommended)

### Build Instructions

```bash
git clone https://github.com/as4395/tempograph.git
cd tempograph
go build -o tempograph ./src
```

## Running the Engine

```bash
./tempograph --http :8080
```

Example: Insert a node

```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"type":"node", "id":"user123", "label":"User", "properties":{"name":"Alice"}}' \
  http://localhost:8080/graph?ts=2025-07-01T12:00:00Z
```

Example: Query a snapshot

```bash
curl http://localhost:8080/snapshot?ts=2025-07-01T12:00:00Z
```

Example: View change history of a node

```bash
curl http://localhost:8080/history/user123
```

Example: View graph diff between two points in time

```bash
curl "http://localhost:8080/diff?from=2025-07-01T00:00:00Z&to=2025-07-14T00:00:00Z"
```

## Temporal API Endpoints

| Method | Endpoint              | Description                              |
|--------|-----------------------|------------------------------------------|
| POST   | `/graph`              | Insert or update node/edge at timestamp  |
| DELETE | `/graph`              | Delete node/edge at timestamp            |
| GET    | `/snapshot`           | Retrieve graph state at a specific time  |
| GET    | `/diff`               | Compute diff between two time snapshots  |
| GET    | `/history/{id}`       | Show full mutation history for entity    |

All endpoints that mutate or read time-sensitive data accept a `ts=` parameter in RFC3339 format (e.g., `2025-07-14T22:00:00Z`).

## Design Principles

- Immutable event log: All graph changes are written as time-stamped events
- Snapshot reconstruction: Any past graph state can be rebuilt from logs
- Modular design: Separation of concerns between API, logic, storage
- Fully in Go: Zero dependencies on external databases or languages

## Roadmap

- CLI query interface for local environments
- Graph export to JSON, GraphViz, or CSV
- Time-based metrics and analytics support
- Persistent backend integration (e.g., Badger, Pebble)

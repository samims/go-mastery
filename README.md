# Go Mastery

A curated collection of hands-on experiments, notes, and mini-projects to deepen my expertise in the Go programming language.  
This repository showcases learning progress through practical examples, organized by topic.

---

## 🎯 Purpose

The primary goal of this repository is **mastery through practice**.  
Rather than learning passively, I explore Go’s core concepts by writing small, focused programs—each demonstrating a specific idea or pattern.

### Objectives:
- Build a **solid foundation** in Go’s internals and idioms.
- Document learnings with **clear, reproducible experiments**.
- Create a **reference hub** for future projects and interviews.
- Demonstrate **structured, disciplined learning** aligned with industry best practices.

---

## 📂 Repository Structure

```text
go-mastery/
├── cmd/                # Self-contained experiments and runnable examples
│   ├── 01_escape_analysis/
│   │   ├── 1_escape.go
│   │   ├── 1_escape_test.go
│   │   └── main.go
│   └── 02_sync_pool/
│       ├── 2_sync_pool.go
│       └── 2_sync_pool_test.go
├── pkg/                # Reusable packages (if/when needed)
├── notebook/           # Notes, reflections, and key takeaways
│   └── LAB-NOTEBOOK.md
├── Makefile            # Common commands (run, test, bench, tidy)
└── go.mod
```

- Each subdirectory in `cmd/` is self-contained, with its own `main.go` (entry point) and supporting files.
- Tests (`*_test.go`) accompany examples to validate and explore behavior.
- The `notebook/` directory includes written reflections and insights from each experiment.

---

## 🚀 Running an Example

You can run an example by navigating into its directory:

```bash
cd cmd/01_escape_analysis
go run .
```

Or use the provided Makefile shortcuts:

```bash
make run01     # Run escape analysis example
make bench01   # Run benchmarks for escape analysis
```

---

## 📖 Topics Covered

- **Escape analysis** – [cmd/01_escape_analysis](cmd/01_escape_analysis)
- **sync.Pool usage and performance** – [cmd/02_sync_pool](cmd/02_sync_pool)

More topics will be added as learning progresses, including:

- Concurrency patterns
- Interfaces and type systems
- Memory management
- Microservices best practices

---

## 💡 Philosophy

I believe mastery is earned through **iterative deep dives**: start small, question everything, write code, measure results, and reflect.  
This repo represents that philosophy — a continuous journey to sharpen my Go skills while upholding production-quality standards.

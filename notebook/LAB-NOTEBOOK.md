### Day 1 – Escape Analysis
Code: [01_escape_analysis/](../cmd/01_escape_analysis/)
- Result:  676.1 ns/op  →  676.1 ns/op after fixing heap escape


### Day 2 – `sync.Pool` & Buffer Reuse
Code: [02_sync_pool/](../cmd/02_sync_pool/)

- ✅ **Proper Put**: `16.7 ns/op`, `48 B/op`, `2 allocs/op`
- ❌ **Wrong Put**: `458.4 ns/op`, `3145 B/op`, `6 allocs/op`

**Key Insight:**  
Reusing objects via `sync.Pool` works **only** when you return the *same* (reset) object.  
Putting a fresh object each time defeats the pool — increases allocations, memory usage, and runtime.

> ✍️ Always `reset` + `put` the same object back.  
> Don’t just allocate a new one — that pollutes the pool.

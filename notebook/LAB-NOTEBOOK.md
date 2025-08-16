### Day 1 – Escape Analysis
📁 [01_escape_analysis/](../cmd/01_escape_analysis)

- ⏱️ Result: `676.1 ns/op` → `676.1 ns/op` after fixing heap escape
- ✅ Confirmed: value now stays on the stack (no unnecessary heap allocation)

---

### Day 2 – `sync.Pool` & Buffer Reuse
📁 [02_sync_pool/](../cmd/02_sync_pool)

- ✅ **Proper Put**: `16.7 ns/op`, `48 B/op`, `2 allocs/op`
- ❌ **Wrong Put**: `458.4 ns/op`, `3145 B/op`, `6 allocs/op`

**Key Insight:**  
`sync.Pool` only works when you return the **same (reset)** object.  
Returning a fresh allocation defeats the purpose — increasing GC pressure and memory use.

> ✍️ Always `reset` + `Put()` the same object.  
> ❌ Don't allocate a new one before putting — it pollutes the pool.

📁 [handler](../cmd/02_sync_pool/handler)
```bash
go test -bench . -benchmem ../cmd/02_sync_pool/handler
```

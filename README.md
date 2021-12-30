`Read()` will read up to `len(p)` into `p`, when possible.
After a `Read()` call, `n` may be less then `len(p)`.
Upon error, `Read()` may still return `n` bytes in buffer `p`. For instance, reading from a TCP socket that is abruptly closed. Depending on your use, you may choose to keep the bytes in p or retry.
When a `Read()` exhausts available data, a reader may return a `non-zero n` and `err=io.EOF`. However, depending on implementation, a reader may choose to return a `non-zero n` and `err = nil` at the end of stream. In that case, any subsequent reads must return `n=0`, `err=io.EOF`.
Latly, a call to `Read()` that returns `n=0` and `err=nil` does not mean EOF as the next call to Read() may return more data.`
---
title: "How to use async-std"
date: 2020-04-12T14:08:30+08:00
draft: true

tags: ["rust", "async-std"]
---

Now, rust has two ways to write async functions, `async-std` and `tokio`.
Let's first start with `async-std`.

### Add dependencies in Cargo.toml

> Note that `attributes` is used for `#[async_std::main]`.

Filename: Cargo.toml

```sh
[dependencies]
async-std = { version = "1.6.5", features = ["attributes"] }
```

### Write functions

Filename: src/hello.rs

> `async main` requires return a `std::io::Result`.

```rust
#[async_std::main]
async fn main() -> std::io::Result<()> {
    println!("hello world!");
    Ok(())
}
```

### Build your code

```bash
cargo run 
```


title = "Introducing Spin"
template = "main"
date = "2022-03-14T00:22:56Z"
[extra]
url = "https://github.com/fermyon/spin/blob/main/docs/content/index.md"

---

Spin is a framework for building and running event-driven microservice applications
with WebAssembly (Wasm) components. With Spin, we’re trying to make it easier to get
started with using WebAssembly on the server so that we can all take advantage of the
security, portability, and speed WebAssembly provides when it comes to running
microservices.

### Structure of a Spin Application

1. A Spin application manifest (`spin.toml`) file which defines where your WebAssembly
   components live and what triggers them.
2. One or more WebAssembly _components_.

Spin executes the component(s) as a result of events being generated by the trigger(s)
defined in the `spin.toml` file.

### Example Spin Application

The following illustrates how to define an HTTP application.

#### HTTP Handler

This `hello_world` function written in Rust defines a component that takes an `http::Request`and returns a `Result<http::Response>`.

```rust
#[http_component]​
fn hello_world(_req: http::Request) -> Result<http::Response> {​
    Ok(http::Response::builder()​
        .status(200)​
        .body(Some("Hello, Fermyon!".into()))?)​
}​
```

#### Spin Manifest

Once the code is compiled to a WebAssembly component, it can be referenced in a `spin.toml`
file to create an HTTP application like below.

```toml
spin_version = "1"
name = "spin-hello-world"
trigger = { type = "http", base = "/" }
version = "1.0.0"

[[component]]
id = "hello"
source = "<path to compiled Wasm module>"
[component.trigger]
route = "/hello"
```

#### Running a Spin Application

Running this application with the `spin` CLI is as simple as using the `spin up` command.
Because a trigger type of `http` is specified in the `spin.toml`, `spin up` will start
a web server. Any time a request is made on the `/hello` route, it will invoke the
`hello_world` function. Adding another component is as simple as adding another `[[component]]`
stanza to the `spin.toml` file.

In the next section, we will [take Spin for a spin](/quickstart).

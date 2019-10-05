# todo-go-wasm

JS + wasm in Go (1.13.1)

*Disclaimer:* It's not an example of how it should be done.
My goal was to play with wasm API in Go and to implement as much logic as I can in wasm.

**Go WebAssembly**

https://github.com/golang/go/wiki/WebAssembly

**Configure IntelliJ IDEA**

IntelliJ IDEA -> Preferences -> Languages & Frameworks -> Go -> Build Tags & Vendoring

```
OS: js
Arch: wasm
```

https://github.com/golang/go/wiki/Configuring-GoLand-for-WebAssembly

**Go Wasm tutorials**

* [Go WebAssembly Tutorial - Building a Calculator Tutorial](https://tutorialedge.net/golang/go-webassembly-tutorial/)
* [Compiling Go to WebAssembly](https://www.sitepen.com/blog/compiling-go-to-webassembly/)
    Examples of which libraries to import in order to compile Go
* [Some changes in how to register go functions](https://stackoverflow.com/a/56469260)
    Use `FuncOf` instead of `NewCallback`
* [Go 1.11: WebAssembly for the gophers](https://medium.zenika.com/go-1-11-webassembly-for-the-gophers-ae4bb8b1ee03)
    Good explanation on different API methods of `js.Value`

**How to install Go**

https://golang.org/doc/install

const go = new Go();

let inst = null;

WebAssembly
    .instantiateStreaming(fetch("main.wasm"), go.importObject)
    .then((result) => {
        inst = result.instance;
    });

document
    .getElementById("runButton")
    .addEventListener('click', () => {
        go.run(inst);
    });
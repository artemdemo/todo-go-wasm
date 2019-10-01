const go = new Go();

let inst = null;

WebAssembly
    .instantiateStreaming(fetch('main.wasm'), go.importObject)
    .then((result) => {
        go.run(result.instance);
    });

document
    .getElementById('runButton')
    .addEventListener('click', () => {
        console.log(addToDo('test title', true));
    });
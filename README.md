# TinyCanvas
TinyGo bindings to use the HTML5 Canvas - convenience utility basically for myself.

=== All of this is still very much WIP. ===

Basically, I often use WASM to do some data visualization work and I find that
working with TinyGo and the HTML5 Canvas to be really handy. This is a collection
of utilities / tools I often use, kind of cleaned up in case anyone else finds use in them.

## Demo
There's a small live demo at <https://ewaldhorn.github.io/tinycanvas/> if you like.

### TinyCanvas
Contains the canvas components along with some Pixel utilities.

### DOM
Contains some DOM utilities, handy for this type of work.

### Zed
I use <https://zed.dev/> as my main editor most of the time. There's a convenient `zed.sh`
file included to start it with the right environment variables to enable code completion etc.

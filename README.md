# TinyCanvas v0.0.8
TinyGo-based in-memory canvas that can output to the HTML5 Canvas.

It's not entirely based on calling the HTML5 Canvas functions directly, but rather
creates an in-memory buffer that gets copied and displayed by the canvas. This tool
was built to help migrate old Turbo Pascal programs to something a bit more modern,
with the intention of sticking as closesly to the old output as possible.

<br>
=== All of this is still very much WIP. ===
<br><br>

This is a collection of utilities / tools I use, kind of cleaned up in case anyone
else finds a use for them.

## Demo
There's a small live demo at <https://ewaldhorn.github.io/tinycanvas/> if you like.

### TinyCanvas
Contains the canvas components along with some Pixel utilities. You can also get direct
access to the underlying graphics context if you look carefully...

### Colour
Contains colour components. I ended up extracting it because it can be useful on its own.

### Zed
I use <https://zed.dev/> as my main editor most of the time. There's a convenient `zed.sh`
file included to start it with the right environment variables to enable code completion etc.

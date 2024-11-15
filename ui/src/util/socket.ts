export let socket = new WebSocket("ws://localhost:9091/ws");

socket.onclose = (event) => {
    socket = new WebSocket("ws://localhost:9091/ws")
}

socket.onerror = (event) => {
    socket = new WebSocket("ws://localhost:9091/ws")
}
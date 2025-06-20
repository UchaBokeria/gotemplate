var socket = new WebSocket("ws://localhost:9999/");

socket.onopen = function(event) {
    console.log("HotReload connection established.");
};

socket.onmessage = function(event) {
    if (event.data === "reload") {
        console.log("Reloading page...");
        location.reload();
    }
};

socket.onclose = function(event) {
    console.log("WebSocket connection closed.");
};
const WebSocket = require('ws');

const ws = new WebSocket("ws://localhost:8080/api/v1/ws");

ws.on('open', () => console.log("Connected to WS"));
ws.on('message', (msg) => console.log("Received:", msg.toString()));
ws.on('close', () => console.log("Disconnected"));
ws.on('error', (err) => console.error("WS error:", err));

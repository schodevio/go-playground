---
url: https://www.youtube.com/watch?v=JuUAEYLkGbM&list=PL0xRBLFXXsP7-0IVCmoo2FEWBrQzfH2l8&index=11
title: How To Build A Chat And Data Feed With WebSockets In Golang?
description: In this Golang tutorial, I'm going to teach you how to build a simple chat application and data feed by building our own custom WebSockets server in Golang.
---

# WebSocket Chat & Data Feed Server

A simple WebSocket server in Go that supports a multi-client chat and a real-time orderbook data feed.

## Features

- **Chat** (`/ws`): Clients connect and broadcast messages to all other connected clients.
- **Orderbook feed** (`/orderbookfeed`): Pushes simulated orderbook data to connected clients every 2 seconds.

## Running

```bash
go run main.go
```

Server listens on `:8080`.

## Endpoints

| Endpoint | Description |
|---|---|
| `ws://localhost:8080/ws` | Chat — messages are broadcast to all connected clients |
| `ws://localhost:8080/orderbookfeed` | Orderbook feed — streams timestamped data every 2s |

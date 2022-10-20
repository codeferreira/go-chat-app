const socket = new WebSocket("ws://localhost:3333/ws")

function connect() {
  socket.onopen = () => {
    console.log("Succesfully Connected!")
  }

  socket.onmessage = (msg) => {
    console.log(msg)
  }

  socket.onclose = (event) => {
    console.log("Socket Closed Connection", event)
  }

  socket.onerror = (error) => {
    console.log("Socket Error: ", error)
  }
}

function sendMessage(message) {
  console.log("Sending message: ", message)

  socket.send(message)
}

export { connect, sendMessage }
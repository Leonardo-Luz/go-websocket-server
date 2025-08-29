const socket = new WebSocket('ws://localhost:8080/ws');

socket.onmessage = (event) => {
  alert('received from the server: ' + event.data);
}

const clickHandler = () => {
  const messageEl = document.getElementsByName('message')[0];

  if (messageEl) {
    const message = messageEl.value;
    socket.send(message);
  } else {
    alert("Write your message before sending it...")
  }
}


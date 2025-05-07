const socket = new WebSocket('ws://localhost:8080/ws');

socket.onmessage = (event) => {
	alert('received from the server: ' + event.data);
}

const clickHandler = () => {
	socket.send('hello from client!');
}


const express = require('express');
const app = express();
const server = require('http').createServer(app);

app.use(express.static(__dirname));

const io = require('socket.io')(server);
io.on('connection', (socket) => {
  console.log('New client connected');
  var userName;
  socket.on('joined', (who) => {
    userName = who;
    console.log(`${userName} joined`);
    io.emit('message', `System: ${who} joined the chat`);
  });
  socket.on('message', (msg) => {
    console.log(`Received message from ${userName}`);
    io.emit('message', `${userName}: ${msg}`);
  });
});

server.listen(3000);
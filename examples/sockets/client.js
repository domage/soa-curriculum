var net = require('net');

var client = new net.Socket();
client.connect(5454, '127.0.0.1', function() {
	console.log('Connected');

	client.write(JSON.stringify({
        Qwe: "hello from Node"
    }));
});

client.on('data', function(data) {
	console.log('Received: ' + data);
	client.destroy(); // kill client after server's response
});

client.on('close', function() {
	console.log('Connection closed');
});
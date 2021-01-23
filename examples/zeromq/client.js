const zmq = require("zeromq")

const readline = require("readline");
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

async function run() {
  const sock = new zmq.Request

  sock.connect("tcp://127.0.0.1:8089")
  console.log("Producer bound to port 8089")

  rl.question("String to reverse ", async str => {
      console.log(`Sending ${str}`);
      await sock.send(str);
      const [result] = await sock.receive();

      console.log(result.toString())

      process.exit(0);
  });
}

run()

const zmq = require("zeromq")

const reverse = (s) => {
    return s.split("").reverse().join("");
}

async function run() {
  const sock = new zmq.Reply

  await sock.bind("tcp://127.0.0.1:8089")

  for await (const [msg] of sock) {
    await sock.send(reverse(`${msg}`))
  }
}

run()

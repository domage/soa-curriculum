var PROTO_PATH = __dirname + "/service.proto";

var grpc = require("@grpc/grpc-js");
var protoLoader = require("@grpc/proto-loader");
var packageDefinition = protoLoader.loadSync(PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});
var serviceProto = grpc.loadPackageDefinition(packageDefinition);

/**
 * Implements the SayHello RPC method.
 */
function sayHello(call, callback) {
  callback(null, { message: "Hello " + call.request.name });
}

function sayAnotherHello(call, callback) {
  callback(null, { message: "Hello ANOTHER " + call.request.name });
}

/**
 * Starts an RPC server that receives requests for the Greeter service at the
 * sample server port
 */
function main() {
  var server = new grpc.Server();

  server.addService(serviceProto.Greeter.service, { sayHello, sayAnotherHello });
  server.bindAsync(
    "0.0.0.0:5050",
    grpc.ServerCredentials.createInsecure(),
    () => {
      server.start();
    }
  );
}

main();

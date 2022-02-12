const fs = require("fs");
const yaml = require("yaml");
const msgpack = require("msgpack");
var avro = require("avro-js");
var protobuf = require("protobufjs");

const main = async () => {
  const json = fs.readFileSync("./json.txt");
  const parsed = JSON.parse(json);

  parsed['id'] = 2;
  console.log(parsed);

  const yamlFile = fs.readFileSync("./yaml.txt");
  const parsedYaml = yaml.parse(yamlFile.toString());

  console.log(parsedYaml);
  parsedYaml.ID = "asd";
  console.log(parsedYaml);

  const binary = fs.readFileSync("./msgpack.txt");
  var msg = msgpack.unpack(binary);

  console.log(msg);
  msg.ID = "asd";
  console.log(msg);

  var type = avro.parse("./schema.avsc");
  var obj = type.fromBuffer(fs.readFileSync("./avro.txt"));

  console.log("From avro: ", obj);

  protobuf.load("./test.proto").then(function (root) {
    var Person = root.lookupType("Person");
    var message = Person.decode(fs.readFileSync("./pb.txt"));

    console.log("Protobuf: ", message);
  });
};

main();

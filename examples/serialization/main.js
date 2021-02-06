const fs = require("fs");
const yaml = require("yaml");
const msgpack = require('msgpack');

const main = async () => {
    const json = fs.readFileSync("./json.txt");
    const parsed = JSON.parse(json);

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
}

main();

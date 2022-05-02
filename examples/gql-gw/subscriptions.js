const { Kafka, CompressionTypes, CompressionCodecs } = require("kafkajs");
const { PubSub } = require("apollo-server");
const SnappyCodec = require("kafkajs-snappy");

CompressionCodecs[CompressionTypes.Snappy] = SnappyCodec;

const pubsub = new PubSub();

const kafka = new Kafka({
  clientId: "gateway",
  brokers: ["localhost:29092"],
});

const consumer = kafka.consumer({ groupId: "gateway" });

const KafkaPublicTopic = "public";

const run = async () => {
  // Consuming
  await consumer.connect();
  await consumer.subscribe({ topic: KafkaPublicTopic });

  await consumer.run({
    eachMessage: async ({ topic, partition, message }) => {
      const key = message.key.toString();
      const value = message.value.toString();

      try {
        pubsub.publish(key, { ...JSON.parse(value) });
      } catch (err) {
        console.error("Unable to parse response as JSON", key, value);
      }
    },
  });
};

const getRegistry = (name) => registries[name];

const listenSubscriptions = () => {
  run().catch(console.error);
};

module.exports = {
  pubsub,
  listenSubscriptions,
  getRegistry,
};

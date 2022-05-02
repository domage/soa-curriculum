const { Kafka } = require("kafkajs");

const reverse = (s) => {
  return s.split("").reverse().join("");
};

const kafka = new Kafka({
  clientId: "gateway",
  brokers: ["localhost:29092"],
});
const producer = kafka.producer();

const order = {
  id: 1,
  name: "Test order",
  products: [
    {
      id: 1,
    },
  ],
};

const resolvers = {
  Query: {
    order: () => order,
  },
  Mutation: {
    createOrder: async () => {
      await producer.connect();
      await producer.send({
        topic: "public",
        messages: [
          {
            key: "orderCreated",
            value: JSON.stringify({ orderCreated: { ...order } }),
          },
        ],
      });

      return order;
    },
  },
};

module.exports = { resolvers };

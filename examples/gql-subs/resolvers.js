const { PubSub } = require("apollo-server");
const pubsub = new PubSub();

const reverse = (s) => {
  return s.split("").reverse().join("");
};

function wait(milliseconds) {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}

const resolvers = {
  Query: {
    test: async () => {
      await wait(3000);
      return "test string";
    },
  },
  Mutation: {
    reverse: async (_, { str }) => {
      const msg = {
        string: reverse(str),
        madeAt: new Date(),
      };

      pubsub.publish("STRING_REVERSED", { stringReversed: msg });

      return msg;
    },
  },
  Subscription: {
    stringReversed: {
      subscribe: () => pubsub.asyncIterator(["STRING_REVERSED"]),
    },
  },
};

module.exports = { resolvers };

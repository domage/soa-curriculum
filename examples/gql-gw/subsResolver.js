const { withFilter } = require("graphql-subscriptions");
const { pubsub } = require("./subscriptions");
const { request } = require('graphql-request');

const resolvers = {
  Query: {
    dummy: () => "subs",
  },
  Order: {
    products: async (order) => {
      const result = await request(
        "http://localhost:5554/graphql",
        "query { productsByOrder(orderID: " +
          order.id +
          ") { id name price } }"
      );

      console.log('Resolverd:', result);

      return result.productsByOrder;
    },
  },
  Subscription: {
    orderCreated: {
      subscribe: withFilter(
        () => pubsub.asyncIterator("orderCreated"),
        (payload, variables) => {
          console.log(payload);

          if (!variables.name) {
            return true;
          }
          const idMatched =
            `${payload.orderCreated.data.name}` === variables.name;

          return idMatched;
        }
      ),
    },
  },
};

module.exports = { resolvers };

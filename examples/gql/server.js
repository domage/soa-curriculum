const { ApolloServer } = require("apollo-server-express");
const express = require("express");
const { json } = require("body-parser");
const {
  calculateCost,
  costDirective,
  extractCost,
} = require("@pipedrive/graphql-query-cost");
const DataLoader = require("dataloader");

const app = express();

const { typeDefs } = require("./schema");
const { resolvers } = require("./resolvers");

const batchUsersFetch = (keys) => {
  console.log("Keys to load:", keys);
  return Promise.resolve(
    keys.map((key) => {
      return {
        id: key,
        firstName: "John",
        lastName: "Doe",
        linkedUsers: []
      };
    })
  );
};

const { costMap, cleanSchema } = extractCost(typeDefs);

(async () => {
  const server = new ApolloServer({
    typeDefs: cleanSchema,
    resolvers,
    dataSources: () => {
      return {
        usersLoader: new DataLoader((keys) => batchUsersFetch(keys)),
      };
    },
  });
  const router = express.Router();

  router.use(json());
  router.use((req, res, next) => {
    const costLimitPerOperation = {
      defaultCost: 1,
      maximumCost: 5000,
    };

    const { query, variables, operationName } = req.body;
    const cost = calculateCost(query, cleanSchema, {
      defaultCost: costLimitPerOperation.defaultCost,
      costMap,
      variables,
    });

    console.log("Query complexity:", cost);

    next();
  });

  await server.start();

  // The `listen` method launches a web server.
  app.use(router);
  server.applyMiddleware({ app });
  app.listen(4000, () => {
    console.log(`🚀 Server ready at http://localhost:4000`);
  });
})();

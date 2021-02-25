const { ApolloServer } = require("apollo-server");
const { buildFederatedSchema } = require('@apollo/federation');

const { typeDefs } = require("./schema");
const { resolvers } = require("./resolvers");

const server = new ApolloServer({
  schema: buildFederatedSchema([{ typeDefs, resolvers }]),
});

// The `listen` method launches a web server.
server.listen(5656).then(({ url }) => {
  console.log(`🚀  Server ready at ${url}`);
});

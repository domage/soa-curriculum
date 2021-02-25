const { gql } = require('apollo-server');

const typeDefs = gql`
  extend type Product @key(fields: "id") {
    id: ID! @external
  }

  type Order {
    id: ID!
    name: String
    products: [Product]
  }

  type Query {
    order: Order
  }
`;

module.exports = { typeDefs };

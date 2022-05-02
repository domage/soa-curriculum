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

  input ProductInput {
    id: ID!
  }

  input OrderInput {
    name: String
    products: [ProductInput]
  }

  type Message {
    string: String
    madeAt: String
  }

  type Query {
    order: Order
  }

  type Mutation {
    createOrder(input: OrderInput): Order
  }
`;

module.exports = { typeDefs };

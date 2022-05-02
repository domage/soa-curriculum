const { gql } = require('apollo-server');

const typeDefs = gql`
  type Message {
    string: String
    madeAt: String
  }

  type Product @key(fields: "id") {
    id: ID!
    name: String
    price: Int
  }

  type Query {
    product: Product
    productsByOrder(orderID: ID!): [Product]
  }

  type Mutation {
    reverse(str: String!): Message
  }
`;

module.exports = { typeDefs };

const { gql } = require('apollo-server');

const typeDefs = gql`
  type Product {
    id: ID!
    name: String
    price: Int
  }

  type Order {
    id: ID!
    name: String
    products: [Product]
  }

  type Query {
    dummy: String
  }

  type Subscription {
    orderCreated(name: String): Order
  }
`;

module.exports = { typeDefs };

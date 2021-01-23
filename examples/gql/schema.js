const { gql } = require('apollo-server');

const typeDefs = gql`
  type Message {
    string: String
    madeAt: String
  }

  type Query {
    test: String
  }

  type Mutation {
    reverse(str: String!): Message
  }
`;

module.exports = { typeDefs };

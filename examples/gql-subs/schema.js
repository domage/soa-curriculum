const { gql } = require('apollo-server');

const typeDefs = gql`
  type Message {
    string: String
    madeAt: String
  }

  type ReverseMessage {
    madeAt: String
    string: String
  }

  type Query {
    test: String
    string2(str: String): String
  }

  type Mutation {
    reverse(str: String!): Message
  }

  type Subscription {
    stringReversed: ReverseMessage
  }
`;

module.exports = { typeDefs };

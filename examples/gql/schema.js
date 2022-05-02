const { gql } = require('apollo-server');
const queryCost = require('@pipedrive/graphql-query-cost');

// const typeDefs = gql`
const typeDefs = `
  ${queryCost.costDirective}
  type Message {
    string: String
    madeAt: String
  }

  type User {
    id: ID!
    firstName: String
    lastName: String
    linkedUsers: [User] @cost(db: 10)
    linkedUsersPaged(amount: Int!): [User]
  }

  type Query {
    test: String
    user(id: ID!): User @cost(
      complexity: 10
    )
    usersByID(ids: [ID!]!): [User]
  }

  type Mutation {
    reverse(str: String!): Message
  }
`;

module.exports = { typeDefs };

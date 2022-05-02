const reverse = (s) => {
  return s.split("").reverse().join("");
};

function wait(milliseconds) {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}

const resolvers = {
  User: {
    linkedUsers: (user, _, { dataSources }) => {
        return [dataSources.usersLoader.load(1)]
    },
    linkedUsersPaged: (user, { amount }, { dataSources }) => {
        return Array(amount).fill(0).map((_, i) => dataSources.usersLoader.load(i));
    },
  },
  Query: {
    test: () => "test string",
    user: (_, { id }, { dataSources }) => dataSources.usersLoader.load(id),
    usersByID: (_, { ids }, { dataSources }) => [dataSources.usersLoader.load(0)],
  },
  Mutation: {
    reverse: async (_, args) => {
      const msg = {
        string: reverse(args.str),
        madeAt: new Date(),
      };

      await wait(3000);
      return msg;
    },
  },
};

module.exports = { resolvers };

const reverse = (s) => {
  return s.split("").reverse().join("");
};

function wait(milliseconds) {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}

const product = {
  id: 1,
  name: "Test product",
  price: 123,
};

const resolvers = {
  Product: {
    __resolveReference(user, { fetchUserById }) {
      return product;
    },
  },
  Query: {
    product: () => product,
  },
  Mutation: {
    reverse: async (_, { str }) => {
      const msg = {
        string: reverse(str),
        madeAt: new Date(),
      };

      return msg;
    },
  },
};

module.exports = { resolvers };

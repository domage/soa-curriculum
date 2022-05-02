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
    __resolveReference(user, params) {
      console.log(user, params);
      return product;
    },
  },
  Query: {
    product: () => product,
    productsByOrder: async (root, { orderID }, { fetchProductsByOrder }) => {
      return [product];
    }
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

const reverse = (s) => {
  return s.split("").reverse().join("");
};

function wait(milliseconds) {
  return new Promise((resolve) => setTimeout(resolve, milliseconds));
}

const order = {
  id: 1,
  name: "Test order",
  products: [
    {
      id: 1
    }
  ]
};

const resolvers = {
  Query: {
    order: () => order,
  },
};

module.exports = { resolvers };

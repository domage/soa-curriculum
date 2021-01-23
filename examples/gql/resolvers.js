const reverse = (s) => {
    return s.split("").reverse().join("");
}

function wait(milliseconds) {
    return new Promise(resolve => setTimeout(resolve, milliseconds));
}

const resolvers = {
    Query: {
      test: () => 'test string',
    },
    Mutation: {
        reverse: async (_, args) => {
            const msg = {
                string: reverse(args.str),
                madeAt: new Date()
            };

            await wait(3000);
            return msg;            
        }
    }
};

module.exports = { resolvers };
  
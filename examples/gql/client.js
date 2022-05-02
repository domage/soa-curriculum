const fetch = require('cross-fetch');
const { ApolloClient, InMemoryCache, createHttpLink } = require('@apollo/client');

const { TEST_QUERY, REVERSE_MUTATION } = require('./gql');

const cache = new InMemoryCache();
const link = createHttpLink({
  uri: 'http://localhost:4000/graphql',
  fetch
});

const client = new ApolloClient({
  // Provide required constructor fields
  cache: cache,
  link: link
});

const run = async () => {
    const res = await client.query({
        query: TEST_QUERY
    });

    console.log(TEST_QUERY);
    console.log(res);

    const res2 = await client.mutate({
        mutation: REVERSE_MUTATION,
        variables: {
            str: 'wololo'
        }
    });

    console.log(res2);
}

run();

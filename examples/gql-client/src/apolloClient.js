import {
    ApolloClient,
    ApolloLink,
    HttpLink,
    InMemoryCache,
    split
} from '@apollo/client';
import { SubscriptionClient } from 'subscriptions-transport-ws';
import { WebSocketLink } from '@apollo/client/link/ws';
import { onError } from '@apollo/client/link/error';
import { getMainDefinition } from '@apollo/client/utilities';

const webSocketClient = new SubscriptionClient(
    `ws://localhost:5656/graphql`,
    {
        reconnect: true,
        timeout: 10000, 
    }
);
const webSocketLink = new WebSocketLink(webSocketClient);

const httpLink = new HttpLink({
    uri: `http://localhost:5656/graphql`
});

const errorLink = onError(({ graphQLErrors, networkError }) => {
    // Send Error stats back to server
    if (graphQLErrors) {
        console.log('graphQLErrors', graphQLErrors);
    }

    if (networkError) {
        console.log(`[Network error]: ${networkError}`);
    }
});

const link = split(
    ({ query }) => {
        const definition = getMainDefinition(query);
        return (
            definition.kind === 'OperationDefinition' &&
            definition.operation === 'subscription'
        );
    },
    webSocketLink,
    ApolloLink.from([errorLink, httpLink])
);

export const apolloClient = new ApolloClient({
    link,
    defaultOptions: {
        query: {
            fetchPolicy: 'network-only'
        }
    },
    cache: new InMemoryCache()
});

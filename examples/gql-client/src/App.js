import { ApolloProvider } from "@apollo/client";

import { apolloClient } from './apolloClient';
import { StringReserver } from "./StringReverser";

import 'antd/dist/antd.css';

function App() {
  return (
    <ApolloProvider client={apolloClient}>
      <StringReserver />
    </ApolloProvider>
  );
}

export default App;

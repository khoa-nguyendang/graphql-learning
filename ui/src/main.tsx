import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import React from "react";
import ReactDOM from "react-dom/client";
import { env } from '../env';
import App from "./App.tsx";
import "./index.css";

const client = new ApolloClient({
  uri: `${env.REACT_APP_HOST_IP_ADDRESS}/graphql`,
  cache: new InMemoryCache(),
});
ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <App />
    </ApolloProvider>
  </React.StrictMode>
);

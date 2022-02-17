import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';
import ApiClientRest from './ApiClient/ApiClientRest'
import reportWebVitals from './reportWebVitals';

import { ApiBuilderWorker } from './MockApi/worker';
import { MockApiServer } from './MockApi/api';
import mockDbFabric from './MockApi/db';
import mockHandlersFabric from './MockApi/handlers';

import './index.css';

// TODO: start mockApiServer only if in debug environment
// TODO: use deferred mounting in debug environment
const mockApiServer = new MockApiServer(new ApiBuilderWorker(), mockDbFabric, mockHandlersFabric);
mockApiServer.start();

const apiClient = new ApiClientRest();

ReactDOM.render(
  <React.StrictMode>
    <App apiClient={apiClient} />
  </React.StrictMode>,
  document.getElementById('root'),
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();

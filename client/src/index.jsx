import React from 'react';
import ReactDOM from 'react-dom';

import App from './App';
import ApiClientRest from './ApiClient/ApiClientRest'
import reportWebVitals from './reportWebVitals';

import './index.css';

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

import React from 'react';
import ReactDOM from 'react-dom';
import { Provider } from 'react-redux';

import { BrowserRouter } from 'react-router-dom';
import axios from 'axios';

import * as serviceWorker from './serviceWorker';
import App from './components/placeholders/App';
import storeFactory from './store';

import './index.css';
import 'bootstrap/dist/css/bootstrap.css';

axios.defaults.baseURL = process.env.KN_BE_API;
axios.defaults.timeout = 60 * 1000; // 60 seconds
const store = storeFactory();

window.store = store;

ReactDOM.render(
  <Provider store={store}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </Provider>,
  document.getElementById('root'),
);

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();

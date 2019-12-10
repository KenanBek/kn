import React from 'react';
import './App.css';
import Navigation from './Navigation';
import Main from './Main';
import Errors from '../containers/Errors'

function App() {
  return (
    <div>
      <Errors />
      <Navigation />
      <Main />
    </div>
  );
}

export default App;

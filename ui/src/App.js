import React from 'react';
import Dashboard from './views/Dashboard'
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
       <Dashboard/>
      </header>
    </div>
  );
}

export default App;

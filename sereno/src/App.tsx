import React from 'react';
import pfp from './static/pfp.jpg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={pfp} className="pfp" alt="logo" />
        <p>
          Sereno Miguel Dominguez
        </p>
      </header>
    </div>
  );
}

export default App;

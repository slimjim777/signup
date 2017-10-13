import React, { Component } from 'react';
import Form from './Form';
import './App.css';

class App extends Component {
  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src="/static/images/football.png" className="App-logo" alt="logo" />
          <h1 className="App-title">Monday Night Football Sign-up</h1>
        </header>
        <Form />
      </div>
    );
  }
}

export default App;

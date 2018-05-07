import React, { Component } from 'react';
import Event from './Event';
import Form from './Form';
import {getSection} from './common'
import './App.css';
import 'react-datepicker/dist/react-datepicker.css';

class App extends Component {

  routes (section) {
    switch(section) {
      case 'core':
        return <Event />
      default:
        return <Form />
    }
  }

  render() {
    var section = getSection();

    return (
      <div className="App">
        <header className="App-header">
          <h1 className="App-title">Life Church Worship</h1>
        </header>

        {this.routes(section)}

      </div>
    );
  }
}

export default App;

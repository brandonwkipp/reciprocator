import React, { Component } from 'react';

const MidiParser = require('midi-parser-js');

MidiParser.debug = true;
class Engine extends Component {
  constructor(props) {
    super(props);
    this.state = {
      json: {},
    };

    this.parse = this.parse.bind(this);
  }

  parse() {
    MidiParser.parse(document.getElementById('filereader'), (obj) => {
      this.setState({
        json: obj,
      });
    });
  }

  render() {
    const { json } = this.state;
    console.log(json);

    return (
      <>
        <input
          id="filereader"
          type="file"
          onInput={() => this.parse()}
        />
      </>
    );
  }
}

export default Engine;

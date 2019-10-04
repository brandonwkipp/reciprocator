import React, { Component } from 'react';
import { FormGroup, Input } from 'reactstrap';

const MidiParser = require('midi-parser-js');

MidiParser.debug = true;
class Engine extends Component {
  constructor(props) {
    super(props);
    this.state = {
      key: 0,
      json: {},
      notes: [],
    };

    this.findNegative = this.findNegative.bind(this);
    this.parse = this.parse.bind(this);
  }

  findNegative() {
    const { key } = this.state;
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
        <FormGroup>
          <input
            id="filereader"
            type="file"
            onInput={() => this.parse()}
          />
        </FormGroup>
        <FormGroup>
          <Input
            type="select"
            name="keySelect"
            id="keySelect"
          >
            <option value={0}>C</option>
            <option value={1}>C♯ / D♭</option>
            <option value={2}>D</option>
            <option value={3}>D♯ / E♭</option>
            <option value={4}>E</option>
            <option value={5}>F</option>
            <option value={6}>F♯ / G♭</option>
            <option value={7}>G</option>
            <option value={8}>G♯ / A♭</option>
            <option value={9}>A</option>
            <option value={10}>A♯ / B♭</option>
            <option value={11}>B</option>
          </Input>
        </FormGroup>
      </>
    );
  }
}

export default Engine;

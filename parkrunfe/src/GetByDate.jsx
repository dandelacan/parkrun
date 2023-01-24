import React, { useState } from "react";
import "react-day-picker/dist/style.css";

class GetByDate extends React.Component {
  constructor() {
    super();
    this.state = {
      date: "",
    };
  }

  handleChange = (date) => {
    this.setState({ date: date });
  };

  render() {
    const path = `https://harkinpark.run?date=`;
    return (
      <div>
        <label for="date">Select date:</label>
        <input
          onChange={(e) => this.handleChange(e.target.value)}
          type="date"
          name="date"
          id="date"
        ></input>
        <a href={path + this.state.date}>Get Results</a>
      </div>
    );
  }
}

export default GetByDate;

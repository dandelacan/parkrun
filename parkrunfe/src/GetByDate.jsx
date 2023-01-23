import React, { useState } from "react";
import "react-day-picker/dist/style.css";
// import { useNavigate } from "react-router-dom";

class GetByDate extends React.Component {
  constructor() {
    super();
    this.state = {
      date: "",
    };
  }

  // navigate = useNavigate()
  // routeChange = ()=>{
  //     const path = `http://ec2-13-41-189-141.eu-west-2.compute.amazonaws.com:3000/times?date=${this.state.date}`
  //     this.navigate(path)
  // }

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

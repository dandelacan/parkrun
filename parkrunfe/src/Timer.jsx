import React from "react";
var TimeFormat = require("hh-mm-ss");

const SECOND = 1000;
const MINUTE = SECOND * 60;
const HOUR = MINUTE * 60;

class Timer extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      timeElapsed: 0,
      startTime: 0,
      startAt: 0,
      timerStarted: false,
      intervalID: null,
    };
  }

  componentDidMount() {
    const startTime = parseInt(localStorage.getItem("time"));
    const timerStarted = localStorage.getItem("timerStarted") === "true";
    if (timerStarted) this.startTimer();
    this.setState({ startTime, timerStarted });
  }

  setLocalStorageTime(time) {
    localStorage.setItem("time", time.toString());
    localStorage.setItem("timerStarted", "true");
  }

  startTimer() {
    const intervalID = setInterval(
      () =>
        this.setState({
          timeElapsed: Date.now() - this.state.startTime,
        }),
      1000
    );
    this.setState({ intervalID });
  }

  handleStartClick = async () => {
    await this.setState({ startTime: Date.now(), timerStarted: true });
    this.setLocalStorageTime(this.state.startTime);
    this.startTimer();
  };

  handleStartClickAt = async () => {
    await this.setState({
      startTime: Date.now() - this.state.startAt,
      timerStarted: true,
    });
    this.setLocalStorageTime(this.state.startTime);
    this.startTimer();
  };

  handleChaneStartAt = (time) => {
    const timeInMS = TimeFormat.toMs(time);
    this.setState({ startAt: timeInMS });
  };

  handleResetTimer = () => {
    clearInterval(this.state.intervalID);
    this.setState({ timeElapsed: 0, timerStarted: false });
    localStorage.setItem("timerStarted", "false");
  };

  render() {
    return (
      <div>
        <div className="timer">
          {Object.entries({
            Hours: (this.state.timeElapsed / HOUR) % 24,
            Minutes: (this.state.timeElapsed / MINUTE) % 60,
            Seconds: (this.state.timeElapsed / SECOND) % 60,
          }).map(([label, value]) => (
            <div key={label} className="col-4">
              <div className="box">
                <p>{`${Math.floor(value)}`.padStart(2, "0")}</p>
                <span className="text">{label}</span>
              </div>
            </div>
          ))}
        </div>
        <div>
          <button onClick={this.handleStartClick}>Start Timer</button>
          <button onClick={this.handleResetTimer}>Reset Timer</button>
        </div>
        <div>
          <button onClick={this.handleStartClickAt} a>
            Start Timer at:
          </button>
          <input
            type="time"
            step="1"
            onChange={(e) => this.handleChaneStartAt(e.target.value)}
          />
        </div>
      </div>
    );
  }
}

export default Timer;

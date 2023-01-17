import React from "react";


const SECOND = 1000;
const MINUTE = SECOND * 60;
const HOUR = MINUTE * 60;

class Timer extends React.Component{
    constructor(props) {
        super(props);
        this.state = { 
          timeElapsed: 0,
          startTime : null 

        };
       
      }  
        
      handleStartClick= ()=> {
        this.setState({startTime : Date.now()})
        setInterval( () =>
          this.setState({
            timeElapsed: Date.now() - this.state.startTime
          })
        , 1000);
      }

    render(){
        return (
          <div>
            <div className="timer">
                {Object.entries({
                    
                    Hours: (this.state.timeElapsed / HOUR) % 24,
                    Minutes: ( this.state.timeElapsed/ MINUTE) % 60,
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
            <button onClick={this.handleStartClick}>Start Timer</button>
            </div>
        );
    }
}

export default Timer
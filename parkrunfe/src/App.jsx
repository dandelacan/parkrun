import "./App.css";

import React from "react";
import Html5QrcodePlugin from "./Html5QrcodePlugin.jsx";
import ResultContainerPlugin from "./ResultContainerPlugin.jsx";
import Timer from "./Timer";
import GetByDate from "./GetByDate";

class App extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      decodedResults: [],
    };
    this.childRef = React.createRef();

    // This binding is necessary to make `this` work in the callback.
    this.onNewScanResult = this.onNewScanResult.bind(this);
    this.clearTimes = this.clearTimes.bind(this);
  }

  componentDidMount() {
    const decodedResults = JSON.parse(localStorage.getItem("results"));
    this.setState({ decodedResults });
  }

  render() {
    return (
      <div className="App">
        <section className="App-section">
          <div className="App-section-title"> Parkrun barcode scanner</div>
          <br />
          <br />
          <br />
          <Html5QrcodePlugin
            fps={10}
            qrbox={250}
            disableFlip={false}
            qrCodeSuccessCallback={this.onNewScanResult}
          />
          <Timer ref={this.childRef} />
          <ResultContainerPlugin
            results={this.state.decodedResults}
            clearTimes={this.clearTimes}
          />
          <GetByDate />
        </section>
      </div>
    );
  }

  async onNewScanResult(decodedText, decodedResult) {
    const childElement = this.childRef.current;
    const timeElapsed = childElement.state.timeElapsed;

    await this.setState((state, props) => {
      state.decodedResults.push({
        decodedResult: decodedResult,
        time: timeElapsed,
      });
      return state;
    });

    localStorage.setItem("results", JSON.stringify(this.state.decodedResults));
  }

  clearTimes() {
    this.setState({ decodedResults: [] });

    localStorage.setItem("results", []);
  }
}

export default App;

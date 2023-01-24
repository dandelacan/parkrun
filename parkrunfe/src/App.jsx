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
      decodedResults: [
        { decodedResult: { decodedText: "12634" }, time: 100000 },
        { decodedResult: { decodedText: "56878" }, time: 300000 },
      ],
    };
    this.childRef = React.createRef();

    // This binding is necessary to make `this` work in the callback.
    this.onNewScanResult = this.onNewScanResult.bind(this);
    this.clearTimes = this.clearTimes.bind(this);
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

  onNewScanResult(decodedText, decodedResult) {
    const childElement = this.childRef.current;
    const timeElapsed = childElement.state.timeElapsed;

    this.setState((state, props) => {
      state.decodedResults.push({
        decodedResult: decodedResult,
        time: timeElapsed,
      });
      return state;
    });
  }

  clearTimes() {
    this.setState({ decodedResults: [] });
  }
}

export default App;
import './App.css';

import React, { useRef } from 'react';
import Html5QrcodePlugin from './Html5QrcodePlugin.jsx'
import ResultContainerPlugin from './ResultContainerPlugin.jsx'
import Timer from './Timer';


class App extends React.Component {


  constructor(props) {
    super(props);
    this.state = {
      decodedResults: [{decodedResult:{decodedText:"12634"}, time:100000}, {decodedResult:{decodedText:"56878"}, time:300000}]
    }
    this.childRef = React.createRef()
   
    // This binding is necessary to make `this` work in the callback.
    this.onNewScanResult = this.onNewScanResult.bind(this);
    
  }


  render() {
    return (
      <div className="App">
        <section className="App-section">
          <div className="App-section-title"> Html5-qrcode React demo</div>
          <br />
          <br /> 
          <br />
          <Html5QrcodePlugin 
            fps={10}
            qrbox={250}
            disableFlip={false}
            qrCodeSuccessCallback={this.onNewScanResult}/>
            <Timer ref={this.childRef} />
          <ResultContainerPlugin results={this.state.decodedResults} />
          <button onClick={this.addManually}>test</button>
      
        </section>
      </div>
    );
  }

  onNewScanResult(decodedText, decodedResult) {
    const childElement = this.childRef.current
    console.log(
      "App [result]", decodedResult, childElement.state.timeElapsed);
      
    // let decodedResults = this.state.decodedResults;
    // decodedResults.push(decodedResult);
    this.setState((state, props) => {
      
      state.decodedResults.push({decodedResult: decodedResult, time: childElement.state.timeElapsed});
      return state;
    });
  }

  addManually(){
    const childElement = this.childRef.current
    const time = childElement.state.timeElapsed
    this.setState((state, props)=>{
      state.decodedResults.push({decodedResult:{decodedText:"test"}, time:time})
    })
  }



  
}

export default App;

import "./App.css";
import React from "react";
import submitTimes from "./httpFunction";

var TimeFormat = require("hh-mm-ss");

function filterResults(results) {
  let filteredResults = [];
  for (var i = 0; i < results.length; ++i) {
    if (i === 0) {
      filteredResults.push(results[i]);
      continue;
    }

    if (
      results[i].decodedResult.decodedText !==
      results[i - 1].decodedResult.decodedText
    ) {
      filteredResults.push(results[i]);
    }
  }
  return filteredResults;
}

class ResultContainerTable extends React.Component {
  render() {
    var results = filterResults(this.props.data);
    return (
      <div>
        <table className={"Qrcode-result-table"}>
          <thead>
            <tr>
              <td>#</td>
              <td>Barcode ID </td>
              <td>Time</td>
            </tr>
          </thead>
          <tbody>
            {results.map((result, i) => {
              console.log(result);
              return (
                <tr key={i}>
                  <td>{i}</td>
                  <td>{result.decodedResult.decodedText}</td>
                  <td>{TimeFormat.fromMs(result.time, "hh:mm:ss")}</td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    );
  }
}

class ResultContainerPlugin extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      statusMessage: "",
    };
  }
  clearTimes = this.props.clearTimes;
  async onSubmiit(results) {
    const status = await submitTimes(results);
    console.log(status);

    if (status == 200) {
      this.setState({ statusMessage: "sucsessfully uploaded" });
      this.clearTimes();
    } else {
      this.setState({ statusMessage: "Upload unsucsessfull please try again" });
    }
  }

  render() {
    let results = filterResults(this.props.results);
    return (
      <div className="Result-container">
        <div className="Result-header">Scanned results ({results.length})</div>
        <div className="Result-section">
          <ResultContainerTable data={this.props.results} />
          <button id="upload-button" onClick={() => this.onSubmiit(results)}>
            Upload times
          </button>
          <div>
            <p className="float-left">{this.state.statusMessage}</p>
          </div>
        </div>
      </div>
    );
  }
}

export default ResultContainerPlugin;

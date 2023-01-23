import axios from "axios";

async function submitTimes(results) {
  const timesArr = [];
  results.forEach((result) => {
    timesArr.push({
      Time: result.time.toString(),
      BarcodeID: result.decodedResult.decodedText,
    });
  });

  var config = {
    method: "post",
    url: "https://harkinpark.run",
    headers: {
      "Content-Type": "application/json",
    },
    data: JSON.stringify(timesArr),
  };

  const response = await axios(config);
  return response.status;
}

export default submitTimes;

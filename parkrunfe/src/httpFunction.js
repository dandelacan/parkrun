import axios from "axios"

async function submitTimes(results){
    const timesArr = []
    results.forEach(result=>{
        timesArr.push({Time: result.time.toString(), BarcodeID:result.decodedResult.decodedText})
    })

    var config = {
        method: 'post',
        url: 'http://ec2-13-41-189-141.eu-west-2.compute.amazonaws.com:3000/times',
        headers: { 
          'Content-Type': 'application/json'
        },
        data : JSON.stringify(timesArr)
      };

    const response = await axios(config)
}

export default submitTimes
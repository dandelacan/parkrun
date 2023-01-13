import { useEffect, useImperativeHandle, useMemo, useState } from "react";

const SECOND = 1000;
const MINUTE = SECOND * 60;
const HOUR = MINUTE * 60;

export const Timer = () => {
    
    const [time, setTime] = useState(0);
    
    const [start, setStart] = useState(false)
    const onClick = () => {
        setStart(true)
}
    useEffect(() => {
        let interval = null

        if(start){
            interval = setInterval(()=>{
                setTime(prevTime => prevTime + 10)
            }, 10)
        }else{
            clearInterval(interval)
        }
    
    }, [start]);



    return (
        <div className="timer">
            {Object.entries({
                Hours: (time / HOUR) % 24,
                Minutes: (time / MINUTE) % 60,
                Seconds: (time / SECOND) % 60,
            }).map(([label, value]) => (
                <div key={label} className="col-4">
                    <div className="box">
                        <p>{`${Math.floor(value)}`.padStart(2, "0")}</p>
                        <span className="text">{label}</span>
                    </div>
                </div>
            ))}
             <div>
                <button onClick={onClick}>Start Timer</button>
                </div>
        </div>
    );
};
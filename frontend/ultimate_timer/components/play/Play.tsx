import React from 'react';
import axios from 'axios';
import Box from '@material-ui/core/Box';
import { useTimer } from "react-timer-hook";
import presetURL from "../../config/settings";
import secondToMinute from "../../lib/second_to_minute";

// interface Props {
//   id: string;
// }

// interface iPreset {
//   id: string;
//   name: string,
//   display_order: number,
//   loop_count: number,
//   waits_confirm_each: boolean,
//   waits_confirm_last: boolean,
//   timer_unit: {
//     durations: number,
//     order: number,
//   },
// }

// type TimeObj = {
//   hour: number,
//   min: number,
//   sec: number,
// }

export const Play: React.FC<Props> = ({ id }) => {
//   const [preset, setPreset] = React.useState<iPreset>();
//   let interval: Timer = React.useRef();

//   const url: string = presetURL + id;
//   let remainingTime: string = '00:00:00';

//   const startTimer = (sec: number): void => {
//     interval = setInterval(() => {
//       const cvtedTime: object = secondToMinute(sec--);
//       const fmtedTime: string = String(cvtedTime.hour) + ':' + String(cvtedTime['min']) + ':' + String(cvtedTime['sec'])
//       remainingTime = fmtedTime
//     }, 1000)
//   }

//   React.useEffect(() => {
//     axios
//       .get<iPreset>(url)
//       .then((response) => {
//         setPreset(response.data);
//       });
//   }, []);

//   if (preset === null || preset === undefined) {
//     window.location.href = "/";
//   }
//   return (
//     <Box className="timer-container">
//       <h1>
//         {preset.name}
//       </h1>
//       <div>
//         {/* TODO want to show remaning sets instead */}
//         Loop: {preset?.loop_count} times
//       </div>
//       <div>
//         Waits Confirm for each timer unit: {preset?.waits_confirm_each === true ? "YES" : "NO"}
//       </div>
//       <div>
//         Waits Confirm for the last timer unit: {preset?.waits_confirm_last === true ? "YES" : "NO"}
//       </div>
//       <div>
//         {remainingTime}
//       </div>
//     </Box>
//   )
// }

// function MyTimer({ expiryTimestamp }: { expiryTimestamp: number }) {
  const {
    seconds,
    minutes,
    hours,
    days,
    isRunning,
    start,
    pause,
    resume,
    restart,
  } = useTimer({
    expiryTimestamp,
    onExpire: () => console.warn("onExpire called"),
  });

  return (
    <div style={{ textAlign: "center" }}>
      <h1>react-timer-hook </h1>
      <p>Timer Demo</p>
      <div style={{ fontSize: "100px" }}>
        <span>{days}</span>:<span>{hours}</span>:<span>{minutes}</span>:
        <span>{seconds}</span>
      </div>
      <p>{isRunning ? "Running" : "Not running"}</p>
      <button onClick={start}>Start</button>
      <button onClick={pause}>Pause</button>
      <button onClick={resume}>Resume</button>
      <button
        onClick={() => {
          // Restarts to 5 minutes timer
          const time = new Date();
          time.setSeconds(time.getSeconds() + 300);
          restart(time as unknown as number);
        }}
      >
        Restart
      </button>
    </div>
  );
}

export default function App() {
  const time = new Date();
  time.setSeconds(time.getSeconds() + 600); // 10 minutes timer
  return (
    <div>
      <MyTimer expiryTimestamp={time as unknown as number} />
    </div>
  );
}

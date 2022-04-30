import React from 'react';
import axios from 'axios';
import { useRouter, NextRouter } from 'next/router'
import { useTimer } from "react-timer-hook";
import presetURL from "../../config/settings";

interface Props {
  expiryTimestamp: Date;
  id: string;
}

interface iPreset {
  id: string;
  name: string,
  display_order: number,
  loop_count: number,
  waits_confirm_each: boolean,
  waits_confirm_last: boolean,
  timer_unit: {
    duration: number,
    order: number,
  }[],
}

export const Play: React.FC<Props> = ({ expiryTimestamp, id }) => {
  const [preset, setPreset] = React.useState<iPreset>();
  const router: NextRouter = useRouter()

  const url: string = presetURL + id;
  React.useEffect(() => {
    axios
      .get<iPreset>(url)
      .then((response) => {
        setPreset(response.data);
      })
      .catch((error) => {
        alert(error.message);
      });
  }, []);

  // React.useEffect(() => {
  //   if (preset?.timer_unit === undefined || preset?.timer_unit === null) {
  //     router.push("/");
  //   }
  // }, [])

  let remainingTime: Date = new Date();
  // TODO: fix ide error below
  // remainingTime.setSeconds(remainingTime.getSeconds() + preset?.timer_unit?.duration ?? 0);
  remainingTime.setSeconds(remainingTime.getSeconds() + 600);
  // let remainingTime: number = 600;

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
          const time = new Date();
          // time.setSeconds(time.getSeconds() + preset?.timer_unit.duration);
          time.setSeconds(time.getSeconds() + 600);
          restart(time);
        }}
      >
        Restart
      </button>
    </div>
  );
}

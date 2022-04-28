import React from 'react';
import axios from 'axios';
import Box from '@material-ui/core/Box';
import presetURL from "../../config/settings";
import secondToMinute from "../../lib/second_to_minute";

interface Props {
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
    durations: number,
    order: number,
  },
}

type TimeObj = {
  hour: number,
  min: number,
  sec: number,
}

export const Play: React.FC<Props> = ({ id }) => {
  const [preset, setPreset] = React.useState<iPreset>();
  let interval: Timer = React.useRef();

  const url: string = presetURL + id;
  let remainingTime: string = '00:00:00';

  const startTimer = (sec: number): void => {
    interval = setInterval(() => {
      const cvtedTime: object = secondToMinute(sec--);
      const fmtedTime: string = String(cvtedTime.hour) + ':' + String(cvtedTime['min']) + ':' + String(cvtedTime['sec'])
      remainingTime = fmtedTime
    }, 1000)
  }

  React.useEffect(() => {
    axios
      .get<iPreset>(url)
      .then((response) => {
        setPreset(response.data);
      });
  }, []);

  if (preset === null || preset === undefined) {
    window.location.href = "/";
  }
  return (
    <Box className="timer-container">
      <h1>
        {preset.name}
      </h1>
      <div>
        {/* TODO want to show remaning sets instead */}
        Loop: {preset?.loop_count} times
      </div>
      <div>
        Waits Confirm for each timer unit: {preset?.waits_confirm_each === true ? "YES" : "NO"}
      </div>
      <div>
        Waits Confirm for the last timer unit: {preset?.waits_confirm_last === true ? "YES" : "NO"}
      </div>
      <div>
        {remainingTime}
      </div>
    </Box>
  )
}

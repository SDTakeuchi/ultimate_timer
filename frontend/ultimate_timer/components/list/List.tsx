import React from 'react';
import axios from 'axios';
import presetURL from "../../config/settings";
import { TimerCard } from './Card';
import Box from '@material-ui/core/Box';

interface ResPresets {
  id: string;
  CreatedAt: Date;
  UpdatedAt: Date;
  name: string,
  display_order: number,
  loop_count: number,
  waits_confirm_each: boolean,
  waits_confirm_last: boolean,
  timer_unit: null,
}

export const TimerList: React.FC = () => {
  const names: string[] = ['Tabata Timer', '9min', '5.5min'];
  const [preset, setPreset] = React.useState(null);
  const url = presetURL;
  
  React.useEffect(() => {
    axios.get<ResPresets>(url).then((response) => {
      setPreset(response.data);
    });
  }, []);

  return (<div>
    <Box>
      {preset.map((value, _) => {
        return <TimerCard name={value} />
      })}
    </Box>
  </div>
)};

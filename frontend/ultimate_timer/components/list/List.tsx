import React from 'react';
import axios from 'axios';
// import presetURL from "../../config/settings";
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
  timer_unit?: {
    durations?: number,
    order?: number,
  },
}

export const TimerList: React.FC = () => {
  // const names: string[] = ['Tabata Timer', '9min', '5.5min'];
  const defaultProps: ResPresets[] = [];
  const [presets, setPresets] = React.useState<ResPresets[]>(defaultProps);
  const url = 'http://localhost/api/presets/';

  React.useEffect(() => {
    axios
      .get<ResPresets[]>(url)
      .then((response) => {
        console.log(response);
        setPresets(response.data);
      });
  }, []);

  if (presets) {
    return (<div>
      <Box>
        {presets.map((preset) => {
          return <TimerCard name={preset.name} />
        })}
      </Box>
    </div>
    )
  } else {
    return null;
  }
};

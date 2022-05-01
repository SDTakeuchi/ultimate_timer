import React from 'react';
import axios from 'axios';
import presetURL from "../../config/settings";
import { TimerCard } from './Card';
import Box from '@material-ui/core/Box';

interface iPresets {
  id: string;
  CreatedAt: Date;
  UpdatedAt: Date;
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

export const TimerList: React.FC = () => {
  const defaultProps: iPresets[] = [];
  const [presets, setPresets] = React.useState<iPresets[]>(defaultProps);
  const url: string = presetURL;

  React.useEffect(() => {
    axios
      .get<iPresets[]>(url)
      .then((response) => {
        setPresets(response.data);
      })
      .catch((error) => {
        alert(error.message);
      });
  }, []);

  presets.sort((a: presets, b: presets) => (a.display_order > b.display_order) ? 1 : -1);

  if (presets) {
    return (<div>
      <Box>
        {presets.map((preset) => {
          return <TimerCard
                    name={preset.name}
                    display_order={preset.display_order}
                    time={preset.timer_unit[0].duration}
                    id={preset.id}
                  />;
        })}
      </Box>
    </div>
    )
  } else {
    return null;
  }
};

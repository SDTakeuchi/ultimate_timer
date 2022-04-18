import React from 'react';
import axios from 'axios';
import presetURL from "../../config/settings";
import Box from '@material-ui/core/Box';

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

export const Play: React.FC = () => {
  const defaultProps: iPresets[] = [];
  const [presets, setPresets] = React.useState<iPresets[]>(defaultProps);
  const url: string = presetURL;

  React.useEffect(() => {
    axios
      .get<iPresets[]>(url)
      .then((response) => {
        setPresets(response.data);
      });
  }, []);

  return (
    <Box>
    </Box>
  )
}

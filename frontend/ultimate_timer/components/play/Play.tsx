import React from 'react';
import axios from 'axios';
import Box from '@material-ui/core/Box';
import presetURL from "../../config/settings";

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

export const Play: React.FC<Props> = ({ id }) => {
  const [preset, setPreset] = React.useState<iPreset>();
  const url: string = presetURL;

  React.useEffect(() => {
    axios
      .get<iPreset>(url)
      .then((response) => {
        setPreset(response.data);
      });
  }, []);

  return (
    <Box>
      
    </Box>
  )
}

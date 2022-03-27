import React from 'react';
import { TimerCard } from './Card';
import Box from '@material-ui/core/Box';

export const TimerList: React.FC = () => {
  const names: string[] = ['Tabata Timer', '9min', '5.5min'];
  
  return (<div>
    <Box>
      {names.map((value, _) => {
        return <TimerCard name={value} />
      })}
    </Box>
  </div>
)};

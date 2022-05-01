import React from 'react';
import axios from 'axios';
import { useRouter, NextRouter } from 'next/router'
import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import { makeStyles, createStyles, Theme } from '@material-ui/core/styles';
import IconButton from '@material-ui/core/IconButton';
import PlayCircleOutlineIcon from '@material-ui/icons/PlayCircleOutline';
import secondToMinute from '../../lib/second_to_minute';
import zeroPadding from '../../lib/zfill'
import presetURL from "../../config/settings";

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexDirection: 'column',
      alignItems: 'center',
      '& > *': {
        margin: theme.spacing(1),
      },
    },
  }),
);

interface Props {
  // expiryTimestamp: Date;
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
  // TODO: put interface below
  const [remainingTime, setRemainingTime] = React.useState<string>('00:00:00');
  const [isRunning, setIsRunning] = React.useState<boolean>(false);
  let interval: Timer = React.useRef();
  const classes = useStyles();

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


  const startTimer = (sec: number): void => {
    setIsRunning(true);
    interval = setInterval(() => {
      const cvtedTime: object = secondToMinute(sec--);
      const fmtedTime: string = 
        zeroPadding(cvtedTime['hour'], 2) + ':' + 
        zeroPadding(cvtedTime['min'], 2) + ':' + 
        zeroPadding(cvtedTime['sec'], 2);
      setRemainingTime(fmtedTime);
    }, 1000)
  }

  return (
    <div className="timer-container">
      <h1>
        {preset?.name}
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
      
      <div className={classes.root}>
        {/* <ButtonGroup color="primary" aria-label="outlined primary button group">
          <Button>One</Button>
          <Button>Two</Button>
          <Button>Three</Button>
        </ButtonGroup>
        <ButtonGroup variant="contained" color="primary" aria-label="contained primary button group">
          <Button>One</Button>
          <Button>Two</Button>
          <Button>Three</Button>
        </ButtonGroup> */}
        <ButtonGroup color="primary" aria-label="outlined primary button group">
          <Button
            disabled={isRunning}
            onClick={() => startTimer(preset?.timer_unit[0].duration)}>
            Play
          </Button>
          <Button>Pause</Button>
          {/* <Button>Resume</Button> */}
          <Button>Restart</Button>
        </ButtonGroup>
      </div>
    </div>
  )
}

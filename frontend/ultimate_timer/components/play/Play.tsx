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
// import playAudio from '../../lib/play_audio';
import { convertCompilerOptionsFromJson } from 'typescript';

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
  const [remainingTimeInt, setRemainingTimeInt] = React.useState<number>(0);
  const [isRunning, setIsRunning] = React.useState<boolean>(false);
  let interval: Timer = React.useRef();
  const classes = useStyles();

  const url: string = presetURL + id;
  React.useEffect(() => {
    axios
      .get<iPreset>(url)
      .then((response) => {
        setPreset(response.data);
        setRemainingTimeInt(response.data.timer_unit[0].duration);
        const cvtedTime_: object = secondToMinute(remainingTimeInt);
        const fmtedTime_: string = 
          zeroPadding(cvtedTime_['hour'], 2) + ':' + 
          zeroPadding(cvtedTime_['min'], 2) + ':' + 
          zeroPadding(cvtedTime_['sec'], 2);
        setRemainingTime(fmtedTime_);
      })
      .catch((error) => {
        alert(error.message);
      });
  }, []);

  const restartTimer = (): void => {
    setRemainingTimeInt(preset?.timer_unit[0].duration);
    setIsRunning(true);
  }
  
  const audioPlay = (): void => {
    const audio: HTMLAudioElement = new Audio('https://audio-previews.elements.envatousercontent.com/files/148785970/preview.mp3?response-content-disposition=attachment%3B+filename%3D%22RZFWLXE-bell-hop-bell.mp3%22');
    audio.play();
  }

  React.useEffect(() => {
    let interval_: Timer = null;

    if (isRunning) {
      interval_ = setInterval(() => {
        if (remainingTimeInt === 0) {
          setIsRunning(false);
          audioPlay();
          const cvtedTime: object = secondToMinute(preset?.timer_unit[0].duration);
          const fmtedTime: string = 
            zeroPadding(cvtedTime['hour'], 2) + ':' + 
            zeroPadding(cvtedTime['min'], 2) + ':' + 
            zeroPadding(cvtedTime['sec'], 2);
          setRemainingTime(fmtedTime);
          clearInterval(interval_);
        }
        setRemainingTimeInt(remainingTimeInt => remainingTimeInt - 1);
        const cvtedTime: object = secondToMinute(remainingTimeInt);
        const fmtedTime: string = 
          zeroPadding(cvtedTime['hour'], 2) + ':' + 
          zeroPadding(cvtedTime['min'], 2) + ':' + 
          zeroPadding(cvtedTime['sec'], 2);
        setRemainingTime(fmtedTime);
      }, 1000)
    } else {
      clearInterval(interval_);
    }
    return () => clearInterval(interval_);
  }, [isRunning, remainingTimeInt, remainingTime])

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
      <h2>
        {remainingTime}
      </h2>
      <div className={classes.root}>
        <ButtonGroup color="primary" aria-label="outlined primary button group">
          <Button
            disabled={isRunning}
            onClick={() => setIsRunning(true)} >
            Play
          </Button>
          <Button
            disabled={!isRunning}
            onClick={() => setIsRunning(false)} >
            Pause
          </Button>
          <Button
            disabled={isRunning}
            onClick={() => restartTimer()} >
            Restart
          </Button>
        </ButtonGroup>
      </div>
    </div>
  )
}

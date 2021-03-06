import Link from 'next/link'
import React from 'react';
import axios from 'axios';
import { makeStyles } from '@material-ui/core/styles';
import Card from '@material-ui/core/Card';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import Button from '@material-ui/core/Button';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import DeleteForeverIcon from '@material-ui/icons/DeleteForever';
import PlayCircleFilledWhiteIcon from '@material-ui/icons/PlayCircleFilledWhite';
import EditIcon from '@material-ui/icons/Edit';
import presetURL from '../../config/settings';
import { secondToMinute } from '../../lib/second_to_minute';
import zeroPadding from '../../lib/zfill';

import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';


const useStyles = makeStyles({
  root: {
    minWidth: 275,
    maxWidth: '100%',
  },
  title: {
    fontSize: 14,
  },
  pos: {
    marginBottom: 12,
  },
  alignRight: {
    marginLeft: 'auto',
  }
});

interface Props {
  name: string;
  display_order: number;
  time: number;
  id: string;
}

interface iDeletedPreset {
  error: string;
}

export const TimerCard: React.FC<Props> = ({ name, display_order, time, id }) => {
  const classes = useStyles();
  const [open, setOpen] = React.useState(false);

  const handleClickOpen = (): void => {
    setOpen(true);
  };

  const handleClose = (): void => {
    setOpen(false);
  };

  const cvtedTime: object = secondToMinute(time);
  const fmtedTime: string = 
    zeroPadding(cvtedTime['hour'], 2) + ':' + 
    zeroPadding(cvtedTime['min'], 2) + ':' + 
    zeroPadding(cvtedTime['sec'], 2);

  const deletePreset = (): void => {
    const deleteURL: string = presetURL + id;
    axios
      .delete<iDeletedPreset>(deleteURL)
      .then((response) => {
        if (response.status === 204) {
          const deletedTimer: HTMLElement | null = document.getElementById(id);
          if (deletedTimer !== null) {
            deletedTimer.remove();
          }
        } else {
          alert('Delete failed');
        }
        handleClose();
      });
  }

  return (
    <div className="timer-card" id={id}>
      <Card className={classes.root} variant="outlined">
        <CardContent>
          <Typography className={classes.title} color="textSecondary" gutterBottom>
            Word of the Day
          </Typography>
          <Typography variant="h5" component="h2">
            {name}
          </Typography>
          <Typography className={classes.pos} color="textSecondary">
            adjective
          </Typography>
          <Typography variant="body2" component="p">
            {fmtedTime}
          </Typography>
        </CardContent>
        <CardActions>
          <div className={classes.alignRight}>

            <Link href={`/timer/play/${encodeURIComponent(id)}`} passHref>
              <IconButton size="medium">
                <PlayCircleFilledWhiteIcon />
              </IconButton>
            </Link>

            <Link href={`/timer/edit/${encodeURIComponent(id)}`} passHref>
              <IconButton size="medium">
                <EditIcon />
              </IconButton>
            </Link>

            <IconButton size="medium" className={classes.alignRight} onClick={handleClickOpen}>
              <DeleteForeverIcon />
            </IconButton>

          </div>
        </CardActions>
      </Card>
      <Dialog
        open={open}
        onClose={handleClose}
        aria-labelledby="alert-dialog-title"
        aria-describedby="alert-dialog-description"
      >
        <DialogTitle id="alert-dialog-title">{"??????????????????????????????"}</DialogTitle>
        <DialogContent>
          <DialogContentText id="alert-dialog-description">
            ??????????????????????????????
          </DialogContentText>
        </DialogContent>
        <DialogActions>
          <Button onClick={handleClose} color="primary">
            ???????????????
          </Button>
          <Button onClick={deletePreset} color="primary" autoFocus>
            ??????
          </Button>
        </DialogActions>
      </Dialog>
    </div>
  );
}

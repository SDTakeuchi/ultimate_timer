import React from 'react';
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
}

export const TimerCard: React.FC<Props> = ({name}) => {
  const classes = useStyles();

  return (
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
          well meaning and kindly.
          <br />
          {'"a benevolent smile"'}
        </Typography>
      </CardContent>
      <CardActions>
        <div className={classes.alignRight}>
          <IconButton size="medium">
            <PlayCircleFilledWhiteIcon />
          </IconButton>
          <IconButton size="medium">
            <EditIcon />
          </IconButton>
          <IconButton size="medium" className={classes.alignRight}>
            <DeleteForeverIcon />
          </IconButton>
        </div>
      </CardActions>
    </Card>
  );
}

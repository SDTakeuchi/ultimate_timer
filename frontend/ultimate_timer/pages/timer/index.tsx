import type { NextPage } from 'next';
import Link from 'next/link';
import React from 'react';
import { NameForm } from '../../components/form/PresetForm'
import { Button } from '@material-ui/core'
import Typography from '@material-ui/core/Typography';
import { TimerList } from '../../components/list/List'
import Head from '../../components/common/Head'
// import { ListItem } from '@mui/material';

const App: NextPage = () => {
  return <div>
    <Head />
    <p>Create your own custom timer presets</p>
    <Link href="/timer/create">
      <Button
        variant="contained"
        color="primary">
        <a>Create A New Preset</a>
      </Button>
    </Link>
    <TimerList />
  </div>
};

export default App;
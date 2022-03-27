import React from 'react';
import type { NextPage } from 'next'
import axios from "axios";
import presetURL from "../../config/settings";
import Head from '../../components/common/Head';
import { NameForm } from '../../components/form/PresetForm';
import Button from '@material-ui/core/Button';


const CreateTimerPage: NextPage = () => {
  const postUrl = presetURL + 'create/'

  function createPreset() {
    axios
      .post(postUrl, {
        name: "Hello World!",
        body: "This is a new post."
      })
      .then(() => {
        alert('preset created!');
      })
      .catch(() => { });
  }

  return <div>
    <Head />
    <NameForm onSubmit={({ timerName }) => {
      console.log(timerName);
    }} />
    <Button onClick={createPreset}>Create Preset</Button>
  </div>
}


export default CreateTimerPage;

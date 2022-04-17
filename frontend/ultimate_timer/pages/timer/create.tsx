import React from 'react';
import type { NextPage } from 'next'
import axios from "axios";
import presetURL from "../../config/settings";
import Head from '../../components/common/Head';
import { NameForm } from '../../components/form/PresetForm';
import Button from '@material-ui/core/Button';


const CreateTimerPage: NextPage = () => {
  const postUrl = presetURL

  function createPreset() {
    axios
      .post(postUrl, {
        name: "Hello World!",
        body: "This is a new post."
      })
      .then(() => {
        alert('preset created!');
      })
      .catch(() => {
        alert('ugh');});
  }

  return <div>
    <Head />
    <NameForm onSubmit={({ name }) => {
      console.log(name);
    }} />
    <Button onClick={createPreset}>Create Preset</Button>
  </div>
}


export default CreateTimerPage;

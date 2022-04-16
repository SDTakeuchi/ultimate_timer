import React from 'react';
import axios from 'axios';
import { Button, TextField } from '@material-ui/core';
import { Formik, Form } from 'formik';
import presetURL from '../../config/settings'
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';


interface Props {
  onSubmit: (values: iPresetForm) => void;
}

interface iPresetForm {
  name: string,
  loopCount: number,
  waitsConfirmEach: boolean,
  waitsConfirmLast: boolean,
  // timerUnit: {
  //   durations: number,
  //   order: number,
  // }[],
}

export const NameForm: React.FC<Props> = ({ onSubmit }) => {
  const [state, setState] = React.useState({
    waitsConfirmEach: false,
    waitsConfirmLast: true,
  });

  const handleChangeSwitch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setState({ ...state, [event.target.name]: event.target.checked });
  };

  const postPreset = () => {
    axios
      .post(presetURL, {
        // name: value.name,
        loop_count: "This is a new post."
      })
      .then(() => {
        alert('preset created!');
      })
      .catch(() => { });
  };

  return (
    <Formik
      initialValues={{
        name: '' ,
        loopCount: 0,
        waitsConfirmEach: false,
        waitsConfirmLast: true
      }}
      onSubmit={values => {
        onSubmit(values);
      }}
    >
      {({ values, handleChange, handleBlur }) => (
        <Form>
          <div>
            <TextField
              variant="outlined"
              label="Timer Preset Name"
              name="name"
              value={values.name}
              onChange={handleChange}
              onBlur={handleBlur}
            />
          </div>
          <div>
            <TextField
              variant="outlined"
              label="loop count"
              name="loopCount"
              type="number"
              value={values.loopCount}
              onChange={handleChange}
              onBlur={handleBlur}
            />
          </div>
          <div>
            <FormGroup row>
              <FormControlLabel
                control={
                  <Switch
                    checked={state.waitsConfirmEach}
                    onChange={handleChangeSwitch}
                    color="primary"
                    name="waitsConfirmEach"
                    value={values.waitsConfirmEach}
                    inputProps={{ 'aria-label': 'primary checkbox' }}
                  />
                }
                label="Waits confirm for each timer"
              />
              <FormControlLabel
                control={
                  <Switch
                    checked={state.waitsConfirmLast}
                    onChange={handleChangeSwitch}
                    color="primary"
                    name="waitsConfirmLast"
                    value={values.waitsConfirmLast}
                    inputProps={{ 'aria-label': 'primary checkbox' }}
                  />
                }
                label="Waits confirm for the last timer"
              />
            </FormGroup>
          </div>
          <div>
            <Button
              type="submit"
              variant="contained"
              color="primary"
              onClick={postPreset}
            >
              Submit
            </Button>
          </div>
          <pre>
            {JSON.stringify(values, null, 2)}
          </pre>
        </Form>
      )}
    </Formik>
  );
};
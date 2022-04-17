import React from 'react';
import axios from 'axios';
import { Button, TextField } from '@material-ui/core';
import { Formik, Form, FieldArray } from 'formik';
import presetURL from '../../config/settings'
import FormGroup from '@material-ui/core/FormGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Switch from '@material-ui/core/Switch';
import { Navigate } from 'react-router-dom'

interface Props {
  onSubmit: (values: iPresetForm) => void;
}

interface iPresetForm {
  name: string,
  loop_count: number,
  display_order: number,
  waits_confirm_each: boolean,
  waits_confirm_last: boolean,
  timer_unit: {
    duration: number,
    order: number,
  }[],
}

export const NameForm: React.FC<Props> = ({ onSubmit }) => {
  const [state, setState] = React.useState({
    waits_confirm_each: false,
    waits_confirm_last: true,
  });

  const handleChangeSwitch = (event: React.ChangeEvent<HTMLInputElement>) => {
    setState({ ...state, [event.target.name]: event.target.checked });
  };

  return (
    <Formik
      initialValues={{
        name: '' ,
        loop_count: 0,
        display_order: 1,
        waits_confirm_each: false,
        waits_confirm_last: true
      }}
      onSubmit={values => {
        axios
          .post(presetURL, values)
          .then((response) => {
            console.log(response);
            console.log(response.data);
            return <Navigate replace to="/timer" />;
            alert('preset created!');
          })
          .catch((response) => {
            console.log(response);
            console.log(response.data);
          });
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
              name="loop_count"
              type="number"
              value={values.loop_count}
              onChange={handleChange}
              onBlur={handleBlur}
            />
          </div>
          <div>
            <TextField
              variant="outlined"
              label="display order"
              name="display_order"
              type="number"
              value={values.display_order}
              onChange={handleChange}
              onBlur={handleBlur}
            />
          </div>
          <div>
            <FormGroup row>
              <FormControlLabel
                control={
                  <Switch
                    checked={state.waits_confirm_each}
                    onChange={handleChangeSwitch}
                    color="primary"
                    name="waits_confirm_each"
                    value={values.waits_confirm_each}
                    inputProps={{ 'aria-label': 'primary checkbox' }}
                  />
                }
                label="Waits confirm for each timer"
              />
              <FormControlLabel
                control={
                  <Switch
                    checked={state.waits_confirm_last}
                    onChange={handleChangeSwitch}
                    color="primary"
                    name="waits_confirm_last"
                    value={values.waits_confirm_last}
                    inputProps={{ 'aria-label': 'primary checkbox' }}
                  />
                }
                label="Waits confirm for the last timer"
              />
            </FormGroup>
          </div>
          <FieldArray
            name="timer_unit"
            render={arrayHelpers => (
              <div>
                {values.timer_unit && values.timer_unit.length > 0 ? (
                  values.timer_unit.map((timerUnit, index) => (
                    <div key={index}>
                      <TextField
                        variant="outlined"
                        label="order"
                        name={`timer_unit.${index}.order`}  // YEAH
                        type="number"
                        // value={values.timer_unit.order}
                        onChange={handleChange}
                        onBlur={handleBlur}
                      />
                      <TextField
                        variant="outlined"
                        label="duration"
                        name={`timer_unit.${index}.duration`}
                        type="number"
                        // value={values.timer_unit.duration}
                        onChange={handleChange}
                        onBlur={handleBlur}
                      />
                      <button
                        type="button"
                         onClick={() => arrayHelpers.remove(index)}
                      >
                        -
                      </button>
                      <button
                        type="button"
                         onClick={() => arrayHelpers.insert(index, '')}
                      >
                        +
                      </button>
                    </div>
                  ))
                ) : (
                  <button type="button" onClick={() => arrayHelpers.push('')}>
                     {/* show this when user has removed all friends from the list */}
                    Add a unit
                  </button>
                )}
              </div>
            )}
          />
          <div>
            <Button
              type="submit"
              variant="contained"
              color="primary"
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
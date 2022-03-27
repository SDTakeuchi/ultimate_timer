import React from 'react';
import { Button, TextField } from '@material-ui/core'
import { Formik, Form } from 'formik'

interface Values {
  timerName: string;
  unit1: number;
}

interface Props {
  onSubmit: (values: Values) => void;
}

export const NameForm: React.FC<Props> = ({ onSubmit }) => {
  return (
    <Formik
      initialValues={{ timerName: '' , unit1: 0}}
      onSubmit={values => {
        onSubmit(values);
      }}
    >
      {({ values, handleChange, handleBlur }) => (
        <Form>
          <div>
            <TextField
              variant="outlined"
              label="Timer Name"
              name="timerName"
              value={values.timerName}
              onChange={handleChange}
              onBlur={handleBlur}
            />
          </div>
          <div>
            <TextField
              variant="outlined"
              label="Unit 1"
              name="unit1"
              type="number"
              value={values.unit1}
              onChange={handleChange}
              onBlur={handleBlur}
            />
          </div>
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
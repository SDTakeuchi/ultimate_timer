import axios from "axios";
import presetURL from "../config/settings"

const deletePreset = (presetID: string) => {
  const url: string = presetURL + "delete/" +presetID;
  let errMsg: string | null = null;

  axios
    .delete(`${url}`)
    .then(() => {
      alert("Preset deleted!");
    })
    .catch(error => {
      errMsg = error;
    });

  return errMsg;
}

export default deletePreset
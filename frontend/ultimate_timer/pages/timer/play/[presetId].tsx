import React from 'react';
import axios from 'axios';
import type { NextPage } from 'next';
import { useRouter } from "next/router";
import { Play } from "../../../components/play/Play"
import presetURL from "../../../config/settings";

const TimerDetail: NextPage = () => {
  const router = useRouter();
  const query = router.query;

  // const url: string = presetURL + presetId;
  // React.useEffect(() => {
  //   axios
  //     .get<iPreset>(url)
  //     .then((response) => {
  //       setPreset(response.data);
  //     })
  //     .catch((error) => {
  //       alert(error.message);
  //     });
  // }, []);
  // const time = new Date();
  // time.setSeconds(time.getSeconds() + preset?.timer_unit[0]?.duration);

  return (
    <div>
      <Play id={query.presetId} />
    </div>
  );
};

export default TimerDetail;

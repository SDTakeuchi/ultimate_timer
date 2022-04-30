import type { NextPage } from 'next';
import { useRouter } from "next/router";
import {Play} from "../../../components/play/Play"

const TimerDetail: NextPage = () => {
  const router = useRouter();
  const { presetId } = router.query;
  const time = new Date();
  time.setSeconds(time.getSeconds() + 600);

  return (
    <div>
      <Play expiryTimestamp={time} id={presetId}/>
    </div>
  );
};

export default TimerDetail;

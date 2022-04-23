import type { NextPage } from 'next';
import { useRouter } from "next/router";
import {Play} from "../../../components/play/Play"

const TimerDetail: NextPage = () => {
  const router = useRouter();
  const presetId: string = String(router.query);

  return (
    <div>
      <Play id={presetId}/>
    </div>
  );
};

export default TimerDetail;

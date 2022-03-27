import type { NextPage } from 'next'
import { useRouter } from "next/router";


const TimerDetail: NextPage = () => {
  const router = useRouter();
  const {presetId} = router.query;
  return (
    <div>
      <h4>Id input in URL</h4>
      <h1>{presetId}</h1>
    </div>
  );
};

export default TimerDetail;

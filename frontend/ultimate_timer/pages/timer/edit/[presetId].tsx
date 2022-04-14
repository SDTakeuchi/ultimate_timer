import type { NextPage } from 'next'
import { useRouter } from "next/router";

const TimerDetail: NextPage = () => {
  const router = useRouter();
  const {presetId} = router.query;
  console.log(presetId);
  return (
    <div>
      <h4>Edit page</h4>
      <h1>{presetId}</h1>
    </div>
  );
};

export default TimerDetail;

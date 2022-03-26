import type { NextPage } from 'next'
import { useRouter } from "next/router";
import Header from "../../components/common/Header";
import PostRefer from "../../components/post/PostRefer";

const Index: NextPage = () => {
  const router = useRouter();
  // パスパラメータから値を取得
  const { post_id } = router.query;
  return (
    <div>
      <Header />
      <PostRefer postId={post_id} />
    </div>
  );
};

export default Index;
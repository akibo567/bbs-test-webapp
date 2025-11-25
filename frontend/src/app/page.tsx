'use client'

import { useRouter } from "next/navigation";


export default function Home() {
    const router = useRouter();
  return (
    <div>
      <span>簡易チャット/わくわく川柳掲示板トップ(α試作版)</span><br/>
      <span>注:わくわく川柳掲示板は5+7+5=17文字しか書き込めないよ</span>
      <div className="card"
      onClick={() => router.push("/bbs")}
      >わくわく川柳掲示板</div>
      <div className="card"
      onClick={() => router.push("/easychat")}
      >簡易チャット</div>
    </div>
  );
}

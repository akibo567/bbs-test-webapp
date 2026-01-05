"use client";

import { useRouter } from "next/navigation";

export default function Home() {
  const router = useRouter();
  return (
    <div>
      <span>わくわく川柳掲示板トップ(α試作版)</span>
      <br />
      <span>
        注:<s>わくわく川柳掲示板は5+7+5=17文字しか書き込めないよ</s>
        現在はテスト版なので何文字でも書き込めます
      </span>
      <br />
      <span>
        α試作版につき、クライアントがリロードしないと全ての書き込みはクライアントに反映されません！
      </span>
      <div className="card" onClick={() => router.push("/bbs")}>
        わくわく川柳掲示板へENTER
      </div>
    </div>
  );
}

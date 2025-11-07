'use client'

//import Image from "next/image";

import { useState,useEffect } from "react";

interface Type_kakikomi {
  id: number;
  name: string;
  message: string;
}

interface Type_ChatMessage1_props{
  id:number;
  name:string;
  message:string;
}

export default function Home() {

  const [kakikomis,Ckakikomis] = useState<Type_kakikomi[]>([]);

  const fetchPosts = async () => {
    Ckakikomis(
      [
        ...kakikomis,
        {
          id:1,
          name:"三郎",
          message:"こんにちは"
        },
        {
          id:2,
          name:"花子",
          message:"さようなら"
        }
      ]
    )
  };

  const fetchPing = async () => {
      const res = await fetch("/api/ping2", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        // 送信したいデータ
        body: JSON.stringify({
          TEST: "20",
        }),
      });

      const ping_resp = await res.json();
      console.log(ping_resp["message"]);
  };

  useEffect(() => {
    fetchPosts();
    fetchPing();
  }, []);

  return (
    <div>
      <span>楽しい簡易チャット</span>
      {
        kakikomis.map(
          (kakikomi) => (
            <div key={kakikomi.id}>
              <ChatMessage1
              id={kakikomi.id}
              name={kakikomi.name}
              message={kakikomi.message}
              />
            </div>
          )
        )
      }
    {/*<ChatMessage1
      name="次郎"
      message="こんちくわ"
    />*/}
    </div>
  );
}


export function ChatMessage1(props:Type_ChatMessage1_props) {
  return (
    <div>
      <span>名前:{props.name}</span>
      <div>{props.message}</div>
    </div>
  );
}

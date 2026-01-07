"use client";

//import Image from "next/image";

import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";

interface Type_Thread {
  id: number;
  name: string;
  title: string;
  message: string;
}

interface Type_ThreadCard_props {
  id: number;
  name: string;
  title: string;
  message: string;
}

const APIENDPOINT: string = "/api";
const MESSAGE_LENGTH = 17;
const NAME_MAX_LENGTH = 20;
const TITLE_MAX_LENGTH = 20;

const MakeThread = async (
  _name: string,
  _title: string,
  _message: string,
): Promise<void> => {
  const trimmedName = _name.trim();
  const trimmedTitle = _title.trim();
  if (trimmedName.length === 0) {
    alert("名前を入力してください");
    return;
  }
  if (trimmedName.length > NAME_MAX_LENGTH) {
    alert(`名前は${NAME_MAX_LENGTH}文字以内で入力してください`);
    return;
  }
  if (trimmedTitle.length === 0) {
    alert("タイトルを入力してください");
    return;
  }
  if (trimmedTitle.length > TITLE_MAX_LENGTH) {
    alert(`タイトルは${TITLE_MAX_LENGTH}文字以内で入力してください`);
    return;
  }
  const trimmedMessage = _message.replace(/\r?\n/g, "");
  if (trimmedMessage.length !== MESSAGE_LENGTH) {
    alert(`本文は5+7+5丁度の${MESSAGE_LENGTH}文字丁度で入力してください！(字余りは許容されません)`);
    return;
  }
  const res = await fetch(APIENDPOINT + "/bbs/post_thread", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    // 送信したいデータ
    body: JSON.stringify({
      Name: _name,
      Title: _title,
      Message: _message,
    }),
  });

  if (res.status === 200) {
    alert("書き込みました！");
  } else {
    alert("書き込みに失敗しました");
  }
  //const ping_resp = await res.json();
  //console.log(ping_resp["message"]);
};

const LoadThreads = async (): Promise<Type_Thread[]> => {
  const res = await fetch(APIENDPOINT + "/bbs/get_threads", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    // 送信したいデータ
    /*body: JSON.stringify({
          TEST: "20",
        }),*/
  });

  if (res.status !== 200) {
    alert("ロードに失敗しました");
  }

  const respj = await res.json();
  //console.log(respj);

  // const return_messages: Type_Thread[] = [];
  /*respj["threads"].map( (resp) => {
        return_messages = [
          ...return_messages,
          {
            id:resp["id"],
            name:resp["name"],
            title:resp["title"],
          }
        console.log(resp["threads"])
      ]});*/
  //console.log(respj["threads"]);
  return respj["threads"] as Type_Thread[];
};

export default function Home() {
  const [threads, Cthreads] = useState<Type_Thread[]>([]);
  const [input_name, Cinput_name] = useState<string>("");
  const [input_title, Cinput_title] = useState<string>("");
  const [input_message, Cinput_message] = useState<string>("");

  async function fetchPosts(_viewthreadsdata: Type_Thread[]) {
    console.log(_viewthreadsdata);
    Cthreads([..._viewthreadsdata]);
  }

  useEffect(() => {
    //fetchPosts();
    (async () => {
      const data = await LoadThreads();
      fetchPosts(data);
    })();
    //fetchPing();
  }, []);

  return (
    <div>
      <span>わくわく川柳掲示板/トップスレッド</span>
      {threads.map((thread) => (
        <div key={thread.id}>
          <ThreadCard
            id={thread.id}
            name={thread.name}
            title={thread.title}
            message={thread.message}
          />
        </div>
      ))}

      <div>
        <label>名前:</label>
        <textarea
          value={input_name}
          maxLength={NAME_MAX_LENGTH}
          onChange={(e) => {
            Cinput_name(e.target.value);
          }}
        />
      </div>
      <div>
        <label>タイトル:</label>
        <textarea
          value={input_title}
          maxLength={TITLE_MAX_LENGTH}
          onChange={(e) => {
            Cinput_title(e.target.value);
          }}
        />
      </div>

      <div>
        <label>本文:</label>
        <textarea
          value={input_message}
          onChange={(e) => {
            Cinput_message(e.target.value);
          }}
        />
      </div>

      <button
        onClick={() => {
          MakeThread(input_name, input_title, input_message);
          //alert(input_message);
        }}
      >
        送信
      </button>
    </div>
  );
}

export function ThreadCard(props: Type_ThreadCard_props) {
  const router = useRouter();

  return (
    <div
      className="card"
      onClick={() => router.push("/bbs/thread/" + props.id)}
    >
      <span>{props.title ? props.title : "名無しスレ"}</span>
    </div>
  );
}

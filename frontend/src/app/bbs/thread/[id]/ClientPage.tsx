"use client";

import { useState, useEffect } from "react";

interface Type_Message {
  id: number;
  name: string;
  message_text: string;
}

interface Type_MessageCard_props {
  id: number;
  name: string;
  message_text: string;
}

const APIENDPOINT: string = "/api";
const MESSAGE_LENGTH = 17;
const NAME_MAX_LENGTH = 20;

const SendMessage = async (
  _name: string,
  _message: string,
  _threadID: number,
): Promise<void> => {
  console.log(_threadID);
  const trimmedName = _name.trim();
  if (trimmedName.length === 0) {
    alert("名前を入力してください");
    return;
  }
  if (trimmedName.length > NAME_MAX_LENGTH) {
    alert(`名前は${NAME_MAX_LENGTH}文字以内で入力してください`);
    return;
  }
  const trimmedMessage = _message.replace(/\r?\n/g, "");
  if (trimmedMessage.length !== MESSAGE_LENGTH) {
    alert(`本文は5+7+5丁度の${MESSAGE_LENGTH}文字丁度で入力してください！(字余りは許容されません)`);
    return;
  }
  const res = await fetch(APIENDPOINT + "/bbs/post_message", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    // 送信したいデータ
    body: JSON.stringify({
      Name: _name,
      ThreadID: _threadID,
      Message_Text: _message,
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

const LoadMessages = async (
  _threadID: number,
): Promise<{ messages: Type_Message[]; title: string }> => {
  if (_threadID === 0) return { messages: [], title: "" };

  const controller = new AbortController();

  const res = await fetch(APIENDPOINT + "/bbs/get_messages", {
    signal: controller.signal,
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      ID: _threadID,
    }),
  });

  if (res.status !== 200) {
    alert("ロードに失敗しました");
  }

  const respj = await res.json();
  //console.log(respj);

  // const return_messages: Type_Message[] = [];
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
  return {
    messages: respj["messages"] as Type_Message[],
    title: respj["title"] as string,
  };
};

export function ClientPage({ pageid }: { pageid: string }) {
  const [messages, Cmessages] = useState<Type_Message[]>([]);
  const [input_name, Cinput_name] = useState<string>("");
  const [input_message, Cinput_message] = useState<string>("");
  const [thread_title, Cthread_title] = useState<string>("");

  async function fetchPosts(_viewmessagesdata: {
    messages: Type_Message[];
    title: string;
  }) {
    console.log(_viewmessagesdata);
    Cmessages([..._viewmessagesdata.messages]);
    Cthread_title(_viewmessagesdata.title);
  }

  useEffect(() => {
    if (!pageid) return;

    (async () => {
      const data = await LoadMessages(Number(pageid));
      fetchPosts(data);
    })();
    //fetchPing();
  }, [pageid]);

  return (
    <div>
      <span>わくわく川柳掲示板/{thread_title}</span>
      {messages.map((message) => (
        <div key={message.id}>
          <MessageCard
            id={message.id}
            name={message.name}
            message_text={message.message_text}
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
          SendMessage(input_name, input_message, Number(pageid));
          //alert(input_message);
        }}
      >
        送信
      </button>
    </div>
  );
}

export function MessageCard(props: Type_MessageCard_props) {
  return (
    <div className="card">
      <div>
        <span>名前：</span>
        <span>{props.name}</span>
      </div>
      <div>
        <span>{props.message_text}</span>
      </div>
    </div>
  );
}

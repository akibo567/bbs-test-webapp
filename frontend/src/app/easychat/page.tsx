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

  const APIENDPOINT:string = "/api";

    const SendKakikomi:Function = async (_name:string,_message:string) => {
      const res = await fetch(APIENDPOINT +"/post_message", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        // 送信したいデータ
        body: JSON.stringify({
          Name: _name,
          Message: _message,
        }),
      });

      await alert(res);
      //const ping_resp = await res.json();
      //console.log(ping_resp["message"]);
  };

   const Loadkakikomi:Function = async () => {
      const res = await fetch(APIENDPOINT + "/get_chat_messages", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        // 送信したいデータ
        /*body: JSON.stringify({
          TEST: "20",
        }),*/
      });

      const respj = await res.json();
      //console.log(respj);
      
      let return_messages:Type_kakikomi[]=[];
      respj["messages"].map( (resp) => {
        return_messages = [
          ...return_messages,
          {
            id:resp["id"],
            name:resp["name"],
            message:resp["message"],
          }
        //console.log(resp["message"])
      ]});
      //console.log(respj["message"]);
      return respj["messages"];
  };

export default function Home() {

  const [kakikomis,Ckakikomis] = useState<Type_kakikomi[]>([]);
  const [input_name,Cinput_name] = useState<string>();
  const [input_message,Cinput_message] = useState<string>();

  async function fetchPosts(_viewpostsdata:Type_kakikomi[]){
    console.log(_viewpostsdata);
    Ckakikomis(
      [
        ..._viewpostsdata,
      ]
    )
  };



  useEffect(() => {
    //fetchPosts();
  (async () => {
    const data = await Loadkakikomi();
    fetchPosts(data);
  })();
    //fetchPing();
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

<div>
    <label>名前:</label>
    <textarea
    value={input_name}
    onChange={(e)=> {Cinput_name(e.target.value)} }
    />
</div>
<div>
    <label>メッセージ:</label>
    <textarea
    value={input_message}
    onChange={(e)=> {Cinput_message(e.target.value)} }
    />
</div>

    <button onClick={()=>{
      SendKakikomi(input_name,input_message);
      //alert(input_message);
    }}>送信</button>
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

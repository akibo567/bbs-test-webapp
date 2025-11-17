'use client'

//import Image from "next/image";

import { useState,useEffect } from "react";

interface Type_Thread {
  id: number;
  name: string;
  title: string;
}

interface Type_ThreadCard_props{
  id:number;
  name:string;
  title:string;
}

  const APIENDPOINT:string = "/api";

    const MakeThread:Function = async (_name:string,_title:string) => {
      const res = await fetch(APIENDPOINT +"/post_threads", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        // 送信したいデータ
        body: JSON.stringify({
          Name: _name,
          Title: _title,
        }),
      });

      await alert(res);
      //const ping_resp = await res.json();
      //console.log(ping_resp["message"]);
  };

   const LoadThreads:Function = async () => {
      const res = await fetch(APIENDPOINT + "/get_threads", {
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
      
      let return_messages:Type_Thread[]=[];
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
      return respj["threads"];
  };

export default function Home() {

  const [threads,Cthreads] = useState<Type_Thread[]>([]);
  const [input_name,Cinput_name] = useState<string>();
  const [input_title,Cinput_title] = useState<string>();

  async function fetchPosts(_viewthreadsdata:Type_Thread[]){
    console.log(_viewthreadsdata);
    Cthreads(
      [
        ..._viewthreadsdata,
      ]
    )
  };



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
      {
        threads.map(
          (thread) => (
            <div key={thread.id}>
              <ThreadCard
              id={thread.id}
              name={thread.name}
              title={thread.title}
              />
            </div>
          )
        )
      }

<div>
    <label>名前:</label>
    <textarea
    value={input_name}
    onChange={(e)=> {Cinput_name(e.target.value)} }
    />
</div>
<div>
    <label>タイトル:</label>
    <textarea
    value={input_title}
    onChange={(e)=> {Cinput_title(e.target.value)} }
    />
</div>

    <button onClick={()=>{
      MakeThread(input_name,input_title);
      //alert(input_message);
    }}>送信</button>
    </div>
  );





}


export function ThreadCard(props:Type_ThreadCard_props) {
  return (
    <div>
      <span>{props.title? props.title: '名無しすれ'}</span>
    </div>
  );
}

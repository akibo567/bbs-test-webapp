import { use } from "react";
import { ClientPage } from "./ClientPage";

export default function Home({ params }: { params: Promise<{ id: string }> }) {
  const { id } = use(params);

  return <ClientPage pageid={id} />;
}

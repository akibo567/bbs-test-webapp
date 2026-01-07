import { ClientPage } from "./ClientPage";

export default function Home({ params }: { params: { id: string } }) {
  return <ClientPage pageid={params.id} />;
}


export default function Home({ params }: { params: { id: string } }) {
  return (
    <div>
     ID は {params.id} だよ
    </div>
  );
}

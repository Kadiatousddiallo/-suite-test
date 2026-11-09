
import { Suspense } from "react";
import ClientDashboard from "./ClientDashboard";

async function ClientPageContent({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  return <ClientDashboard clientId={id} />;
}

export default function ClientPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  return (
    <Suspense
      fallback={
        <div className="flex min-h-screen items-center justify-center">
          Chargement du tableau de bord...
        </div>
      }
    >
      <ClientPageContent params={params} />
    </Suspense>
  );
}

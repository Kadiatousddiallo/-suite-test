
"use client";

import { useEffect, useState } from "react";

import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from "chart.js";

import type { ChartOptions, TooltipItem } from "chart.js";
import { Bar } from "react-chartjs-2";

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

// Types des données reçues de l'API

type PlatformStats = {
  platform: string;
  publications: number;
  impressions: number;
  interactions: number;
  engagement_rate: number;
};

type StatsResponse = {
  client_id: number;
  client_name: string;
  total: {
    publications: number;
    impressions: number;
    interactions: number;
    engagement_rate: number;
  };
  by_platform: PlatformStats[];
};

type Publication = {
  id: number;
  platform: string;
  title: string;
  published_at: string;
  impressions: number;
  interactions: number;
  engagement_rate: number;
};

// Fonctions et constantes

const formatNumber = (value: number) =>
  new Intl.NumberFormat("fr-FR").format(value);

const platformNames: Record<string, string> = {
  instagram: "Instagram",
  facebook: "Facebook",
  tiktok: "TikTok",
  linkedin: "LinkedIn",
};

const chartColors = [
  "#3B82F6",
  "#EC4899",
  "#111827",
  "#06B6D4",
];

// Tableau de bord

export default function ClientDashboard({
  clientId,
}: {
  clientId: string;
}) {
  const [stats, setStats] = useState<StatsResponse | null>(null);
  const [publications, setPublications] = useState<Publication[]>([]);

  const [fromDate, setFromDate] = useState("2026-10-01");
  const [toDate, setToDate] = useState("2026-10-31");

  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  // Chargement des statistiques et des publications

  async function loadDashboard(from: string, to: string) {
    setLoading(true);
    setError("");

    try {
      const [statsResponse, publicationsResponse] = await Promise.all([
        fetch(
          `/api/clients/${clientId}/stats?from=${from}&to=${to}`,
          { cache: "no-store" }
        ),

        fetch(`/api/clients/${clientId}/publications/top?limit=3`, {
          cache: "no-store",
        }),
      ]);

      const statsData = await statsResponse.json();
      const publicationsData = await publicationsResponse.json();

      if (!statsResponse.ok) {
        throw new Error(
          statsData.error || "Impossible de charger les statistiques."
        );
      }

      if (!publicationsResponse.ok) {
        throw new Error(
          publicationsData.error || "Impossible de charger les publications."
        );
      }

      setStats(statsData);
      setPublications(publicationsData);
    } catch (err) {
      setError(
        err instanceof Error ? err.message : "Une erreur est survenue."
      );
    } finally {
      setLoading(false);
    }
  }

  // Chargement initial

  useEffect(() => {
    loadDashboard("2026-10-01", "2026-10-31");
  }, [clientId]);

  // Filtrage par dates

  function handleFilter(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();

    if (fromDate > toDate) {
      setError("La date de début doit être antérieure à la date de fin.");
      return;
    }

    loadDashboard(fromDate, toDate);
  }

  // Données du graphique Chart.js

  const chartData = {
    labels:
      stats?.by_platform.map(
        (platform) =>
          platformNames[platform.platform] ?? platform.platform
      ) ?? [],

    datasets: [
      {
        label: "Taux d'engagement (%)",

        data:
          stats?.by_platform.map(
            (platform) => platform.engagement_rate
          ) ?? [],

        backgroundColor: chartColors,
        borderRadius: 8,
        borderSkipped: false as const,
      },
    ],
  };

  // Options du graphique

  const chartOptions: ChartOptions<"bar"> = {
    responsive: true,
    maintainAspectRatio: false,

    plugins: {
      legend: {
        display: true,
        position: "top",
      },

      title: {
        display: false,
      },

      tooltip: {
        callbacks: {
          label: (context: TooltipItem<"bar">) =>
            `Engagement : ${(context.parsed.y ?? 0).toFixed(2)} %`,
        },
      },
    },

    scales: {
      y: {
        beginAtZero: true,

        title: {
          display: true,
          text: "Taux d'engagement (%)",
        },

        ticks: {
          callback: (value) => `${value} %`,
        },
      },

      x: {
        title: {
          display: true,
          text: "Plateformes",
        },
      },
    },
  };

  // Affichage du tableau de bord

  return (
    <main className="min-h-screen bg-slate-50 px-4 py-8 text-slate-900 sm:px-8">
      <div className="mx-auto max-w-6xl">

        {/* En-tête */}

        <header className="mb-8">
          <p className="text-sm font-semibold uppercase tracking-wider text-blue-600">
            SUITE · Reporting social media
          </p>

          <h1 className="mt-2 text-3xl font-bold">
            Tableau de bord
          </h1>

          <p className="mt-2 text-slate-500">
            {stats?.client_name ?? `Client #${clientId}`}
          </p>
        </header>

        {/* Filtre des dates */}

        <form
          onSubmit={handleFilter}
          className="mb-8 flex flex-wrap items-end gap-4 rounded-xl border border-slate-200 bg-white p-5 shadow-sm"
        >
          <div>
            <label className="mb-2 block text-sm font-medium">
              Date de début
            </label>

            <input
              type="date"
              value={fromDate}
              onChange={(event) => setFromDate(event.target.value)}
              className="rounded-lg border border-slate-300 px-3 py-2"
            />
          </div>

          <div>
            <label className="mb-2 block text-sm font-medium">
              Date de fin
            </label>

            <input
              type="date"
              value={toDate}
              onChange={(event) => setToDate(event.target.value)}
              className="rounded-lg border border-slate-300 px-3 py-2"
            />
          </div>

          <button
            type="submit"
            disabled={loading}
            className="rounded-lg bg-blue-600 px-5 py-2.5 font-medium text-white hover:bg-blue-700 disabled:opacity-50"
          >
            {loading ? "Chargement..." : "Filtrer"}
          </button>
        </form>

        {/* Chargement */}

        {loading && (
          <p className="py-12 text-center text-slate-500">
            Chargement des statistiques...
          </p>
        )}

        {/* Gestion des erreurs */}

        {!loading && error && (
          <div className="rounded-xl border border-red-200 bg-red-50 p-5 text-red-700">
            {error}
          </div>
        )}

        {/* Résultats */}

        {!loading && !error && stats && (
          <>
            {/* Indicateurs principaux */}

            <section className="mb-8 grid gap-4 sm:grid-cols-2 lg:grid-cols-4">
              <MetricCard
                label="Publications"
                value={formatNumber(stats.total.publications)}
              />

              <MetricCard
                label="Impressions"
                value={formatNumber(stats.total.impressions)}
              />

              <MetricCard
                label="Interactions"
                value={formatNumber(stats.total.interactions)}
              />

              <MetricCard
                label="Engagement moyen"
                value={`${stats.total.engagement_rate.toFixed(2)} %`}
              />
            </section>

            {/* Graphique Chart.js */}

            <section className="mb-8 rounded-xl border border-slate-200 bg-white p-6 shadow-sm">
              <h2 className="mb-6 text-xl font-bold">
                Engagement par plateforme
              </h2>

              {stats.by_platform.length === 0 ? (
                <p className="py-10 text-center text-slate-500">
                  Aucune donnée disponible pour cette période.
                </p>
              ) : (
                <div className="relative h-80 w-full">
                  <Bar data={chartData} options={chartOptions} />
                </div>
              )}
            </section>

            {/* Top 3 des publications */}

            <section className="rounded-xl border border-slate-200 bg-white p-6 shadow-sm">
              <h2 className="mb-5 text-xl font-bold">
                Top 3 des publications
              </h2>

              {publications.length === 0 ? (
                <p className="text-slate-500">
                  Aucune publication trouvée.
                </p>
              ) : (
                <div className="space-y-4">
                  {publications.map((publication, index) => (
                    <article
                      key={publication.id}
                      className="flex flex-col justify-between gap-3 rounded-lg border border-slate-100 p-4 sm:flex-row sm:items-center"
                    >
                      <div>
                        <p className="text-xs font-semibold uppercase text-blue-600">
                          #{index + 1} ·{" "}
                          {platformNames[publication.platform] ??
                            publication.platform}
                        </p>

                        <h3 className="mt-1 font-semibold">
                          {publication.title}
                        </h3>

                        <p className="mt-1 text-sm text-slate-500">
                          {formatNumber(publication.impressions)} impressions
                          {" · "}
                          {formatNumber(publication.interactions)} interactions
                        </p>
                      </div>

                      <p className="text-lg font-bold text-green-700">
                        {publication.engagement_rate.toFixed(2)} %
                      </p>
                    </article>
                  ))}
                </div>
              )}
            </section>
          </>
        )}
      </div>
    </main>
  );
}

// Composant des cartes statistiques

function MetricCard({
  label,
  value,
}: {
  label: string;
  value: string;
}) {
  return (
    <article className="rounded-xl border border-slate-200 bg-white p-5 shadow-sm">
      <p className="text-sm text-slate-500">{label}</p>
      <p className="mt-2 text-2xl font-bold">{value}</p>
    </article>
  );
}

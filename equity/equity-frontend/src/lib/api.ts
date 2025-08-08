export async function getEquities(filter?: string) {
  let url = `${process.env.NEXT_PUBLIC_API_URL}/api/equities`;
  if (filter) url += `?filter=${filter}`;
  const res = await fetch(url, { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch equities");
  return res.json();
}

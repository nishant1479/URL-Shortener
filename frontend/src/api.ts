export async function getBackendData(): Promise<any> {
  const res = await fetch("http://localhost:3000/api/data"); // Replace with your real endpoint
  if (!res.ok) throw new Error("Failed to fetch data");
  return await res.json();
}

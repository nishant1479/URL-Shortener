export async function getBackendData(): Promise<any> {
  const res = await fetch("http://localhost:8080/api/data"); // Updated to correct backend port
  if (!res.ok) throw new Error("Failed to fetch data");
  return await res.json();
}

export async function shortenUrl(originalUrl: string, validForMinutes: number = 1440): Promise<string> {
  const res = await fetch("http://localhost:8080/shorten", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ original_url: originalUrl, valid_for_minutes: validForMinutes })
  });
  if (!res.ok) throw new Error("Failed to shorten URL");
  const data = await res.json();
  return data.short_url;
}

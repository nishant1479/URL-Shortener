import { useState } from "react";
import { shortenUrl } from "./api";

function App() {
  // State for URL shortener form
  const [inputUrl, setInputUrl] = useState("");
  const [shortUrl, setShortUrl] = useState<string | null>(null);
  const [formError, setFormError] = useState<string | null>(null);
  const [formLoading, setFormLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setFormError(null);
    setShortUrl(null);
    if (!inputUrl.trim()) {
      setFormError("Please enter a URL.");
      return;
    }
    setFormLoading(true);
    try {
      const result = await shortenUrl(inputUrl.trim());
      setShortUrl(result);
    } catch (err: any) {
      setFormError(err.message || "Failed to shorten URL");
    } finally {
      setFormLoading(false);
    }
  };

  return (
    <div style={{ padding: "2rem", fontFamily: "Arial", maxWidth: 600, margin: "0 auto" }}>
      <h1>🔗 URL Shortener</h1>
      <form onSubmit={handleSubmit} style={{ marginBottom: "2rem" }}>
        <input
          type="url"
          placeholder="Enter a long URL..."
          value={inputUrl}
          onChange={e => setInputUrl(e.target.value)}
          style={{ width: "70%", padding: "0.5rem", fontSize: 16 }}
          required
        />
        <button type="submit" style={{ padding: "0.5rem 1rem", marginLeft: 8, fontSize: 16 }} disabled={formLoading}>
          {formLoading ? "Shortening..." : "Shorten"}
        </button>
      </form>
      {formError && <p style={{ color: "red" }}>{formError}</p>}
      {shortUrl && (
        <div style={{ marginBottom: "2rem" }}>
          <strong>Shortened URL:</strong> <a href={shortUrl} target="_blank" rel="noopener noreferrer">{shortUrl}</a>
        </div>
      )}
    </div>
  );
}

export default App;

import CryptoDashboard from "@/components/CryptoDashboard";

export default function TrendingTokens() {
  return (
    <CryptoDashboard 
      endpoint="/ws/trending" 
      title="Trending & AI Assets" 
    />
  );
}

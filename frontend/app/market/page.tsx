import CryptoDashboard from "@/components/CryptoDashboard";

export default function MarketTokens() {
  return (
    <CryptoDashboard 
      endpoint="/ws/market" 
      title="L1 & Layer 2 Market" 
    />
  );
}

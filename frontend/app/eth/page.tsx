import CryptoDashboard from "@/components/CryptoDashboard";

export default function EthTokens() {
  return (
    <CryptoDashboard 
      endpoint="/ws/eth" 
      title="Ethereum Ecosystem" 
    />
  );
}
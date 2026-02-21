import CryptoDashboard from "@/components/CryptoDashboard";

export default function SolTokens() {
  return (
    <CryptoDashboard 
      endpoint="/ws/sol" 
      title="Solana Ecosystem" 
    />
  );
}
import CryptoDashboard from "@/components/CryptoDashboard";

export default function Home() {
  return (
    <CryptoDashboard 
      endpoint="/ws/main" 
      title="Core Assets" 
    />
  );
}

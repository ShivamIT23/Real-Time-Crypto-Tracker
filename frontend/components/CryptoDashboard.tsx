"use client";

import { useEffect, useState, useRef } from "react";

interface CryptoDashboardProps {
  endpoint: string;
  title: string;
}

export default function CryptoDashboard({ endpoint, title }: CryptoDashboardProps) {
  const [prices, setPrices] = useState<Record<string, number>>({});
  const [prevPrices, setPrevPrices] = useState<Record<string, number>>({});
  const [status, setStatus] = useState<"connecting" | "connected" | "disconnected">("connecting");
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    const connect = () => {
      ws.current = new WebSocket(`ws://localhost:8080${endpoint}`);

      ws.current.onopen = () => setStatus("connected");
      ws.current.onclose = () => {
        setStatus("disconnected");
        setTimeout(connect, 3000); // Reconnect after 3 seconds
      };

      ws.current.onmessage = (event) => {
        const data = JSON.parse(event.data);
        setPrices((prev) => {
          setPrevPrices(prev);
          return { ...prev, ...data };
        });
      };
    };

    connect();

    return () => {
      if (ws.current) ws.current.close();
    };
  }, [endpoint]);

  return (
    <div className="max-w-6xl mx-auto p-8">
      <div className="flex items-center justify-between mb-12">
        <h1 className="text-4xl font-bold tracking-tight bg-gradient-to-r from-white to-white/60 bg-clip-text text-transparent">
          {title}
        </h1>
        <div className="flex items-center gap-2 px-4 py-2 bg-white/5 rounded-full border border-white/10">
          <div className={`w-2 h-2 rounded-full ${status === 'connected' ? 'bg-emerald-500 shadow-[0_0_10px_#10b981]' : 'bg-red-500'}`} />
          <span className="text-xs font-medium uppercase tracking-widest text-white/60">{status}</span>
        </div>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {Object.entries(prices).length === 0 ? (
          <div className="col-span-full py-20 text-center glass-card">
            <p className="text-white/40 animate-pulse">Waiting for price feeds...</p>
          </div>
        ) : (
          Object.entries(prices).map(([coin, price]) => {
            const prevPrice = prevPrices[coin];
            const direction = prevPrice === undefined || price === prevPrice ? 'stable' : price > prevPrice ? 'up' : 'down';
            
            return (
              <div key={coin} className="glass-card p-6 flex flex-col gap-4">
                <div className="flex justify-between items-start">
                  <div>
                    <h2 className="text-sm font-bold uppercase tracking-widest text-white/40 mb-1">{coin}</h2>
                    <p className="text-4xl font-mono font-bold tracking-tight">
                      ${price.toLocaleString(undefined, { minimumFractionDigits: 2, maximumFractionDigits: 8 })}
                    </p>
                  </div>
                  <div className={`text-xs px-2 py-1 rounded-md bg-white/5 border border-white/10 ${direction === 'up' ? 'price-up' : direction === 'down' ? 'price-down' : 'text-white/40'}`}>
                    {direction === 'up' ? '▲ LIVE' : direction === 'down' ? '▼ LIVE' : '• LIVE'}
                  </div>
                </div>
                <div className="h-[2px] w-full bg-white/5 overflow-hidden rounded-full">
                   <div className={`h-full transition-all duration-1000 ${direction === 'up' ? 'bg-emerald-500 w-full' : direction === 'down' ? 'bg-red-500 w-full' : 'bg-white/10 w-0'}`} />
                </div>
              </div>
            );
          })
        )}
      </div>
    </div>
  );
}

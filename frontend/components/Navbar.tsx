"use client";

import Link from "next/link";
import { usePathname } from "next/navigation";

export default function Navbar() {
  const pathname = usePathname();

  const links = [
    { href: "/", label: "Main", icon: "🌐" },
    { href: "/market", label: "L1 Market", icon: "📊" },
    { href: "/eth", label: "ETH Tokens", icon: "💎" },
    { href: "/sol", label: "SOL Tokens", icon: "⚡" },
    { href: "/trending", label: "Trending", icon: "🔥" },
  ];

  return (
    <nav className="sticky top-0 z-50 py-6 px-8 flex justify-center">
      <div className="bg-white/5 backdrop-blur-xl border border-white/10 p-1.5 rounded-full flex flex-wrap justify-center gap-1 shadow-2xl">
        {links.map((link) => {
          const isActive = pathname === link.href;
          return (
            <Link
              key={link.href}
              href={link.href}
              className={`nav-link flex items-center gap-2 text-xs md:text-sm font-medium whitespace-nowrap ${
                isActive ? "active text-white" : "text-white/60 hover:text-white/90"
              }`}
            >
              <span>{link.icon}</span>
              {link.label}
            </Link>
          );
        })}
      </div>
    </nav>
  );
}

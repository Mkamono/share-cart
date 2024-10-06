"use client";
import { buttonVariants } from "@/components/ui/button";
import Link from "next/link";
import { usePathname } from "next/navigation";

function NavigationItem({
	href,
	children,
	className,
}: {
	href: "/home" | "/explore" | "/notification";
	children: string;
	className?: string;
}) {
	const pathname = usePathname();
	const isActive = pathname === href;
	const linkClassName = isActive
		? `${buttonVariants({ variant: "default" })} ${className}`
		: `${buttonVariants({ variant: "ghost" })} ${className}`;
	return (
		<Link
			href={
				isActive
					? href
					: { pathname: href, search: new URLSearchParams({}).toString() }
			}
			className={linkClassName}
		>
			{children}
		</Link>
	);
}

export function Footer() {
	return (
		<nav className="flex" aria-label="フッターナビゲーション">
			<NavigationItem href="/home" className="flex-1">
				Home
			</NavigationItem>
			<NavigationItem href="/explore" className="flex-1">
				Explore
			</NavigationItem>
			<NavigationItem href="/notification" className="flex-1">
				Notification
			</NavigationItem>
		</nav>
	);
}

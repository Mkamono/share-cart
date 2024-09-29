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

export default function ExampleClientComponent() {
	return (
		<section className="flex">
			<NavigationItem href="/home" className="flex-auto">
				Home
			</NavigationItem>
			<NavigationItem href="/explore" className="flex-auto">
				Explore
			</NavigationItem>
			<NavigationItem href="/notification" className="flex-auto">
				Notification
			</NavigationItem>
		</section>
	);
}

import type { MetaFunction } from "@remix-run/node";
import { Outlet } from "@remix-run/react";
import { Footer } from "~/routes/_top/Footer";

export const meta: MetaFunction = () => {
	return [
		{ title: "Share Cart" },
		{ name: "description", content: "Let's Share Your Cart!" },
	];
};

export default function Index() {
	return (
		<div className="flex flex-col h-dvh">
			<div className="flex-auto overflow-auto p-2 pb-0">
				<Outlet />
			</div>
			<div className="sticky bottom-0">
				<Footer />
			</div>
		</div>
	);
}

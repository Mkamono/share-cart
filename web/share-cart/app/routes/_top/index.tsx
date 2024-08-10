import { type MetaFunction } from "@remix-run/node";
import { Outlet, useLocation } from "@remix-run/react";
import Footer from "~/routes/_top/footer";

export const meta: MetaFunction = () => {
	return [
		{ title: "Share Cart" },
		{ name: "description", content: "Let's Share Your Cart!" },
	];
};

export default function Index() {
	const location = useLocation();
	return (
		<div className="flex flex-col min-h-screen">
			<div>location.pathname : {location.pathname}</div>
			<Outlet />
			<div className="flex-grow"></div>
			<Footer />
		</div>
	);
}

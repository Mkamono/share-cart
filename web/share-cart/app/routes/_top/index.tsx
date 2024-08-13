import { Box } from "@radix-ui/themes";
import { type MetaFunction } from "@remix-run/node";
import { Outlet, useLocation } from "@remix-run/react";
import { Footer } from "~/routes/_top/Footer";

export const meta: MetaFunction = () => {
	return [
		{ title: "Share Cart" },
		{ name: "description", content: "Let's Share Your Cart!" },
	];
};

export default function Index() {
	const location = useLocation();
	return (
		<Box className="flex flex-col h-dvh">
			<Box className="flex-grow overflow-auto p-2 pb-0">
				<Outlet />
			</Box>
			<Box className="sticky bottom-0">
				<Footer />
			</Box>
		</Box>
	);
}

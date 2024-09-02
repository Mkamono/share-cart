import { NavLink } from "@remix-run/react";
import type { ReactNode } from "react";
import { $path } from "remix-routes";

const navContent = (
	isActive: boolean,
	isPending: boolean,
	children: ReactNode,
) => {
	if (isPending) {
		return <div className="text-center text-gray-400 py-2">{children}</div>;
	}
	if (isActive) {
		return (
			<div className="text-center border-b-4 border-blue-500 py-2">
				{children}
			</div>
		);
	}
	return <div className="text-center text-gray-400 py-2">{children}</div>;
};

export const Footer = () => {
	return (
		<div className="flex">
			<NavLink to={$path("/home")} className="flex-1 text-center">
				{({ isActive, isPending }) => navContent(isActive, isPending, "Home")}
			</NavLink>
			<NavLink to={$path("/explore")} className="flex-1 text-center">
				{({ isActive, isPending }) =>
					navContent(isActive, isPending, "Explore")
				}
			</NavLink>
			<NavLink to={$path("/notification")} className="flex-1 text-center">
				{({ isActive, isPending }) =>
					navContent(isActive, isPending, "Notification")
				}
			</NavLink>
		</div>
	);
};

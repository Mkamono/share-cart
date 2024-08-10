import { TabNav } from "@radix-ui/themes";
import { useLocation } from "@remix-run/react";
import { $path } from "remix-routes";

function Footer() {
	const pathname = useLocation().pathname;
	return (
		<TabNav.Root className="flex flex-row justify-center">
			<div className="flex-1">
				<TabNav.Link
					href={$path("/home")}
					active={pathname === $path("/home")}
					className="w-full"
				>
					Home
				</TabNav.Link>
			</div>
			<div className="flex-1">
				<TabNav.Link
					href={$path("/explore")}
					active={pathname === $path("/explore")}
					className="w-full"
				>
					Explore
				</TabNav.Link>
			</div>
			<div className="flex-1">
				<TabNav.Link
					href={$path("/notification")}
					active={pathname === $path("/notification")}
					className="w-full"
				>
					Notification
				</TabNav.Link>
			</div>
		</TabNav.Root>
	);
}
export default Footer;

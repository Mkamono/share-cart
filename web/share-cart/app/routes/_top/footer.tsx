import { BellIcon, HomeIcon, RocketIcon } from "@radix-ui/react-icons";
import { Box, TabNav, Text } from "@radix-ui/themes";
import { useLocation } from "@remix-run/react";
import { $path } from "remix-routes";

export const Footer = () => {
	const pathname = useLocation().pathname;
	return (
		<TabNav.Root className="flex flex-row justify-center bg-white">
			<Box className="flex-1">
				<TabNav.Link
					href={$path("/home")}
					active={pathname === $path("/home")}
					className="w-full"
				>
					<HomeIcon />
					<Text size="1" ml="1">
						Home
					</Text>
				</TabNav.Link>
			</Box>
			<Box className="flex-1">
				<TabNav.Link
					href={$path("/explore")}
					active={pathname === $path("/explore")}
					className="w-full"
				>
					<RocketIcon />
					<Text size="1" ml="1">
						Explore
					</Text>
				</TabNav.Link>
			</Box>
			<Box className="flex-1">
				<TabNav.Link
					href={$path("/notification")}
					active={pathname === $path("/notification")}
					className="w-full"
				>
					<BellIcon />
					<Text size="1" ml="1">
						Notification
					</Text>
				</TabNav.Link>
			</Box>
		</TabNav.Root>
	);
};

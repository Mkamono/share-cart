import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { LoadingSpinner } from "@/components/ui/loading-spinner";
import { getSession } from "@auth0/nextjs-auth0";
import Link from "next/link";
import { Suspense } from "react";
import { ModeToggle } from "./ModeToggle";
import { UserDropdownMenu } from "./UserDropdownMenu";

export default function Header() {
	return (
		<Suspense fallback={<Loading />}>
			<UserProfile />
		</Suspense>
	);
}

function Loading() {
	return (
		<div className="h-full flex">
			<LoadingSpinner className="m-auto" />
		</div>
	);
}

async function UserProfile() {
	const session = await getSession();
	if (!session) {
		return (
			<div className="flex items-baseline px-4">
				<Link
					href={"/home"}
					className="text-2xl font-semibold leading-none tracking-tight py-4"
				>
					Share Cart
				</Link>
				<div className="flex-auto" />
				<ModeToggle />
				<form action="/api/auth/login" method="GET">
					<Button type="submit">Login</Button>
				</form>
			</div>
		);
	}
	const user = session.user;
	return (
		<div className="flex items-baseline px-4">
			<Link
				href={"/home"}
				className="text-2xl font-semibold leading-none tracking-tight py-4"
			>
				Share Cart
			</Link>
			<div className="flex-auto" />
			<div>
				<div className="flex">
					<div className="flex-auto" />
					<ModeToggle />
					<UserDropdownMenu>
						<Avatar>
							<AvatarImage src={user.picture} />
							<AvatarFallback>{user.name[0]}</AvatarFallback>
						</Avatar>
					</UserDropdownMenu>
				</div>
				<p className="text-right">{user.name}</p>
			</div>
		</div>
	);
}

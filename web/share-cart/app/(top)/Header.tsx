import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { getSession } from "@auth0/nextjs-auth0";
import { Suspense } from "react";
import { ModeToggle } from "./ModeToggle";

export default function Header() {
	return (
		<Suspense fallback={<Loading />}>
			<UserProfile />
		</Suspense>
	);
}

function Loading() {
	return <div>Loading...</div>;
}

async function UserProfile() {
	const session = await getSession();
	if (!session) {
		return (
			<p>
				<a href="/api/auth/login">Login</a>
			</p>
		);
	}
	const user = session.user;
	return (
		<div className="flex">
			<ModeToggle />
			<div>
				<Avatar>
					<AvatarImage src={user.picture} />
					<AvatarFallback>{user.name[0]}</AvatarFallback>
				</Avatar>
				<p>{user.name}</p>
				<p>
					<a href="/api/auth/logout">Logout</a>
				</p>
			</div>
		</div>
	);
}

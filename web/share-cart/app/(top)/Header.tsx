import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { getSession } from "@auth0/nextjs-auth0";
import Link from "next/link";
import { ModeToggle } from "./ModeToggle";
import { UserDropdownMenu } from "./UserDropdownMenu";

export function Header() {
	return <UserProfile />;
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
				<Button type="submit">
					<a href="/api/auth/login">
						Login
					</a>
				</Button>
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

import { Form } from "@remix-run/react";
import { $path } from "remix-routes";

export default function Logout() {
	return (
		<Form action={$path("/auth/auth0/logout")} method="post">
			<button type="submit">Logout with Auth0</button>
		</Form>
	);
}

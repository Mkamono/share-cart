import { Form } from "@remix-run/react";

export default function Logout() {
	return (
		<Form action="/auth/auth0/logout" method="post">
			<button type="submit">Logout with Auth0</button>
		</Form>
	);
}

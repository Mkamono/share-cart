import { Form } from "@remix-run/react";

export default function Login() {
	return (
		<Form action="/auth/auth0/login" method="post">
			<button type="submit">Login with Auth0</button>
		</Form>
	);
}

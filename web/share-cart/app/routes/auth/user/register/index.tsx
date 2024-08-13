import { ActionFunction, ActionFunctionArgs } from "@remix-run/node";
import { Form, redirect } from "@remix-run/react";
import { $path } from "remix-routes";
import { authenticator } from "~/service/auth.server";
import { shareCartClient } from "~/service/client";

export async function loader({ request }: { request: Request }) {
	const jwt = await authenticator.isAuthenticated(request);
	if (!jwt) {
		return redirect($path("/"));
	}
	return {};
}

export const action: ActionFunction = async ({
	request,
}: ActionFunctionArgs) => {
	const jwt = await authenticator.isAuthenticated(request);
	const client = shareCartClient(jwt?.accessToken);
	const formData = await request.formData();
	const data = await client.POST("/sign-up", {
		body: { name: formData.get("name")?.toString() },
	});

	if (data.error) {
		console.error(data.error);
		return redirect($path("/auth/user/register"));
	}

	console.log(data);
	return redirect($path("/"));
};

export default function SignUp() {
	return (
		<Form action={$path("/auth/user/register")} method="post">
			<input type="text" name="name" id="name" />
			<button type="submit">Sign Up</button>
		</Form>
	);
}

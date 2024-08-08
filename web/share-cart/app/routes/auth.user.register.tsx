import { ActionFunction, ActionFunctionArgs } from "@remix-run/node";
import { Form, redirect } from "@remix-run/react";
import createClient from "openapi-fetch";
import type { paths } from "~/models/schema";
import { authenticator } from "~/service/auth.server";

export async function loader({ request }: { request: Request }) {
	const jwt = await authenticator.isAuthenticated(request);
	if (!jwt) {
		return redirect("/auth/login");
	}
	return {};
}

export const action: ActionFunction = async ({
	request,
}: ActionFunctionArgs) => {
	const jwt = await authenticator.isAuthenticated(request);
	const client = createClient<paths>({ baseUrl: process.env.API_HOST });
	const formData = await request.formData();
	const data = await client.POST("/sign-up", {
		params: { header: { Authorization: jwt?.accessToken ?? "" } },
		body: { name: formData.get("name")?.toString() },
	});

	if (data.error) {
		console.error(data.error);
		return redirect("/auth/user/register");
	}

	console.log(data);
	return redirect("/");
};

export default function SignUp() {
	return (
		<Form action="/auth/user/register" method="post">
			<input type="text" name="name" id="name" />
			<button type="submit">Sign Up</button>
		</Form>
	);
}

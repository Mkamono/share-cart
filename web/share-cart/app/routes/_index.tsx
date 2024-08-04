import { type MetaFunction, json } from "@remix-run/node";
import { Form, useLoaderData } from "@remix-run/react";
import createClient from "openapi-fetch";
import type { paths } from "~/models/schema";
import { authenticator } from "~/service/auth.server";
import { config } from "~/service/config.server";
import { jwtMiddleware } from "~/service/middleware";
import Login from "./login";


export async function loader({ request }: { request: Request }) {
	const client = createClient<paths>({ baseUrl: process.env.API_HOST });
	const jwt = await authenticator.isAuthenticated(request);
	client.use(jwtMiddleware(jwt?.accessToken ?? ""));
	const { data } = await client.GET("/test");
	return json({
		testData: data,
		config,
		jwt,
	});
}

export const meta: MetaFunction = () => {
	return [
		{ title: "New Remix App" },
		{ name: "description", content: "Welcome to Remix!" },
	];
};

export default function Index() {
	const data = useLoaderData<typeof loader>();
	return (
		<div className="font-sans p-4">
			<h1 className="text-3xl">Welcome to Remix</h1>
			<ul className="list-disc mt-4 pl-6 space-y-2">
				<li>
					<p>{data.testData?.message}</p>
					{data.config.NODE_ENV === "development" && (
						<p className="text-red-700">You are in development mode</p>
					)}
					<Login />
					<Form action="/auth/logout" method="post">
						<button type="submit">Logout with Auth0</button>
					</Form>
					{data.jwt ? (
						<div>
							<p className="text-green-700">You are logged in</p>
							<p>{data.jwt.profile.id}</p>
							<p>{data.jwt.profile.displayName}</p>
							{data.jwt?.profile.emails?.map((email) => (
								<p key={email.value}>{email.value}</p>
							))}
							<p>{data.jwt.accessToken}</p>
						</div>
					) : (
						<p className="text-red-700">You are not logged in</p>
					)}
				</li>
			</ul>
		</div>
	);
}

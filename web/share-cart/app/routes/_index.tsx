import { type MetaFunction, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import createClient from "openapi-fetch";
import type { paths } from "~/models/schema";
import { authenticator } from "~/service/auth.server";
import { config } from "~/service/config.server";
import { getSession } from "~/service/session.server";
import Login from "./login";

export async function loader({ request }: { request: Request }) {
	const user = await authenticator.isAuthenticated(request);
	const email = user?.profile.emails?.[0].value;
	const session = await getSession();
	const client = createClient<paths>({ baseUrl: process.env.API_HOST });
	const { data } = await client.GET("/test", {});
	return json({
		testData: data,
		config,
		session,
		user,
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
					{data.session ? (
						<div>
							<p className="text-green-700">You are logged in</p>
							<p>{data.user?.profile.id}</p>
							<p>{data.user?.profile.displayName}</p>
							{data.user?.profile.emails?.map((email) => (
								<p key={email.value}>{email.value}</p>
							))}
						</div>
					) : (
						<p className="text-red-700">You are not logged in</p>
					)}
				</li>
			</ul>
		</div>
	);
}

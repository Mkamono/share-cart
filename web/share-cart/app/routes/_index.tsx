import { type MetaFunction, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import createClient from "openapi-fetch";
import type { paths } from "~/models/schema";

export async function loader() {
	const client = createClient<paths>({ baseUrl: process.env.API_HOST });
	const { data } = await client.GET("/test", {});
	return json(data);
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
					<a
						className="text-blue-700 underline visited:text-purple-900"
						target="_blank"
						href="https://remix.run/start/quickstart"
						rel="noreferrer"
					>
						5m Quick Start
					</a>
					<p>{data.message}</p>
				</li>
				<li>
					<a
						className="text-blue-700 underline visited:text-purple-900"
						target="_blank"
						href="https://remix.run/start/tutorial"
						rel="noreferrer"
					>
						30m Tutorial
					</a>
				</li>
				<li>
					<a
						className="text-blue-700 underline visited:text-purple-900"
						target="_blank"
						href="https://remix.run/docs"
						rel="noreferrer"
					>
						Remix Docs
					</a>
				</li>
			</ul>
		</div>
	);
}

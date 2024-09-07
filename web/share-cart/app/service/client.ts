import createClient from "openapi-fetch";

import type { paths } from "~/models/schema";
import { config } from "~/service/config.server";

export const shareCartClient = (accessToken: string | undefined) => {
	const client = createClient<paths>({ baseUrl: config.API_HOST });
	client.use({
		async onRequest({ request, options }) {
			request.headers.set("Authorization", `Bearer ${accessToken ?? "empty"}`);
			return request;
		},
	});
	return client;
};

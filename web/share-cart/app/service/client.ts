import createClient from "openapi-fetch";

import type { paths } from "~/models/schema";
import { config } from "~/service/config.server";

export const shareCartClient = (accessToken: string | undefined) => {
	return createClient<paths>({
		baseUrl: config.API_HOST,
		fetch: async (
			input: string | URL | globalThis.Request,
			init?: RequestInit,
		): Promise<Response> => {
			return fetch(input, {
				...init,
				headers: new Headers({
					...init?.headers,
					Authorization: `Bearer ${accessToken ?? "empty"}`,
				}),
			});
		},
	});
};

"use server";
import type { paths } from "@/lib/share-cart";
import { shareCartClient } from "@/lib/share-cart-client";
import { getSession } from "@auth0/nextjs-auth0";

export async function getMarkets(): Promise<
	Result<
		paths["/market"]["get"]["responses"]["200"]["content"]["application/json"]
	>
> {
	const session = await getSession();
	const client = shareCartClient(session?.accessToken);
	const res = await client.GET("/market");
	if (!res.response.ok) {
		return {
			data: [],
			error: res.error?.message || "Failed to get markets",
		};
	}

	return {
		data: res.data ?? [],
		error: undefined,
	};
}

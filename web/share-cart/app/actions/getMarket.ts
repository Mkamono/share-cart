"use server";
import type { paths } from "@/lib/share-cart";
import { shareCartClient } from "@/lib/share-cart-client";
import { getSession } from "@auth0/nextjs-auth0";

export type Input = {
	id: string;
};

export async function getMarket(
	input: Input,
): Promise<
	Result<
		paths["/market/{marketId}"]["get"]["responses"]["200"]["content"]["application/json"]
	>
> {
	const session = await getSession();
	const client = shareCartClient(session?.accessToken);
	const res = await client.GET("/market/{marketId}", {
		params: { path: { marketId: input.id } },
	});

	const emptyData = {
		id: "",
		name: "",
		description: "",
		images: [],
	};

	if (!res.response.ok) {
		return {
			data: emptyData,
			error: res.error?.message || "Failed to get market",
		};
	}

	return {
		data: res.data ?? emptyData,
		error: undefined,
	};
}

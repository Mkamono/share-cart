"use server";
import { shareCartClient } from "@/lib/share-cart-client";
import { getSession } from "@auth0/nextjs-auth0";

export type Input = {
	id: string;
};

export async function deleteMarket(input: Input): Promise<Result<string>> {
	const session = await getSession();
	const client = shareCartClient(session?.accessToken);
	const res = await client.DELETE("/market/{marketId}", {
		params: { path: { marketId: input.id } },
	});
	if (!res.response.ok) {
		return {
			data: "Failed to delete market",
			error: res.error?.message || "Failed to delete market",
		};
	}
	return {
		data: "Market deleted successfully",
		error: undefined,
	};
}

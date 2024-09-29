"use server";
import { shareCartClient } from "@/lib/share-cart-client";
import { getSession } from "@auth0/nextjs-auth0";

export type Input = {
	id: string;
};

export async function getMarket(input: Input): Promise<Result<undefined>> {
	const session = await getSession();
	const client = shareCartClient(session?.accessToken);
	const res = await client.GET("/market/{marketId}", {
		params: { path: { marketId: input.id } },
	});
	if (!res.response.ok) {
		return {
			data: undefined,
			error: res.error?.message,
		};
	}
	return {
		data: undefined,
		error: undefined,
	};
}

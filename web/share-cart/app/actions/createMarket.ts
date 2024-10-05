"use server";
import { shareCartClient } from "@/lib/share-cart-client";
import { getSession } from "@auth0/nextjs-auth0";
import { revalidateShareCartCache } from "./revalidatePath";

type Input = {
	name: string;
	description: string;
};

export async function createMarket(input: Input): Promise<Result<string>> {
	const session = await getSession();
	const client = shareCartClient(session?.accessToken);
	const res = await client.POST("/market", { body: input });
	if (!res.response.ok) {
		return {
			data: "Failed to create market",
			error: res.error?.message || "Failed to create market",
		};
	}

	revalidateShareCartCache("/market");
	return {
		data: res.data?.id ?? "",
		error: undefined,
	};
}

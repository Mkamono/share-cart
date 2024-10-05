"use server";
import { config } from "@/lib/config";
import type { paths } from "@/lib/share-cart";
import { revalidatePath } from "next/cache";

export async function revalidateShareCartCache(path: keyof paths) {
	revalidatePath(`${config.API_HOST}${path}`);
}

"use server";
import { LoadingSpinner } from "../custom-ui/loading-spinner";

export async function LoadingPage() {
	return (
		<div className="h-full flex">
			<LoadingSpinner className="m-auto" />
		</div>
	);
}

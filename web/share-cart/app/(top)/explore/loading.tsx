import { LoadingSpinner } from "@/components/ui/loading-spinner";

export default function Loading() {
	return (
		<div className="h-full flex">
			<LoadingSpinner className="m-auto" />
		</div>
	);
}

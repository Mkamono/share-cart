import { getMarket } from "@/app/actions/getMarket";
import { ErrorPage } from "@/components/page/ErrorPage";

export default async function Home() {
	const res = await getMarket({ id: "tes" });

	if (res.error) {
		return <ErrorPage />;
	}

	return (
		<div className="h-full border-4">
			<h1>test</h1>
		</div>
	);
}

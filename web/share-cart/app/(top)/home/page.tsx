import { getMarkets } from "@/app/actions/getMarkets";
import { ErrorPage } from "@/components/page/ErrorPage";
import { Market } from "./Market";

export default async function Home() {
	const res = await getMarkets();
	if (res.error) {
		return <ErrorPage errorMessage="マーケットが見つかりませんでした" />;
	}
	const markets = res.data;
	return (
		<div className="flex flex-col space-y-4">
			{markets.map((market) => (
				<div key={market.id}>
					<Market market={market} />
				</div>
			))}
		</div>
	);
}

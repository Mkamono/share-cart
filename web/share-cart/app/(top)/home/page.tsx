import { getMarkets } from "@/app/actions/getMarkets";
import { Market } from "./Market";

export default async function Home() {
	const res = await getMarkets();
	if (res.error) {
		return <div>{res.error}</div>;
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

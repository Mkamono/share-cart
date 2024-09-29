import { getMarkets } from "@/app/actions/getMarkets";

export default async function Home() {
	const res = await getMarkets();
	if (res.error) {
		return <div>{res.error}</div>;
	}
	const markets = res.data;
	return (
		<div className="flex flex-col space-y-4">
			{markets.map((market) => (
				<div key={market.id} className="p-4 rounded-md">
					<h2 className="text-lg font-bold">{market.name}</h2>
					<p>{market.description}</p>
				</div>
			))}
		</div>
	);
}

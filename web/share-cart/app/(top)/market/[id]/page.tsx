import { getMarket } from "@/app/actions/getMarket";
import { DeleteMarketButton } from "./DeleteMarketButton";

export default async function Home({ params }: { params: { id: string } }) {
	const res = await getMarket({ id: params.id });

	if (res.error) {
		return <h1>{res.error}</h1>;
	}

	const { data: market } = res;

	return (
		<div>
			<DeleteMarketButton id={market.id} />
			<h1>{market.name}</h1>
			<p>{market.description}</p>
		</div>
	);
}

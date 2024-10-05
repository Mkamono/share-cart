import { getMarket } from "@/app/actions/getMarket";
import { ErrorPage } from "@/components/page/ErrorPage";
import { DeleteMarketButton } from "./DeleteMarketButton";

export default async function Home({ params }: { params: { id: string } }) {
	const res = await getMarket({ id: params.id });

	if (res.error) {
		return <ErrorPage errorMessage="マーケットが見つかりませんでした" />;
	}

	const { data: market } = res;

	return (
		<div>
			<DeleteMarketButton id={market.id} />
			<h1>{market.name || "マーケット名なし"}</h1>
			<p>{market.description || "説明なし"}</p>
		</div>
	);
}

import { getMarket } from "@/app/actions/getMarket";

export default async function Home({ params }: { params: { id: string } }) {
	const res = await getMarket({ id: params.id });

	if (res.error) {
		return <h1>{res.error}</h1>;
	}

	const { data } = res;

	return (
		<div>
			<h1>{data.name}</h1>
			<p>{data.description}</p>
		</div>
	);
}

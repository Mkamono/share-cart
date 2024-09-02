import { json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { authenticator } from "~/service/auth.server";
import { shareCartClient } from "~/service/client";
import { MarketCard } from "./MarketCard";

export async function loader({ request }: { request: Request }) {
	const jwt = await authenticator.isAuthenticated(request);
	const client = shareCartClient(jwt?.accessToken);
	const { data } = await client.GET("/market");
	return json({ markets: data });
}

const testImageURL =
	"https://images.unsplash.com/photo-1617050318658-a9a3175e34cb?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=600&q=80";


export const Home = () => {
	const { markets } = useLoaderData<typeof loader>();
	return (
		<div>
			<h2 className="text-2xl font-bold tracking-tight text-gray-900">
				Your Markets
			</h2>
			<div className="grid grid-cols-2 gap-4">
				{markets?.map((market) => (
					<div key={market.id} className="rounded-md">
						<MarketCard market={{ ...market, imageURL: testImageURL }} />
					</div>
				))}
			</div>
		</div>
	);
};

export default Home;

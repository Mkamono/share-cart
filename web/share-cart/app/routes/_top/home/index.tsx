import { ActionFunctionArgs, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { useState } from "react";
import { authenticator } from "~/service/auth.server";
import { shareCartClient } from "~/service/client";
import { CreateNewMarketDialog, marketName } from "./CreateNewMarketModal";
import { MarketCard } from "./MarketCard";

export async function loader({ request }: { request: Request }) {
	const jwt = await authenticator.isAuthenticated(request);
	const client = shareCartClient(jwt?.accessToken);
	const { data } = await client.GET("/market");
	return json({ markets: data });
}

export async function action({ request }: ActionFunctionArgs) {
	const formData = await request.formData();
	const jwt = await authenticator.isAuthenticated(request);
	const client = shareCartClient(jwt?.accessToken);

	const { data } = await client.POST("/market", {
		body: {
			name: formData.get(marketName)?.toString() ?? "",
			description: formData.get("description")?.toString() ?? "",
		},
	});
	return json({ message: data });
}

const defaultImageURL =
	"https://cdn.pixabay.com/photo/2021/01/01/12/44/concert-5878452_640.jpg";

export default function Home() {
	const { markets } = useLoaderData<typeof loader>();
	const [open, setOpen] = useState(false);
	return (
		<div>
			<div className="flex py-2">
				<h2 className="text-2xl font-bold tracking-tight text-gray-900">
					Your Markets
				</h2>
				<div className="flex-grow" />
				<button
					className="bg-blue-500 text-white px-4 rounded-lg"
					onClick={() => {
						setOpen(true);
					}}
				>
					New Market
				</button>
				<CreateNewMarketDialog
					open={open}
					handleClose={() => {
						setOpen(false);
					}}
				/>
			</div>
			<div className="grid grid-cols-2 gap-4">
				{markets?.map((market) => (
					<div key={market.id} className="rounded-md">
						<MarketCard
							market={{
								...market,
								imageURL: market.images[0] || defaultImageURL,
							}}
						/>
					</div>
				))}
			</div>
		</div>
	);
}

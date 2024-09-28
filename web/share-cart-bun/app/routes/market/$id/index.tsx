import type { ActionFunctionArgs, LoaderFunctionArgs } from "@remix-run/node";
import { json, redirect } from "@remix-run/node";
import { Form, Link, useLoaderData } from "@remix-run/react";
import { authenticator } from "~/service/auth.server";
import { shareCartClient } from "~/service/client";

// marketの読み込み
export async function loader({ params, request }: LoaderFunctionArgs) {
	if (!params.id) {
		return redirect("/home");
	}
	const jwt = await authenticator.isAuthenticated(request);
	const client = shareCartClient(jwt?.accessToken);
	const { data } = await client.GET("/market/{marketId}", {
		params: { path: { marketId: params.id } },
	});

	if (!data) {
		return redirect("/home");
	}

	return json({ market: data });
}

// marketの削除
// TODO: 編集
export async function action({ params, request }: ActionFunctionArgs) {
	const formData = await request.formData();
	const jwt = await authenticator.isAuthenticated(request);
	const client = shareCartClient(jwt?.accessToken);

	if (formData.get("_method") == "delete") {
		if (!params.id) {
			return redirect("/home");
		}
		const res = await client.DELETE("/market/{marketId}", {
			params: { path: { marketId: params.id } },
		});
		if (!res.response.ok) {
			return json({ message: "Failed to delete market" }, { status: 500 });
		}
		return redirect("/home");
	}
	return redirect("/home");
}

export default function Home() {
	const { market } = useLoaderData<typeof loader>();
	return (
		<>
			<div className="flex py-2">
				<Link to="/home">
					<button className="bg-blue-500 text-white px-4 rounded-lg">
						Back
					</button>
				</Link>
				<div className="flex-grow" />
				<Form method="post">
					<input type="hidden" name="_method" value="delete" />
					<button
						type="submit"
						className="bg-red-500 text-white px-4 rounded-lg"
					>
						Delete
					</button>
				</Form>
			</div>
			<MarketCard market={market} />
		</>
	);
}

type MarketCardProps = {
	market: {
		id: string;
		name: string;
		description: string;
		images: string[];
	};
};

const noImageURL = "https://placehold.jp/150x150.png";
const defaultImageURL =
	"https://cdn.pixabay.com/photo/2021/01/01/12/44/concert-5878452_640.jpg";

function MarketCard(props: MarketCardProps) {
	if (!props.market.images.length) {
		props.market.images = [defaultImageURL];
	}

	return (
		<div>
			<div className="relative">
				<img
					src={props.market.images[0]}
					className="object-cover w-full h-32"
					alt="market"
					onError={(e) => {
						(e.target as HTMLImageElement).src = noImageURL;
					}}
				/>
				<div className="absolute inset-x-0 bottom-0 h-1/3 bg-gradient-to-b from-transparent to-black rounded-b-2xl">
					<p className="text-gray-200 absolute left-4 bottom-2">
						<strong className="text-xl">{props.market.name}</strong>
					</p>
				</div>
			</div>
			<p className="text-gray-500">{props.market.description}</p>
		</div>
	);
}

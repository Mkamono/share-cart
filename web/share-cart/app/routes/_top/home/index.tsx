import { Box, Flex } from "@radix-ui/themes";
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
		<Flex gap="2" wrap="wrap" pb="9">
			{markets?.map((market) => (
				<Box key={market.id} className="flex-grow w-28">
					<MarketCard
						key={market.id}
						market={{ ...market, imageURL: testImageURL }}
					/>
				</Box>
			))}
		</Flex>
	);
};

export default Home;

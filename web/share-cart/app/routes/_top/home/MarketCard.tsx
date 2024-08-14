import { Avatar, Card, Inset, Strong, Text } from "@radix-ui/themes";

type MarketCardProps = {
	market: {
		id: number;
		name: string;
		imageURL: string;
		description: string;
	};
};

const noImageURL =
	"https://placehold.jp/eeeeee/cccccc/240x140.png?text=No%20Image";

export const MarketCard = (props: MarketCardProps) => {
	return (
		<Card size="2">
			<Inset clip="padding-box" side="top" pb="current" className="h-44">
				<Avatar
					src={props.market.imageURL}
					alt={props.market.name}
					fallback={
						<img
							src={noImageURL}
							alt={props.market.name}
							className="w-full h-full object-cover"
						/>
					}
					className="w-full h-full object-cover"
				/>
			</Inset>
			<Text as="p" size="1">
				<Strong>{props.market.name}</Strong>
			</Text>
		</Card>
	);
};

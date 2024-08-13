import * as Avatar from "@radix-ui/react-avatar";
import { Card, Inset, Strong, Text } from "@radix-ui/themes";

type MarketCardProps = {
	market: {
		id: number;
		name: string;
		imageURL: string;
		description: string;
	};
};

export const MarketCard = (props: MarketCardProps) => {
	return (
		<Card size="2">
			<Inset clip="padding-box" side="top" pb="current" className="h-44">
				<Avatar.Root className="AvatarRoot">
					<Avatar.Image
						className="AvatarImage"
						src={props.market.imageURL}
						alt={props.market.name}
						style={{
							display: "block",
							objectFit: "cover",
							width: "100%",
							height: "100%",
							backgroundColor: "var(--gray-5)",
						}}
					/>
					<Avatar.Fallback className="AvatarFallback" delayMs={500}>
						<img
							className="AvatarImage"
							src="https://placehold.jp/eeeeee/cccccc/240x180.png?text=No%20Image"
							alt={props.market.name}
							style={{
								display: "block",
								objectFit: "cover",
								width: "100%",
								height: "100%",
								backgroundColor: "var(--gray-5)",
							}}
						/>
					</Avatar.Fallback>
				</Avatar.Root>
			</Inset>
			<Text as="p" size="2">
				<Strong>{props.market.name}</Strong>
			</Text>
		</Card>
	);
};

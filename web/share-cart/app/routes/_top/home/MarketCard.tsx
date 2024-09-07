type MarketCardProps = {
	market: {
		id: string;
		name: string;
		imageURL: string;
		description: string;
	};
};
const noImageURL = "https://placehold.jp/150x150.png";

export const MarketCard = (props: MarketCardProps) => {
	return (
		<div>
			<img
				src={props.market.imageURL}
				className="object-cover w-full h-32"
				alt="market"
				onError={(e) => {
					(e.target as HTMLImageElement).src = noImageURL;
				}}
			/>
			<p>
				<strong>{props.market.name}</strong>
			</p>
			<p className="text-gray-500">{props.market.description}</p>
		</div>
	);
};

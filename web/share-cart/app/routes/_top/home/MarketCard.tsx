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
		<div>
			<img
				src={props.market.imageURL}
				className="object-cover w-full h-32"
				alt="market"
			/>
			<p>{props.market.name}</p>
		</div>
	);
};

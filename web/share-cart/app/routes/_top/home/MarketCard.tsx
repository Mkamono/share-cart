type MarketCardProps = {
	market: {
		id: number;
		name: string;
		imageURL: string;
		description: string;
	};
};
const defaultImageURL = "https://cdn.pixabay.com/photo/2021/01/01/12/44/concert-5878452_640.jpg";

export const MarketCard = (props: MarketCardProps) => {
	return (
		<div>
			<img
				src={props.market.imageURL}
				className="object-cover w-full h-32"
				alt="market"
				onError={(e) => {
					(e.target as HTMLImageElement).src =
						defaultImageURL;
				}}
			/>
			<p>
				<strong>{props.market.name}</strong>
			</p>
		</div>
	);
};

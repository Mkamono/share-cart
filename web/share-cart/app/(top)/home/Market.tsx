import {
	Card,
	CardDescription,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import Image from "next/image";
import Link from "next/link";

type Props = {
	market: {
		id: string;
		name: string;
		description: string;
		images: string[];
	};
};

const altImage = "https://via.placeholder.com/300x200/?text=No+Image";

export function Market(props: Props) {
	const { market } = props;
	return (
		<Link href={`/market/${market.id}`}>
			<Card className="overflow-hidden">
				<Image
					src={market.images[0] || altImage}
					className="w-full h-full object-contain"
					width={300}
					height={200}
					alt=""
				/>
				<CardHeader>
					<CardTitle>{market.name}</CardTitle>
					<CardDescription>{market.description}</CardDescription>
				</CardHeader>
			</Card>
		</Link>
	);
}

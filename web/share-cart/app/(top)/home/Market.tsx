import {
	Card,
	CardContent,
	CardDescription,
	CardFooter,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import Link from "next/link";

type Props = {
	market: {
		id: string;
		name: string;
		description: string;
		images: string[];
	};
};

export function Market(props: Props) {
	const { market } = props;
	return (
		<Link href={`/market/${market.id}`}>
			<Card>
				<CardHeader>
					<CardTitle>{market.name}</CardTitle>
					<CardDescription>{market.description}</CardDescription>
				</CardHeader>
				<CardContent>
					<p>Card Content</p>
				</CardContent>
				<CardFooter>
					<p>Card Footer</p>
				</CardFooter>
			</Card>
		</Link>
	);
}

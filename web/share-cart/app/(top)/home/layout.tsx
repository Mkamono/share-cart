import { CreateMarketDialogButton } from "./Button"; // server component
import { CreateMarketDialog } from "./Dialog"; // client component

export default function TopPageLayout({
	children,
}: { children: React.ReactNode }) {
	return (
		<section className="relative h-full p-2">
			<div className="overflow-auto h-full pb-20">
				{children}
			</div>
			<div className="absolute bottom-4 right-4">
				<CreateMarketDialog>
					<CreateMarketDialogButton />
				</CreateMarketDialog>
			</div>
		</section>
	);
}

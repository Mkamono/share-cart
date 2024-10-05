import { CreateMarketDialogButton } from "./Button"; // server component
import { CreateMarketDialog } from "./Dialog"; // client component

export default function TopPageLayout({
	children,
}: { children: React.ReactNode }) {
	return (
		<section className="relative h-full flex flex-col">
			<div className="flex-auto overflow-auto">
				{children}
				<div className="h-16" />
			</div>
			<div className="absolute bottom-4 right-4">
				<CreateMarketDialog>
					<CreateMarketDialogButton />
				</CreateMarketDialog>
			</div>
		</section>
	);
}

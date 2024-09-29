import { CreateMarketDialogButton } from "./Button";
import { CreateMarketDialog } from "./Dialog";

export default function TopPageLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<section className="relative h-full flex flex-col">
			<div className="flex-auto overflow-auto">
				{children}
				<div className="h-16" />
			</div>
			<div className="absolute bottom-0 right-0">
				<CreateMarketDialog>
					<CreateMarketDialogButton />
				</CreateMarketDialog>
			</div>
		</section>
	);
}

import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";

export function CreateMarketDialogButton({
	className,
}: { className?: string }) {
	const classes = `rounded-full ${className}`;

	return (
		<Button variant="default" size="icon" className={classes}>
			<Plus className="h-4 w-4" />
		</Button>
	);
}

import { Button } from "@/components/ui/button";
import { Plus } from "lucide-react";

export function CreateMarketDialogButton() {
	return (
		<Button variant="default" size="icon" className="rounded-full">
			<Plus className="h-4 w-4" />
		</Button>
	);
}

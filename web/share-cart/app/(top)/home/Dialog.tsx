import { Button } from "@/components/ui/button";
import {
	Dialog,
	DialogClose,
	DialogContent,
	DialogDescription,
	DialogFooter,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "@/components/ui/dialog";
import type React from "react";
import { CreateMarketForm, CreateMarketFormSubmitButton } from "./Form";

export type DialogProps = {
	triggerButton: React.ReactNode;
};

export function CreateMarketDialog({
	children, // dialog button
}: { children: React.ReactNode }) {
	return (
		<Dialog>
			<DialogTrigger asChild>{children}</DialogTrigger>
			<DialogContent>
				<DialogHeader>
					<DialogTitle>マーケットを作成する</DialogTitle>
					<DialogDescription>
						マーケットを作成して他のユーザーとシェア
					</DialogDescription>
				</DialogHeader>
				<CreateMarketForm />
				<DialogFooter className="flex-row">
					<DialogClose asChild>
						<Button type="button" variant="secondary" className="flex-1">
							キャンセル
						</Button>
					</DialogClose>
					<CreateMarketFormSubmitButton variant="default" className="flex-1">
						作成
					</CreateMarketFormSubmitButton>
				</DialogFooter>
			</DialogContent>
		</Dialog>
	);
}

"use client";

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
import { useRef } from "react";
import { CreateMarketForm, CreateMarketFormSubmitButton } from "./Form";

export type DialogProps = {
	triggerButton: React.ReactNode;
};

export function CreateMarketDialog({
	children, // dialog button
}: { children: React.ReactNode }) {
	const closeDialogRef = useRef<HTMLButtonElement>(null);
	const afterSubmit = () => {
		closeDialogRef.current?.click();
	};

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
				<CreateMarketForm afterSubmit={afterSubmit} />
				<DialogFooter className="flex-row gap-4">
					<DialogClose asChild>
						<Button type="button" variant="secondary" className="flex-1">
							キャンセル
						</Button>
					</DialogClose>
					<CreateMarketFormSubmitButton variant="default" className="flex-1">
						作成
					</CreateMarketFormSubmitButton>
					<DialogClose hidden ref={closeDialogRef} />
				</DialogFooter>
			</DialogContent>
		</Dialog>
	);
}

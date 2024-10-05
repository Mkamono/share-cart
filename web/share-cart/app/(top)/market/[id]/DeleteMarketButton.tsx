"use client";
import { deleteMarket } from "@/app/actions/deleteMarket";
import { Button } from "@/components/ui/button";
import { useToast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";
import { useState } from "react";

export function DeleteMarketButton({ id }: { id: string }) {
	const { toast } = useToast();
	const [buttonVariant, setButtonVariant] = useState<"link" | "default" | "destructive" | "outline" | "secondary" | "ghost" | null>(
		"destructive",
	);
	const router = useRouter();

	const handleClick = async () => {
		setButtonVariant("outline");
		deleteMarket({ id }).then((res) => {
			if (res.error) {
				toast({
					title: "Failed to delete market",
					description: res.error,
					variant: "destructive",
				});
			} else {
				router.push("/home");
			}
		});
	};
	return (
		<Button variant={buttonVariant} onClick={handleClick}>
			削除
		</Button>
	);
}

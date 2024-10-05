"use client";
import { deleteMarket } from "@/app/actions/deleteMarket";
import { Button } from "@/components/ui/button";
import { useToast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";
import { useState } from "react";

export function DeleteMarketButton({ id }: { id: string }) {
	const { toast } = useToast();
	const [buttonDisabled, setButtonDisabled] = useState(false);
	const router = useRouter();

	const handleClick = async () => {
		setButtonDisabled(true);
		const res = await deleteMarket({ id });
		if (res.error) {
			toast({
				title: "Failed to delete market",
				description: res.error,
				variant: "destructive",
			});
			setButtonDisabled(false);
		} else {
			router.push("/home");
		}
	};
	return (
		<Button variant="destructive" disabled={buttonDisabled} onClick={handleClick}>
			削除
		</Button>
	);
}

"use client";
import { deleteMarket } from "@/app/actions/deleteMarket";
import { Button } from "@/components/ui/button";
import { useToast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";

export function DeleteMarketButton({ id }: { id: string }) {
	const { toast } = useToast();
	const router = useRouter();

	const handleClick = () => {
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
		<Button variant="destructive" onClick={handleClick}>
			削除
		</Button>
	);
}

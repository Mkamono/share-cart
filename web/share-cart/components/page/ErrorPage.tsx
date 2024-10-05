"use client";
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
} from "@/components/ui/card";
import { House, MoveLeft, RefreshCcw } from "lucide-react";
import { useRouter } from "next/navigation";
import { useRef } from "react";
import {
	ReportBugForm,
	ReportBugSubmitButton,
} from "../custom-ui/report-bug-form";
import { Button } from "../ui/button";
import {
	Dialog,
	DialogClose,
	DialogContent,
	DialogDescription,
	DialogFooter,
	DialogHeader,
	DialogTitle,
	DialogTrigger,
} from "../ui/dialog";

export function ErrorPage({ errorMessage }: { errorMessage?: string }) {
	const router = useRouter();

	return (
		<div className="h-full p-8 content-center">
			<Card className="bg-secondary">
				<CardHeader>
					<CardTitle>Error</CardTitle>
					<CardDescription>
						{errorMessage || "エラーが発生しました"}
					</CardDescription>
				</CardHeader>
				<CardContent>
					<Button className="w-full" onClick={() => router.refresh()}>
						<RefreshCcw className="mr-4" />
						ページを再読込する
					</Button>
					<Button className="w-full mt-4" onClick={() => router.back()}>
						<MoveLeft className="mr-4" />
						戻る
					</Button>
					<Button className="w-full mt-4" onClick={() => router.push("/")}>
						<House className="mr-4" />
						トップページへ
					</Button>
					<ErrorReportForm>
						<Button className="w-full mt-4" variant="outline">
							不具合を報告する
						</Button>
					</ErrorReportForm>
				</CardContent>
			</Card>
		</div>
	);
}

function ErrorReportForm({ children }: { children: React.ReactNode }) {
	const closeDialogRef = useRef<HTMLButtonElement>(null);
	const afterSubmit = () => {
		closeDialogRef.current?.click();
	};
	return (
		<Dialog>
			<DialogTrigger asChild>{children}</DialogTrigger>
			<DialogContent>
				<DialogHeader>
					<DialogTitle>不具合を報告する</DialogTitle>
					<DialogDescription>
						送信された情報は、問題の特定と解決に役立てられます。
					</DialogDescription>
				</DialogHeader>

				<ReportBugForm afterSubmit={afterSubmit} />

				<DialogFooter className="flex-row gap-4">
					<DialogClose asChild className="flex-1">
						<Button type="button" variant="secondary">
							キャンセル
						</Button>
					</DialogClose>
					<ReportBugSubmitButton className="flex-1">送信</ReportBugSubmitButton>
					<DialogClose hidden ref={closeDialogRef} />
				</DialogFooter>
			</DialogContent>
		</Dialog>
	);
}

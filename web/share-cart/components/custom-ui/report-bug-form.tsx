"use client";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form";
import { zodResolver } from "@hookform/resolvers/zod";
import { usePathname } from "next/navigation";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button, type ButtonProps } from "../ui/button";
import { Input } from "../ui/input";
import { Textarea } from "../ui/textarea";

const FormSchema = z.object({
	currentURL: z.string(),
	report: z.string().min(10, {
		message: "10文字以上で入力してください",
	}),
});

const reportBugFormID = "report-bug";

export function ReportBugForm({
	afterSubmit,
}: { afterSubmit?: () => void } = {}) {
	const pathname = usePathname();
	const form = useForm<z.infer<typeof FormSchema>>({
		resolver: zodResolver(FormSchema),
		defaultValues: {
			currentURL: pathname,
			report: "",
		},
	});

	function onSubmit(data: z.infer<typeof FormSchema>) {
		console.log(data);
		if (afterSubmit) {
			afterSubmit();
		}
	}

	return (
		<Form {...form}>
			<form
				onSubmit={form.handleSubmit(onSubmit)}
				className="w-full space-y-6"
				id={reportBugFormID}
			>
				<FormField
					control={form.control}
					name="currentURL"
					render={({ field }) => (
						<FormItem>
							<FormLabel>現在のURL</FormLabel>
							<FormControl>
								<Input readOnly disabled {...field} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="report"
					render={({ field }) => (
						<FormItem>
							<FormLabel>報告内容</FormLabel>
							<FormControl>
								<Textarea
									placeholder="マーケットの取得時に画像が表示されず、詳細画面に遷移できない"
									{...field}
								/>
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
			</form>
		</Form>
	);
}

export function ReportBugSubmitButton({
	children,
	...props
}: ButtonProps &
	React.RefAttributes<HTMLButtonElement> & { children?: React.ReactNode }) {
	return (
		<Button type="submit" form={reportBugFormID} {...props}>
			{children}
		</Button>
	);
}

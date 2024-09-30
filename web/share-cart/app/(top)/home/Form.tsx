"use client";
import { createMarket } from "@/app/actions/createMarket";
import { Button, type ButtonProps } from "@/components/ui/button";
import {
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useForm } from "react-hook-form";
import { z } from "zod";

const FormSchema = z.object({
	name: z.string().min(2, {
		message: "Username must be at least 2 characters.",
	}),
	description: z.string().max(100, {
		message: "Description must be at most 100 characters.",
	}),
	apiMessage: z.string(), // hidden field: API response message
});

const createMarketFormID = "create-market";

export function CreateMarketForm({
	afterSubmit,
}: { afterSubmit?: () => void } = {}) {
	const router = useRouter();
	const form = useForm<z.infer<typeof FormSchema>>({
		resolver: zodResolver(FormSchema),
		defaultValues: {
			name: "",
			description: "",
			apiMessage: "",
		},
	});

	function onSubmit(data: z.infer<typeof FormSchema>) {
		createMarket(data).then((res) => {
			if (res.error) {
				form.setValue("apiMessage", "");
				form.setError("apiMessage", {
					type: "manual",
					message: res.error,
				});
				return;
			}

			if (afterSubmit) {
				afterSubmit();
			}
		});
	}

	return (
		<Form {...form}>
			<form
				onChange={() => form.setValue("apiMessage", "")}
				onSubmit={form.handleSubmit(onSubmit)}
				className="w-full space-y-6"
				id={createMarketFormID}
			>
				<FormField
					control={form.control}
					name="name"
					render={({ field }) => (
						<FormItem>
							<FormLabel>マーケット名</FormLabel>
							<FormControl>
								<Input placeholder="Market Name" {...field} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="description"
					render={({ field }) => (
						<FormItem>
							<FormLabel>マーケットの説明</FormLabel>
							<FormControl>
								<Input placeholder="Description" {...field} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<FormField
					control={form.control}
					name="apiMessage"
					render={({ field }) => (
						<FormItem>
							<FormControl>
								<Input type="hidden" {...field} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>
				<h1 className="text-green-600">{form.getValues("apiMessage")}</h1>
			</form>
		</Form>
	);
}

export function CreateMarketFormSubmitButton({
	children,
	...props
}: ButtonProps &
	React.RefAttributes<HTMLButtonElement> & { children?: React.ReactNode }) {
	return (
		<Button type="submit" form={createMarketFormID} {...props}>
			{children}
		</Button>
	);
}

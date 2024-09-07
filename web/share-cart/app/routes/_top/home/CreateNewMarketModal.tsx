import { Dialog, DialogBackdrop, DialogPanel } from "@headlessui/react";
import { Form } from "@remix-run/react";

type Props = {
	open: boolean;
	handleClose: () => void;
};

// https://tailwindui.com/components/application-ui/overlays/modal-dialogs#component-b6812b6c13fff16861f2645c4100ae5b
export function CreateNewMarketDialog(props: Props) {
	const { open, handleClose } = props;
	return (
		<Dialog open={open} onClose={handleClose} className="relative z-10">
			<DialogBackdrop
				transition
				className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity data-[closed]:opacity-0 data-[enter]:duration-300 data-[leave]:duration-200 data-[enter]:ease-out data-[leave]:ease-in"
			/>
			<div className="fixed inset-0 z-10 w-screen overflow-y-auto">
				<div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
					<DialogPanel
						transition
						className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all data-[closed]:translate-y-4 data-[closed]:opacity-0 data-[enter]:duration-300 data-[leave]:duration-200 data-[enter]:ease-out data-[leave]:ease-in sm:my-8 sm:w-full sm:max-w-lg data-[closed]:sm:translate-y-0 data-[closed]:sm:scale-95 w-full"
					>
						<CreateNewMarketForm handleClose={handleClose} />
					</DialogPanel>
				</div>
			</div>
		</Dialog>
	);
}

export const marketName = "marketName";
export const marketDescription = "marketDescription";

// https://tailwindui.com/components/application-ui/forms/sign-in-forms#component-766a0bf1b8800d383b6c5b77ef9c626c
function CreateNewMarketForm(props: { handleClose: () => void }) {
	return (
		<div className="flex min-h-full flex-1 flex-col justify-center px-6 py-12 lg:px-8">
			<div className="sm:mx-auto sm:w-full sm:max-w-sm">
				<svg
					xmlns="http://www.w3.org/2000/svg"
					fill="none"
					viewBox="0 0 24 24"
					strokeWidth={1.5}
					stroke="currentColor"
					className="size-20 w-full"
				>
					<path
						strokeLinecap="round"
						strokeLinejoin="round"
						d="M13.5 21v-7.5a.75.75 0 0 1 .75-.75h3a.75.75 0 0 1 .75.75V21m-4.5 0H2.36m11.14 0H18m0 0h3.64m-1.39 0V9.349M3.75 21V9.349m0 0a3.001 3.001 0 0 0 3.75-.615A2.993 2.993 0 0 0 9.75 9.75c.896 0 1.7-.393 2.25-1.016a2.993 2.993 0 0 0 2.25 1.016c.896 0 1.7-.393 2.25-1.015a3.001 3.001 0 0 0 3.75.614m-16.5 0a3.004 3.004 0 0 1-.621-4.72l1.189-1.19A1.5 1.5 0 0 1 5.378 3h13.243a1.5 1.5 0 0 1 1.06.44l1.19 1.189a3 3 0 0 1-.621 4.72M6.75 18h3.75a.75.75 0 0 0 .75-.75V13.5a.75.75 0 0 0-.75-.75H6.75a.75.75 0 0 0-.75.75v3.75c0 .414.336.75.75.75Z"
					/>
				</svg>
				<h2 className="mt-4 text-2xl font-bold leading-9 tracking-tight text-gray-900 text-center">
					マーケットをつくる
				</h2>
			</div>
			<div className="mt-10 sm:mx-auto sm:w-full sm:max-w-sm">
				<Form method="post" onSubmit={props.handleClose}>
					<div>
						<label
							htmlFor={marketName}
							className="block text-sm font-medium leading-6 text-gray-900"
						>
							マーケットの名前
						</label>
						<div className="mt-2">
							<input
								id={marketName}
								name={marketName}
								type="text"
								required
								className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
							/>
						</div>
					</div>

					<div>
						<div className="flex items-center justify-between">
							<label
								htmlFor={marketDescription}
								className="block text-sm font-medium leading-6 text-gray-900"
							>
								マーケットの説明
							</label>
						</div>
						<div className="mt-2">
							<input
								id={marketDescription}
								name={marketDescription}
								type="text"
								className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
							/>
						</div>
					</div>

					<div>
						<button
							type="submit"
							className="flex w-full justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600 mt-2"
						>
							作成
						</button>
					</div>
				</Form>
			</div>
		</div>
	);
}

import Footer from "./Footer";

export default function TopPageLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<section className="flex flex-col h-dvh">
			<div className="flex-auto overflow-auto p-2 pb-0">{children}</div>
			<div className="sticky bottom-0">
				<Footer />
			</div>
		</section>
	);
}

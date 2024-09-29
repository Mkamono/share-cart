import Footer from "./Footer";
import Header from "./Header";

export default function TopPageLayout({
	children,
}: {
	children: React.ReactNode;
}) {
	return (
		<section className="flex flex-col h-dvh">
			<div className="sticky top-0">
				<Header />
			</div>
			<div className="flex-1 p-2 pb-0 overflow-auto">{children}</div>
			<div className="sticky bottom-0">
				<Footer />
			</div>
		</section>
	);
}

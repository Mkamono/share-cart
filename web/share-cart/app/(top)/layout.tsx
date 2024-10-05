import { Footer } from "./Footer";
import { Header } from "./Header";

export default function TopPageLayout({
	children,
}: { children: React.ReactNode }) {
	return (
		<section className="flex flex-col h-dvh">
			<header>
				<Header />
			</header>
			<div className="flex-1 px-2 overflow-auto">{children}</div>
			<footer>
				<Footer />
			</footer>
		</section>
	);
}

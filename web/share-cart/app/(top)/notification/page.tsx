export default async function Home() {
	// 三秒待つ
	await new Promise((resolve) => setTimeout(resolve, 3000));

	return (
		<div>
			<h1>Notification Share Cart</h1>
		</div>
	);
}

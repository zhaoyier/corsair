import * as React from "react";

const Index: React.FC = () => {
	React.useEffect(() => {
		window.location.href = "/order.html";
	}, []);
	return <section>这里什么都没有</section>;
};

export default Index;

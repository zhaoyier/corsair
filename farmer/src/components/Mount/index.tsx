import * as React from "react";
import * as ReactDOM from "react-dom";
import Layout from "components/AdminLayout";

// import "iconfont/iconfont.css";
// import "styles/common.scss";
import "antd/dist/antd.min.css";

function render(App: any) {
	const target: HTMLElement = document.getElementById("container");
	ReactDOM.unmountComponentAtNode(target);
	ReactDOM.render(App, target);
}

function Mount(AppElement: any, hideMenu?: boolean, hidePadding?: boolean) {
	const App = (
		<Layout hideMenu={hideMenu} hidePadding={hidePadding}>
			<AppElement />
		</Layout>
	);

	render(App);
}

export default Mount;

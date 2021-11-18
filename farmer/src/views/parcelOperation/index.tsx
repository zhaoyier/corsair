import * as React from "react";
import { Switch, Route, HashRouter } from "react-router-dom";
import zhCN from "antd/es/locale/zh_CN";
import { ConfigProvider } from "antd";

import Loadable from "react-loadable";
import { MyLoadingComponent } from "./Components/Common/Loading";

const ParcelList = Loadable({
	loader: () => import("./Components/ParcelList"),
	loading: MyLoadingComponent
});

const ParcelDetail = Loadable({
	loader: () => import("./Components/ParcelDetail"),
	loading: MyLoadingComponent
});

class ParcelOperation extends React.Component {
	render() {
		return (
			<ConfigProvider locale={zhCN}>
				<HashRouter {...History}>
					<Switch>
						<Route path="/" exact component={ParcelList} />
						<Route path="/detail" component={ParcelDetail} />
					</Switch>
				</HashRouter>
			</ConfigProvider>
		);
	}
}

export default ParcelOperation;

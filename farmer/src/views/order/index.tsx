import * as React from "react";
import { HashRouter, Switch, Route } from "react-router-dom";
import { StoreProvider } from "./store/helper/storeProvider";
import store from "./store";
import Loadable from "react-loadable";
import { Loading } from "./Loading";
const AsyncList = Loadable({ loader: () => import("./list"), loading: Loading });
const AsyncDetail = Loadable({ loader: () => import("./detail"), loading: Loading });
export default function() {
	return (
		<StoreProvider value={store}>
			<HashRouter {...History}>
				<Switch>
					<Route path="/detail" component={AsyncDetail} />
					<Route path="/" exact component={AsyncList} />
				</Switch>
			</HashRouter>
		</StoreProvider>
	);
}

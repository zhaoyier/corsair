import * as React from "react";
import { observer } from "mobx-react-lite";
import Loadable from "react-loadable";
import { Divider } from "antd";
import { MyLoadingComponent } from "../../Common/Loading";

const BaseInfo = Loadable({
	loader: () => import("./BaseInfo"),
	loading: MyLoadingComponent
});

const ShippingRecord = Loadable({
	loader: () => import("./ShippingRecord"),
	loading: MyLoadingComponent
});

const OrderInfo = Loadable({
	loader: () => import("./OrderInfo"),
	loading: MyLoadingComponent
});

const GoodReceiveInfo = Loadable({
	loader: () => import("./GoodReceiveInfo"),
	loading: MyLoadingComponent
});

const CostInfo = Loadable({
	loader: () => import("./CostInfo"),
	loading: MyLoadingComponent
});

const DivMargin = () => (
	<Divider style={{ background: "#dfdfdf", marginBottom: 20, marginTop: 20  }} />
);

const BaseInfos = () => {
	return (
		<div>
			<BaseInfo />
			<DivMargin />
			<GoodReceiveInfo />
			<DivMargin />
			<ShippingRecord />
			<DivMargin />
			<OrderInfo />
			<DivMargin />
			<CostInfo />
		</div>
	);
};

export default observer(BaseInfos);

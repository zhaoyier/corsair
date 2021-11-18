import * as React from "react";
import { Tabs } from "antd";
const { TabPane } = Tabs;
import { BaseInfo } from "./baseInfo";
import { PayInfo } from "./payinfo";
import { OrderLog } from "./orderLog";
export function OrderInfo() {
	function callback(key) {
		console.log(key);
	}
	return (
		<div style={{ marginTop: 20 }}>
			<Tabs centered tabBarGutter={250} defaultActiveKey="1" onChange={callback} size="large">
				<TabPane tab="基础信息" key="basInfo">
					<BaseInfo />
				</TabPane>
				<TabPane tab="支付信息" key="payInfo">
					<PayInfo />
				</TabPane>
				<TabPane tab="订单日志" key="logs">
					<OrderLog />
				</TabPane>
			</Tabs>
		</div>
	);
}

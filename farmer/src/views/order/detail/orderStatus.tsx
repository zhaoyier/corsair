import * as React from "react";
import { Row, Descriptions } from "antd";
import { useObserver } from "mobx-react-lite";
import { store } from "../store/helper/useStore";
const OrderStatusStyle = {
	height: 27,
	lineHeight: "27px",
	marginTop: "20px",
	backgroundColor: "#fff",
	width: 700
};
export function OrderStatusShow() {
	const orderStore = store().orderStore;
	const commonStore = store().commonStore;
	const status =
		commonStore.orderStatusList &&
		orderStore.order &&
		orderStore.order.orderStatus &&
		commonStore.orderStatusList.orderStatusList.find(
			i => i.statusCode === orderStore.order.orderStatus
		).description;
	return useObserver(() => (
		<Row style={{ ...OrderStatusStyle }}>
			<Descriptions>
				<Descriptions.Item label="订单号">
					{(orderStore.order && orderStore.order.orderNo) || ""}
				</Descriptions.Item>
				<Descriptions.Item label="订单状态">{status}</Descriptions.Item>
			</Descriptions>
			{/* <Col span={10} offset={1}>
				<strong>{`订单号  ${(orderStore.orderStatus && orderStore.order.orderNo) ||
					""}`}</strong>
			</Col>
			<Col span={10}>
				<strong>{`订单状态  ${(orderStore.orderStatus &&
					orderStore.orderStatus.description) ||
					""}`}</strong>
			</Col> */}
		</Row>
	));
}

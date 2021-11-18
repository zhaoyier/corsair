import * as React from "react";
import { Row, Col, Table } from "antd";
import { store } from "../../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
const Colitem = {
	span: 8
};

export function DrliveryInfo() {
	const orderStore = store().orderStore;
	const commonStore = store().commonStore;
	const columns = [
		{
			title: "包裹号",
			dataIndex: "packageCode",
			render: id => (
				<a
					onClick={() =>
						window.open(
							`${location.origin}/parcelOperation.html#/detail?packageCode=${id}`
						)
					}>
					{id}
				</a>
			)
		},
		{
			title: "发货ETA",
			dataIndex: "eta"
		},
		{
			title: "包裹状态",
			dataIndex: "packageStatus",
			render: status =>
				commonStore.parcelStatusList &&
				commonStore.parcelStatusList.parcelStatusList.find(i => i.statusCode === status) &&
				commonStore.parcelStatusList.parcelStatusList.find(i => i.statusCode === status)
					.description
		},
		{
			title: "预约取货时间",
			dataIndex: "pickingTime"
		},
		{
			title: "签收时间",
			dataIndex: "signingTime"
		}
	];
	return useObserver(() => (
		<div style={{ lineHeight: 1.9 }}>
			<Row>
				<Col {...Colitem}>
					<span>取货方式:</span> {orderStore.delivery && orderStore.delivery.deliveryMethod}
				</Col>
				<Col {...Colitem}><span>收件人:</span> {orderStore.delivery && orderStore.delivery.name}</Col>
				<Col {...Colitem}><span>收件电话:</span> {orderStore.delivery && orderStore.delivery.phone}</Col>
				<Col {...Colitem}>
					<span>配送单号:</span>{" "}
					{orderStore.delivery &&
						!!+orderStore.delivery.shipmentId &&
						orderStore.delivery.shipmentId}
				</Col>
				<Col {...Colitem}>
					<span>详细地址: </span>{orderStore.delivery && orderStore.delivery.addressDetail}
				</Col>
				<Col {...Colitem}><span>收件城市: </span>{orderStore.delivery && orderStore.delivery.city}</Col>
				<Col {...Colitem}><span>收件省份（州）: </span>{orderStore.delivery && orderStore.delivery.state}</Col>
				<Col {...Colitem}><span>zipcode: </span>{orderStore.delivery && orderStore.delivery.zipcode}</Col>
			</Row>
			<Row style={{ paddingTop: "10px" }}>
				<Col span={24}>
					<Table
						size="small"
						columns={columns}
						dataSource={orderStore.delivery && orderStore.delivery.packages}
						rowKey={record => record.packageCode}
						pagination={false}
					/>
				</Col>
			</Row>
		</div>
	));
}

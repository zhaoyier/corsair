import * as React from "react";
import { Col, Row } from "antd";
import { store } from "../../store/helper/useStore";
import { formatUnixTime } from "utils/time";
const ColItem = {
	span: 6
};
export function PayDetail() {
	const orderStore = store().orderStore;
	const orderPayInfo = orderStore.orderPayInfo;
	return (
		<Row justify="space-between" style={{ lineHeight: 3 }}>
			<Col {...ColItem}>支付方式: {orderPayInfo && orderPayInfo.payWay}</Col>
			<Col {...ColItem}>支付币种: {orderPayInfo && orderPayInfo.currencyCode}</Col>
			<Col {...ColItem}>
				支付时间:{" "}
				{orderPayInfo &&
					!!Number(orderPayInfo.payDate) &&
					formatUnixTime(orderPayInfo.payDate)}
			</Col>
			{/* <Col {...ColItem}>
				支付确认时间:{" "}
				{orderPayInfo &&
					!!Number(orderPayInfo.confirmPayDate) &&
					formatUnixTime(orderPayInfo.confirmPayDate)}
			</Col> */}
			<Col {...ColItem}>支付号: {orderPayInfo && orderPayInfo.billNo}</Col>
		</Row>
	);
}

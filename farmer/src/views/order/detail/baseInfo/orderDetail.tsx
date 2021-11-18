import * as React from "react";
import { Col, Row } from "antd";
import { CatAddServerModal } from "./catAddServerModal";
const ColItem = {
	span: 6
};
import { store } from "../../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { formatUnixTime } from "utils/time";
import { formatWeightUnit } from "utils/util";

export function OrderDetail() {
	const orderStore = store().orderStore;
	const [showAddServerModal, setShowAddServerModal] = React.useState<boolean>(false);
	return useObserver(() => (
		<Row justify="space-between" style={{ lineHeight: 1.9 }}>
			<Col {...ColItem}><span>会员名称: </span>{orderStore.order && orderStore.order.nickName}</Col>
			<Col {...ColItem}>
			<span>订单创建时间:</span> {formatUnixTime(orderStore.order && orderStore.order.orderDate)}
			</Col>
			<Col {...ColItem}>
			<span>转运快递公司: </span>{orderStore.order && orderStore.order.logisticsName}
			</Col>
			<Col {...ColItem}><span>转运快递单号: </span>{orderStore.order && orderStore.order.logisticsNo}</Col>
			<Col {...ColItem}><span>发货仓库: </span>{orderStore.order && orderStore.order.warehouseName}</Col>
			<Col {...ColItem}>
			<span>已存储天数: </span>{orderStore.order && orderStore.order.storeDays}&nbsp;天
			</Col>
			<Col {...ColItem}>
			<span>增值服务:</span>
				{orderStore.order && orderStore.order.addedService && (
					<a onClick={() => setShowAddServerModal(true)}> 查看</a>
				)}
			</Col>
			<Col {...ColItem}>
			<span>实际运输方式:</span> {orderStore.order && orderStore.order.shipmentType}
			</Col>
			<Col {...ColItem}>
			<span>运输方式备注:</span> {orderStore.order && orderStore.order.shipmentRemark}
			</Col>
			<Col {...ColItem}>
			<span>实际重量:</span> {formatWeightUnit(orderStore.order && orderStore.order.actualWeight)}
			</Col>
			<Col {...ColItem}>
			<span>体积重量:</span> {formatWeightUnit(orderStore.order && orderStore.order.volume)}
			</Col>
			<Col {...ColItem}>
			<span>计费重量: </span>{formatWeightUnit(orderStore.order && orderStore.order.chargeableWeight)}
			</Col>
			<Col {...ColItem}><span>长</span>: {orderStore.order && `${orderStore.order.length}(cm)`}</Col>
			<Col {...ColItem}><span>宽</span>: {orderStore.order && `${orderStore.order.width}(cm)`}</Col>
			<Col {...ColItem}><span>高</span>: {orderStore.order && `${orderStore.order.height}(cm)`}</Col>
			<Col {...ColItem}>
			<span>可选运输方式</span>:{orderStore.order && orderStore.order.availableShipmentType}
			</Col>
			{showAddServerModal && (
				<CatAddServerModal handleClose={() => setShowAddServerModal(false)} />
			)}
		</Row>
	));
}

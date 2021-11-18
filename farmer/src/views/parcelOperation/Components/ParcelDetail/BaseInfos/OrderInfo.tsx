import * as React from "react";
import { observer } from "mobx-react-lite";
import { Table } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { formatPrice } from "utils/util";

const styles = require("./index.scss");

const getColumns = () => {
	const columns = [
		{
			title: "订单号",
			dataIndex: "orderNumber",
			render: (t, r) => (
				<a rel="opener" href={`order.html#/detail?orderId=${r.orderId}`} target="_blank">
					{t}
				</a>
			)
		},
		{ title: "报关品类", dataIndex: "itemType" },
		{ title: "数量", dataIndex: "qty" },
		{ title: "申报价值", dataIndex: "declaredValue", render: t => formatPrice(t) }
	];
	return columns;
};

const OrderInfo = () => {
	const { parcelDetailStore } = useStores();
	const { record } = parcelDetailStore;
	const { items } = record;

	return (
		<section>
			<div className={styles.header}>订单信息</div>
			<Table
				rowKey="orderId"
				columns={getColumns()}
				dataSource={items || []}
				pagination={false}
			/>
		</section>
	);
};

export default observer(OrderInfo);

import * as React from "react";
import { Table } from "antd";
import { store } from "../../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { formatPrice } from "utils/util";

const columns = [
	{
		title: "订单Item号",
		dataIndex: "itemId"
	},
	{
		title: "一级报关品类",
		dataIndex: "fstCategoryName"
	},
	{
		title: "二级报关品类",
		dataIndex: "secCategoryName"
	},
	{
		title: "数量",
		dataIndex: "qty"
	},
	{
		title: "申报单价",
		dataIndex: "declaredAmount",
		render: text => formatPrice(text)
	}
];
export function ProductInfo() {
	const orderStore = store().orderStore;
	return useObserver(() => (
		<Table
			size="small"
			columns={columns}
			dataSource={orderStore.order && orderStore.order.item}
			rowKey={record => record.itemId}
			pagination={false}
		/>
	));
}

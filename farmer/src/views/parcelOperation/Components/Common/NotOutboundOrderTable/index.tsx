import * as React from "react";
import { Table } from "antd";

interface ModalNotOutboundOrderTableProps {
	dataSource: any[];
}
const styles = require("./index.scss");

export default class NotOutboundOrderTable extends React.Component<
	ModalNotOutboundOrderTableProps,
	{}
> {
	getColumns = () => {
		const column = [
			{
				title: "订单号",
				dataIndex: "orderNum",
				key: "orderNum",
				render: (t, r) => (
					<a
						rel="opener"
						href={`order.html#/detail?orderId=${r.orderId}`}
						target="_blank">
						{t}
					</a>
				)
			},
			{ title: "订单状态", dataIndex: "orderStatus", key: "orderStatus" },
			{ title: "商品数量", dataIndex: "quantity", key: "quantity" },
			{ title: "发货仓库", dataIndex: "warehouse", key: "warehouse" }
		];
		return column;
	};

	render() {
		const { dataSource } = this.props;
		return (
			<section>
				<h4 className={styles.header}>未出库订单列表</h4>
				<Table
					rowKey="orderNum"
					// scroll={{ y: 250 }}
					columns={this.getColumns()}
					pagination={false}
					dataSource={dataSource}
				/>
			</section>
		);
	}
}

import * as React from "react";
import { Col, Row, Table } from "antd";
import { store } from "../../store/helper/useStore";
import { formatUnixTime } from "utils/time";

export function OrderLog() {
	const orderStore = store().orderStore;
	const logInfo = orderStore.logInfo;
	const orderLogColums = [
		{
			title: "创建时间",
			dataIndex: "logDate",
			render: text => formatUnixTime(text)
		},
		{
			title: "内容说明",
			dataIndex: "logContent"
		},
		{
			title: "作者",
			dataIndex: "operator"
		}
	];
	return (
		<Row style={{ marginTop: 30 }}>
			<Col span={24}>
				<Table
					columns={orderLogColums}
					dataSource={logInfo}
					rowKey={record => record.logDate}
					pagination={false}
				/>
			</Col>
		</Row>
	);
}

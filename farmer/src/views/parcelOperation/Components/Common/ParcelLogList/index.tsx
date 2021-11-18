import * as React from "react";
import { Table } from "antd";
import { PackageLogInfo } from "genServices/ezShipOMS/oms";
import { formatUnixTime, TimeFormatTotalReverse } from "utils/time";

interface ParcelLogListProps {
	dataSource: PackageLogInfo[];
}

const getColumns = () => {
	const columns = [
		{
			title: "时间",
			dataIndex: "createDate",
			key: "createDate",
			render: t => formatUnixTime(t, TimeFormatTotalReverse)
		},
		{
			title: "内容",
			dataIndex: "context",
			key: "context",
			// width: 400,
			render: t => (
				<div
					style={{
						width: "92%",
						whiteSpace: "pre-wrap",
						wordWrap: "break-word",
						wordBreak: "break-all"
					}}>
					{t}
				</div>
			)
		},
		{ title: "操作人", dataIndex: "opName", key: "opName" }
	];

	return columns;
};

const ParcelLogList = (props: ParcelLogListProps) => {
	const { dataSource } = props;
	return (
		<Table
			rowKey="packageCode"
			scroll={{ y: 200 }}
			// pagination={false}
			style={{ marginLeft: 0 }}
			columns={getColumns()}
			dataSource={dataSource}
		/>
	);
};

export default ParcelLogList;

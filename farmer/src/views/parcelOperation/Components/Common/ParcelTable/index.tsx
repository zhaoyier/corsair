import * as React from "react";
import { Table } from "antd";
import { CheckboxProps } from "antd/lib/checkbox";

interface ModalParcelTableProps {
	dataSource: any[];
	hiddeHeader?: boolean;
	showHeaderCheckBox?: boolean;
	showCheckBox?: boolean;
	selectedRowKeys?: string[];
	getCheckboxProps?: (record) => Partial<CheckboxProps>;
	setSelectRowKeys?: (selectedRowKeys) => void;
}

const getColumns = () => {
	const column = [
		{
			title: "包裹号",
			dataIndex: "packageNum",
			key: "packageNum",
			width: 200,
			render: t => (
				<a
					rel="opener"
					href={`/parcelOperation.html#/detail?packageCode=${t}`}
					target="_blank">
					{t}
				</a>
			)
		},
		{
			title: "发货ETA",
			dataIndex: "deliverEta",
			key: "deliverEta"
		},
		{
			title: "运输方式",
			dataIndex: "transport",
			key: "transport"
		},
		{
			title: "包裹状态",
			dataIndex: "packageStatusName",
			key: "packageStatusName"
		},
		{
			title: "发货仓库",
			dataIndex: "warehouse",
			key: "warehouse"
		}
	];
	return column;
};

const ParcelTable = (props: ModalParcelTableProps) => {
	const {
		dataSource,
		hiddeHeader,
		showHeaderCheckBox,
		showCheckBox,
		selectedRowKeys,
		setSelectRowKeys,
		getCheckboxProps
	} = props;

	const getRowSelection = () => {
		const rowSelection = {
			selectedRowKeys,
			onChange: (selectedRowKeys, selectedRows) => {
				setSelectRowKeys(selectedRowKeys);
				console.log(`selectedRowKeys: ${selectedRowKeys}`, "selectedRows: ", selectedRows);
			},
			columnTitle: showHeaderCheckBox ? "" : " ",
			getCheckboxProps: record => getCheckboxProps(record),
			hideDefaultSelections: true
			// getCheckboxProps: record => ({
			// 	disabled: record.packageStatus !== 20700 && record.packageStatus !== 14600
			// })
		};
		return showCheckBox ? rowSelection : null;
	};

	return (
		<section>
			{!hiddeHeader && <h4>包裹列表</h4>}
			<Table
				rowKey="packageNum"
				// scroll={{ y: 250 }}
				columns={getColumns()}
				pagination={false}
				rowSelection={getRowSelection()}
				dataSource={dataSource}
			/>
		</section>
	);
};

export default ParcelTable;

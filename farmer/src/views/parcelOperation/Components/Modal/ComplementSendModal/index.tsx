import * as React from "react";
import { Modal, Table, Tooltip } from "antd";
import { PackageDetailResp } from "genServices/ezShipOMS/oms";

export interface ComplementSendModalProps {
	visible: boolean;
	record: PackageDetailResp; // 单条包裹信息
	refresh?: () => void;
	onCancel?: () => void;
	onOk?: (record, cb?: (result, refresh?: () => void) => void) => void;
}
const styles = require("./index.scss");
// 补送  弹窗

const getColumns = (record: PackageDetailResp) => {
	const column = [
		{
			title: "包裹号",
			dataIndex: "packageCode",
			key: "packageCode",
			width: 200,
			render: () => (
				<a
					rel="opener"
					href={`parcelOperation.html#/detail?packageCode=${record.packageCode}`}
					target="_blank">
					{record.packageCode}
				</a>
			)
		},
		{
			title: "订单号",
			dataIndex: "orderNumber",
			width: 200,
			render: (t, r) => (
				<a rel="opener" href={`order.html#/detail?orderId=${r.orderId}`} target="_blank">
					{t}
				</a>
			)
		},
		{
			title: "发货ETA",
			dataIndex: "deliverEta",
			key: "deliverEta",
			render: () => <span>{record.shipmentInfo.dispatchEta}</span>
		},
		{
			title: "包裹状态",
			dataIndex: "packageStatusName",
			key: "packageStatusName",
			render: () => <span>{record.packageStatus}</span>
		},
		{
			title: "发货仓库",
			dataIndex: "warehouse",
			key: "warehouse",
			render: () => <span>{record.warehouse}</span>
		},
		{
			title: "配送单号",
			dataIndex: "warehouse",
			key: "warehouse",
			render: () => <span>{record.shipmentId}</span>
		}
	];
	return column;
};

const ComplementSendModal = (props: ComplementSendModalProps) => {
	const { visible, record, onCancel, onOk, refresh } = props;
	const title = "补送(需要指定当前包裹所对应的订单用于生成订单信息)";
	const [selectedRowKeys, setSelectRowKeys] = React.useState([]);
	const [selectedRows, setSelectRows] = React.useState([]);

	const successRes = result => {
		if (result) {
			Modal.confirm({
				title: "操作成功",
				content: (
					<section className={styles.resultModal}>
						<p>
							生成了一个补送订单：订单号为：
							<span>
								<a
									onClick={() =>
										window.open(
											`order.html#/detail?orderId=${result.newOrderId}`,
											"_blank"
										)
									}>
									{result.newOrderNumber}
								</a>
							</span>
						</p>
						<p>
							生成了一笔新的配送单，配送单号为：<span>{result.newShipmentId}</span>
						</p>
						<p>
							补送包裹号：
							<span>
								<a
									rel="opener"
									href={`/parcelOperation.html#/detail?packageCode=${result.newPackageCode}`}
									target="_blank">
									{result.newPackageCode}
								</a>
							</span>
						</p>
					</section>
				),
				cancelText: "关闭",
				onCancel: () => {
					onCancel();
					refresh();
				},
				okText: "查看",
				onOk: () => {
					onCancel();
					refresh();
					window.open(
						`/parcelOperation.html#/detail?packageCode=${result.newPackageCode}`,
						"_blank"
					);
				}
			});
		}
	};

	const getRowSelection = () => {
		const rowSelection = {
			selectedRowKeys,
			onChange: (selectedRowKeys, selectedRows) => {
				console.log(selectedRowKeys, selectedRows);
				setSelectRowKeys(selectedRowKeys);
				setSelectRows(selectedRows);
				console.log(`selectedRowKeys: ${selectedRowKeys}`, "selectedRows: ", selectedRows);
			},
			columnTitle: <Tooltip title="选择需补送的项">NO</Tooltip>,
			hideDefaultSelections: true
			// getCheckboxProps: record => ({
			// 	disabled: record.packageStatus !== 20700 && record.packageStatus !== 14600
			// })
		};
		return {
			type: "radio",
			...rowSelection
		} as any;
	};

	return (
		<Modal
			visible={visible}
			title={title}
			onCancel={onCancel}
			maskClosable={false}
			width={1000}
			okButtonProps={{
				disabled: selectedRowKeys.length === 0 // 未选择包裹时不允许  确定补送
			}}
			onOk={() =>
				onOk(
					{ packageCode: record.packageCode, orderNumber: selectedRows[0].orderNumber },
					result => successRes(result)
				)
			}
			afterClose={() => setSelectRowKeys([])}>
			<Table
				rowKey="orderId"
				scroll={{ y: 150 }}
				columns={getColumns(record)}
				pagination={false}
				rowSelection={getRowSelection()}
				dataSource={record.items || []}
			/>
		</Modal>
	);
};

export default ComplementSendModal;

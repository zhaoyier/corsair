import * as React from "react";
import { Table, Button, Spin, Popover } from "antd";
import { store } from "../store/helper/useStore";
import { Link } from "react-router-dom";
import { observer } from "mobx-react-lite";
const styles = require("./index.scss");
import { sampleOrderInfo } from "genServices/ezShipOMS/oms";
import { warehouseListArray } from "views/order/constant";
import { toJS } from "mobx";
import { formatUnixTime } from "utils/time";
import { formatPrice } from "utils/util";
import { AddEventModal } from "../detail/operationModal/additionEvent";
import { ChargeBackModal } from "../detail/operationModal/chargeBack";
import { CancelModal } from "../detail/operationModal/cancel";
import { RemarkModal } from "../detail/operationModal/remark";
import { DownOutlined } from "@ant-design/icons";
const OrderTable = () => {
	// const orderListUi = store().orderListUi;
	const orderListStore = store().orderListStore;
	const commonStore = store().commonStore;
	const orderStore = store().orderStore;

	const {
		spinning,
		actionTypeisMulti,
		selectedRowKeys,
		selectedRows,
		showAddModal,
		showBackModal,
		showCancelModal,
		showRemarkModal,
		changeListStore
	} = orderListStore;
	const [currentOrder, setCurrentOrder] = React.useState<sampleOrderInfo>(null);

	function renderRemarkColums(hasRemark, record: sampleOrderInfo) {
		return (
			<div>
				<div>
					<Button
						type="link"
						disabled={+record.orderStatus >= 8150 || +record.orderStatus <= 1}
						onClick={() => {
							setCurrentOrder(record);
							changeListStore("showAddModal", true);
							changeListStore("actionTypeisMulti", false);
							orderStore.initAddedServiceTerm({
								warehouse: record.warehouseId as any
							});
						}}>
						添加附加服务
					</Button>
					<Button
						type="link"
						onClick={() => {
							setCurrentOrder(record);
							changeListStore("actionTypeisMulti", false);
							changeListStore("showBackModal", true);
						}}
						disabled={+record.orderStatus < 6000 || +record.orderStatus >= 8160}>
						退单
					</Button>
				</div>
				<div>
					<Button
						type="link"
						disabled={+record.orderStatus >= 6000 || +record.orderStatus <= 1}
						onClick={() => {
							setCurrentOrder(record);
							changeListStore("actionTypeisMulti", false);
							changeListStore("showCancelModal", true);
						}}>
						取消
					</Button>
					<Button
						type="link"
						// disabled={+record.orderStatus >= 8150 || +record.orderStatus <= 1}
						onClick={() => {
							setCurrentOrder(record);
							changeListStore("actionTypeisMulti", false);
							changeListStore("showRemarkModal", true);
						}}>
						添加备注
					</Button>
					{hasRemark && (
						<Button
							type="link"
							onClick={() =>
								window.open(
									`${location.origin}/order.html#/detail?orderId=${record.orderId}&anchor=remark`,
									"_blank"
								)
							}>
							查看备注
						</Button>
					)}
				</div>
			</div>
		);
	}
	function renderWarehouse(Warehouse: string) {
		return warehouseListArray.find(({ value }) => value === Warehouse);
	}
	function renderStatus(status: string) {
		const orderStatusList = commonStore.orderStatusList;
		if (orderStatusList && orderStatusList.orderStatusList) {
			const result = orderStatusList.orderStatusList.find(({ statusCode }) => {
				if (statusCode === status) {
					return true;
				} else {
					return null;
				}
			});
			return result && result.description;
		} else {
			return null;
		}
	}

	const columns = [
		{
			title: "操作",
			width: 80,
			dataIndex: "hasRemark",
			align: "center",
			render: (t, r) =>
				<Popover placement="bottom" zIndex={90} trigger="click" content={renderRemarkColums(t, r)}>
					<a>操作&nbsp;<DownOutlined /></a>
				</Popover>
		},
		{
			title: "订单号",
			dataIndex: "orderNumber",
			render: (...rest) => (
				<React.Fragment>
					<Link
							// onClick={() => commonStore.clearList()}
							to={`/detail?orderId=${rest[1].orderId}`}
						>
							{rest[0]}
						</Link>
					{/* <a
						onClick={() => {
							// commonStore.clearList("Search");
							window.open(
								`${location.origin}${location.pathname}#/detail?orderId=${rest[1].orderId}`
							);
						}}>
						{rest[0]}
					</a> */}
					{rest[1].hasAdditionalServices && (
						<img
							src={require("../img/addServer.png")}
							className={styles.moreOrderImg}
						/>
					)}
					{rest[1].isAfterSales && (
						<img
							src={require("../img/superOrder.png")}
							className={styles.moreOrderImg}
						/>
					)}
					{rest[1].isDefect && (
						<img src={require("../img/zhuyi.png")} className={styles.moreOrderImg} />
					)}
				</React.Fragment>
			),
			width: 200
		},
		{
			title: "转运运单号",
			width: 150,
			dataIndex: "transportNumber",
			render: (text, r) => <a
				onClick={() => {
					// commonStore.clearList("Search");
					window.open(
						`${location.origin}${location.pathname}#/detail?orderId=${r.orderId}`
					);
				}}
			>{text}</a>
		},
		// {
		// 	title: "订单来源",
		// 	width: 100,
		// 	dataIndex: "orderOrigin",
		// 	render: text => <span>{text ? text.name : null}</span>
		// },
		{
			title: "订单状态",
			width: 100,
			dataIndex: "orderStatus",
			render: text => <span>{renderStatus(text)}</span>
		},
		{
			title: "仓库",
			width: 50,
			dataIndex: "warehouseId",
			render: text => (
				<span>{(renderWarehouse(text) && renderWarehouse(text).label) || ""}</span>
			)
		},
		{
			title: "国家",
			width: 50,
			dataIndex: "catalogCode"
		},
		{
			title: "运输方式",
			width: 80,
			dataIndex: "shipmentType"
		},
		{
			title: "计费重(KG)",
			width: 60,
			dataIndex: "chargWeight",
			render: text => <span>{(+text / 1000).toFixed(2)}</span>
		},
		{
			title: "称重重(KG)",
			width: 60,
			dataIndex: "actualWeight",
			render: text => <span>{(+text / 1000).toFixed(2)}</span>
		},
		{
			title: "体积重(KG)",
			width: 60,
			dataIndex: "volumeWeight",
			render: text => <span>{(+text / 1000).toFixed(2)}</span>
		},
		{
			title: "会员名称",
			width: 100,
			dataIndex: "nickName"
		},
		// {
		// 	title: "CustomerID",
		// 	width: 100,
		// 	dataIndex: "customerId"
		// },
		{
			title: "标识码",
			width: 80,
			dataIndex: "selfHelpCode"
		},
		{
			title: "申报价值",
			width: 80,
			dataIndex: "declaredAmount",
			render: text => <span>{formatPrice(text)}</span>
		},
		{
			title: "创建时间",
			width: 160,
			dataIndex: "createDate",
			render: text => <span>{formatUnixTime(text)}</span>
		}
	];
	const rowSelection = {
		selectedRowKeys: orderListStore.selectedRowKeys,
		onChange: (selectedRowKeys, selectedRows) => {
			orderListStore.changeListStore("selectedRowKeys", selectedRowKeys);
			orderListStore.changeListStore("selectedRows", selectedRows);
			console.log(`selectedRowKeys: ${selectedRowKeys}`, "selectedRows: ", selectedRows);
		}
	};
	return (
		<Spin spinning={spinning}>
			<Table
				columns={columns as any}
				rowSelection={rowSelection}
				dataSource={toJS(orderListStore.orderList)}
				rowKey={record => record.orderId}
				scroll={{ x: 1000, y: 530 }}
				pagination={false}
			/>
			{showAddModal && (
				<AddEventModal
					actionTypeisMulti={actionTypeisMulti}
					orderIdArr={actionTypeisMulti ? selectedRowKeys : [currentOrder.orderId]}
					handleClose={() => changeListStore("showAddModal", false)}
				/>
			)}
			{showBackModal && (
				<ChargeBackModal
					order={currentOrder}
					orderId={currentOrder.orderId}
					handleClose={() => {
						changeListStore("showBackModal", false);
						commonStore.clearList("ChargeBack");
					}}
				/>
			)}
			{showCancelModal && (
				<CancelModal
					actionTypeisMulti={actionTypeisMulti}
					orderArr={actionTypeisMulti ? selectedRows : [currentOrder]}
					orderIdArr={actionTypeisMulti ? selectedRowKeys : [currentOrder.orderId]}
					handleClose={() => changeListStore("showCancelModal", false)}
				/>
			)}
			{showRemarkModal && (
				<RemarkModal
					record={currentOrder}
					orderId={currentOrder && currentOrder.orderId}
					handleClose={() => changeListStore("showRemarkModal", false)}
				/>
			)}
		</Spin>
	);
};

export default observer(OrderTable);

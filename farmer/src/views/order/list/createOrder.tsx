import * as React from "react";
import { Button, message, Pagination } from "antd";
import OrderForm from "./orderForm";
import { SuccessModal } from "./successModel";
import { observer } from "mobx-react-lite";
import { store } from "../store/helper/useStore";
import { BatchOperationType, SearchOrderReq, SiteType } from "genServices/ezShipOMS/oms";
const CreateOrder = () => {
	const orderStore = store().orderStore;
	const [isShowCreateModal, setIsShowCreateModal] = React.useState<boolean>(false);
	const [isShowSuccessModal, setIsShowSuccessModal] = React.useState<boolean>(false);
	const [orderId, setOrderId] = React.useState<string>("");
	const [orderFormType, setOrderFormType] = React.useState<SiteType>(SiteType.SiteTypeEzbuy);

	const orderListUi = store().orderListUi;
	const orderListStore = store().orderListStore;
	const { selectedRowKeys, selectedRows, changeListStore, isBatchExecutable } = orderListStore;

	const onCreate = (value: string) => {
		setOrderId(value);
		setIsShowSuccessModal(true);
		setIsShowCreateModal(false);
	};

	// 批量取消
	const clickCancel = async () => {
		if (selectedRowKeys && selectedRowKeys.length > 0) {
			const resq = await isBatchExecutable({
				orderIds: selectedRowKeys,
				batchOperationType: BatchOperationType.BatchCancellOrder
			});
			if (!resq.result.result) {
				message.error(resq.result.msg, 3);
				// message.error("所勾选的列表中存在已被仓库签收的订单，不支持批量取消，请重新选择！");
			} else {
				changeListStore("actionTypeisMulti", true);
				changeListStore("showCancelModal", true);
			}
		} else {
			message.error("请勾选需要操作的订单！");
		}
	};

	// 批量添加附加服务
	const clickAddAction = async () => {
		if (selectedRowKeys && selectedRowKeys.length > 0) {
			const arr = [...new Set(selectedRows.map(elem => elem.warehouseId))];
			console.log(selectedRows, arr);
			if (arr && arr.length > 1) {
				message.error("勾选的订单仓库需要保持一致");
			} else {
				const resq = await isBatchExecutable({
					orderIds: selectedRowKeys,
					batchOperationType: BatchOperationType.BatchCreateAddedService
				});
				if (!resq.result.result) {
					message.error(resq.result.msg, 3);
					// message.error("所勾选的列表中存在已提交发货的订单，不支持批量添加附加服务，请重新选择！");
				} else {
					changeListStore("actionTypeisMulti", true);
					orderStore.initAddedServiceTerm({
						warehouse: selectedRows[0].warehouseId as any
					});
					changeListStore("showAddModal", true);
				}
			}
		} else {
			message.error("请勾选需要操作的订单！");
		}
	};

	const changePage = (page, pageSize) => {
		orderListUi.setOffset(page - 1);
		orderListUi.setPage(page);
		orderListStore.searchOrderList({
			...orderListStore.searchParams,
			offset: (page - 1) * pageSize + "",
			limit: pageSize + ""
		} as SearchOrderReq);
	};

	return (
		<div
			style={{
				padding: "0px 15px 15px 0",
				display: "flex",
				alignItems: "center",
				justifyContent: "space-between"
			}}>
			<section>
				<Button
					type="primary"
					disabled={selectedRowKeys.length > 0}
					onClick={() => {
						setIsShowCreateModal(true);
						setOrderFormType(SiteType.SiteTypeEzbuy);
					}}>
					创建新马订单
				</Button>
				<Button
					type="primary"
					style={{ marginLeft: 10 }}
					disabled={selectedRowKeys.length > 0}
					onClick={() => {
						setIsShowCreateModal(true);
						setOrderFormType(SiteType.SiteTypeGlobal);
					}}>
					创建国际化订单
				</Button>
				<Button
					type="primary"
					style={{ marginLeft: 10 }}
					disabled={selectedRowKeys.length === 0}
					onClick={clickCancel}>
					批量取消
				</Button>
				<Button
					type="primary"
					style={{ marginLeft: 10 }}
					disabled={selectedRowKeys.length === 0}
					onClick={clickAddAction}>
					批量添加附加服务
				</Button>
			</section>
			<section>
				<Pagination
					current={orderListUi.page}
					total={orderListStore.total}
					showQuickJumper={true}
					showSizeChanger={true}
					pageSize={orderListUi.limit}
					showTotal={total => `共${total}条`}
					onChange={(current, pageSize) => changePage(current, pageSize)}
					onShowSizeChange={(current, pageSize) => changePage(current, pageSize)}
				/>
			</section>
			{isShowCreateModal && (
				<OrderForm
					site={orderFormType}
					visible={isShowCreateModal}
					handleClose={() => setIsShowCreateModal(false)}
					onCreate={onCreate}
				/>
			)}
			{isShowSuccessModal && (
				<SuccessModal handleClose={() => setIsShowSuccessModal(false)} orderId={orderId} />
			)}
		</div>
	);
};

export default observer(CreateOrder);

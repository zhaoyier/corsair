import * as React from "react";
import { Button, Col, Row } from "antd";
import { AddEventModal } from "./operationModal/additionEvent";
import { ChargeBackModal } from "./operationModal/chargeBack";
import { CancelModal } from "./operationModal/cancel";
import { ChangeMoneyModal } from "./operationModal/changeMoney";
import { RemarkModal } from "./operationModal/remark";
import { store } from "../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { warehouseListArray } from "../constant";

const buttonStyle = {
	borderRadius: 6
};
// 这个函数应该传入一个状态 orderUi中的orderStatus用这个状态来控制按钮的置灰,当然还要写一个函数类控制
export function ActiveButton() {
	const orderStore = store().orderStore;
	const commonStore = store().commonStore;
	const [showAddModal, setShowAddModal] = React.useState<boolean>(false);
	const [showBackModal, setShowBackModal] = React.useState<boolean>(false);
	const [showCancelModal, setShowCancelModal] = React.useState<boolean>(false);
	const [showChangeMoneyModal, setShowChangeMoneyModal] = React.useState<boolean>(false);
	const [showRemarkModal, setShowRemarkModal] = React.useState<boolean>(false);

	return useObserver(() => (
		<Row>
			<Col span={3}>
				<Button
					type="primary"
					style={{ ...buttonStyle }}
					disabled={
						+(orderStore.order && orderStore.order.orderStatus) >= 8150 ||
						+(orderStore.order && orderStore.order.orderStatus) <= 1
					}
					onClick={() => {
						setShowAddModal(true);
						orderStore.initAddedServiceTerm({
							warehouse:
								warehouseListArray.find(
									i => i.label === orderStore.order.warehouseName
								) &&
								warehouseListArray.find(
									i => i.label === orderStore.order.warehouseName
								).value
						});
					}}>
					添加附加服务
				</Button>
			</Col>
			<Col span={3}>
				<Button
					type="primary"
					style={{ ...buttonStyle }}
					onClick={() => setShowBackModal(true)}
					disabled={
						+(orderStore.order && orderStore.order.orderStatus) < 6000 ||
						+(orderStore.order && orderStore.order.orderStatus) >= 8160
					}>
					退单
				</Button>
			</Col>
			<Col span={3}>
				<Button
					disabled={
						+(orderStore.order && orderStore.order.orderStatus) >= 6000 ||
						+(orderStore.order && orderStore.order.orderStatus) <= 1
					}
					type="primary"
					style={{ ...buttonStyle }}
					onClick={() => setShowCancelModal(true)}>
					取消
				</Button>
			</Col>
			<Col span={3}>
				<Button
					disabled={
						+(orderStore.order && orderStore.order.orderStatus) >= 8160 ||
						+(orderStore.order && orderStore.order.orderStatus) <= 1
					}
					type="primary"
					style={{ ...buttonStyle }}
					onClick={() => {
						orderStore.initDeclaredAmount(
							{
								orderId: orderStore.orderId
							},
							() => setShowChangeMoneyModal(true)
						);
					}}>
					修改申报金额
				</Button>
			</Col>
			<Col span={3}>
				<Button
					// disabled={
					// 	+(orderStore.order && orderStore.order.orderStatus) >= 8150 ||
					// 	+(orderStore.order && orderStore.order.orderStatus) <= 1
					// }
					type="primary"
					style={{ ...buttonStyle }}
					onClick={() => setShowRemarkModal(true)}>
					添加备注
				</Button>
			</Col>
			{showAddModal && <AddEventModal handleClose={() => setShowAddModal(false)} />}
			{showBackModal && (
				<ChargeBackModal
					handleClose={() => {
						setShowBackModal(false);
						commonStore.clearList("ChargeBack");
					}}
				/>
			)}
			{showCancelModal && <CancelModal handleClose={() => setShowCancelModal(false)} />}
			{showChangeMoneyModal && (
				<ChangeMoneyModal
					showChangeMoneyModal={showChangeMoneyModal}
					handleClose={() => setShowChangeMoneyModal(false)}
				/>
			)}
			{showRemarkModal && (
				<RemarkModal
					record={orderStore.order}
					handleClose={() => setShowRemarkModal(false)}
				/>
			)}
		</Row>
	));
}

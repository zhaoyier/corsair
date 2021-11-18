import * as React from "react";
import { Modal, Form, Input, Radio } from "antd";
import { store } from "../../store/helper/useStore";
import { warehouseListArray, WarehouseCorrespondingCountryArray } from "../../constant";
import { useObserver } from "mobx-react-lite";
import { CancelOrderEnum, UserPlatform, sampleOrderInfo } from "genServices/ezShipOMS/oms";
interface ChargeBackModalProps {
	orderId?: string;
	order?: sampleOrderInfo;
	handleClose: () => void;
}
const layout = {
	labelCol: { span: 6 },
	wrapperCol: { span: 16 }
};
export function ChargeBackModal(props: ChargeBackModalProps) {
	const { handleClose, orderId, order } = props;
	const [form] = Form.useForm();
	const { getFieldValue } = form;
	const [needBack, setNeedBack] = React.useState(true);
	const commonStore = store().commonStore;
	const orderStore = store().orderStore;
	const orderListStore = store().orderListStore;
	const orderListUi = store().orderListUi;
	const onOk = async () => {
		form.submit();
	};
	const warehouseName = orderId
		? warehouseListArray.find(elem => elem.value === order.warehouseId).label
		: orderStore.order.warehouseName;
	const warehouseId = warehouseListArray.findIndex(i => i.label === warehouseName) + "";
	const catalogShow = WarehouseCorrespondingCountryArray.find(i => i.value === warehouseName)
		.labelCN;
	const catalog = WarehouseCorrespondingCountryArray.find(i => i.value === warehouseName).labelEN;
	function onChange(e) {
		// console.log("radio checked", e.target.value);
		setNeedBack(e.target.value);
	}

	function warning(message: string) {
		return Modal.warning({
			title: "提示",
			content: message
		});
	}
	return useObserver(() => (
		<Modal
			title={
				<React.Fragment>
					<span>退单</span>
					<span style={{ color: "#df2c42", fontSize: "12px", marginLeft: 4 }}>
						（转运仓库签收后的包裹处理)
					</span>
				</React.Fragment>
			}
			visible={true}
			onOk={onOk}
			onCancel={handleClose}>
			<Form.Provider
				onFormFinish={async () => {
					const res = await orderStore.cancelOrder({
						orderId: orderId ? orderId : orderStore.orderId,
						cancelType: CancelOrderEnum.CANCEL_ORDER_2, // 退单
						platformId: UserPlatform.UserPlatformOms,
						refunds: {
							isRefunds: getFieldValue("isRefunds"),
							catalog,
							warehouseId,
							name: getFieldValue("name"),
							address: getFieldValue("address"),
							phone: getFieldValue("phone")
						}
					});
					if (res.code === "104") {
						warning(res.message);
					}
					commonStore.clearList("ChargeBack");
					if (orderId) {
						await orderListStore.searchOrderList({
							...orderListStore.searchParams,
							offset: (orderListUi.page - 1) * orderListUi.limit,
							limit: orderListUi.limit
						} as any);
					} else {
						await orderStore.initOrderDetail({
							orderId: orderId ? orderId : orderStore.orderId
						});
					}
					handleClose();
				}}>
				<Form
					{...layout}
					name="basic"
					form={form}
					initialValues={{ remember: true, isRefunds: true }}>
					<Form.Item
						label="是否需求退货"
						name="isRefunds"
						rules={[{ required: true, message: "请选择是否退货！" }]}>
						<Radio.Group onChange={onChange} value={needBack}>
							<Radio value={true}>是</Radio>
							<Radio value={false}>否</Radio>
						</Radio.Group>
					</Form.Item>
					<Form.Item name="warehouseId" label="仓库">
						<div>{warehouseName}</div>
					</Form.Item>
					<Form.Item name="catalog" label="目的地国家">
						<div>{catalogShow}</div>
					</Form.Item>
					{needBack && (
						<React.Fragment>
							<Form.Item
								label="收件人"
								name="name"
								rules={[{ required: needBack, message: "请输入收件人！" }]}>
								<Input />
							</Form.Item>
							<Form.Item
								label="收件地址"
								name="address"
								rules={[{ required: needBack, message: "请输入收件地址！" }]}>
								<Input />
							</Form.Item>
							<Form.Item
								label="收件电话"
								name="phone"
								rules={[{ required: needBack, message: "请输入收件电话！" }]}>
								<Input />
							</Form.Item>
						</React.Fragment>
					)}
				</Form>
			</Form.Provider>
		</Modal>
	));
}

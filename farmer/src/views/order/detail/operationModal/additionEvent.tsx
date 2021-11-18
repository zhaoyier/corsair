import * as React from "react";
import { Modal, Select, Row, Col, Input, message } from "antd";
const { Option } = Select;
import { store } from "../../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { CreateAddedServiceReq } from "genServices/ezShipOMS/oms";
import { addServiceTypeArr } from "../../constant";
import { getUserInfo } from "utils/user";
import { TUser } from "genServices/AdminLoginService";

interface AddEventModalProps {
	orderIdArr?: string[];
	actionTypeisMulti?: boolean;
	handleClose: () => void;
}

export function AddEventModal(props: AddEventModalProps) {
	const orderStore = store().orderStore;
	const orderListStore = store().orderListStore;
	const orderListUi = store().orderListUi;
	const { handleClose, orderIdArr, actionTypeisMulti } = props;
	const [remark, setRemark] = React.useState("");
	const [serviceType, setServiceType] = React.useState([]);
	const [userInfo, setUserInfo] = React.useState<TUser>({ username: "" });
	const [loading, setLoading] = React.useState(false);

	React.useEffect(() => {
		initAddEventModal();
	}, []);

	async function initAddEventModal() {
		const userInfo = getUserInfo();
		setUserInfo(userInfo);
	}
	function handleChange(value) {
		setServiceType(value);
	}

	async function submitAddServer() {
		const { username } = userInfo;
		try {
			setLoading(true);
			let param = {
				createBy: username,
				serviceType,
				remark
			};
			if (actionTypeisMulti) {
				// 批量操作 > 1
				const resq = await orderListStore.batchCreateAddedService({
					orderIds: orderIdArr,
					...param
				});
				message.warn(resq.result.msg, 5);
			} else {
				await orderStore.createAddedService({
					orderId:
						orderIdArr && orderIdArr.length > 0 ? orderIdArr[0] : orderStore.orderId,
					...param
				} as CreateAddedServiceReq);
			}
			if (orderIdArr && orderIdArr.length > 0) {
				await orderListStore.searchOrderList({
					...orderListStore.searchParams,
					offset: (orderListUi.page - 1) * orderListUi.limit,
					limit: orderListUi.limit
				} as any);
			} else {
				await orderStore.initOrderDetail({ orderId: orderStore.orderId });
			}
		} catch (e) {
			console.log(e);
		} finally {
			setLoading(false);
			handleClose();
		}
	}
	return useObserver(() => (
		<Modal
			title="增加附加事件"
			visible={true}
			onOk={submitAddServer}
			onCancel={handleClose}
			okText="确定"
			cancelText="取消"
			confirmLoading={loading}>
			<Row>
				<Col span={6}>服务名称:</Col>
				<Col span={16}>
					<Select mode="multiple" style={{ width: "100%" }} onChange={handleChange}>
						{/* {addServiceTypeArr.map(({ label, value }) => (
							<Option value={value} key={value}>
								{label}
							</Option>
						))} */}
						{(orderStore.addedServiceTerm || []).map((item, index) => (
							<Option value={item} key={index}>
								{addServiceTypeArr.find(i => i.value === item) &&
									addServiceTypeArr.find(i => i.value === item).label}
							</Option>
						))}
					</Select>
				</Col>
			</Row>
			<Row style={{ marginTop: 20 }}>
				<Col span={6}>添加备注:</Col>
				<Col span={16}>
					<Input onChange={e => setRemark(e.target.value)} />
				</Col>
			</Row>
		</Modal>
	));
}

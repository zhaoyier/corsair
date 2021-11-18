import * as React from "react";
import { Modal, message } from "antd";
import { store } from "../../store/helper/useStore";
import { CancelOrderEnum, UserPlatform, sampleOrderInfo } from "genServices/ezShipOMS/oms";
import { getUserInfo } from "utils/user";
import { TUser } from "genServices/AdminLoginService";

interface CancelModalProps {
	orderIdArr?: string[];
	orderArr?: sampleOrderInfo[];
	actionTypeisMulti?: boolean;
	handleClose: () => void;
}
export function CancelModal(props: CancelModalProps) {
	const { handleClose, orderIdArr, orderArr, actionTypeisMulti } = props;
	const orderStore = store().orderStore;
	const orderListStore = store().orderListStore;
	const orderListUi = store().orderListUi;
	const [userInfo, setUserInfo] = React.useState<TUser>({ username: "" });
	const [loading, setLoading] = React.useState(false);

	React.useEffect(() => {
		initAddEventModal();
	}, []);

	async function initAddEventModal() {
		const userInfo = getUserInfo();
		setUserInfo(userInfo);
	}

	const cancel = async () => {
		const { username } = userInfo;
		try {
			setLoading(true);
			const param = {
				cancelType: CancelOrderEnum.CANCEL_ORDER_1,
				platformId: UserPlatform.UserPlatformOms,
				updateBy: username
			};
			if (actionTypeisMulti) {
				// 批量操作
				const resq = await orderListStore.batchCancellOrder({
					orderIds: orderIdArr,
					...param
				});
				message.warn(resq.result.msg, 5);
			} else {
				await orderStore.cancelOrder({
					orderId:
						orderIdArr && orderIdArr.length > 0 ? orderIdArr[0] : orderStore.orderId,
					...param
				});
			}
			if (orderIdArr && orderIdArr.length > 0) {
				await orderListStore.searchOrderList({
					...orderListStore.searchParams,
					offset: (orderListUi.page - 1) * orderListUi.limit,
					limit: orderListUi.limit
				} as any);
			} else {
				await orderStore.initOrderDetail({
					orderId: orderStore.orderId
				});
			}
			handleClose();
		} catch (err) {
			console.log(err);
		} finally {
			setLoading(false);
		}
	};

	return (
		<Modal
			title={
				<React.Fragment>
					<span>取消订单</span>
					<span style={{ color: "#df2c42", fontSize: "12px", marginLeft: 4 }}>
						（转运仓库签收前取消)
					</span>
				</React.Fragment>
			}
			visible={true}
			okText="确定"
			cancelText="取消"
			confirmLoading={loading}
			onOk={cancel}
			onCancel={handleClose}>
			<p>确定要取消当前订单，且取消相应的订单附加服务？</p>
			<p>
				订单号：
				{orderArr && orderArr.length > 0
					? orderArr.map(elem => elem.orderNumber).join(",")
					: orderStore.order && orderStore.order.orderNo}
			</p>
		</Modal>
	);
}

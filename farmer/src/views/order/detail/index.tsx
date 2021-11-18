import * as React from "react";
import { ActiveButton } from "./activeButton";
import { OrderStatusShow } from "./orderStatus";
import { OrderInfo } from "./orderInfo";
import { useObserver } from "mobx-react-lite";
import { store } from "../store/helper/useStore";
import { OrderStatus } from "genServices/ezShipOMS/public";

const styles = require("./index.scss");
function Detail() {
	const orderStore = store().orderStore;
	const commonStore = store().commonStore;
	React.useEffect(() => {
		initDetailPage();
	}, []);
	function initDetailPage() {
		orderStore.initOrderStore();
		initOrderStatus();
	}
	function initOrderStatus() {
		const orderStatusList =
			commonStore.orderStatusList && commonStore.orderStatusList.orderStatusList;
		const code = orderStore.order && orderStore.order.orderStatus;
		const result = filterOrderStatus(orderStatusList, code);
		orderStore.setOrderStatus(result);
	}
	function filterOrderStatus(orderStatusList: OrderStatus[], Code: string) {
		if (!orderStatusList) {
			return null;
		}
		return orderStatusList.find(({ statusCode }) => statusCode === Code);
	}
	return useObserver(() => (
		<div className={styles.index}>
			<ActiveButton />
			<OrderStatusShow />
			<OrderInfo />
		</div>
	));
}

export default Detail;

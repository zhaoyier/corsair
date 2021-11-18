import { OrderListStore } from "./orderList/orderList";
import { OrderListUi } from "./orderList/orderListUi";
import { OrderUi } from "./order/orderUi";
import { OrderStore } from "./order/order";
import { CommonStore } from "./common";

const store = {
	commonStore: new CommonStore(),
	orderListStore: new OrderListStore(),
	orderListUi: new OrderListUi(),
	orderUi: new OrderUi(),
	orderStore: new OrderStore()
};

export default store;

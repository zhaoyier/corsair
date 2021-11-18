import { action, observable } from "mobx";
import { DetailModalType } from "views/order/constant";
import { AddedServiceStatusEnum } from "genServices/ezShipOMS/oms";
import { OrderStatus } from "genServices/ezShipOMS/public";
export class OrderUi {
	@observable orderStatus: OrderStatus; // 订单状态用来控制活动按钮的
	@observable detailModalType: DetailModalType;
	@observable addedServiceStatusEnum: AddedServiceStatusEnum;
	// @observable orderStatus: string = ""; // 订单状态用来控制活动按钮的置灰还有其他的

	@action
	initOrderUi = () => {
		this.initDetailModalType();
	};

	@action
	initDetailModalType = () => {
		this.detailModalType = DetailModalType.none;
	};

	@action
	setDetailModalType = (detailModalType: DetailModalType) => {
		this.detailModalType = detailModalType;
	};
}

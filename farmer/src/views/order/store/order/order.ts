import { action, observable } from "mobx";
import {
	OrderDetail,
	OrderDetailReq,
	OrderInfo,
	DeliveryInfo,
	AllRemarks,
	OrderPayInfo,
	PayInfoResp,
	PayInfoReq,
	OrderLogs,
	OrderLogReq,
	LogInfo,
	CreateAddedService,
	CreateAddedServiceReq,
	CancelOrder,
	CancelOrderReq,
	ModifyDeclaredAmount,
	ModifyDeclaredReq,
	AddRemarks,
	AddRemarkReq,
	CancelAddedService,
	CancelAddedServiceReq,
	OrderAddedService,
	GetAddedServiceReq,
	OrderPakgeLogistics,
	LogisticsReq,
	AddedServiceTermReq,
	UpdateOrderStatus,
	RpcUpdateOrderStatusReq,
	GetAddedServiceTerm,
	AddServiceType,
	GetDeclaredAmount,
	DeclaredAmountReq,
	DeclaredAmountItem
} from "genServices/ezShipOMS/oms";
import {
	ModifyArrivalNoticeDocument,
	ModifyArrivalNoticeReq
} from "genServices/ezShipOMS/external";
import { OrderStatus } from "genServices/ezShipOMS/public";
import { tryCatchResult, tryCatchMessage } from "views/order/utils/tryCatch";
import { getUrlParams } from "utils/url";
import { notification } from "antd";

export class OrderStore {
	constructor() {
		// const orderId = this.getRouterInfo().orderId;
		// this.initOrderDetail({ orderId } as OrderDetailReq);
	}
	@observable orderId: string;
	@observable order: OrderInfo;
	@observable delivery: DeliveryInfo;
	@observable remarks: AllRemarks;
	@observable orderPayInfo: PayInfoResp;
	@observable logInfo: LogInfo[];
	@observable orderStatus: OrderStatus;
	@observable addedServiceTerm: any;
	@observable declaredAmount: DeclaredAmountItem[];

	@observable createBy: string = "";

	getRouterInfo() {
		const params = getUrlParams(window.location.href) as any;
		const orderId = params.orderId || "";
		this.orderId = orderId;
		return { orderId };
	}
	@action
	initOrderStore = async () => {
		const orderId = this.getRouterInfo().orderId;
		// await this.initOrderStatusList();
		this.initOrderDetail({ orderId } as OrderDetailReq);
		this.initOrderPayInfo({ orderId } as PayInfoReq);
		this.initOrderLogs({ orderId } as OrderLogReq);
		// this.initAddedServiceTerm({ warehouse: warehouseListArray.find(i => i.label === this.order.warehouseName) && warehouseListArray.find(i => i.label === this.order.warehouseName).value });
	};

	@action
	initOrderDetail = async ({ orderId }: OrderDetailReq) => {
		const result = await this.getOrderDetail({ orderId });
		this.order = result.order;
		this.delivery = result.delivery;
		this.remarks = result.remarks;
	};
	@action
	initOrderPayInfo = async ({ orderId }: PayInfoReq) => {
		const result = await this.getOrderPayInfo({ orderId });
		this.orderPayInfo = result;
	};
	@action
	initOrderLogs = async ({ orderId }: OrderLogReq) => {
		const result = await this.getOrderLogs({ orderId });
		this.logInfo = result.Logs;
	};
	@action
	initAddedServiceTerm = async ({ warehouse }: AddedServiceTermReq) => {
		const result = await this.getAddedServiceTerm({ warehouse });
		this.addedServiceTerm = result.serviceType.filter(
			i => i !== AddServiceType.Service_invalid
		);
	};
	@action
	initDeclaredAmount = async ({ orderId }: DeclaredAmountReq, cb?: () => void) => {
		const result = await this.getDeclaredAmount({ orderId });
		this.declaredAmount = result.items;
		cb && cb();
	};
	@action
	getOrderDetail = async (orderDetailReq: OrderDetailReq) => {
		return await tryCatchResult(OrderDetail(orderDetailReq), "getOrderDetail");
	};
	// 查看支付信息
	@action
	getOrderPayInfo = async (payInfoReq: PayInfoReq) => {
		return await tryCatchResult(OrderPayInfo(payInfoReq), "getOrderPayInfo");
	};
	@action
	getOrderLogs = async (orderLogReq: OrderLogReq) => {
		return await tryCatchResult(OrderLogs(orderLogReq), "getOrderLogReqInfo");
	};
	// 增加附加服务
	@action
	createAddedService = async (createAddedServiceReq: CreateAddedServiceReq) => {
		return await tryCatchMessage(CreateAddedService(createAddedServiceReq));
	};
	// 到货通知到更新接口  可以把状态更新为已取消
	@action
	modifyArrivalNoticeDocument = async (modifyArrivalNoticeReq: ModifyArrivalNoticeReq) => {
		return await tryCatchResult(
			ModifyArrivalNoticeDocument(modifyArrivalNoticeReq),
			"modifyArrivalNoticeDocument"
		);
	};
	// 取消订单
	@action
	cancelOrder = async (cancelOrderReq: CancelOrderReq) => {
		try {
			const res = await CancelOrder(cancelOrderReq);
			if (res && res.result) {
				if (res.result.code === "104") {
					return res.result;
				} else {
					notification.success({
						message: `${res.result.message || "success"}`
					});
					return res.result.message;
				}
			}
		} catch (e) {
			notification.error({
				message: `${e.message}`,
				description: e.message
			});
			return e.result;
		}
	};
	// 更改订单状态
	@action
	updateOrderStatus = async ({}) => {
		return await tryCatchResult(UpdateOrderStatus({} as RpcUpdateOrderStatusReq));
	};
	// 删除增值服务
	@action
	cancelAddedService = async (cancelAddedServiceReq: CancelAddedServiceReq) => {
		return await tryCatchMessage(CancelAddedService(cancelAddedServiceReq));
	};
	// 查看增值服务
	@action
	orderAddedService = async (getAddedServiceReq: GetAddedServiceReq) => {
		return await tryCatchResult(OrderAddedService(getAddedServiceReq), "orderAddedService");
	};
	// 查看物流轨迹
	@action
	orderPakgeLogistics = async (logisticsReq: LogisticsReq) => {
		return await tryCatchResult(OrderPakgeLogistics(logisticsReq), "orderPakgeLogistics");
	};
	// 修改申报金额
	@action
	modifyDeclaredAmount = async (modifyDeclaredReq: ModifyDeclaredReq) => {
		return await tryCatchMessage(ModifyDeclaredAmount(modifyDeclaredReq));
	};
	// 增加备注
	@action
	addRemarks = async (addRemarkReq: AddRemarkReq) => {
		return await tryCatchMessage(AddRemarks(addRemarkReq));
	};
	@action
	setOrderStatus = (orderStatus: OrderStatus) => {
		this.orderStatus = orderStatus;
	};
	// 获取添加附加服务的服务名称
	@action
	getAddedServiceTerm = async (addedServiceTermReq: AddedServiceTermReq) => {
		return await tryCatchResult(
			GetAddedServiceTerm(addedServiceTermReq),
			"getAddedServiceTerm"
		);
	};
	@action
	getDeclaredAmount = async (declaredAmountReq: DeclaredAmountReq) => {
		return await tryCatchResult(GetDeclaredAmount(declaredAmountReq), "getDeclaredAmount");
	};
}

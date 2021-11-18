import { action, observable } from "mobx";
import {
	SearchOrderReq,
	SearchOrderResp,
	sampleOrderInfo,
	SubmitNewOrder,
	SubmitNewOrderReq,
	IsBatchExecutable,
	IsBatchExecutableReq,
	BatchCreateAddedService,
	BatchCreateAddedServiceReq,
	BatchCancellOrder,
	BatchCancellOrderReq
} from "genServices/ezShipOMS/oms";

import { SearchOrder } from "genServices/fruit/orchard";
import { tryCatchResult } from "views/order/utils/tryCatch";
export class OrderListStore {
	@observable searchOrderResp: SearchOrderResp;
	@observable orderList: sampleOrderInfo[] = [];
	@observable total: number = 0;
	@observable searchParams: SearchOrderReq;
	@observable spinning: boolean = false;

	@observable showAddModal: boolean = false; // 添加附加服务
	@observable showBackModal: boolean = false; // 退单
	@observable showCancelModal: boolean = false; // 取消
	@observable showRemarkModal: boolean = false; // 添加备注

	@observable actionTypeisMulti: boolean = false; // 是否为批量操作

	@observable selectedRowKeys: string[] = [];
	@observable selectedRows: sampleOrderInfo[] = [];

	@action
	changeListStore = (param: string, value) => {
		this[param] = value;
	};

	@action initOrderListStore = async () => {
		this.searchOrderResp = await this.initOrderList();
		this.orderList = this.searchOrderResp.items;
		this.total = +this.searchOrderResp.total;
		this.selectedRowKeys = [];
		this.selectedRows = [];
	};
	@action
	initOrderList = async () => {
		return await tryCatchResult(
			this.getOrderList({ offset: "0", limit: "50" }),
			"initOrderList"
		);
	};

	@action
	saveSearchParams = (searchParams: SearchOrderReq) => {
		this.searchParams = searchParams;
	};
	@action
	searchOrderList = async (searchOrderReq: SearchOrderReq) => {
		this.spinning = true;
		this.searchOrderResp = await tryCatchResult(
			this.getOrderList(searchOrderReq),
			"searchOrderList"
		);
		this.spinning = false;
		this.orderList = this.searchOrderResp.items;
		this.total = +this.searchOrderResp.total;
		this.selectedRowKeys = []; // 清空selectKeys
		this.selectedRows = [];
	};
	// 获取订单列表
	@action
	getOrderList = async (searchOrderReq: SearchOrderReq) => {
		return await tryCatchResult(SearchOrder(searchOrderReq), "getOrderList");
	};
	// 创建订单
	@action
	submitNewOrder = async (submitNewOrderReq: SubmitNewOrderReq) => {
		return await tryCatchResult(SubmitNewOrder(submitNewOrderReq), "submitNewOrder");
	};

	// IsBatchExecutable
	@action
	isBatchExecutable = async (isBatchExecutableReq: IsBatchExecutableReq) => {
		return await tryCatchResult(IsBatchExecutable(isBatchExecutableReq), "IsBatchExecutable");
	};

	// BatchCreateAddedService
	@action
	batchCreateAddedService = async (batchCreateAddedServiceReq: BatchCreateAddedServiceReq) => {
		return await tryCatchResult(
			BatchCreateAddedService(batchCreateAddedServiceReq),
			"BatchCreateAddedService"
		);
	};

	// BatchCancellOrder
	@action
	batchCancellOrder = async (batchCancellOrderReq: BatchCancellOrderReq) => {
		return await tryCatchResult(BatchCancellOrder(batchCancellOrderReq), "BatchCancellOrder");
	};
}

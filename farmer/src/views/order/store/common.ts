import { action, observable } from "mobx";
import {
	OrderStatusListResp,
	DeliveryMethodListResp,
	WarehouseListResp,
	LogisticsListResp,
	ShippingMethodListResp,
	DesCountryCodeListResp,
	ParcelStatusListResp,
	OrderStatus,
	CnHscodeResp
} from "genServices/ezShipOMS/public";
import {
	WarehouseList,
	OrderStatusList,
	DeliveryMethodList,
	LogisticsCOList,
	ShippingMethodList,
	DesCountryCode,
	ParcelStatusList,
	GetCnHscodeList,
	OrderOriginList,
	OrderOrigin
} from "genServices/ezShipOMS/oms";
import { tryCatchResult } from "views/order/utils/tryCatch";
import { WarehouseType } from "genServices/ezShipOrder/submit";

export class CommonStore {
	constructor() {
		this.initCommonStore();
	}

	@observable orderOriginList: OrderOrigin[] = []; // 订单来源列表
	@observable orderStatusList: OrderStatusListResp;
	@observable deliveryMethodList: DeliveryMethodListResp;
	// 搜索用转运仓库
	@observable warehouseListInSearch: WarehouseListResp;
	// 创建用转运仓库
	@observable warehouseListInCreate: WarehouseListResp;
	// 退单用转运仓库
	@observable warehouseListInChargeBack: WarehouseListResp;
	// 搜索用快递公司
	@observable logisticsListInSearch: LogisticsListResp;
	// 创建用快递公司
	@observable logisticsListInCreate: LogisticsListResp;
	@observable shippingMethodList: ShippingMethodListResp;
	@observable desCountryCode: DesCountryCodeListResp;
	@observable parcelStatusList: ParcelStatusListResp;
	@observable cnHscodeList: CnHscodeResp;

	@action
	initCommonStore = () => {
		this.initOrderStatusList();
		this.initOrderOriginList();
		this.initDeliveryMethodList();
		this.initShippingMethodList();
		this.initDesCountryCode();
		this.initParcelStatusList();
	};

	@action
	initOrderStatusList = async () => {
		const result = await this.getOrderStatusList();
		const sortedOrderStatusList = this.sortOrderStatusList(result.orderStatusList);
		this.orderStatusList = { orderStatusList: sortedOrderStatusList }; // 初始好的状态code加状态
	};

	@action
	initOrderOriginList = async () => {
		const result = await this.getOrderOriginList();
		this.orderOriginList = result ? result.origins || [] : [];
	};

	@action.bound
	sortOrderStatusList(orderStatusList: OrderStatus[]): OrderStatus[] {
		return orderStatusList.sort((pre, next) => +pre.statusCode - +next.statusCode);
	}

	@action
	initDeliveryMethodList = async () => {
		this.deliveryMethodList = await this.getDeliveryMethodList();
	};

	@action
	initShippingMethodList = async () => {
		this.shippingMethodList = await this.getShippingMethodList();
	};

	@action
	initDesCountryCode = async () => {
		this.desCountryCode = await this.getDesCountryCode();
	};

	@action
	initParcelStatusList = async () => {
		this.parcelStatusList = await this.getParcelStatusList();
	};

	/**
	 * 获取转运仓库列表
	 * @param use 用途 Search（列表搜索） | Create（创建订单）| ChargeBack（退单）
	 */
	@action
	fetchWarehouseList = async (
		use: "Search" | "Create" | "ChargeBack",
		catalog: string,
		clear?: boolean
	) => {
		if (clear) {
			this[`warehouseListIn${use}`] = undefined;
		} else {
			this[`warehouseListIn${use}`] = await this.getWarehouseList(catalog);
		}
	};

	/**
	 * 获取快递公司列表
	 * @param warehouse 转运仓库
	 */
	@action
	fetchLogisticsCOList = async (
		use: "Search" | "Create",
		warehouse?: WarehouseType,
		clear?: boolean
	) => {
		if (clear) {
			this[`logisticsListIn${use}`] = undefined;
		} else {
			this[`logisticsListIn${use}`] = await this.getLogisticsCOList(warehouse);
		}
	};

	@action
	fetchCnHscodeList = async () => {
		this.cnHscodeList = await this.getCnHscodeList();
	};

	@action
	clearList = async (use: "Search" | "Create" | "ChargeBack") => {
		this[`warehouseListIn${use}`] = undefined;
	};

	// 仓库
	@action
	getWarehouseList = async (catalog: string) => {
		return await tryCatchResult(WarehouseList({ catalog }));
	};
	// 订单来源
	@action
	getOrderOriginList = async () => {
		return await OrderOriginList({});
	};
	// 订单状态
	@action
	getOrderStatusList = async () => {
		return await OrderStatusList({});
	};
	// 派送方式
	@action
	getDeliveryMethodList = async () => {
		return await DeliveryMethodList({});
	};
	// 快递公司
	@action
	getLogisticsCOList = async (warehouse?: WarehouseType) => {
		return await LogisticsCOList({ warehouse });
	};
	// 运输方式
	@action
	getShippingMethodList = async () => {
		return await ShippingMethodList({});
	};
	// 目的地国家
	@action
	getDesCountryCode = async () => {
		return await DesCountryCode({});
	};
	// 包裹状态
	@action
	getParcelStatusList = async () => {
		return await ParcelStatusList({});
	};
	// 获取中国报关分类
	@action
	getCnHscodeList = async () => {
		return await GetCnHscodeList({});
	};
}

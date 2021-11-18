import { HeaderProps } from "components/CommonHeader";
import { SearchForm, ListParam } from "./types";
import { OMSRemarkType, SearchPackageState } from "genServices/ezShipOMS/oms";
import { WarehouseType } from "genServices/ezShipOMS/public";

export const listHeader: HeaderProps = {
	module: "包裹管理",
	current: {
		name: "Ezship包裹查询",
		hash: "list"
	}
};

export const detailHeader = (cb?: () => void): HeaderProps => {
	return {
		module: "包裹管理",
		parent: {
			name: "Ezship包裹查询",
			hash: "list"
		},
		current: {
			name: "包裹详情",
			hash: "detail"
		},
		cb: () => cb()
	};
};

// tab状态
export const TabStatus = {
	[SearchPackageState.SearchPackageStateUnkown]: "全部",
	[SearchPackageState.SearchPackageStateEtaExpired]: "ETA过期"
};

// childTab状态
export const ChildTabStatus = {
	parcelLog: "包裹日志",
	parcelRemark: "查看备注"
};

// 初始化 包裹查询条件
export const initSearchForm: SearchForm = {
	packageCode: "", // 包裹号
	deliveryId: "", // 配送单号
	packageStatue: "全部", // 包裹状态
	orderId: "", // 订单号
	catalog: "", // 目的地国家
	boxNumber: "", // 封箱号
	nickname: "", // 会员昵称
	// email: "", // 会员email
	phone: "", // 收货电话
	transportNumber: "", // 物流单号
	warehouseId: "全部", // 所属仓库
	shipmentType: "全部", // 运输方式
	deliveryType: "全部", // 取货方式
	isRemark: "全部", // 是否有备注
	cabinetNumber: "", // 主单/柜号
	deliverDate: "", // 送货日期
	etaDate: "", // 发货ETA
	searchState: Object.keys(TabStatus)[0], // 查询状态
	origin: "全部",
	billId: "" // 付款单号
};

// 初始化  包裹查询 分页参数
export const initListParam: ListParam = {
	limit: 50,
	offset: 0,
	current: 1,
	total: 0
};

// 所属仓库
export const WarehoseList = {
	[WarehouseType.WarehouseTypeGuangzhou]: "广州",
	[WarehouseType.WarehouseTypeTaiwan]: "台湾",
	[WarehouseType.WarehouseTypeUSA]: "美国",
	[WarehouseType.WarehouseTypeShanghai]: "上海"
}; // 货源地所在仓库
export const WarehoseListMy = {
	[WarehouseType.WarehouseTypeGuangzhou]: "广州",
	[WarehouseType.WarehouseTypeTaiwan]: "台湾",
	[WarehouseType.WarehouseTypeShanghai]: "上海"
}; // MY
// 是否有备注
export const IsRemarkList = {
	[OMSRemarkType.OMSRemarkTypePackage]: "包裹备注",
	[OMSRemarkType.OMSRemarkTypeDelivery]: "配送备注",
	[OMSRemarkType.OMSRemarkTypeNone]: "无备注"
};

export const RemarksTypes = {
	"1": "包裹备注",
	"2": "配送备注",
	"3": "全局包裹备注",
	"4": "全局配送备注"
};

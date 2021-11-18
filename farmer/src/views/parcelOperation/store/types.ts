import { PackageDetail, OMSRemark, PackageLogInfo } from "genServices/ezShipOMS/oms";
import { OrderOriginEnum } from "genServices/ezShipOMS/public";

export interface ArrSelectOption {
	label: string;
	value: string | number;
}

export interface SearchForm {
	packageCode: string; // 包裹号（排他）
	deliveryId: string; // 派送id(配送单号,排他)
	packageStatue: string; // 包裹状态
	orderId: string; // 订单号 (排他)
	catalog: string; // 目的地国家
	boxNumber: string; // 封箱号(排他)
	nickname: string; // 会员昵称
	// email: string; // 会员email
	phone: string; // 收货电话(排他)
	// deliveryCode: string; // 派送商
	transportNumber: string; // 物流单号(排他)
	warehouseId: string; // 所属仓库
	shipmentType: string; // 运输方式
	deliveryType: string; // 取货方式
	isRemark: string; // 是否有备注
	cabinetNumber: string; // 主单/柜号
	deliverDate: string; // 送货日期
	// pickupPeriod: string; // （预约取货时间）
	etaDate: string; // 发货ETA
	searchState: any;
	origin?: OrderOriginEnum | string;
	billId?: string;
}

export interface ListParam {
	limit: number; // 每页查询个数
	offset: number; // 差值
	current: number; // 当前页面
	total: number; // 总数
}

export interface ListResqItem extends PackageDetail {
	isLoading: boolean; // 表格扩展项loading状态
	childTabActiveKey: string; // 表格拓展项 -- tab激活页
	parcelLogList: PackageLogInfo[]; // 表格拓展项 -- 包裹日志表
	remarkList: OMSRemark[]; // 表格拓展项 --- 包裹备注表
}

// 待修改  extends PackageInfo
export interface PackageInfoItem {
	packageCode: string; // 包裹号（排他）
	deliveryId: string; // 派送id(配送单号,排他)
	warehouseId: string; // 所属仓库
	shipmentType: string; // 运输方式
	deliveryType: string; // 取货方式
	actualWeight: string; // 实际重
	chargWeight: string; // 计费重
	etaDate: string; // 发货ETA
	status: string; // 包裹状态
	childTabActiveKey: string; // 子表格活动Tab
	parcelLogList: PackageLogInfo[]; // 子表格 -- 包裹日志
	remarkList: OMSRemark[]; // 子表格 --- 查看备注
	isAftersales: boolean; // 是否补送过包裹
}

export interface UrlParam {
	packageCode: string;
}

export interface CancelData {
	deliveryId?: string;
	parcelNumList?: string[];
}

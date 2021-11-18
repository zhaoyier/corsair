import { WarehouseType, RemarkTypeEnum } from "genServices/ezShipOMS/public";
import {
	AddServiceType,
	AddedServiceStatusEnum,
	OrderRegTypeEnum
} from "genServices/ezShipOMS/oms";
import { HeaderProps } from "components/CommonHeader";
export const listHeader: HeaderProps = {
	module: "订单管理",
	current: {
		name: "Ezship订单查询",
		hash: "list"
	}
};

export const orderStatusArray = [
	{
		label: "全部",
		value: "all"
	},
	{
		label: "部分",
		value: "some"
	}
];
export const orderRegArray = [
	{
		label: "用户昵称",
		value: OrderRegTypeEnum.OrderRegTypeEnumNickName
	},
	{
		label: "用户ID",
		value: OrderRegTypeEnum.OrderRegTypeEnumUserId
	},
	{
		label: "标识码",
		value: OrderRegTypeEnum.OrderRegTypeEnumRegCode
	}
];

export const orderRegGlobalArray = [
	{
		label: "标识码",
		value: OrderRegTypeEnum.OrderRegTypeEnumRegCode
	}
];

export const catalogListArray = {
	SG: "新加坡",
	MY: "马来"
};
export const warehouseListArray = [
	{
		value: WarehouseType.WarehouseTypeShanghai,
		label: "上海"
	},
	{
		value: WarehouseType.WarehouseTypeGuangzhou,
		label: "广州"
	},
	{
		value: WarehouseType.WarehouseTypeTaiwan,
		label: "台湾"
	},
	{
		value: WarehouseType.WarehouseTypeUSA,
		label: "美国"
	}
];
export const WarehouseCorrespondingCountryArray = [
	{
		value: "上海",
		labelEN: "CN",
		labelCN: "中国"
	},
	{
		value: "广州",
		label: "CN",
		labelCN: "中国"
	},
	{
		value: "台湾",
		label: "TW",
		labelCN: "台湾"
	},
	{
		value: "美国",
		label: "US",
		labelCN: "美国"
	}
];
export const expressCompanyArray = [
	{
		label: "全部",
		value: "all"
	},
	{
		label: "部分",
		value: "some"
	},
	{
		label: "其他",
		value: "other"
	}
];
export const addServiceTypeArr = [
	{
		label: "拍照",
		value: AddServiceType.Service_PHOTOGRAPH
	},
	{
		label: "重打包",
		value: AddServiceType.Service_REPACKAGE
	}
];

export enum DetailModalType {
	addEventModal,
	none
}

export const RemarkTypeArr = [
	{
		label: "前台备注",
		value: RemarkTypeEnum.FRONT_REMARK
	},
	{
		label: "客户备注",
		value: RemarkTypeEnum.CUSTOMER_REMARK
	},
	{
		label: "收货备注",
		value: RemarkTypeEnum.RECEIVING_REMARK
	},
	{
		label: "订单备注",
		value: RemarkTypeEnum.ORDER_REMARK
	}
];

export const addedServiceStatusEnum = [
	{
		label: "新建",
		value: AddedServiceStatusEnum.Pending_Status
	},
	{
		label: "处理中",
		value: AddedServiceStatusEnum.Processing_Status
	},
	{
		label: "已完成",
		value: AddedServiceStatusEnum.Completed_Status
	},
	{
		label: "已取消",
		value: AddedServiceStatusEnum.CANCEL_Status
	}
];

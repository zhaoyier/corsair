import { observable, action, toJS } from "mobx";
import { initSearchForm, initListParam, ChildTabStatus, WarehoseList } from "./constant";
import { SearchForm, ListParam, PackageInfoItem, ListResqItem, ArrSelectOption } from "./types";
import { message } from "antd";
import { sortBy } from "lodash";
import { deleteEmpty } from "utils/util";
import {
	GetPackageRemark,
	SearchPackage,
	ParcelStatusList,
	DesCountryCode,
	WarehouseList,
	ShippingMethodList,
	DeliveryMethodList,
	SearchPackageReq,
	GetPackageLog,
	GetDeliveryRemark,
	OrderOrigin,
	OrderOriginList
} from "genServices/ezShipOMS/oms";
export default class ParcelListStore {
	@observable searchform: SearchForm = initSearchForm; // 查询 参数
	@observable pageParam: ListParam = initListParam; // 分页 参数
	@observable orderOriginList: OrderOrigin[] = [];
	@observable parcelStatusList: ArrSelectOption[] = []; // 包裹状态列表
	@observable desCodeList: string[] = []; // 目的地国家列表
	@observable warehouseList: ArrSelectOption[] = []; // 所属仓库列表
	@observable transportMethodList: ArrSelectOption[] = []; // 运输方式列表
	@observable deliveryMethodList: string[] = []; // 取货方式列表
	// @observable tabActiveKey: string = Object.keys(TabStatus)[0]; // 列表查询状态  ： 全部、ETA过期
	@observable parcelList: PackageInfoItem[] = []; // 包裹列表数据
	@observable listLoading: boolean = false; // 表格加载
	@observable selectedRowKeys: string[] = []; // 表格中选择的rowKeys
	@observable selectedRows: any[] = []; // 表格中选择的rows

	// 获取订单来源列表
	@action
	getOrderOriginList = () => {
		OrderOriginList({})
			.then(resq => {
				let orderOriginList = [];
				if (resq.origins && resq.origins.length > 0) {
					orderOriginList = resq.origins.map(item => ({
						label: item.name,
						value: item.code
					}));
				}
				this.changeStoreData("orderOriginList", orderOriginList);
			})
			.catch(err => message.error(err.message));
	};

	// 获取包裹状态列表
	@action
	getParcelStatusList = () => {
		ParcelStatusList({})
			.then(resq => {
				let parcelStatusList = [];
				if (resq.parcelStatusList && resq.parcelStatusList.length > 0) {
					parcelStatusList = resq.parcelStatusList.map(item => ({
						label: item.description,
						value: item.statusCode
					}));
				}
				this.changeStoreData("parcelStatusList", parcelStatusList);
			})
			.catch(err => message.error(err.message));
	};

	// 获取目的地国家列表
	@action
	getDestCountryList = () => {
		DesCountryCode({})
			.then(resq => {
				this.changeStoreData("desCodeList", resq.desCodeList || []);
			})
			.catch(err => message.error(err.message));
	};

	// 获取所属仓库列表
	@action
	getWarehouseList = () => {
		const { catalog } = this.searchform;
		WarehouseList({ catalog })
			.then(resq => {
				let warehouseList = [];
				if (resq.warehouse && resq.warehouse.length > 0) {
					warehouseList = resq.warehouse.map((item, index) => ({
						label: WarehoseList[item],
						value: item,
						key: String(index)
					}));
				}
				this.changeStoreData("warehouseList", warehouseList);
			})
			.catch(err => message.error(err.message));
	};

	// 获取运输方式列表
	@action
	getTransportMethodList = () => {
		ShippingMethodList({})
			.then(resq => {
				let transportMethodList = [];
				if (resq.deliveryMethodList && resq.deliveryMethodList.length > 0) {
					transportMethodList = resq.deliveryMethodList.map(item => ({
						label: item.description,
						value: item.shippingCode
					}));
				}
				this.changeStoreData("transportMethodList", transportMethodList);
			})
			.catch(err => message.error(err.message));
	};

	// 获取取货方式列表
	@action
	getPickupMethodList = () => {
		DeliveryMethodList({})
			.then(resq => {
				this.changeStoreData("deliveryMethodList", resq.deliveryMethodList || []);
			})
			.catch(err => message.error(err.message));
	};

	// 修改Store数据
	@action
	changeStoreData = (
		param: string,
		value,
		subParam?: string,
		secParam?: string,
		cb?: () => void
	) => {
		if (subParam) {
			if (secParam) {
				this[param][subParam][secParam] = value;
			} else {
				this[param][subParam] = value;
			}
		} else {
			this[param] = value;
		}
		if (cb) {
			cb();
		}
	};

	// 获取 查询接口参数
	private getSearchParam = (offset: number, limit: number) => {
		const selectArr = [
			"catalog",
			"origin",
			"packageStatue",
			"warehouseId",
			"shipmentType",
			"deliveryType",
			"isRemark"
		];
		const param: SearchPackageReq = deleteEmpty(this.searchform, selectArr);
		return {
			...param,
			offset,
			max: limit
		} as any;
	};

	// 查询包裹列表
	@action
	searchParcelList = (offset: number = 0, limit: number = 50) => {
		const param = this.getSearchParam(offset, limit);
		this.changeStoreData("listLoading", true);
		// 清除选中项
		this.changeStoreData("selectedRowKeys", []);
		this.changeStoreData("selectedRows", []);
		// 调用后端接口查询 列表
		// 1.将childTabActiveKey 字段加入 Object.keys(ChildTabStatus)[0]
		// 2.添加 parcelLogList, remarkList 字段 []
		SearchPackage(param)
			.then(resq => {
				if (resq.result.code === "0") {
					let list = [];
					if (resq.items && resq.items.length > 0) {
						resq.items.forEach((elem: ListResqItem) => {
							elem.isLoading = false;
							elem.childTabActiveKey = Object.keys(ChildTabStatus)[0];
							elem.parcelLogList = [];
							elem.remarkList = [];
							list.push(elem);
						});
					}
					this.changeStoreData("parcelList", list || []);
					this.changeStoreData("pageParam", resq.total || 0, "total");
					this.changeStoreData("listLoading", false);
				} else {
					message.error(resq.result.message);
					this.changeStoreData("listLoading", false);
				}
			})
			.catch(err => {
				message.error(err.message);
				this.changeStoreData("listLoading", false);
			});
	};

	// 分页 查询
	@action
	changePage = (current: number = 1, pageSize: number = 50) => {
		const offset = Number(current) > 1 ? (Number(current) - 1) * pageSize : 0;
		this.pageParam.current = current;
		this.pageParam.limit = pageSize;
		this.pageParam.offset = offset;
		this.searchParcelList(offset, pageSize);
	};

	// 修改表格数据
	@action
	changeParcelTable = (index, param, value) => {
		const list = [...toJS(this.parcelList)];
		list[index][param] = value;
		this.parcelList = [...list];
	};

	// 获取包裹日志List
	@action
	getPackageLog = async (packageCode: string, index: number) => {
		if (index > -1) {
			this.changeParcelTable(index, "isLoading", true);
			await GetPackageLog({ packageCode })
				.then(resq => {
					if (resq.result.code === "0") {
						this.changeParcelTable(index, "parcelLogList", resq.items || []);
						this.changeParcelTable(index, "isLoading", false);
						console.log(index, toJS(this.parcelList)[index]);
					} else {
						message.error(resq.result.message);
					}
				})
				.catch(err => {
					message.error(err.message);
					this.changeParcelTable(index, "isLoading", false);
				});
		} else {
			message.error("包裹号不存在");
		}
	};

	// 获取备注
	@action
	getTotalRemarks = async (packageCode: string, deliveryId: string, index: number) => {
		let packageRemarks = [];
		let deliveryRemarks = [];
		if (index === -1) {
			message.error("包裹号不存在");
			return;
		}
		const packageRemarkResq = await GetPackageRemark({
			packageCode: [packageCode]
		});
		const deliveryRemarkResq = await GetDeliveryRemark({
			deliveryId,
			packageCode: [packageCode]
		});
		if (packageRemarkResq.result && packageRemarkResq.result.code === "0") {
			packageRemarks = packageRemarkResq.remarks || [];
		} else {
			message.error(packageRemarkResq.result.message);
		}
		if (deliveryRemarkResq.result && deliveryRemarkResq.result.code === "0") {
			deliveryRemarks = deliveryRemarkResq.remarks || [];
		} else {
			message.error(deliveryRemarkResq.result.message);
		}
		const sorted = sortBy(packageRemarks.concat(deliveryRemarks), "updateDate").reverse();
		this.changeParcelTable(index, "remarkList", sorted || []);
	};
}

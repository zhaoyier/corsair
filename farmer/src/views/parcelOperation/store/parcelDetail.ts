import { observable, action } from "mobx";
import { UrlParam } from "./types";
import { getUrlParams } from "utils/url";
import {
	PackageDetail,
	PackageDetailResp,
	GetPackageLog,
	GetPackageRemark,
	GetPackageTrail,
	PackageLogInfo,
	OMSRemark,
	PackageTrailInfo,
	GetDeliveryRemark
} from "genServices/ezShipOMS/oms";
import { message, Modal } from "antd";
import { sortBy } from "lodash";

export default class ParcelDetailStore {
	@observable packageCode: string = "";
	@observable record: PackageDetailResp = {};
	@observable activeKey: string = "1";
	@observable parcelLog: PackageLogInfo[] = []; // 包裹日志
	@observable parcelRemarks: OMSRemark[] = []; // 包裹备注
	@observable parcelTrack: PackageTrailInfo[] = []; // 包裹轨迹
	// 修改Store数据
	@action
	changeStoreData = (param: string, value, subParam?: string, secParam?: string) => {
		if (subParam) {
			if (secParam) {
				this[param][subParam][secParam] = value;
			} else {
				this[param][subParam] = value;
			}
		} else {
			this[param] = value;
		}
	};

	@action
	getDetail = () => {
		const params = getUrlParams(window.location.href) as UrlParam;
		this.changeStoreData("packageCode", params.packageCode); // 设置包裹号
		// 查询包裹详情
		PackageDetail({ packageCode: params.packageCode })
			.then(resq => {
				if (resq.result.code === "0") {
					this.changeStoreData("record", resq);
				} else {
					Modal.error({
						title: "查询详情失败！",
						content: resq.result.message
					});
				}
			})
			.catch(err => message.error(err.message));
	};

	// 获取包裹日志
	@action
	getPackageLog = () => {
		GetPackageLog({ packageCode: this.packageCode })
			.then(resq => {
				if (resq.result.code === "0") {
					this.changeStoreData("parcelLog", resq.items || []);
				} else {
					message.error(resq.result.message);
				}
			})
			.catch(err => message.error(err.message));

		this.changeStoreData("parcelLog", []);
	};

	// 获取备注
	@action
	getTotalRemarks = async () => {
		let packageRemarks = [];
		let deliveryRemarks = [];
		const packageRemarkResq = await GetPackageRemark({
			packageCode: [this.packageCode]
		});
		const deliveryRemarkResq = await GetDeliveryRemark({
			deliveryId: this.record.shipmentId,
			packageCode: [this.packageCode]
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
		this.changeStoreData("parcelRemarks", sorted || []);
	};

	// 包裹物流轨迹
	@action
	getPackageTrail = () => {
		GetPackageTrail({ packageCode: this.packageCode })
			.then(resq => {
				this.changeStoreData("parcelTrack", resq.items || []);
			})
			.catch(err => message.error(err.message));
	};
}

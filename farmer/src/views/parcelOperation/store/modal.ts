import { observable, action, toJS } from "mobx";
import { CancelData } from "./types";
import * as _ from "lodash";
import { DeliveryNumModalProps } from "../Components/Modal/DeliveryNumModal";
import { CancelParcelModalProps } from "../Components/Modal/CancelParcelModal";
import { DelayNotificationModalProps } from "../Components/Modal/DelayNotificationModal";
import { EditDeliveryInfoModalProps } from "../Components/Modal/EditDeliveryInfoModal";
import { PODRemarksModalProps } from "../Components/Modal/PODRemarksModal";
import { PrintLabelModalProps } from "../Components/Modal/PrintLabelModal";
import { SplitDeliveryModalProps } from "../Components/Modal/SplitDeliveryModal";
import {
	initDeliveryNumModalData,
	initCancelParcelModalData,
	initComplementSendModalData,
	initDelayNotificationModalData,
	initEditDeliveryInfoModalData,
	initPODRemarksModalData,
	initPrintLabelModalData,
	initSplitDeliveryModalData
} from "../Components/Modal/constant";
import { ComplementSendModalProps } from "../Components/Modal/ComplementSendModal";
import { message } from "antd";
import {
	GetPackageRemark,
	GetDeliveryDetail,
	CancelDeliveryParcel,
	SplitShipmentParcel,
	GetPackageFaceInfo,
	AddPackageRemark,
	AddDeliveryRemark,
	GetDeliveryRemark,
	PackageDetail,
	AftersaleRedelivery
} from "genServices/ezShipOMS/oms";
export default class ModaltStore {
	@observable deliveryNumModalData: DeliveryNumModalProps = initDeliveryNumModalData; // 配送单号 弹窗信息
	@observable complementSendModalData: ComplementSendModalProps = initComplementSendModalData; // 补发
	@observable cancelParcelModalData: CancelParcelModalProps = initCancelParcelModalData; // 取消包裹 弹窗信息
	@observable
	delayNotificationModalData: DelayNotificationModalProps = initDelayNotificationModalData; // 延误通知
	@observable
	editDeliveryInfoModalData: EditDeliveryInfoModalProps = initEditDeliveryInfoModalData; // 编辑派送信息
	@observable pODRemarksModalData: PODRemarksModalProps = initPODRemarksModalData; // 包裹备注信息  派送备注信息
	@observable printLabelModalData: PrintLabelModalProps = initPrintLabelModalData; // 打印标签信息
	@observable splitDeliveryModalData: SplitDeliveryModalProps = initSplitDeliveryModalData; // 拆分派送信息

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

	// 点击 配送包裹号
	@action
	clickDeliveryNum = record => {
		// 配送单号查询包裹列表
		GetDeliveryDetail({ deliveryId: record.deliveryId })
			.then(resq => {
				if (resq.result.code === "0") {
					const deliveryNumModalData = {
						visible: true,
						deliveryId: record.deliveryId,
						parcelList: resq.packages || [],
						notOutboundOrderList: resq.orders || []
					};
					this.changeStoreData("deliveryNumModalData", deliveryNumModalData);
				} else {
					message.error(resq.result.message);
				}
			})
			.catch(err => message.error(err.message));
	};

	// 点击 补送
	@action
	clickComplementSend = record => {
		if (record.isAftersales) {
			message.error("该包裹被补送过了,不可重复补送");
		} else {
			PackageDetail({ packageCode: record.packageCode })
				.then(resq => {
					if (resq.result.code === "0") {
						const record = { ...resq };
						record.items = _.uniqBy(toJS(record.items), "orderId");
						const complementSendModalData = {
							visible: true,
							record
						};
						this.changeStoreData("complementSendModalData", complementSendModalData);
					} else {
						message.error("订单信息获取失败");
					}
				})
				.catch(err => message.error(err.message));
		}
	};
	// 确认 补发
	@action
	confirmComplementSend = (param, cb?: (result) => void) => {
		// 调用后端接口  确认补发
		AftersaleRedelivery(param)
			.then(resq => {
				if (resq.result.result) {
					this.changeStoreData("complementSendModalData", initComplementSendModalData);
					if (cb) {
						cb(resq.info);
					}
				} else {
					message.error(resq.result.msg);
				}
			})
			.catch(err => message.error(err.message));
	};

	// 点击 强制取消包裹
	@action
	clickCancelParcel = record => {
		if (record.isLocked) {
			message.error("配送单已被锁定，无法取消包裹");
			return;
		}
		// 调用接口获取 配送单下的   包裹列表
		GetDeliveryDetail({ deliveryId: record.deliveryId })
			.then(resq => {
				if (resq.result.code === "0") {
					const cancelParcelModalData = {
						visible: true,
						record,
						parcelList: resq.packages || [],
						orders: resq.orders || []
					};
					this.changeStoreData("cancelParcelModalData", cancelParcelModalData);
				} else {
					message.error(resq.result.message);
				}
			})
			.catch(err => message.error(err.message));
	};
	// 确认 强制取消包裹
	@action
	confirmCancelParcel = (cancelData: CancelData, cb?: (result) => void) => {
		const { deliveryId, parcelNumList } = cancelData;
		// 调用后端接口  确认取消包裹
		CancelDeliveryParcel({ shipmentId: deliveryId, parcelCodes: parcelNumList })
			.then(resq => {
				if (String(resq.result.code) === "0") {
					// this.changeStoreData("cancelParcelModalData", initCancelParcelModalData);
					if (cb) {
						cb(resq);
					}
				} else {
					message.error(resq.result.message);
				}
			})
			.catch(err => err.message(err.message));
	};

	// 点击 拆分
	@action
	clickSplitDelivery = record => {
		// 调用接口获取 配送单下的   包裹列表
		GetDeliveryDetail({ deliveryId: record.deliveryId || record.shipmentId })
			.then(resq => {
				if (resq.result.code === "0") {
					if (resq.orders.length === 0 && resq.packages.length === 1) {
						message.error("订单完整且只包含一个包裹，不可拆分");
						return;
					}
					const splitDeliveryModalData = {
						visible: true,
						record,
						parcelList: resq.packages || [],
						orders: resq.orders || []
					};
					this.changeStoreData("splitDeliveryModalData", splitDeliveryModalData);
				} else {
					message.error(resq.result.message);
				}
			})
			.catch(err => message.error(err.message));
	};
	// 确认 拆分
	@action
	confirmSplitDelivery = (splitData: CancelData, cb?: (result) => void) => {
		// 调用接口拆分配送单
		const { deliveryId, parcelNumList } = splitData;
		SplitShipmentParcel({ shipmentId: deliveryId, parcelCodes: parcelNumList })
			.then(resq => {
				if (String(resq.result.code) === "0") {
					// this.changeStoreData("splitDeliveryModalData", initSplitDeliveryModalData);
					if (cb) {
						cb(resq);
					}
				} else {
					message.error(resq.result.message ? resq.result.message : "拆分失败!");
				}
			})
			.catch(err => {
				err.message(err.message ? err.message : "拆分失败!");
			});
	};

	// 点击 修改配送信息
	@action
	clickEditDeliveryInfo = record => {
		// 包裹状态待配送，且未锁定
		if (record.isLocked) {
			message.error("配送单已被锁定，无法修改配送信息");
			return;
		}

		const editDeliveryInfoModalData = {
			visible: true,
			record
		};
		this.changeStoreData("editDeliveryInfoModalData", editDeliveryInfoModalData);
	};

	// 确认 修改配送信息
	@action
	confirmEditDeliveryInfo = (cb?: () => void) => {
		// 调用接口修改配送信息
		this.changeStoreData("editDeliveryInfoModalData", initEditDeliveryInfoModalData);
		if (cb) {
			cb();
		}
	};

	// 点击 打印标签
	@action
	clickPrintLabel = record => {
		// 调用接口获取 配送单下的   包裹列表
		// 获取面单信息
		GetPackageFaceInfo({ packageCode: record.packageCode })
			.then(resq => {
				const printLabelModalData = {
					visible: true,
					record,
					packageFaceInfo: resq
				};
				this.changeStoreData("printLabelModalData", printLabelModalData);
			})
			.catch(err => message.error(err.message));

		const printLabelModalData = {
			visible: true,
			record,
			packageFaceInfo: {
				packageCode: record.packageCode,
				nickname: "客户昵称",
				deliveryCode: "1234",
				deliveryMethod: "3333"
			}
		};
		this.changeStoreData("printLabelModalData", printLabelModalData);
	};

	// 点击 包裹备注 、 配送备注
	@action
	clickPODRemark = async (type: string, selectedRowKeys, selectedRows) => {
		if (selectedRowKeys && selectedRowKeys.length === 0) {
			message.error("请先选择包裹");
			return;
		}
		const isSingle = selectedRowKeys && selectedRowKeys.length === 1;
		let remarkLogs = [];
		let parcelList = [];
		if (isSingle) {
			try {
				if (type === "1") {
					// 查询包裹备注记录列表
					const packageRemarkResq = await GetPackageRemark({
						packageCode: selectedRowKeys
					});
					if (packageRemarkResq.result && packageRemarkResq.result.code === "0") {
						remarkLogs = packageRemarkResq.remarks || [];
					} else {
						message.error(packageRemarkResq.result.message);
					}
				}

				if (type === "2") {
					// 查询配送备注记录列表
					const deliveryId = selectedRows[0].deliveryId || selectedRows[0].shipmentId;
					const deliveryRemarkResq = await GetDeliveryRemark({
						deliveryId,
						packageCode: selectedRowKeys
					});
					if (deliveryRemarkResq.result && deliveryRemarkResq.result.code === "0") {
						remarkLogs = deliveryRemarkResq.remarks || [];
						parcelList = deliveryRemarkResq.packageCode || [];
					} else {
						message.error(deliveryRemarkResq.result.message);
					}
				}
			} catch (err) {
				remarkLogs = [];
				parcelList = [];
				console.log(err);
			}
		}
		const pODRemarksModalData = {
			visible: true,
			type,
			remarkLogs,
			parcelList,
			selectedRows,
			selectedRowKeys
		};
		this.changeStoreData("pODRemarksModalData", pODRemarksModalData);
	};

	// 确认  添加备注
	@action
	confirmAddRemark = (remark: string, refresh?: () => void) => {
		// 调用接口添加备注
		const { type, selectedRowKeys, selectedRows } = this.pODRemarksModalData;

		if (type === "1") {
			// 添加  包裹备注
			AddPackageRemark({ remark, packageCode: selectedRowKeys })
				.then(resq => {
					if (resq.result.code === "0") {
						this.changeStoreData("pODRemarksModalData", initPODRemarksModalData);
						if (refresh) {
							refresh();
						}
					} else {
						message.error(resq.result.message);
					}
				})
				.catch(err => message.error(err.message));
		}

		if (type === "2") {
			const deliverys = selectedRows.map(elem => {
				return {
					deliveryId: elem.deliveryId || elem.shipmentId,
					packageCode: elem.packageCode
				};
			});
			// 添加 派送备注
			AddDeliveryRemark({ remark, deliverys })
				.then(resq => {
					if (resq.result.code === "0") {
						this.changeStoreData("pODRemarksModalData", initPODRemarksModalData);
						if (refresh) {
							refresh();
						}
					} else {
						message.error(resq.result.message);
					}
				})
				.catch(err => message.error(err.message));
		}
	};

	// 点击 延误通知
	@action
	clickDelayNotification = async (selectedRowKeys, selectedRows) => {
		if (selectedRowKeys && selectedRowKeys.length === 0) {
			message.error("请先选择包裹");
			return;
		}
		const delayNotificationModalData = {
			visible: true,
			record: {},
			selectedRows,
			selectedRowKeys
		};
		this.changeStoreData("delayNotificationModalData", delayNotificationModalData);
	};

	// 确认  添加通知
	@action
	confirmAddNotice = (refresh?: () => void) => {
		// 调用接口修改配送信息
		this.changeStoreData("delayNotificationModalData", initDelayNotificationModalData);
		if (refresh) {
			refresh();
		}
	};
}

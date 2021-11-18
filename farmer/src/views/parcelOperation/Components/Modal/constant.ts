import { DeliveryNumModalProps } from "./DeliveryNumModal";
import { CancelParcelModalProps } from "./CancelParcelModal";
import { SplitDeliveryModalProps } from "./SplitDeliveryModal";
import { EditDeliveryInfoModalProps } from "./EditDeliveryInfoModal";
import { ComplementSendModalProps } from "./ComplementSendModal";
import { PrintLabelModalProps } from "./PrintLabelModal";
import { DelayNotificationModalProps } from "./DelayNotificationModal";
import { PODRemarksModalProps } from "./PODRemarksModal";

// 点击派送单号弹窗
export const initDeliveryNumModalData: DeliveryNumModalProps = {
	visible: false,
	deliveryId: "",
	parcelList: [],
	notOutboundOrderList: []
};

// 取消包裹
export const initCancelParcelModalData: CancelParcelModalProps = {
	visible: false,
	record: {},
	parcelList: [],
	orders: []
};

// 拆分派送信息
export const initSplitDeliveryModalData: SplitDeliveryModalProps = {
	visible: false,
	record: {},
	parcelList: [],
	orders: []
};

// 编辑派送信息
export const initEditDeliveryInfoModalData: EditDeliveryInfoModalProps = {
	visible: false,
	record: {}
};

// 补发
export const initComplementSendModalData: ComplementSendModalProps = {
	visible: false,
	record: {}
};

// 打印标签信息
export const initPrintLabelModalData: PrintLabelModalProps = {
	visible: false,
	record: {},
	packageFaceInfo: {}
};

// 延误通知
export const initDelayNotificationModalData: DelayNotificationModalProps = {
	visible: false,
	selectedRows: [],
	selectedRowKeys: []
};

// 包裹备注信息  派送备注信息
export const initPODRemarksModalData: PODRemarksModalProps = {
	visible: false,
	type: "",
	remarkLogs: [],
	parcelList: [],
	selectedRows: [],
	selectedRowKeys: []
};

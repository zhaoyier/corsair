import * as React from "react";
import {
	initDeliveryNumModalData,
	initComplementSendModalData,
	initCancelParcelModalData,
	initSplitDeliveryModalData,
	initEditDeliveryInfoModalData,
	initPrintLabelModalData,
	initPODRemarksModalData,
	initDelayNotificationModalData
} from "./constant";
import ModaltStore from "views/parcelOperation/store/modal";
import DeliveryNumModal from "./DeliveryNumModal";
import ComplementSendModal from "./ComplementSendModal";
import CancelParcelModal from "./CancelParcelModal";
import SplitDeliveryModal from "./SplitDeliveryModal";
import EditDeliveryInfoModal from "./EditDeliveryInfoModal";
import PrintLabelModal from "./PrintLabelModal";
import PODRemarksModal from "./PODRemarksModal";
import DelayNotificationModal from "./DelayNotificationModal";

{
	/* 点击派送单号弹窗 */
}
export const getDeliveryNumModal = (store: ModaltStore) => {
	const { deliveryNumModalData, changeStoreData } = store;
	const { visible, deliveryId, parcelList, notOutboundOrderList } = deliveryNumModalData;
	return (
		<DeliveryNumModal
			visible={visible}
			deliveryId={deliveryId}
			parcelList={parcelList}
			notOutboundOrderList={notOutboundOrderList}
			onCancel={() => changeStoreData("deliveryNumModalData", initDeliveryNumModalData)}
		/>
	);
};

{
	/* 补送 */
}
export const getComplementSendModal = (store: ModaltStore, refresh?: () => void) => {
	const { complementSendModalData, changeStoreData, confirmComplementSend } = store;
	const { visible, record } = complementSendModalData;
	return (
		<ComplementSendModal
			visible={visible}
			record={record}
			refresh={refresh}
			onCancel={() => changeStoreData("complementSendModalData", initComplementSendModalData)}
			onOk={(record, cb) => confirmComplementSend(record, result => cb(result))}
		/>
	);
};

{
	/* 取消配送 */
}
export const getCancelParcelModal = (store: ModaltStore, refresh?: () => void) => {
	const { cancelParcelModalData, changeStoreData, confirmCancelParcel } = store;
	const { visible, record, parcelList, orders } = cancelParcelModalData;
	return (
		<CancelParcelModal
			visible={visible}
			record={record}
			parcelList={parcelList}
			orders={orders}
			refresh={refresh}
			onCancel={() => changeStoreData("cancelParcelModalData", initCancelParcelModalData)}
			onOk={(cancelData, cb) => confirmCancelParcel(cancelData, result => cb(result))}
		/>
	);
};

{
	/* 拆分配送单 */
}
export const getSplitDeliveryModal = (store: ModaltStore, refresh?: () => void) => {
	const { splitDeliveryModalData, changeStoreData, confirmSplitDelivery } = store;
	const { visible, record, parcelList, orders } = splitDeliveryModalData;
	return (
		<SplitDeliveryModal
			visible={visible}
			record={record}
			parcelList={parcelList}
			orders={orders}
			refresh={refresh}
			onCancel={() => changeStoreData("splitDeliveryModalData", initSplitDeliveryModalData)}
			onOk={(splitData, cb) => confirmSplitDelivery(splitData, result => cb(result))}
		/>
	);
};

{
	/* 修改配送信息 */
}
export const getEditDeliveryInfoModal = (store: ModaltStore, refresh?: () => void) => {
	const { editDeliveryInfoModalData, changeStoreData, confirmEditDeliveryInfo } = store;
	const { visible, record } = editDeliveryInfoModalData;
	return (
		<div>
			{visible && (
				<EditDeliveryInfoModal
					visible={visible}
					record={record}
					onCancel={() =>
						changeStoreData("editDeliveryInfoModalData", initEditDeliveryInfoModalData)
					}
					onOk={() => confirmEditDeliveryInfo(refresh)}
				/>
			)}
		</div>
	);
};

{
	/* 打印标签 */
}
export const getPrintLabelModal = (store: ModaltStore) => {
	const { printLabelModalData, changeStoreData } = store;
	const { visible, record, packageFaceInfo } = printLabelModalData;
	return (
		<PrintLabelModal
			visible={visible}
			record={record}
			packageFaceInfo={packageFaceInfo}
			onCancel={() => changeStoreData("printLabelModalData", initPrintLabelModalData)}
		/>
	);
};

{
	/* 包裹备注，配送备注 */
}
export const getPODRemarksModal = (store: ModaltStore, refresh?: () => void) => {
	const { pODRemarksModalData, changeStoreData, confirmAddRemark } = store;
	const {
		visible,
		type,
		selectedRows,
		selectedRowKeys,
		remarkLogs,
		parcelList
	} = pODRemarksModalData;
	return (
		<PODRemarksModal
			visible={visible}
			type={type}
			remarkLogs={remarkLogs}
			parcelList={parcelList}
			selectedRows={selectedRows}
			selectedRowKeys={selectedRowKeys}
			onCancel={() => changeStoreData("pODRemarksModalData", initPODRemarksModalData)}
			onOk={remark => confirmAddRemark(remark, refresh)}
		/>
	);
};

{
	/* 延误通知 */
}
export const getDelayNotificationModal = (store: ModaltStore, refresh?: () => void) => {
	const { delayNotificationModalData, changeStoreData, confirmAddNotice } = store;
	const { visible, selectedRows, selectedRowKeys } = delayNotificationModalData;
	return (
		<div>
			{visible && (
				<DelayNotificationModal
					visible={visible}
					selectedRows={selectedRows}
					selectedRowKeys={selectedRowKeys}
					onCancel={() =>
						changeStoreData(
							"delayNotificationModalData",
							initDelayNotificationModalData
						)
					}
					onOk={() => confirmAddNotice(refresh)}
				/>
			)}
		</div>
	);
};

import * as React from "react";
import { observer } from "mobx-react-lite";
import CommonLayout from "components/CommonLayout";
import { Header } from "components/CommonHeader";
import { detailHeader } from "views/parcelOperation/store/constant";
import { useStores } from "views/parcelOperation/hooks";
import {
	getComplementSendModal,
	getSplitDeliveryModal,
	getPrintLabelModal,
	getPODRemarksModal,
	getDelayNotificationModal
} from "../Modal";
import DetailHeader from "./DetailHeader";
import DetailTab from "./DetailTab";

const ParcelDetail = () => {
	const { parcelDetailStore, modalStore } = useStores();
	const { getDetail, changeStoreData } = parcelDetailStore;

	React.useEffect(() => {
		getDetail();
		changeStoreData("activeKey", "1");
	}, []);

	return (
		<CommonLayout
			header={<Header {...detailHeader(() => changeStoreData("activeKey", "1"))} />}>
			<DetailHeader />
			<DetailTab />
			{/* 补送 */}
			{getComplementSendModal(modalStore, () => getDetail())}
			{/* 拆分配送单 */}
			{getSplitDeliveryModal(modalStore, () => getDetail())}
			{/* 打印标签 */}
			{getPrintLabelModal(modalStore)}
			{/* 包裹备注，配送备注 */}
			{getPODRemarksModal(modalStore, () => getDetail())}
			{/* 延误通知 */}
			{getDelayNotificationModal(modalStore, () => getDetail())}
		</CommonLayout>
	);
};

export default observer(ParcelDetail);

import * as React from "react";
import { observer } from "mobx-react-lite";
import { listHeader } from "views/parcelOperation/store/constant";
import { useStores } from "views/parcelOperation/hooks";
import {
	getDeliveryNumModal,
	getComplementSendModal,
	getCancelParcelModal,
	getSplitDeliveryModal,
	getEditDeliveryInfoModal,
	getPrintLabelModal,
	getPODRemarksModal,
	getDelayNotificationModal
} from "../Modal";
import CommonLayout from "components/CommonLayout";
import { Header } from "components/CommonHeader";
import SearchForm from "./SearchForm";
import ListTab from "./ListTab";
import TabHeaderBtn from "./TabHeaderBtn";

const ParcelList = () => {
	const { parcelListStore, modalStore } = useStores();
	const {
		getOrderOriginList,
		getParcelStatusList,
		getDestCountryList,
		getWarehouseList,
		getTransportMethodList,
		getPickupMethodList,
		pageParam,
		searchParcelList
	} = parcelListStore;
	const { offset, limit } = pageParam;

	// 初始化调用的方法
	React.useEffect(() => {
		getOrderOriginList(); // 获取订单来源列表
		getParcelStatusList(); // 获取包裹状态列表
		getDestCountryList(); // 获取目的地国家列表
		getWarehouseList(); // 获取所属仓库列表
		getTransportMethodList(); // 获取运输方式列表
		getPickupMethodList(); // 获取取货方式列表
	}, []);

	return (
		<CommonLayout style={{ padding: 0 }} header={<Header {...listHeader} />} showBar={false}>
			<SearchForm />
			<CommonLayout
				style={{ margin: 0, padding: 0 }}
				header={<TabHeaderBtn />}
				showBar={false}>
				<ListTab />
			</CommonLayout>
			{/* 点击派送单号弹窗 */}
			{getDeliveryNumModal(modalStore)}
			{/* 补送 */}
			{getComplementSendModal(modalStore, () => searchParcelList(offset, limit))}
			{/* 取消配送 */}
			{getCancelParcelModal(modalStore, () => searchParcelList(offset, limit))}
			{/* 拆分配送单 */}
			{getSplitDeliveryModal(modalStore, () => searchParcelList(offset, limit))}
			{/* 修改配送信息 */}
			{getEditDeliveryInfoModal(modalStore, () => searchParcelList(offset, limit))}
			{/* 打印标签 */}
			{getPrintLabelModal(modalStore)}
			{/* 包裹备注，配送备注 */}
			{getPODRemarksModal(modalStore, () => searchParcelList(offset, limit))}
			{/* 延误通知 */}
			{getDelayNotificationModal(modalStore, () => searchParcelList(offset, limit))}
		</CommonLayout>
	);
};
export default observer(ParcelList);

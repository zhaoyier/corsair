import * as React from "react";
import { observer } from "mobx-react-lite";
import Loadable from "react-loadable";
import { Tabs } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { MyLoadingComponent } from "../../Common/Loading";

const { TabPane } = Tabs;
const styles = require("./index.scss");

const BaseInfos = Loadable({
	loader: () => import("../BaseInfos"),
	loading: MyLoadingComponent
});

const ParcelRemarks = Loadable({
	loader: () => import("../../Common/ParcelRemarks"),
	loading: MyLoadingComponent
});

const ParcelTrack = Loadable({
	loader: () => import("../ParcelTrack"),
	loading: MyLoadingComponent
});

const ParcelLogList = Loadable({
	loader: () => import("../../Common/ParcelLogList"),
	loading: MyLoadingComponent
});

const getTab = (title: string) => {
	return <div className={styles.title}>{title}</div>;
};

const DetailTab = () => {
	const { parcelDetailStore } = useStores();
	const {
		activeKey,
		changeStoreData,
		parcelLog,
		parcelRemarks,
		getDetail,
		getPackageLog,
		getTotalRemarks,
		getPackageTrail
	} = parcelDetailStore;

	const changeTab = (activeKey: string) => {
		changeStoreData("activeKey", activeKey);
		if (activeKey === "1") {
			getDetail();
		}
		if (activeKey === "2") {
			getPackageLog();
		}
		if (activeKey === "3") {
			getTotalRemarks();
		}
		if (activeKey === "4") {
			getPackageTrail();
		}
	};

	return (
		<Tabs
			centered
			activeKey={activeKey}
			onChange={activeKey => changeTab(activeKey)}
			tabBarGutter={100}>
			<TabPane tab={getTab("基础信息")} key="1">
				<div className={styles.content}>
					<BaseInfos />
				</div>
			</TabPane>
			<TabPane tab={getTab("包裹日志")} key="2">
				<ParcelLogList dataSource={parcelLog} />
			</TabPane>
			<TabPane tab={getTab("包裹备注")} key="3">
				<ParcelRemarks dataSource={parcelRemarks} />
			</TabPane>
			<TabPane tab={getTab("包裹轨迹")} key="4">
				<ParcelTrack />
			</TabPane>
		</Tabs>
	);
};

export default observer(DetailTab);

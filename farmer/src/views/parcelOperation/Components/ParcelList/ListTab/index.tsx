import * as React from "react";
import { Tabs } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { TabStatus } from "views/parcelOperation/store/constant";
import { observer } from "mobx-react-lite";
import Loadable from "react-loadable";
import { MyLoadingComponent } from "../../Common/Loading";

const { TabPane } = Tabs;

const TabTable = Loadable({
	loader: () => import("../TabTable"),
	loading: MyLoadingComponent
});

const ListTab = () => {
	const { parcelListStore } = useStores();
	const { searchform, changeStoreData, changePage } = parcelListStore;

	return (
		<Tabs
			size="small"
			tabBarStyle={{ marginBottom: 10 }}
			activeKey={searchform.searchState}
			onChange={activeKey =>
				changeStoreData("searchform", activeKey, "searchState", null, () => changePage())
			}>
			{Object.keys(TabStatus).map(item => (
				<TabPane tab={TabStatus[item]} key={item}>
					<TabTable />
				</TabPane>
			))}
		</Tabs>
	);
};

export default observer(ListTab);

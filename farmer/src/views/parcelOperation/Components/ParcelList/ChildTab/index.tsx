import * as React from "react";
import { Tabs } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { observer } from "mobx-react-lite";
import { ChildTabStatus } from "views/parcelOperation/store/constant";
import { toJS } from "mobx";
import Loadable from "react-loadable";
import { MyLoadingComponent } from "../../Common/Loading";

const { TabPane } = Tabs;

const ParcelLogList = Loadable({
	loader: () => import("../../Common/ParcelLogList"),
	loading: MyLoadingComponent
});

const ParcelRemarks = Loadable({
	loader: () => import("../../Common/ParcelRemarks"),
	loading: MyLoadingComponent
});

interface ChildTabProps {
	index: number;
}

const ChildTab = (props: ChildTabProps) => {
	const { parcelListStore } = useStores();
	const { changeParcelTable, parcelList } = parcelListStore;
	const keys = Object.keys(ChildTabStatus);
	const tabs = Object.values(ChildTabStatus);
	console.log(toJS(parcelList[props.index]));
	console.log(props.index, toJS(parcelList[props.index].parcelLogList));
	return (
		<Tabs
			type="card"
			activeKey={parcelList[props.index].childTabActiveKey}
			onChange={activeKey => changeParcelTable(props.index, "childTabActiveKey", activeKey)}>
			<TabPane tab={tabs[0]} key={keys[0]}>
				<ParcelLogList dataSource={toJS(parcelList[props.index].parcelLogList)} />
			</TabPane>
			<TabPane tab={tabs[1]} key={keys[1]}>
				<ParcelRemarks dataSource={toJS(parcelList[props.index].remarkList)} />
			</TabPane>
		</Tabs>
	);
};

export default observer(ChildTab);

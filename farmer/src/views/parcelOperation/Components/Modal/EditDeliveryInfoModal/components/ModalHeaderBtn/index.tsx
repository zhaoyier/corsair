import * as React from "react";
import { observer } from "mobx-react-lite";
import { Button } from "antd";
import { useModalStores } from "../../store/hooks";
import { DeliveryTypes, allPickTime } from "../../store/constant";

const styles = require("./index.scss");

const ModalHeaderBtn = props => {
	const store = useModalStores();
	const { changeField } = props;
	const {
		deliveryTypeId,
		changeStoreData,
		onRegionChange,
		getDeliveryInfo,
		neiStationList,
		stationList,
		pickUpDatesList,
		isForcedChange,
		periodsList,
		addressList
	} = store;

	const checkSelect = () => {
		// 区域
		if (neiStationList.length === 0) {
			changeStoreData("neighbourhoodStations", null);
			changeField({ neighbourhoodStations: null });
		}
		// 站点名称
		if (stationList.length === 0) {
			changeStoreData("stations", null);
			changeField({ stations: null });
		}
		// 送货日期
		if (pickUpDatesList.length === 0) {
			changeStoreData("pickUpDates", []);
			changeField({ pickUpDates: [] });
		}
		// 送货时段
		if (isForcedChange ? allPickTime.length === 0 : periodsList.length === 0) {
			changeStoreData("PickupPeriods", []);
			changeField({ PickupPeriods: [] });
		}
		// 地址
		if (addressList.length === 0) {
			changeStoreData("addresses", null);
			changeField({ addresses: null });
		}
	};

	return (
		<section className={styles.tabHeaderBtns}>
			{Object.keys(DeliveryTypes).map(item => (
				<Button
					key={DeliveryTypes[item].value}
					type={DeliveryTypes[item].value === deliveryTypeId ? "primary" : "default"}
					className={styles.btn}
					onClick={() => {
						changeStoreData("deliveryTypeId", DeliveryTypes[item].value);
						if (DeliveryTypes[item].fields.includes("neighbourhoodStations")) {
							onRegionChange(DeliveryTypes[item].value);
						}
						getDeliveryInfo(DeliveryTypes[item].value);
						checkSelect();
					}}>
					{DeliveryTypes[item].label}
				</Button>
			))}
		</section>
	);
};

export default observer(ModalHeaderBtn);

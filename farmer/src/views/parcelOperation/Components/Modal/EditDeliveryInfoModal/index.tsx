import * as React from "react";
import { Modal, Form, Spin } from "antd";
import { observer } from "mobx-react-lite";
import { useModalStores } from "./store/hooks";
import { formatWeight } from "utils/util";
import ModalHeaderBtn from "./components/ModalHeaderBtn";
import ModalContent from "./components/ModalContent";
import ParcelTable from "../../Common/ParcelTable";
import { toJS } from "mobx";
import { TimeFormatToDayReverse, formatUnixTime } from "utils/time";

/**
 * EZSHIP  修改配送单信息  弹窗
 * https://www.tapd.cn/51528997/markdown_wikis/?#1151528997001057668
 */
export interface EditDeliveryInfoModalProps {
	visible: boolean;
	record: any; // 当前包裹信息
	onCancel?: () => void;
	onOk?: () => void;
}
// 修改配送信息  弹窗

const styles = require("./index.scss");

const EditDeliveryInfoModal = (props: EditDeliveryInfoModalProps) => {
	const { visible, record, onCancel, onOk } = props;
	const [modalVisible, setModalVisible] = React.useState(visible);

	const [form] = Form.useForm();
	const store = useModalStores();
	const { loading, disabled, info, getDefaultDeliveryInfo, updateDelivery } = store;
	const deliveryTypeInfo = { ...toJS(store.deliveryTypeInfo) };

	React.useEffect(() => {
		getDefaultDeliveryInfo(record);
	}, [modalVisible]);

	const getTitle = () => {
		return (
			<div className={styles.title}>
				修改配送号：{record.deliveryId}
				<span>
					总重量：
					{info && info.weight ? formatWeight(info.weight) : 0}
					kg
				</span>
			</div>
		);
	};
	React.useEffect(() => {
		resetForm();
	});

	const resetForm = () => {
		const { isForcedChange } = store;
		const {
			addressToName,
			addressToPhone,
			zipCode,
			street,
			block,
			companyName,
			buildingName
		} = deliveryTypeInfo;

		const pickUpDates = deliveryTypeInfo.pickUpDates || [];
		const PickupPeriods = deliveryTypeInfo.PickupPeriods || [];
		const address = deliveryTypeInfo.addresses || [];
		const neighbourhoodStations = deliveryTypeInfo.neighbourhoodStations || [];
		const stations = deliveryTypeInfo.stations || [];
		const neighbourhoodStationsValue =
			neighbourhoodStations && neighbourhoodStations.length > 0
				? neighbourhoodStations[0].stationId
				: null;
		const stationsValue = stations && stations.length > 0 ? stations[0].stationId : null;
		const pickUpDatesValue =
			pickUpDates.length > 0 ? formatUnixTime(pickUpDates[0], TimeFormatToDayReverse) : null;
		const PickupPeriodsValue = PickupPeriods.length > 0 ? PickupPeriods[0] : null;
		const addressValue =
			address && address.length > 0 && address[0].addressId !== "0"
				? address[0].addressId
				: null;
		const stationAddressValue =
			stations && stations.length > 0 ? stations[0].stationAddress : null;

		const values = {
			neighbourhoodStations: neighbourhoodStationsValue,
			stations: stationsValue,
			pickUpDates: pickUpDatesValue,
			PickupPeriods: PickupPeriodsValue,
			isForcedChange,
			addresses: addressValue,
			addressToName,
			addressToPhone,
			zipCode,
			street,
			block,
			companyName,
			buildingName,
			stationAddress: stationAddressValue
		};
		form.setFieldsValue(values);
	};

	const onFinish = values => {
		console.log("Received values of form: ", values);
		updateDelivery(() => {
			onOk();
		});
	};

	const title = getTitle();

	const changeField = values => {
		form.setFieldsValue(values);
	};

	return (
		<Modal
			visible={modalVisible}
			title={title}
			onCancel={() => {
				onCancel();
				setModalVisible(false);
			}}
			onOk={() => form.submit()}
			width={900}
			cancelText="关闭"
			okButtonProps={{
				disabled
			}}
			destroyOnClose={true}>
			<Spin spinning={loading}>
				<ModalHeaderBtn changeField={changeField} />
				<Form
					form={form}
					name="advanced_search"
					className={styles["ant-advanced-search-form"]}
					onFinish={onFinish}>
					<ModalContent changeField={changeField} />
				</Form>
				<ParcelTable dataSource={info && info.items ? info.items : []} />
			</Spin>
		</Modal>
	);
};

export default observer(EditDeliveryInfoModal);

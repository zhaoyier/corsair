import * as React from "react";
import { observer } from "mobx-react-lite";
import { Form, Row, Col, Input, Checkbox, Select } from "antd";
import { useModalStores } from "../../store/hooks";
import EditDeliveryModalStore from "../../store/store";
import { getDeliveryTypeById, allPickTime } from "../../store/constant";
import { formatUnixTime, TimeFormatToDayReverse } from "utils/time";
import { fromPairs } from "lodash";
import { toJS } from "mobx";

const styles = require("./index.scss");

interface FormItemProps {
	name: string;
	label?: string;
	tag?: string;
	children: any;
	visible?: boolean;
	rules?: any[];
	index?: number;
	span?: number;
}

const getFormItem = (props: FormItemProps) => {
	const Label = label => <div className={styles.labelWidth}>{label}</div>;
	const long = props.name === "addresses" || props.name === "stationAddress";
	return (
		<Col span={props.span ? props.span : 12} key={props.index}>
			{props.label ? (
				<Form.Item
					name={props.name}
					label={Label(props.label)}
					rules={props.rules}
					labelCol={{ xs: { span: 24 }, sm: { span: long ? 4 : 8 } }}
					wrapperCol={{ xs: { span: 24 }, sm: { span: long ? 20 : 16 } }}>
					{props.children}
				</Form.Item>
			) : (
				<Form.Item name={props.name} noStyle>
					{props.children}
				</Form.Item>
			)}
		</Col>
	);
};

const FloorUnit = ({ store }) => {
	const { deliveryTypeInfo, changeStoreData, disabled } = store;
	return (
		<div>
			<Input
				placeholder="请输入楼层"
				disabled={disabled}
				value={deliveryTypeInfo.unitStart}
				style={{ width: "110px" }}
				onChange={e => changeStoreData("deliveryTypeInfo", e.target.value, "unitStart")}
			/>
			&nbsp;-&nbsp;
			<Input
				placeholder="请输入单元"
				disabled={disabled}
				value={deliveryTypeInfo.unitEnd}
				style={{ width: "110px" }}
				onChange={e => changeStoreData("deliveryTypeInfo", e.target.value, "unitEnd")}
			/>
		</div>
	);
};

const getVisible = (name, store: EditDeliveryModalStore) => {
	const { deliveryTypeId } = store;
	const arr = getDeliveryTypeById(deliveryTypeId).fields;
	const find = arr.find(elem => elem === name);
	return Boolean(find);
};

const getPickUpPeriodSelectList = (deliveryTypeId, isForcedChange, periodsList) => {
	const item = getDeliveryTypeById(deliveryTypeId).result;
	const isHome = item === "infoHome";
	if (isForcedChange && isHome) {
		return allPickTime.map(item => ({
			label: item.periodName,
			value: item.periodName
		}));
	} else {
		if (periodsList && periodsList.length > 0) {
			return periodsList.map(item => ({
				label: item.periodName,
				value: item.periodName
			}));
		}
		return [];
	}
};

const getFields = (store: EditDeliveryModalStore, changeField: (values) => void) => {
	const {
		disabled,
		deliveryTypeId,
		changeStoreData,
		neiStationList,
		stationList,
		pickUpDatesList,
		isForcedChange,
		periodsList,
		addressList,
		onRegionChange,
		onStationChange,
		onPickUpDateChange
	} = store;
	const deliveryTypeInfo = { ...toJS(store.deliveryTypeInfo) };
	const pickUpDates = deliveryTypeInfo.pickUpDates || [];
	const PickupPeriods = deliveryTypeInfo.PickupPeriods || [];
	const address = deliveryTypeInfo.addresses || [];
	const neighbourhoodStations = deliveryTypeInfo.neighbourhoodStations || [];
	const stations = deliveryTypeInfo.stations || [];
	const neiStationSelectList = neiStationList.map(item => ({
		label: item.stationName,
		value: item.stationId
	}));
	const pickUpDateSelectList = (pickUpDatesList || []).map(item => ({
		label: formatUnixTime(item, TimeFormatToDayReverse),
		value: item
	}));
	const pickUpPeriodSelectList = getPickUpPeriodSelectList(
		deliveryTypeId,
		isForcedChange,
		periodsList
	);
	const chooseAddressList = addressList.map(item => ({
		label: item.addressName,
		value: item.addressId
	}));
	const stationSelectList = stationList.map(item => ({
		label: item.stationName,
		value: item.stationId
	}));

	const neighbourhoodStationsValue =
		neighbourhoodStations && neighbourhoodStations.length > 0
			? neighbourhoodStations[0].stationId
			: null;
	const stationsValue = stations && stations.length > 0 ? stations[0].stationId : null;
	const pickUpDatesValue =
		pickUpDates.length > 0 ? formatUnixTime(pickUpDates[0], TimeFormatToDayReverse) : null;
	const PickupPeriodsValue = PickupPeriods.length > 0 ? PickupPeriods[0] : null;
	const addressValue =
		address && address.length > 0 && address[0].addressId !== "0" ? address[0].addressId : null;
	const stationAddressValue = stations && stations.length > 0 ? stations[0].stationAddress : null;
	const fieldArr: FormItemProps[] = [
		// 区域
		{
			name: "neighbourhoodStations",
			label: "区域",
			visible: getVisible("neighbourhoodStations", store),
			rules: [{ required: true, message: "请输入区域" }],
			children: (
				<div>
					<Select
						style={{ width: "100%" }}
						disabled={disabled}
						value={neighbourhoodStationsValue}
						onChange={value => {
							const deliveryTypeInfoNew = { ...deliveryTypeInfo };
							changeStoreData("pickUpDatesList", []);
							changeStoreData("periodsList", []);
							const neighbourhoodStations = neiStationList.filter(
								item => item.stationId === value
							);
							deliveryTypeInfoNew.neighbourhoodStations = neighbourhoodStations;
							changeStoreData("deliveryTypeInfo", deliveryTypeInfoNew);
							changeField({ pickUpDatesList: [], periodsList: [] });
							changeField({ neighbourhoodStations });
							onRegionChange(value);
						}}
						allowClear={true}
						filterOption={(input, option: any) =>
							option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
						}>
						{neiStationSelectList.map(o => (
							<Select.Option value={o.value} key={o.value} title={o.label}>
								{o.label}
							</Select.Option>
						))}
					</Select>
				</div>
			)
		},
		// 站点名称
		{
			name: "stations",
			label: "站点名称",
			visible: getVisible("stations", store),
			rules: [{ required: true, message: "请输入站点名称" }],
			children: (
				<div>
					<Select
						style={{ width: "230px" }}
						disabled={disabled}
						value={stationsValue}
						defaultValue={stationsValue}
						onChange={value => {
							const deliveryTypeInfoNew = { ...deliveryTypeInfo };
							const stations = stationList.filter(item => item.stationId === value);
							deliveryTypeInfoNew.stations = stations;
							deliveryTypeInfoNew.pickUpDates = [];
							deliveryTypeInfoNew.PickupPeriods = [];
							changeStoreData("pickUpDatesList", []);
							changeStoreData("periodsList", []);
							changeStoreData("deliveryTypeInfo", deliveryTypeInfoNew);
							changeField({ stations, pickUpDates: [], PickupPeriods: [] });
							// changeField({ pickUpDates: [], PickupPeriods: [] });
							onStationChange();
						}}
						allowClear={true}
						filterOption={(input, option: any) =>
							option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
						}>
						{stationSelectList.map(o => (
							<Select.Option value={o.value} key={o.value} title={o.label}>
								{o.label}
							</Select.Option>
						))}
					</Select>
				</div>
			)
		},
		// 送货日期
		{
			name: "pickUpDates",
			label: "送货日期",
			visible: getVisible("pickUpDates", store),
			children: (
				<div>
					<Select
						style={{ width: "230px" }}
						disabled={disabled}
						value={pickUpDatesValue}
						onChange={value => {
							changeStoreData("deliveryTypeInfo", new Array(value), "pickUpDates");
							changeStoreData("deliveryTypeInfo", [], "PickupPeriods");
							changeField({ pickUpDates: new Array(value), PickupPeriods: [] });
							onPickUpDateChange();
						}}
						allowClear={true}
						filterOption={(input, option: any) =>
							option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
						}>
						{pickUpDateSelectList.map(o => (
							<Select.Option value={o.value} key={o.value} title={o.label}>
								{o.label}
							</Select.Option>
						))}
					</Select>
				</div>
			)
		},
		// 送货时段
		{
			name: "PickupPeriods",
			label: "送货时段",
			visible: getVisible("PickupPeriods", store),
			children: (
				<div>
					<Select
						style={{ width: "230px" }}
						value={PickupPeriodsValue}
						defaultValue={PickupPeriodsValue}
						disabled={disabled}
						onChange={value => {
							changeStoreData("deliveryTypeInfo", new Array(value), "PickupPeriods");
							changeField({ PickupPeriods: new Array(value) });
						}}
						allowClear={true}
						filterOption={(input, option: any) =>
							option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
						}>
						{pickUpPeriodSelectList.map(o => (
							<Select.Option value={o.value} key={o.value} title={o.label}>
								{o.label}
							</Select.Option>
						))}
					</Select>
				</div>
			)
		},
		// 强制大小货
		{
			name: "isForcedChange",
			span: 24,
			visible: getVisible("isForcedChange", store),
			children: (
				<div>
					<Checkbox
						checked={isForcedChange}
						defaultChecked={isForcedChange}
						disabled={disabled}
						style={{ paddingLeft: 520, marginBottom: 15 }}
						onChange={e => {
							const isForcedChange = e.target.checked;
							if (!isForcedChange) {
								deliveryTypeInfo.PickupPeriods = [];
								changeField({ PickupPeriods: [] });
							}
							changeField({ isForcedChange });
							changeStoreData("isForcedChange", e.target.checked);
						}}>
						强制大小货
					</Checkbox>
				</div>
			)
		},
		// 地址
		{
			name: "addresses",
			label: "地址",
			span: 24,
			visible: getVisible("addresses", store),
			// rules: [{ required: true, message: "请选择地址" }],
			children: (
				<div>
					<Select
						style={{ width: "640px" }}
						disabled={disabled}
						defaultValue={addressValue}
						value={addressValue}
						onChange={value => {
							const deliveryTypeInfoNew = { ...toJS(deliveryTypeInfo) };
							const add = toJS(addressList.find(item => item.addressId === value));
							const addresses = add ? new Array(add) : [];
							deliveryTypeInfoNew.addresses = addresses;
							const homeAddInfo = add
								? fromPairs(
										add.addressName.split(",").map(x => x.trim().split(":"))
										// tslint:disable-next-line:indent
								  )
								: {};
							fieldArr.forEach(item => {
								if (item.tag && Object.keys(homeAddInfo).includes(item.tag)) {
									if (item.tag === "单元") {
										const unit = homeAddInfo[item.tag].split("-");
										deliveryTypeInfoNew.unitStart = unit[0].replace(/#/, "");
										deliveryTypeInfoNew.unitEnd = unit[1];
										changeField({
											unitStart: deliveryTypeInfoNew.unitStart,
											unitEnd: deliveryTypeInfoNew.unitEnd
										});
									} else {
										deliveryTypeInfoNew[item.name] = homeAddInfo[item.tag];
										changeField({ [item.name]: homeAddInfo[item.tag] } as any);
									}
								}
							});
							changeField({ addresses });
							changeStoreData("deliveryTypeInfo", deliveryTypeInfoNew);
						}}
						allowClear={true}
						filterOption={(input, option: any) =>
							option.props.children.toLowerCase().indexOf(input.toLowerCase()) >= 0
						}>
						{chooseAddressList.map(o => (
							<Select.Option value={o.value} key={o.value} title={o.label}>
								{o.label}
							</Select.Option>
						))}
					</Select>
				</div>
			)
		},
		// 请输入收件人姓名
		{
			name: "addressToName",
			label: "收件人姓名",
			tag: "收货人",
			visible: getVisible("addressToName", store),
			rules: [{ required: true, message: "请输入收件人姓名" }],
			children: (
				<div>
					<Input
						placeholder="请输入收件人姓名"
						style={{ width: "230px" }}
						disabled={disabled}
						defaultValue={deliveryTypeInfo.addressToName}
						value={deliveryTypeInfo.addressToName}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "addressToName");
							changeField({ addressToName: e.target.value });
						}}
					/>
				</div>
			)
		},
		// 收件人电话
		{
			name: "addressToPhone",
			label: "收件人电话",
			tag: "电话",
			visible: getVisible("addressToPhone", store),
			rules: [{ required: true, message: "请输入收件人电话" }],
			children: (
				<div>
					<Input
						placeholder="请输入收件人电话"
						style={{ width: "230px" }}
						disabled={disabled}
						defaultValue={deliveryTypeInfo.addressToPhone}
						value={deliveryTypeInfo.addressToPhone}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "addressToPhone");
							changeField({ addressToPhone: e.target.value });
						}}
					/>
				</div>
			)
		},
		// 邮编
		{
			name: "zipCode",
			label: "邮编",
			tag: "邮编",
			visible: getVisible("zipCode", store),
			rules: [{ required: true, message: "请输入邮编" }],
			children: (
				<div>
					<Input
						placeholder="请输入邮编"
						style={{ width: "230px" }}
						disabled={disabled}
						defaultValue={deliveryTypeInfo.zipCode}
						value={deliveryTypeInfo.zipCode}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "zipCode");
							changeField({ zipCode: e.target.value });
						}}
					/>
				</div>
			)
		},
		// 街道
		{
			name: "street",
			label: "街道",
			tag: "街道",
			visible: getVisible("street", store),
			rules: [{ required: true, message: "请输入街道" }],
			children: (
				<div>
					<Input
						placeholder="请输入街道"
						style={{ width: "230px" }}
						disabled={disabled}
						defaultValue={deliveryTypeInfo.street}
						value={deliveryTypeInfo.street}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "street");
							changeField({ street: e.target.value });
						}}
					/>
				</div>
			)
		},
		// 大牌
		{
			name: "block",
			label: "大牌",
			tag: "大牌",
			visible: getVisible("block", store),
			children: (
				<div>
					<Input
						placeholder="请输入大牌"
						style={{ width: "230px" }}
						disabled={disabled}
						defaultValue={deliveryTypeInfo.block}
						value={deliveryTypeInfo.block}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "block");
							changeField({ block: e.target.value });
						}}
					/>
				</div>
			)
		},
		// 楼层-单元
		{
			name: "unitStart",
			label: "楼层-单元",
			tag: "单元",
			visible: getVisible("unitStart", store),
			children: <FloorUnit store={store} />
		},
		// 公司名
		{
			name: "companyName",
			label: "公司名",
			tag: "公司",
			visible: getVisible("companyName", store),
			children: (
				<div>
					<Input
						placeholder="请输入公司名"
						style={{ width: "230px" }}
						disabled={disabled}
						defaultValue={deliveryTypeInfo.companyName}
						value={deliveryTypeInfo.companyName}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "companyName");
							changeField({ companyName: e.target.value });
						}}
					/>
				</div>
			)
		},
		// 建筑名称
		{
			name: "buildingName",
			label: "建筑名称",
			tag: "建筑物",
			visible: getVisible("buildingName", store),
			children: (
				<div>
					<Input
						placeholder="请输入建筑名称"
						style={{ width: "230px" }}
						disabled={disabled}
						defaultValue={deliveryTypeInfo.buildingName}
						value={deliveryTypeInfo.buildingName}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "buildingName");
							changeField({ buildingName: e.target.value });
						}}
					/>
				</div>
			)
		},
		// 站点地址
		{
			name: "stationAddress",
			label: "站点地址",
			visible: getVisible("stationAddress", store),
			span: 24,
			children: (
				<div>
					<Input.TextArea
						placeholder="请输入站点地址"
						defaultValue={stationAddressValue}
						value={stationAddressValue}
						disabled
						style={{ width: "640px" }}
						onChange={e => {
							changeStoreData("deliveryTypeInfo", e.target.value, "stationAddress");
							changeField({ stationAddress: e.target.value });
						}}
					/>
				</div>
			)
		}
	];
	const children = fieldArr.map((elem, index) =>
		elem.visible
			? getFormItem({
					name: elem.name,
					label: elem.label,
					children: elem.children,
					index,
					rules: elem.rules,
					span: elem.span
					// tslint:disable-next-line:indent
			  })
			: null
	);
	return children;
};
const ModalContent = props => {
	const store = useModalStores();

	return <Row gutter={24}>{getFields(store, props.changeField)}</Row>;
};

// tslint:disable-next-line:max-file-line-count
export default observer(ModalContent);

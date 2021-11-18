import * as React from "react";
import { Form, Button, Input, DatePicker, Divider, Row, Col, FormItemProps } from "antd";
import { DownOutlined, UpOutlined } from "@ant-design/icons";
import { useStores } from "views/parcelOperation/hooks";
import { IsRemarkList, initSearchForm } from "views/parcelOperation/store/constant";
import { observer } from "mobx-react-lite";
import ArrSelect from "../../Common/ArrSelect";
import ObjectSelect from "../../Common/ObjectSelect";
import { toJS } from "mobx";
import { dataString } from "utils/time";
import { DateGap } from "genServices/ezShipOMS/oms";
interface FormItemPropsType extends FormItemProps {
	children: any;
	index?: number;
	span?: number;
}

const styles = require("./index.scss");

const getFormItem = (props: FormItemPropsType) => {
	const Label = label => <div className={styles.labelWidth}>{label}</div>;
	return (
		<Col span={props.span ? props.span : 6} key={props.index}>
			{props.label ? (
				<Form.Item rules={props.rules} name={props.name} label={Label(props.label)}>
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

const SearchForm = () => {
	const [expand, setExpand] = React.useState(false);
	const [form] = Form.useForm();
	const { parcelListStore } = useStores();
	const { changePage, changeStoreData, searchform, listLoading } = parcelListStore;
	const onFinish = values => {
		changeStoreData("searchform", DateGap.DateGapLatestThreeMonth, "dateGap", "", () => changePage());
		console.log("Received values of form: ", values);
	};

	React.useEffect(() => {
		changeStoreData("searchform", searchform);
	}, []);

	const clearOtherParam = (
		param: string,
		value,
		isOnly: boolean = false,
		isDate = false,
		cb?: () => void
	) => {
		let newSearchform = {} as any;
		if (isOnly) {
			newSearchform = { ...initSearchForm };
			newSearchform[param] = value;
			form.setFieldsValue(newSearchform);
		} else {
			newSearchform = { ...parcelListStore.searchform };
			if (isDate) {
				newSearchform[param] = value ? String(value.unix()) : null;
			} else {
				newSearchform[param] = value;
			}
		}
		changeStoreData("searchform", newSearchform, "", "", cb);
	};

	const getFields = () => {
		const {
			searchform,
			orderOriginList,
			parcelStatusList,
			desCodeList,
			warehouseList,
			transportMethodList,
			deliveryMethodList,
			getWarehouseList
		} = parcelListStore;
		const isPackageCode = !!searchform.packageCode;
		const isDeliveryId = !!searchform.deliveryId;
		const isTransportNumber = !!searchform.transportNumber;
		const comDisabled = isPackageCode || isDeliveryId || isTransportNumber;
		const fieldArr: FormItemPropsType[] = [
			{
				name: "origin",
				label: "包裹来源",
				children: (
					<ArrSelect
						names={orderOriginList}
						disabled={comDisabled}
						value={searchform.origin}
						placeholder="请选择包裹状态"
						onChange={value => clearOtherParam("origin", value)}
					/>
				)
			},
			// 包裹号 --- 排他， 不与其他筛选项组合
			{
				name: "packageCode",
				label: "包裹号",
				children: (
					<Input
						placeholder="请输入包裹号"
						value={searchform.packageCode}
						defaultValue={searchform.packageCode}
						disabled={isDeliveryId || isTransportNumber}
						onChange={e => clearOtherParam("packageCode", e.target.value, true)}
					/>
				)
			},
			// 配送单号 -- 排他， 不与其他筛选项组合
			{
				name: "deliveryId",
				label: "配送单号",
				children: (
					<Input
						placeholder="请输入配送单号"
						value={searchform.deliveryId}
						defaultValue={searchform.deliveryId}
						disabled={isPackageCode || isTransportNumber}
						onChange={e => clearOtherParam("deliveryId", e.target.value, true)}
					/>
				)
			},
			{
				name: "packageStatue",
				label: "包裹状态",
				children: (
					<ArrSelect
						names={parcelStatusList}
						disabled={comDisabled}
						value={searchform.packageStatue}
						placeholder="请选择包裹状态"
						onChange={value => clearOtherParam("packageStatue", value)}
					/>
				)
			},
			{
				name: "orderId",
				label: "订单号",
				children: (
					<Input
						placeholder="请输入订单号"
						value={searchform.orderId}
						defaultValue={searchform.orderId}
						disabled={comDisabled}
						onChange={e => clearOtherParam("orderId", e.target.value)}
					/>
				)
			},
			{
				name: "catalog",
				label: "目的地国家",
				children: (
					<ArrSelect
						names={desCodeList}
						notHasAll={true}
						disabled={comDisabled}
						value={toJS(searchform.catalog)}
						placeholder="请选择目的地国家"
						onChange={value =>
							clearOtherParam("catalog", value, false, false, () =>
								getWarehouseList()
							)
						}
					/>
				)
			},
			{
				name: "boxNumber",
				label: "封箱号",
				children: (
					<Input
						placeholder="请输入封箱号"
						value={searchform.boxNumber}
						defaultValue={searchform.boxNumber}
						disabled={comDisabled}
						onChange={e => clearOtherParam("boxNumber", e.target.value)}
					/>
				)
			},
			{
				name: "nickname",
				label: "会员名称",
				children: (
					<Input
						placeholder="请输入会员名称"
						value={searchform.nickname}
						defaultValue={searchform.nickname}
						disabled={comDisabled}
						onChange={e => clearOtherParam("nickname", e.target.value)}
					/>
				)
			},
			{
				name: "deliverDate",
				label: "送货日期",
				children: (
					<DatePicker
						allowClear
						value={searchform.deliverDate ? dataString(searchform.deliverDate) : null}
						defaultValue={
							searchform.deliverDate ? dataString(searchform.deliverDate) : null
						}
						style={{ width: "100%" }}
						disabled={comDisabled}
						onChange={date => clearOtherParam("deliverDate", date, false, true)}
					/>
				)
			},
			{
				name: "phone",
				label: "收货电话",
				children: (
					<Input
						placeholder="请输入收货电话"
						value={searchform.phone}
						defaultValue={searchform.phone}
						disabled={comDisabled}
						onChange={e => clearOtherParam("phone", e.target.value)}
					/>
				)
			},
			{
				name: "billId",
				label: "付款单号",
				rules: [{ pattern: /^[0-9]*$/, message: "请输入纯数字" }],
				children: (
					<Input
						placeholder="请输入付款单号"
						value={searchform.billId}
						defaultValue={searchform.billId}
						disabled={comDisabled}
						onChange={e => clearOtherParam("billId", e.target.value)}
					/>
				)
			},
			{
				name: "etaDate",
				label: "发货ETA",
				children: (
					<DatePicker
						allowClear
						style={{ width: "100%" }}
						disabled={comDisabled}
						value={searchform.etaDate ? dataString(searchform.etaDate) : null}
						defaultValue={searchform.etaDate ? dataString(searchform.etaDate) : null}
						onChange={date => clearOtherParam("etaDate", date, false, true)}
					/>
				)
			},
			// 物流单号  --排他， 不与其他筛选项组合
			{
				name: "transportNumber",
				label: "物流单号",
				children: (
					<Input
						placeholder="请输入物流单号"
						disabled={isPackageCode || isDeliveryId}
						value={searchform.transportNumber}
						defaultValue={searchform.transportNumber}
						onChange={e => clearOtherParam("transportNumber", e.target.value, true)}
					/>
				)
			},
			{
				name: "warehouseId",
				label: "所属仓库",
				children: (
					<ArrSelect
						names={warehouseList}
						placeholder="请选择所属仓库"
						disabled={comDisabled}
						value={searchform.warehouseId}
						onChange={value => clearOtherParam("warehouseId", value)}
					/>
				)
			},
			{
				name: "shipmentType",
				label: "运输方式",
				children: (
					<ArrSelect
						names={transportMethodList}
						placeholder="请选择运输方式"
						disabled={comDisabled}
						value={searchform.shipmentType}
						onChange={value => clearOtherParam("shipmentType", value)}
					/>
				)
			},
			{
				name: "deliveryType",
				label: "取货方式",
				children: (
					<ArrSelect
						names={deliveryMethodList}
						placeholder="请选择取货方式"
						disabled={comDisabled}
						value={searchform.deliveryType}
						onChange={value => clearOtherParam("deliveryType", value)}
					/>
				)
			},
			{
				name: "isRemark",
				label: "是否有备注",
				children: (
					<ObjectSelect
						names={IsRemarkList}
						disabled={comDisabled}
						value={searchform.isRemark}
						onChange={value => clearOtherParam("isRemark", value)}
					/>
				)
			},
			{
				name: "cabinetNumber",
				label: "主单/柜号",
				children: (
					<Input
						placeholder="请输入主单/柜号"
						disabled={comDisabled}
						value={searchform.cabinetNumber}
						defaultValue={searchform.cabinetNumber}
						onChange={e => clearOtherParam("cabinetNumber", e.target.value)}
					/>
				)
			}
		];
		const shortArr = fieldArr.slice(0, 12);
		const arr = expand ? fieldArr : shortArr;
		arr.push({
			span: 12,
			children: (<React.Fragment>
				<Button
					type="primary"
					htmlType="submit"
					className={styles.formBtn}
					disabled={listLoading}>
					查询
				</Button>
				<Button
					style={{ marginLeft: 8 }}
					onClick={() => {
						form.resetFields();
						changeStoreData("searchform", initSearchForm);
					}}>
					重置
				</Button>
				<Button style={{ margin: "0px 8px" }} type="primary" onClick={() => changeStoreData("searchform", DateGap.DateGapBeforeThreeMonth, "dateGap", "", () => changePage())} > 查询归档 </Button>
				<a
					style={{ marginLeft: 8, fontSize: 12 }}
					onClick={() => setExpand(!expand)}>
					&emsp;{expand ? "收起" : "展开"}&nbsp;
					{expand ? <UpOutlined /> : <DownOutlined />}
				</a>
			</React.Fragment>)
		});
		const children = arr.map((elem, index) =>
			getFormItem({
				name: elem.name,
				label: elem.label,
				children: elem.children,
				index,
				span: elem.span,
				rules: elem.rules
			})
		);
		return children;
	};

	return (
		<Form
			form={form}
			name="advanced_search"
			className={styles["ant-advanced-search-form"]}
			onFinish={onFinish}>
			<div className={styles.outer}>
				<Row>{getFields()}</Row>
			</div>

			<Divider style={{ background: "#ddd", marginBottom: 0, marginTop: 10 }} />
		</Form>
	);
};

export default observer(SearchForm);

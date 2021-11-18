import * as React from "react";
import { Form, Input, Button, Divider, DatePicker, Row, Col, Select, message } from "antd";
import { warehouseListArray } from "../constant";
import { SearchOutlined, UpOutlined, DownOutlined } from "@ant-design/icons";
import { store } from "../store/helper/useStore";
import { DateGap, SearchOrderReq } from "genServices/ezShipOMS/oms";
import * as moment from "moment";
import { useObserver } from "mobx-react-lite";
import { WarehouseType } from "genServices/ezShipOrder/submit";
const { Option } = Select;

export default function Search() {
	const [form] = Form.useForm();
	const orderListStore = store().orderListStore;
	const orderListUi = store().orderListUi;
	const commonStore = store().commonStore;
	const [showAllSearchForm, setShowAllSearchForm] = React.useState<boolean>(false);
	const { resetFields } = form;
	function handleFormView() {
		setShowAllSearchForm(!showAllSearchForm);
	}
	const onFinish = values => {
		searchLog(DateGap.DateGapLatestThreeMonth, values);
	};


	const searchLog = (dateGap: DateGap = DateGap.DateGapNone, formValues?) => {
		const values = formValues ? formValues : form.getFieldsValue();
		const { orderNumber, transportNumber, packageCode, billId, invoiceNumber } = values;
		const exclusiveArr = [orderNumber, transportNumber, packageCode, billId, invoiceNumber];
		if (exclusiveArr.filter(i => i !== undefined && i !== "").length > 1) {
			message.error("仅支持一个排他搜索条件，请重置！");
			return;
		}
		orderListUi.setOffset(0);
		orderListUi.setPage(1);
		filterFormValue(values);
		let createDateStart, createDateEnd, payDateStart, payDateEnd;
		createDateStart =
			values.createDate &&
			moment(values.createDate[0])
				.startOf("date")
				.unix();
		createDateEnd =
			values.createDate &&
			moment(values.createDate[1])
				.startOf("date")
				.unix();
		payDateStart =
			values.payDate &&
			moment(values.payDate[0])
				.startOf("date")
				.unix();
		payDateEnd =
			values.payDate &&
			moment(values.payDate[1])
				.startOf("date")
				.unix();
		if (createDateStart === createDateEnd) {
			createDateEnd = createDateStart + 86399;
		}
		if (payDateStart === payDateEnd) {
			payDateEnd = payDateStart + 86399;
		}
		const req = {
			...values,
			createDateStart,
			createDateEnd,
			payDateStart,
			payDateEnd,
			offset: orderListUi.offset,
			limit: orderListUi.limit,
			createDate: undefined,
			payDate: undefined,
			dateGap
		} as SearchOrderReq;
		saveSearchParams(req);
		orderListStore.searchOrderList(req);
	}

	function filterFormValue(values: any) {
		for (let [key, value] of Object.entries(values)) {
			if (!value) {
				delete values[key];
			}
		}
	}
	const initSearchForm = () => {
		resetFields();
		commonStore.fetchWarehouseList("Search", undefined, true);
		commonStore.fetchLogisticsCOList("Search", undefined, true);
	};

	function saveSearchParams(values: SearchOrderReq) {
		orderListStore.saveSearchParams({ ...values });
	}

	const onFinishFailed = errorInfo => {
		console.log("Failed:", errorInfo);
	};

	const initFormValues = {};
	const Label = label => <div style={{ width: 75 }}>{label}</div>;

	return useObserver(() => (
		<React.Fragment>
			<Form
				name="basic"
				form={form}
				initialValues={initFormValues}
				onFinish={onFinish}
				onFinishFailed={onFinishFailed}>
				<Row>
					<Col span={6}>
						<Form.Item label={Label("订单来源")} name="orderOrigin">
							<Select allowClear>
								{commonStore.orderOriginList &&
									commonStore.orderOriginList.length > 0 &&
									commonStore.orderOriginList.map(({ code, name }) => (
										<Option value={code} key={code}>
											{name}
										</Option>
									))}
							</Select>
						</Form.Item>
					</Col>
					<Col span={6}>
						<Form.Item label={Label("订单号")} name="orderNumber">
							<Input allowClear />
						</Form.Item>
					</Col>
					<Col span={6}>
						<Form.Item label={Label("转运运单号")} name="transportNumber">
							<Input allowClear />
						</Form.Item>
					</Col>
					<Col span={6} >
						<Form.Item label={Label("会员名称")} name="nickName">
							<Input allowClear />
						</Form.Item>
					</Col>
				
					<Col span={6}>
						<Form.Item label={Label("付款单号")} name="billId">
							<Input allowClear />
						</Form.Item>
					</Col>		
					<Col span={6}>
						<Form.Item label={Label("包裹号")} name="packageCode">
							<Input allowClear />
						</Form.Item>
					</Col>
					<Col span={6}>
						<Form.Item label={Label("目的地国家")} name="catalog">
							<Select
								allowClear
								onChange={e => {
									commonStore.fetchWarehouseList("Search", e as string);
									resetFields(["warehouseType"]);
									resetFields(["logisticsName"]);
									commonStore.fetchLogisticsCOList("Search", undefined, true);
								}}>
								{commonStore.desCountryCode &&
									commonStore.desCountryCode.desCodeList.map(value => (
										<Option value={value} key={value}>
											{value}
										</Option>
									))}
							</Select>
						</Form.Item>
					</Col>
					<Col span={6}>
						<Form.Item label={Label("转运仓库")} name="warehouseType">
							<Select
								allowClear
								disabled={!commonStore.warehouseListInSearch}
								onChange={e => {
									if (e) {
										commonStore.fetchLogisticsCOList(
											"Search",
											e as WarehouseType
										);
									} else {
										commonStore.fetchLogisticsCOList("Search", undefined, true);
									}
									resetFields(["logisticsName"]);
								}}>
								{commonStore.warehouseListInSearch &&
									commonStore.warehouseListInSearch.warehouse.map(value => (
										<Option value={value} key={value}>
											{warehouseListArray.find(i => i.value === value).label}
										</Option>
									))}
							</Select>
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("快递公司")} name="logisticsName">
							<Select allowClear disabled={!commonStore.logisticsListInSearch}>
								{commonStore.logisticsListInSearch &&
									commonStore.logisticsListInSearch.logisticsList.map(value => (
										<Option value={value} key={value}>
											{value}
										</Option>
									))}
							</Select>
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("运输方式")} name="shipmentType">
							<Select allowClear>
								{commonStore.shippingMethodList &&
									commonStore.shippingMethodList.deliveryMethodList.map(value => (
										<Option value={value.shippingCode} key={value.shippingCode}>
											{value.description}
										</Option>
									))}
							</Select>
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("派送方式")} name="deliveryMethod">
							<Select allowClear>
								{commonStore.deliveryMethodList &&
									commonStore.deliveryMethodList.deliveryMethodList.map(value => (
										<Option value={value} key={value}>
											{value}
										</Option>
									))}
							</Select>
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("订单状态")} name="orderStatus">
							<Select allowClear>
								{commonStore.orderStatusList &&
									commonStore.orderStatusList.orderStatusList.map(
										({ statusCode, description }) => (
											<Option value={statusCode} key={statusCode}>
												{description}
											</Option>
										)
									)}
							</Select>
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("CustomerID")} name="customerId">
							<Input allowClear />
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("标识码")} name="selfHelpCode">
							<Input allowClear />
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("创建日期")} name="createDate" >
							<DatePicker.RangePicker style={{ width: "100%" }} allowClear />
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("付款日期")} name="payDate">
							<DatePicker.RangePicker style={{ width: "100%" }} allowClear />
						</Form.Item>
					</Col>
					<Col span={6} style={showAllSearchForm ? {} : { display: "none" }}>
						<Form.Item label={Label("发货单号")} name="invoiceNumber">
							<Input allowClear />
						</Form.Item>
					</Col>
					<Col span={12}>
						&emsp;&emsp;
						<Button type="primary" htmlType="submit">
							<SearchOutlined />
							查询
						</Button>
						<Button style={{ margin: "0px 10px", width: 100 }} onClick={initSearchForm}>
							重置
						</Button>
						<Button style={{ margin: "0px 10px" }}  type="primary" onClick={() => searchLog(DateGap.DateGapBeforeThreeMonth)}>查询归档 </Button>
						<a style={{ color: "rgba(51, 143, 248, 1)" }} onClick={handleFormView}>
							{showAllSearchForm ? (
								<span>
									收起 <UpOutlined />
								</span>
							) : (
								<span>
									展开 <DownOutlined />{" "}
								</span>
							)}
						</a>
					</Col>
				</Row>
			</Form>
			<Divider style={{ background: "#ddd", marginBottom: 10, margin: 10 }} />
		</React.Fragment>
	));
}

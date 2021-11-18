import * as React from "react";
import { Col, Form, Modal, Input, Row, Select, notification, message, Radio } from "antd";
import {
	warehouseListArray,
	addServiceTypeArr,
	orderRegArray,
	catalogListArray,
	orderRegGlobalArray
} from "../constant";
import { store } from "../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { ItemList } from "./itemList";
import {
	OrderRegTypeEnum,
	SubmitNewOrder,
	AddServiceType,
	SubmitNewOrderRespCode,
	SiteType
} from "genServices/ezShipOMS/oms";
import { WarehouseType } from "genServices/ezShipOMS/public";
import { checkEmpty, priceToInt } from "utils/util";
import { ItemType } from "./item";
// import { isUAT } from "utils/url";
// import { checkGroupId } from "utils/auth";
const styles = require("./index.scss");

// https://www.tapd.cn/51528997/markdown_wikis/show/#1151528997001013342

const { Option } = Select;
const { TextArea } = Input;

interface OrderFormProps {
	handleClose: () => void;
}

const layout = {
	labelCol: { span: 4 },
	wrapperCol: { span: 19 }
};

interface OrderFormProps {
	visible: boolean;
	site: SiteType;
	handleClose: () => void;
	onCreate: (value: string) => void;
}
const OrderForm: React.FC<OrderFormProps> = ({ handleClose, onCreate, visible, site }) => {
	const [form] = Form.useForm();
	const [item, setItem] = React.useState<ItemType[]>([]);
	const [confirmLoading, setConfirmLoading] = React.useState<boolean>(false);
	const [rendvisible, setVisible] = React.useState(visible);
	const [isActiveChange, setIsActiveChange] = React.useState(false);
	// const { resetFields } = form;
	const commonStore = store().commonStore;
	const orderStore = store().orderStore;

	const isGlobal = site === SiteType.SiteTypeGlobal;

	const catalogList = Object.keys(catalogListArray);

	React.useEffect(() => {
		if (isGlobal) {
			changeCatalog(null);
		} else {
			const catalog = getInitCatalog();
			if (catalog) {
				changeCatalog(catalog);
			}
		}
	}, [rendvisible]);

	// const groupId = isUAT() ? "10308" : "10660";
	const isAm =false
	// checkGroupId(groupId); // 是美国用户

	const getInitCatalog = () => {
		let catalog = null;
		if (
			commonStore.desCountryCode &&
			commonStore.desCountryCode.desCodeList &&
			commonStore.desCountryCode.desCodeList.length > 0
		) {
			// 美国 - 目的地国家： 默认新加坡
			// const findSG = commonStore.desCountryCode.desCodeList.find(elem => elem === "SG");
			// catalog = isAm && findSG ? "SG" : commonStore.desCountryCode.desCodeList[0];
		}
		return catalog;
	};

	const getInitWarehouse = () => {
		let warehouse = null;
		if (
			commonStore.warehouseListInCreate &&
			commonStore.warehouseListInCreate.warehouse &&
			commonStore.warehouseListInCreate.warehouse.length > 0
		) {
			const currentWarehouse = form.getFieldValue("warehouse");
			// 判断当前的仓库在不在 仓库列表中 -- 在就不变，不在就变成列表第一个值
			// 美国 - 货源地仓库： 默认美国
			// 全球化默认选 广州
			const findAm =
				isAm &&
				commonStore.warehouseListInCreate.warehouse.includes(
					WarehouseType.WarehouseTypeUSA
				);
			const guangzhou = commonStore.warehouseListInCreate.warehouse.includes(
				WarehouseType.WarehouseTypeGuangzhou
			);
			if (findAm) {
				warehouse = WarehouseType.WarehouseTypeUSA;
			} else {
				if (
					currentWarehouse &&
					commonStore.warehouseListInCreate.warehouse.includes(currentWarehouse)
				) {
					warehouse = currentWarehouse;
				} else {
					if (isGlobal && guangzhou) {
						warehouse = WarehouseType.WarehouseTypeGuangzhou;
					} else {
						warehouse = commonStore.warehouseListInCreate.warehouse[0];
					}
				}
			}
		}
		return warehouse;
	};

	const getInitLogisticsName = () => {
		let logisticsName = "其他";
		if (
			commonStore.logisticsListInCreate &&
			commonStore.logisticsListInCreate.logisticsList &&
			commonStore.logisticsListInCreate.logisticsList.length > 0
		) {
			logisticsName = commonStore.logisticsListInCreate.logisticsList[0];
		}
		return logisticsName;
	};

	// 美国 - 订单的识别类型的选择，默认不选择，然后基于收货关键字来进行自动调整，同时支持手动调整。
	// 其他国家 - 订单识别类型默认 用户昵称
	const initValues = {
		["catalog"]: getInitCatalog(),
		["orderRegType"]: isGlobal
			? OrderRegTypeEnum.OrderRegTypeEnumRegCode
			: isAm
			? null
			: OrderRegTypeEnum.OrderRegTypeEnumNickName,
		["warehouse"]: getInitWarehouse(),
		["logisticsName"]: getInitLogisticsName()
	};

	// 修改国家
	const changeCatalog = async (catalog: string) => {
		// 根据国家修改转运仓库
		await commonStore.fetchWarehouseList("Create", catalog);
		const warehouse = getInitWarehouse();
		changeWarehouse(warehouse);
	};

	// 修改转运仓库
	const changeWarehouse = async warehouse => {
		// 根据转运仓库，初始化  附加服务
		await orderStore.initAddedServiceTerm({ warehouse });
		// 根据仓库获取快递公司
		await commonStore.fetchLogisticsCOList("Create", warehouse);
		// 选择仓库之后，快递公司默认加载第一家。
		const logisticsName = getInitLogisticsName();
		// 订单识别类型调整到收货关键字下方，非美国仓库之外，依然默认选中用户昵称 为标识类型
		// 1. 美国用户 -- 标识类型 1.1有值则保持不变， 1.2没有值则  显示 null
		// 2. 其他国家用户 -- 表示类型  1.1有值则保持不变， 1.2没有值则 使用用户昵称
		const currentOrderRegType = form.getFieldValue("orderRegType");
		const orderRegType = isAm
			? currentOrderRegType
				? currentOrderRegType
				: null
			: currentOrderRegType
			? currentOrderRegType
			: OrderRegTypeEnum.OrderRegTypeEnumNickName;
		form.setFieldsValue({
			warehouse,
			logisticsName,
			orderRegType: isGlobal ? currentOrderRegType : orderRegType
		});
	};

	// 修改收货关键字
	const changePickupKey = pickupKey => {
		if (isAm && !isActiveChange) {
			if (pickupKey) {
				let orderRegType = null;
				const orderRegTypeEnumRegCodeReg = /^[a-zA-Z]{5}$/; // 标识码
				const orderRegTypeEnumUserIdReg = /^[0-9]*$/; // 用户ID
				if (orderRegTypeEnumRegCodeReg.test(pickupKey)) {
					orderRegType = orderRegArray[2].value;
				} else if (orderRegTypeEnumUserIdReg.test(pickupKey)) {
					orderRegType = orderRegArray[1].value;
				} else {
					orderRegType = orderRegArray[0].value;
				}
				form.setFieldsValue({ orderRegType });
			} else {
				form.setFieldsValue({ orderRegType: null });
			}
		}
	};

	const handleOk = () => {
		form.submit();
	};
	function handleCancel() {
		commonStore.clearList("Create");
		commonStore.fetchLogisticsCOList("Create", undefined, true);
		handleClose();
	}

	async function onFinish(values) {
		if (item.length !== 0) {
			const findEmpty = item.find(
				elem => !elem.category || !elem.categoryId || checkEmpty(elem.qty)
			);
			if (findEmpty) {
				message.error("请将item内容填写完整");
			} else {
				setConfirmLoading(true);
				const param = {
					catalog: values.catalog,
					orderRegType: values.orderRegType,
					pickupKey: values.pickupKey,
					logisticsNumber: values.logisticsNumber.trim(),
					warehouse: values.warehouse,
					logisticsName:
						values.logisticsName === "其他"
							? values.otherLogisticsName
							: values.logisticsName,
					isNeedPhotos: values.addServeType
						? values.addServeType.some(i => i === AddServiceType.Service_PHOTOGRAPH)
						: false,
					isREPack: values.addServeType
						? values.addServeType.some(i => i === AddServiceType.Service_REPACKAGE)
						: false,
					remark: values.remark,
					declared_amount: values.declared_amount
						? priceToInt(values.declared_amount)
						: undefined,
					items: (item || []).map((v: ItemType) => ({
						categoryId: v.categoryId,
						qty: v.qty
					})),
					site
				};

				SubmitNewOrder({ ...param, isFirstTime: true })
					.then(resp => {
						if (resp.code === SubmitNewOrderRespCode.Success) {
							setConfirmLoading(false);
							onCreate(resp.orderNumber);
						} else {
							Modal.confirm({
								title: "确认国家和用户名",
								content: (
									<div>
										其他国家存在<span className={styles.red}>相同昵称</span>
										的用户，请确认<span className={styles.red}>用户和国家</span>
										是否正确！
									</div>
								),
								okText: "确认无误",
								cancelText: "返回检查",
								onCancel: () => setConfirmLoading(false),
								onOk: () => {
									SubmitNewOrder({ ...param, isFirstTime: false })
										.then(resp => {
											setConfirmLoading(false);
											if (resp.code === SubmitNewOrderRespCode.Success) {
												onCreate(resp.orderNumber);
											}
										})
										.catch(err => {
											setConfirmLoading(false);
											notification.error({ message: err.message });
										});
								}
							});
						}
					})
					.catch(err => {
						setConfirmLoading(false);
						notification.error({ message: err.message });
					});
			}
		} else {
			message.error("item为空,请添加item后重试！", 1);
		}
	}

	function onFinishFailed(errorInfo) {
		console.log("Failed:", errorInfo);
	}
	return useObserver(() => (
		<Modal
			title={isGlobal ? "新建ezship国际化订单" : "新建ezship新马订单"}
			okText="确认"
			cancelText="取消"
			visible={true}
			onOk={handleOk}
			onCancel={handleCancel}
			width={700}
			confirmLoading={confirmLoading}
			afterClose={() => setVisible(false)}>
			<Row justify="center">
				<Col span={24}>
					<Form
						{...layout}
						form={form}
						initialValues={initValues}
						onFinish={onFinish}
						onFinishFailed={onFinishFailed}>
						{/** 目的地国家 */}
						{!isGlobal && (
							<Form.Item
								label="目的地国家"
								name="catalog"
								rules={[{ required: true, message: "请选择目的地国家" }]}>
								<Radio.Group
									className={styles.btnRadio}
									buttonStyle="solid"
									onChange={e => changeCatalog(e.target.value)}>
									{catalogList &&
										catalogList.length > 0 &&
										catalogList.map(value => (
											<Radio.Button value={value} key={value}>
												{catalogListArray[value]}
											</Radio.Button>
										))}
								</Radio.Group>
							</Form.Item>
						)}
						{/** 转运仓库 */}
						{/* 国家化仓库  上海和广州 */}
						<Form.Item
							label="转运仓库"
							name="warehouse"
							rules={[{ required: true, message: "请输入转运仓库!" }]}>
							<Radio.Group
								disabled={!commonStore.warehouseListInCreate}
								onChange={e => changeWarehouse(e.target.value)}>
								{commonStore.warehouseListInCreate &&
									commonStore.warehouseListInCreate.warehouse &&
									commonStore.warehouseListInCreate.warehouse.length > 0 &&
									commonStore.warehouseListInCreate.warehouse.map(value => (
										<Radio value={value} key={value}>
											{warehouseListArray.find(i => i.value === value).label}
										</Radio>
									))}
							</Radio.Group>
						</Form.Item>
						{/** 快递公司 */}
						<Form.Item
							label="快递公司"
							name="logisticsName"
							rules={[{ required: true, message: "请输入快递公司" }]}>
							<Select disabled={!commonStore.logisticsListInCreate}>
								{commonStore.logisticsListInCreate &&
									commonStore.logisticsListInCreate.logisticsList &&
									commonStore.logisticsListInCreate.logisticsList
										.concat(["其他"])
										.map(value => (
											<Option value={value} key={value}>
												{value}
											</Option>
										))}
							</Select>
						</Form.Item>
						<Form.Item noStyle shouldUpdate>
							{({ getFieldValue }) => {
								return getFieldValue("logisticsName") === "其他" ? (
									<Form.Item
										name="otherLogisticsName"
										label=" "
										colon={false}
										rules={[{ required: true }]}>
										<Input placeholder="请填写其他快递公司名称" allowClear />
									</Form.Item>
								) : null;
							}}
						</Form.Item>
						{/** 转运运单号 */}
						<Form.Item
							label="转运运单号"
							name="logisticsNumber"
							rules={[{ required: true, message: "请输入转运运单号" }]}>
							<Input placeholder="请填写" allowClear />
						</Form.Item>
						{/** 收货关键字 */}
						<Form.Item
							label="收货关键字"
							name="pickupKey"
							rules={[
								{ required: true, message: "请输入收货关键字" },
								{
									pattern: isGlobal ? /^[A-Z0-9]*$/ : null,
									message: "请输入字母标识码: 可输入大写字母和数字"
								}
							]}>
							<Input
								placeholder="请填写"
								allowClear
								onChange={e => (isGlobal ? null : changePickupKey(e.target.value))}
							/>
						</Form.Item>
						{/** 订单识别类型 */}
						<Form.Item
							label="订单识别类型"
							name="orderRegType"
							rules={[{ required: true, message: "请选择订单识别类型" }]}>
							<Radio.Group onChange={() => setIsActiveChange(true)}>
								{(isGlobal ? orderRegGlobalArray : orderRegArray).map(
									({ label, value }) => (
										<Radio value={value} key={value}>
											{label}
										</Radio>
									)
								)}
							</Radio.Group>
						</Form.Item>

						{/** 附加服务 */}
						<Form.Item label="附加服务" name="addServeType">
							<Select
								mode="multiple"
								style={{ width: "100%" }}
								disabled={!orderStore.addedServiceTerm}>
								{(orderStore.addedServiceTerm || []).map((item, index) => (
									<Option value={item} key={index}>
										{addServiceTypeArr.find(i => i.value === item) &&
											addServiceTypeArr.find(i => i.value === item).label}
									</Option>
								))}
							</Select>
						</Form.Item>
						{/** 备注说明 */}
						<Form.Item
							label="备注说明"
							name="remark"
							rules={[{ min: 5 }, { whitespace: true }]}>
							<TextArea placeholder="请输入至少五个字符" allowClear />
						</Form.Item>
						{/** 总申报金额 */}
						{/* <Form.Item
							label={
								<div>
									总申报金额
									<br />
									<span style={{ color: "#d9051c", marginTop: -2 }}>
										(目的地国家当地币)
									</span>
								</div>
							}
							name="declared_amount"
							className={styles.declaredAmount}>
							<InputNumber
								style={{ width: "100%" }}
								max={99999.99}
								min={0}
								step={0.01}
								parser={value => limitDecimals(value)}
							/>
						</Form.Item> */}
						<ItemList save={setItem} />
					</Form>
				</Col>
			</Row>
		</Modal>
	));
};

export default OrderForm;

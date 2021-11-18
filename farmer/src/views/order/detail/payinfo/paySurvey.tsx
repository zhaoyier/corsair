import * as React from "react";
import { Tabs, Table, Descriptions } from "antd";
const { TabPane } = Tabs;
const styles = require("./index.scss");
import { store } from "../../store/helper/useStore";
import { formatPrice, formatWeightUnit } from "utils/util";

function callback(key) {
	console.log(key);
}
const surveyColumns = [
	{
		title: "订单号",
		dataIndex: "orderNumber",
		render: (t, r) => (
			<a rel="opener" href={`order.html#/detail?orderId=${r.orderId}`} target="_blank">
				{t}
			</a>
		)
	},
	{
		title: "国际运费",
		dataIndex: "internationalShippingFee",
		render: text => formatPrice(text)
	},
	{
		title: "本地派送",
		dataIndex: "localDeliveryFee",
		render: text => formatPrice(text)
	},
	{
		title: "保险费",
		dataIndex: "insuranceFee",
		render: text => formatPrice(text)
	},
	{
		title: "附加服务费用",
		dataIndex: "photoServiceFee",
		render: (text, record) => (
			<span>{formatPrice(Number(text) + Number(record.repackageServiceFee) + "")}</span>
		)
	},
	{
		title: "GST",
		dataIndex: "gst",
		render: text => formatPrice(text)
	},
	{
		title: "逾期仓储费",
		dataIndex: "storageFee",
		render: text => formatPrice(text)
	},
	{
		title: "折扣金额",
		dataIndex: "shippingDiscount",
		render: (text, record) => (
			<span>{formatPrice(Number(text) + Number(record.couponDiscount) + "")}</span>
		)
	},
	{
		title: "订单申报价值",
		dataIndex: "declaredValue",
		render: text => formatPrice(text)
	},
	{
		title: "称重重量",
		dataIndex: "actualWeight",
		render: val => formatWeightUnit(val)
	},
	{
		title: "体积重量",
		dataIndex: "volumeWeight",
		render: val => formatWeightUnit(val)
	},
	{
		title: "计费重量",
		dataIndex: "chargeableWeight",
		render: val => formatWeightUnit(val)
	},
	{
		title: "订单实付总额",
		dataIndex: "totalFee",
		render: text => formatPrice(text)
	}
];
const discountInfoColumns = [
	{
		title: "订单号",
		dataIndex: "orderNumber",
		render: (t, r) => (
			<a rel="opener" href={`order.html#/detail?orderId=${r.orderId}`} target="_blank">
				{t}
			</a>
		)
	},
	{
		title: "折扣名称",
		dataIndex: "discountName"
	},
	{
		title: "折扣金额",
		dataIndex: "discountAmount",
		render: text => formatPrice(text)
	}
];
export function PaySurvey() {
	const orderStore = store().orderStore;
	return (
		<div className={styles.index}>
			<Tabs onChange={callback} type="card">
				<TabPane tab="概况" key="most">
					<Table
						columns={surveyColumns}
						rowClassName={record =>
							record.orderNumber === orderStore.order.orderNo ? styles.rowBg : ""
						}
						dataSource={
							orderStore.orderPayInfo && orderStore.orderPayInfo.fees
								? orderStore.orderPayInfo.fees
								: []
						}
						rowKey={record => record.orderNumber}
						pagination={false}
					/>
				</TabPane>
				<TabPane tab="折扣信息" key="discount" style={{ width: 400 }}>
					<Table
						columns={discountInfoColumns}
						rowClassName={record =>
							record.orderNumber === orderStore.order.orderNo ? styles.rowBg : ""
						}
						dataSource={orderStore.orderPayInfo && orderStore.orderPayInfo.discounts}
						rowKey={record => record.orderNumber}
						pagination={false}
					/>
				</TabPane>
			</Tabs>
			<br />
			{orderStore.orderPayInfo && (
				<div className={styles.rigthFloat}>
					<Descriptions
						style={{ width: 400 }}
						title="总支付账单合计："
						size="small"
						column={1}
						className={styles.setlabelwidth}>
						<Descriptions.Item label="包裹计费重：">
							{formatWeightUnit(orderStore.orderPayInfo.totalChargeWeight)}
						</Descriptions.Item>
						{/* <Descriptions.Item label="总运费：">
							{formatPrice(orderStore.orderPayInfo.totalShippingFee)}
						</Descriptions.Item> */}
						<Descriptions.Item label="总国际运费：">
							{formatPrice(orderStore.orderPayInfo.totalInternationalShippingFee)}
						</Descriptions.Item>
						<Descriptions.Item label="总派送费：">
							{formatPrice(orderStore.orderPayInfo.totalLocalDeliveryFee)}
						</Descriptions.Item>
						<Descriptions.Item label="总折扣：">
							{formatPrice(orderStore.orderPayInfo.totalDiscountFee)}
						</Descriptions.Item>
						<Descriptions.Item label="仓库附加费：">
							{formatPrice(orderStore.orderPayInfo.totalStorageFee)}
						</Descriptions.Item>
						<Descriptions.Item label="总申报金额：">
							{formatPrice(orderStore.orderPayInfo.totalDeclaredValue)}
						</Descriptions.Item>
						<Descriptions.Item label="总税费：">
							{formatPrice(orderStore.orderPayInfo.totalGst)}
						</Descriptions.Item>
						<Descriptions.Item label="账单实付总金额：">
							{formatPrice(orderStore.orderPayInfo.totalActualPaymentFee)}
						</Descriptions.Item>
					</Descriptions>
				</div>
			)}
		</div>
	);
}

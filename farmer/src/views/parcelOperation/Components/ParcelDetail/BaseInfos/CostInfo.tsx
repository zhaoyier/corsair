import * as React from "react";
import { observer } from "mobx-react-lite";
import { Table } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { formatPrice, formatWeightUnit } from "utils/util";

const styles = require("./index.scss");

const getColumns = () => {
	const columns = [
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
			title: "国际运费实付",
			dataIndex: "internationalShippingFeeActual",
			render: text => formatPrice(text)
		},
		{
			title: "本地派送费",
			dataIndex: "localDeliveryFee",
			render: text => formatPrice(text)
		},
		{
			title: "逾期仓储费",
			dataIndex: "storageFee",
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
			title: "保险",
			dataIndex: "insuranceFee",
			render: text => formatPrice(text)
		},
		{
			title: "折扣总额",
			dataIndex: "shippingDiscount",
			render: (text, record) => (
				<span>{formatPrice(Number(text) + Number(record.couponDiscount) + "")}</span>
			)
		},
		{
			title: "包裹订单总费用", // 订单实付总额
			dataIndex: "totalFee",
			render: text => formatPrice(text)
		},
		{
			title: "计费重量（kg）",
			dataIndex: "chargeableWeight",
			render: text => formatWeightUnit(text)
		},
		{
			title: "币种",
			dataIndex: "currencyCode"
		}
	];
	return columns;
};

const CostInfo = () => {
	const { parcelDetailStore } = useStores();
	const { record } = parcelDetailStore;
	const { parcelFee } = record;

	const renderSummaryRow = () => {
		return (
			<React.Fragment>
				<Table.Summary.Row>
					<Table.Summary.Cell index={0}>总计：</Table.Summary.Cell>
					<Table.Summary.Cell index={1} />
					<Table.Summary.Cell index={2}>
						{parcelFee && formatPrice(parcelFee.totalInternationalShippingFeeActual)}
					</Table.Summary.Cell>
					<Table.Summary.Cell index={3}>
						{parcelFee && formatPrice(parcelFee.totalLocalDeliveryFee)}
					</Table.Summary.Cell>
					<Table.Summary.Cell index={4}>
						{parcelFee && formatPrice(parcelFee.totalStorageFee)}
					</Table.Summary.Cell>
					<Table.Summary.Cell index={5}>
						{parcelFee && formatPrice(parcelFee.totalPhotoServiceFee)}
					</Table.Summary.Cell>
					<Table.Summary.Cell index={6}>
						{parcelFee && formatPrice(parcelFee.totalGst)}
					</Table.Summary.Cell>
					<Table.Summary.Cell index={7}>
						{parcelFee && formatPrice(parcelFee.totalInsuranceFee)}
					</Table.Summary.Cell>
					<Table.Summary.Cell index={8} />
					<Table.Summary.Cell index={9}>
						{parcelFee && formatPrice(parcelFee.totalFee)}
					</Table.Summary.Cell>
					<Table.Summary.Cell index={10}>
						{parcelFee && formatWeightUnit(parcelFee.totalChargeWeight)}
					</Table.Summary.Cell>
				</Table.Summary.Row>
			</React.Fragment>
		);
	};

	return (
		<section>
			<div className={styles.header}>包裹费用信息</div>
			<Table
				rowKey="orderId"
				columns={getColumns()}
				dataSource={(parcelFee && parcelFee.orderFees) || []}
				pagination={false}
				summary={renderSummaryRow}
			/>
		</section>
	);
};

export default observer(CostInfo);

import * as React from "react";
import { OrderDetail } from "./orderDetail";
import { ProductInfo } from "./productInfo";
import { DrliveryInfo } from "./drliveryInfo";
import { OrderRemark } from "./orderRemark";
import { Divider } from "antd";
const styles = require("./index.scss");

const DivMargin = () => (
	<Divider style={{ background: "#dfdfdf", marginBottom: 20, marginTop: 20  }} />
);

export function BaseInfo() {
	return (
		<div className={styles.baseinfo}>
			<div className={styles.header} style={{marginTop: 0}}>订单明细</div>
			<OrderDetail />
			<DivMargin />
			<div className={styles.header}>商品信息</div>
			<ProductInfo />
			<DivMargin />
			<div className={styles.header}>配送信息</div>
			<DrliveryInfo />
			<DivMargin />
			<div className={styles.header}>订单备注</div>
			<OrderRemark />
		</div>
	);
}

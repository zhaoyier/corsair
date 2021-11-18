import * as React from "react";
import { observer } from "mobx-react-lite";
import * as cs from "classnames";
import { Row, Col } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { formatWeight } from "utils/util";

const styles = require("./index.scss");

const BaseInfo = () => {
	const { parcelDetailStore } = useStores();
	const { record } = parcelDetailStore;
	const {
		cabinetNumber,
		boxNumber,
		// warehouseId,
		warehouse,
		packageType,
		transport,
		actualWeight,
		volumeWeight,
		chargWeight,
		length,
		width,
		height
	} = record;
	return (
		<section>
			<div className={cs(styles.header, styles.first)}>基础信息</div>
			<Row gutter={[16, 4]}>
				<Col span={6}>
					<span>主单号/柜号：</span>
					{cabinetNumber}
				</Col>
				<Col span={6}>
					<span>封箱号：</span>
					{boxNumber}
				</Col>
				<Col span={6}>
					<span>发货仓库：</span>
					{warehouse}
				</Col>
				<Col span={6}>
					<span>包裹类型：</span>
					{packageType}
				</Col>
				<Col span={6}>
					<span>运输方式：</span>
					{transport}
				</Col>
				<Col span={6}>
					<span>包裹重量：</span>
					{formatWeight(actualWeight)}KG
				</Col>
				<Col span={6}>
					<span>体积重量：</span>
					{formatWeight(volumeWeight)}KG
				</Col>
				<Col span={6}>
					<span>计费重量：</span>
					{formatWeight(chargWeight)}KG
				</Col>
				<Col span={6}>
					<span>长(cm)：</span>
					{length}
				</Col>
				<Col span={6}>
					<span>宽(cm)：</span>
					{width}
				</Col>
				<Col span={6}>
					<span>高(cm)：</span>
					{height}
				</Col>
			</Row>
		</section>
	);
};

export default observer(BaseInfo);

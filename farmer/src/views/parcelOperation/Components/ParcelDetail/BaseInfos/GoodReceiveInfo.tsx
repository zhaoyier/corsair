import * as React from "react";
import { observer } from "mobx-react-lite";
import { Row, Col } from "antd";
import { useStores } from "views/parcelOperation/hooks";

const styles = require("./index.scss");

const BaseInfo = () => {
	const { parcelDetailStore } = useStores();
	const { record } = parcelDetailStore;
	const { deliveryInfo } = record;
	if (!deliveryInfo) {
		return <div>未发现收货人信息</div>;
	}
	const { userName, userPhone, deliveryAddress } = deliveryInfo;
	return (
		<section>
			<div className={styles.header}>收货信息</div>
			<Row gutter={[16, 4]}>
				<Col span={6}>
					<span>收件人：</span> {userName}
				</Col>
				<Col span={18}>
					<span>收件人电话：</span>
					{userPhone}
				</Col>
				<Col span={24}>
					<span>收件人地址：</span>
					{deliveryAddress}
				</Col>
			</Row>
		</section>
	);
};

export default observer(BaseInfo);

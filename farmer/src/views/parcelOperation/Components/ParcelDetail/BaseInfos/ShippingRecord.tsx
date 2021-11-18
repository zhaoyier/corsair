import * as React from "react";
import { observer } from "mobx-react-lite";
import { Col, Row, Button, Modal } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { formatUnixTime, TimeFormatTotalReverse, TimeFormatToDayReverse } from "utils/time";

const styles = require("./index.scss");

const showSign = (url: string, visible: boolean, setVisible) => {
	return (
		<Modal title="查看签名" visible={visible} footer={null} onCancel={() => setVisible(false)}>
			<img className={styles.sign} src={url} />
		</Modal>
	);
};

const BaseInfo = () => {
	const { parcelDetailStore } = useStores();
	const { record } = parcelDetailStore;
	const { shipmentInfo } = record;
	const [visible, setVisible] = React.useState(false);
	if (!shipmentInfo) {
		return <div>未发现运送记录信息</div>;
	}
	const {
		transportName,
		transportNumber,
		outgoingDate,
		toDestDate,
		pickupTime,
		pickupPeriod,
		signDate,
		signUrl,
		dispatchEta,
		deliveryMethod
	} = shipmentInfo;
	return (
		<section>
			<div className={styles.header}>运送记录</div>
			<Row gutter={[16, 4]}>
				<Col span={6}>
					<span>物流商：</span>
					{transportName}
				</Col>
				<Col span={6}>
					<span>物流单号：</span>
					{transportNumber}
				</Col>
				<Col span={12}>
					<span>派送方式：</span>
					{deliveryMethod}
				</Col>
				<Col span={6}>
					<span>出库时间：</span>
					{formatUnixTime(outgoingDate, TimeFormatTotalReverse)}
				</Col>
				<Col span={12}>
					<span>到达目的地时间：</span>
					{formatUnixTime(toDestDate, TimeFormatTotalReverse)}
				</Col>
				<Col span={6}>
					<span>预约取货时间段：</span>
					{pickupPeriod ? pickupPeriod : null}
				</Col>
				<Col span={6}>
					<span>发货ETA：</span>
					{dispatchEta}
				</Col>
				<Col span={12}>
					<span>预约取货时间：</span>
					{formatUnixTime(pickupTime, TimeFormatToDayReverse)}
				</Col>
				<Col span={24}>
					<span>签收时间：</span>
					{formatUnixTime(signDate, TimeFormatTotalReverse)}&emsp;&emsp;
					{signUrl && signUrl.length > 0 && (
						<Button type="link" onClick={() => setVisible(true)}>
							查看签名
						</Button>
					)}
				</Col>
			</Row>
			{showSign(signUrl, visible, setVisible)}
		</section>
	);
};

export default observer(BaseInfo);

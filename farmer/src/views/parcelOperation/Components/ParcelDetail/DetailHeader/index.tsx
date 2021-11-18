import * as React from "react";
import { useStores } from "views/parcelOperation/hooks";
import { observer } from "mobx-react-lite";
import { Button, Row, Col } from "antd";
import { OrderOriginEnum } from "genServices/ezShipOMS/public";

const styles = require("./index.scss");

const DetailHeader = () => {
	const { parcelDetailStore, modalStore } = useStores();
	const { record } = parcelDetailStore;
	const {
		clickPrintLabel,
		clickComplementSend,
		clickDelayNotification,
		clickSplitDelivery,
		clickPODRemark
	} = modalStore;
	const { packageCode, packageStatus, shipmentId, regionCode } = record;
	const isGlobal =
		record.packageOrigin &&
		record.packageOrigin.code === OrderOriginEnum.OrderOriginEnumEzshipGlobal;
	const disabled = isGlobal || regionCode === "MY";
	return (
		<section className={styles.detailHeader}>
			<section className={styles.detailHeaderBtns}>
				<Button
					type="primary"
					disabled={disabled}
					className={styles.btn}
					onClick={() => clickPrintLabel(record)}>
					打印标签
				</Button>
				<Button
					type="primary"
					disabled={disabled}
					className={styles.btn}
					onClick={() => clickComplementSend(record)}>
					补送
				</Button>
				<Button
					type="primary"
					disabled={disabled}
					className={styles.btn}
					onClick={() => clickPODRemark("1", [record.packageCode], [record])}>
					包裹备注
				</Button>
				<Button
					type="primary"
					disabled={disabled}
					className={styles.btn}
					onClick={() => clickPODRemark("2", [record.packageCode], [record])}>
					配送备注
				</Button>
				<Button
					type="primary"
					disabled={disabled}
					className={styles.btn}
					onClick={() => clickDelayNotification([record.packageCode], [record])}>
					延误通知
				</Button>
				<Button
					type="primary"
					disabled={disabled}
					className={styles.btn}
					onClick={() => clickSplitDelivery(record)}>
					拆分配送
				</Button>
			</section>
			<h4>
				<Row gutter={24}>
					<Col span={5}>包裹号：{packageCode}</Col>
					<Col span={5}>包裹状态：{packageStatus}</Col>
					<Col span={5}>配送单号：{shipmentId}</Col>
				</Row>
			</h4>
		</section>
	);
};

export default observer(DetailHeader);

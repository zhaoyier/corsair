import * as React from "react";
import { Modal, Button } from "antd";
import ParcelTable from "../../Common/ParcelTable";
import NotOutboundOrderTable from "../../Common/NotOutboundOrderTable";

export interface DeliveryNumModalProps {
	visible: boolean;
	deliveryId: string; // 配送单号
	parcelList: any[]; // 包裹列表
	notOutboundOrderList: any[]; // 未出库订单列表
	onCancel?: () => void;
}

const DeliveryNumModal = (props: DeliveryNumModalProps) => {
	const { visible, deliveryId, parcelList, notOutboundOrderList, onCancel } = props;
	const title = `配送单号：${deliveryId}`;

	return (
		<Modal
			visible={visible}
			title={title}
			onCancel={onCancel}
			width={900}
			footer={<Button onClick={onCancel}>关闭</Button>}>
			<ParcelTable dataSource={parcelList} />
			<NotOutboundOrderTable dataSource={notOutboundOrderList} />
		</Modal>
	);
};
export default DeliveryNumModal;

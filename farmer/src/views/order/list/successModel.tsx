import * as React from "react";
import { CopyToClipboard } from "react-copy-to-clipboard";
import { CheckCircleTwoTone, CopyOutlined } from "@ant-design/icons";
import { Button, Col, Row, Modal, message } from "antd";
interface SuccessModelProps {
	handleClose: () => void;
	orderId: string;
}

export function SuccessModal(props: SuccessModelProps) {
	const { handleClose, orderId } = props;
	function copyOrderId() {
		message.success("复制成功");
	}
	return (
		<Modal visible={true} footer={null} onCancel={handleClose}>
			<Row>
				<Col style={{ lineHeight: 3 }}>
					<CheckCircleTwoTone twoToneColor="#52c41a" />
				</Col>
				<Col offset={1}>
					<strong style={{ lineHeight: 3 }}>订单创建成功!</strong>
				</Col>
			</Row>
			<Row>
				<Col offset={2}>
					<span style={{ color: "rgba(146,146,146,1)" }}>订单号:</span>
					<strong>{orderId}</strong>
				</Col>
				<Col offset={1}>
					<CopyToClipboard text={orderId} onCopy={copyOrderId}>
						<CopyOutlined />
					</CopyToClipboard>
				</Col>
			</Row>
			<Row>
				<Col offset={20}>
					<Button type="primary" onClick={handleClose}>
						确定
					</Button>
				</Col>
			</Row>
		</Modal>
	);
}

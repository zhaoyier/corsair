import * as React from "react";
import { Modal } from "antd";
import { observer } from "mobx-react-lite";
import { useModalStores } from "./store/hooks";
import CreateDelayNotification from "./create";
import CommonNotification from "./common";

export interface DelayNotificationModalProps {
	visible: boolean;
	selectedRows: any[];
	selectedRowKeys: string[];
	onCancel?: () => void;
	onOk?: () => void;
}
// 延误通知  弹窗

const DelayNotificationModal = (props: DelayNotificationModalProps) => {
	const store = useModalStores();
	const {
		getDelayNoticeTemplate,
		getDelayNotice,
		createDelayNotice,
		mailChecked,
		currentNoticeTemplate,
		currentMailTemplate
	} = store;
	const { visible, onCancel, onOk, selectedRowKeys } = props;
	const title = "延误通知";
	const isSingle = selectedRowKeys && selectedRowKeys.length === 1;
	const [modalVisible, setModalVisible] = React.useState(visible);

	const getDisabled = () => {
		if (isSingle) {
			if (!currentNoticeTemplate.content) {
				return true;
			}
		} else {
			if (!currentNoticeTemplate.content) {
				return true;
			}
			if (mailChecked && !currentMailTemplate.content) {
				return true;
			}
		}
		return false;
	};

	const disabled = getDisabled();

	React.useEffect(() => {
		getDelayNoticeTemplate();
		if (isSingle) {
			getDelayNotice(selectedRowKeys[0]);
		}
	}, [modalVisible]);

	const handleOK = () => {
		createDelayNotice(selectedRowKeys, () => {
			onOk();
		});
	};

	const handleCancel = () => {
		onCancel();
		setModalVisible(false);
	};

	return (
		<Modal
			visible={modalVisible}
			title={title}
			width={700}
			okButtonProps={{
				disabled: Boolean(disabled)
			}}
			onCancel={handleCancel}
			onOk={handleOK}>
			{!isSingle ? <CreateDelayNotification /> : <CommonNotification />}
		</Modal>
	);
};

export default observer(DelayNotificationModal);

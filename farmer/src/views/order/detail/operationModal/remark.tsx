import { Modal, Input, Radio, Select, Form } from "antd";
import * as React from "react";

import { UpFile } from "./upFile";
import { RemarkTypeArr } from "../../constant";
import { store } from "../../store/helper/useStore";
import { RemarkTypeEnum } from "genServices/ezShipOMS/public";
import { useObserver } from "mobx-react-lite";
import { sampleOrderInfo } from "genServices/ezShipOMS/oms";

const { TextArea } = Input;
const Option = Select.Option;
const layout = {
	labelCol: { span: 6 },
	wrapperCol: { span: 16 }
};

interface RemarkModalProps {
	orderId?: string;
	record?: sampleOrderInfo;
	handleClose: () => void;
}

export function RemarkModal(props: RemarkModalProps) {
	const { handleClose, orderId, record } = props;
	const [form] = Form.useForm();
	const orderStore = store().orderStore;
	const orderListStore = store().orderListStore;
	const orderListUi = store().orderListUi;
	const [remarkType, setRemarkType] = React.useState<RemarkTypeEnum>(RemarkTypeEnum.ORDER_REMARK);
	const [enclosureLink, setEnclosureLink] = React.useState<string[]>([]);

	function onOk() {
		form.submit();
	}

	function onChange(value) {
		setRemarkType(value);
	}

	async function handleSubmit() {
		await orderStore.addRemarks({
			orderId: orderId ? orderId : orderStore.orderId,
			remarks: {
				remarkType,
				content: form.getFieldValue("content"),
				isReply: form.getFieldValue("isReply"),
				enclosureLink
			}
		});
		if (orderId) {
			await orderListStore.searchOrderList({
				...orderListStore.searchParams,
				offset: (orderListUi.page - 1) * orderListUi.limit,
				limit: orderListUi.limit
			} as any);
		} else {
			await orderStore.initOrderDetail({ orderId: orderId ? orderId : orderStore.orderId });
		}
		handleClose();
	}

	return useObserver(() => (
		<Modal title="新增备注" visible={true} onOk={onOk} onCancel={handleClose}>
			<Form.Provider onFormFinish={handleSubmit}>
				<Form
					name="remark"
					form={form}
					initialValues={{ remarkType: RemarkTypeEnum.ORDER_REMARK, isReply: true }}
					{...layout}>
					<Form.Item
						label="备注类别"
						name="remarkType"
						rules={[{ required: true, message: "请选择备注类别" }]}>
						<Select value={remarkType} onChange={onChange}>
							{RemarkTypeArr.map(({ label, value }) => (
								<Option
									value={value}
									key={value}
									disabled={
										(+record.orderStatus >= 8150 || +record.orderStatus <= 1) &&
										value !== RemarkTypeEnum.ORDER_REMARK
									}>
									{label}
								</Option>
							))}
						</Select>
					</Form.Item>
					<Form.Item
						label="备注内容"
						name="content"
						rules={[{ required: true, message: "请输入备注内容" }]}>
						<TextArea placeholder="包裹包含XX敏感品不能进行XX方式运输" />
					</Form.Item>
					{remarkType === RemarkTypeEnum.FRONT_REMARK && (
						<Form.Item label="需回复?" name="isReply">
							<Radio.Group>
								<Radio value={true}>是</Radio>
								<Radio value={false}>否</Radio>
							</Radio.Group>
						</Form.Item>
					)}
					<Form.Item
						label="附件"
						name="enclosureLink"
						extra="支持扩展名：.png .jpg .jpeg 单张不超过5M">
						<UpFile setUpFileLinkArr={setEnclosureLink} />
					</Form.Item>
				</Form>
			</Form.Provider>
		</Modal>
	));
}

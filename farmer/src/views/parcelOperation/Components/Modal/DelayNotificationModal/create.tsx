import * as React from "react";
import { Row, Col, Select, Input, Checkbox } from "antd";
import { observer } from "mobx-react-lite";
import { useModalStores } from "./store/hooks";

interface FormItemProps {
	label: string;
	children: any;
	index?: number;
	hasChecked?: boolean;
}

const styles = require("./index.scss");
/**
 *  批量 添加延误通知
 */

const getChangedValue = (list, value) => {
	const find = list && list.length > 0 && list.find(elem => elem.title === value);
	return find;
};

const getChangedContent = (chosed, value) => {
	const newChoosed = { ...chosed };
	newChoosed.content = value;
	return newChoosed;
};

const CreateDelayNotification = () => {
	const store = useModalStores();
	const {
		changeStoreData,
		noticeTemplateList,
		currentNoticeTemplate,
		mailTemplateList,
		currentMailTemplate,
		mailChecked
	} = store;

	const getFormItem = (formItemprops: FormItemProps) => {
		return (
			<Col span={24} key={formItemprops.index}>
				<div className={styles.formItem}>
					<div className={styles.label}>
						{formItemprops.hasChecked && (
							<Checkbox
								className={styles.checkbx}
								checked={mailChecked}
								onChange={(e: any) =>
									changeStoreData("mailChecked", e.target.checked)
								}
							/>
						)}
						{formItemprops.label}
					</div>
					<div className={styles.children}>{formItemprops.children}</div>
				</div>
			</Col>
		);
	};

	const getFields = () => {
		const fieldArr: FormItemProps[] = [
			{
				label: "通知模板",
				children: (
					<Select
						allowClear
						style={{ width: "100%" }}
						value={currentNoticeTemplate.title}
						onChange={value =>
							changeStoreData(
								"currentNoticeTemplate",
								getChangedValue(noticeTemplateList, value)
							)
						}>
						{noticeTemplateList &&
							noticeTemplateList.length > 0 &&
							noticeTemplateList.map((item, index) => (
								<Select.Option key={index} value={item.title}>
									{item.title}
								</Select.Option>
							))}
					</Select>
				)
			},
			{
				label: "通知内容",
				children: <Input.TextArea value={currentNoticeTemplate.content} disabled />
			},
			{
				label: "邮件模板",
				hasChecked: true,
				children: (
					<Select
						style={{ width: "100%" }}
						value={currentMailTemplate.title}
						onChange={value =>
							changeStoreData(
								"currentMailTemplate",
								getChangedValue(mailTemplateList, value)
							)
						}>
						{mailTemplateList &&
							mailTemplateList.length > 0 &&
							mailTemplateList.map((item, index) => (
								<Select.Option key={index} value={item.title}>
									{item.title}
								</Select.Option>
							))}
					</Select>
				)
			},
			{
				label: "邮件内容",
				children: (
					<Input.TextArea
						value={currentMailTemplate.content}
						onChange={e =>
							changeStoreData(
								"currentMailTemplate",
								getChangedContent(currentMailTemplate, e.target.value)
							)
						}
					/>
				)
			}
		];
		const children = fieldArr.map((elem, index) =>
			getFormItem({
				label: elem.label,
				children: elem.children,
				index,
				hasChecked: elem.hasChecked
			})
		);
		return children;
	};

	return (
		<div className={styles["ant-advanced-search-form"]}>
			<Row gutter={24}>{getFields()}</Row>
		</div>
	);
};

export default observer(CreateDelayNotification);

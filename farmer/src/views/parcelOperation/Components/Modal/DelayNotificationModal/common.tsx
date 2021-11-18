import * as React from "react";
import { Row, Input, Select, Col } from "antd";
import { observer } from "mobx-react-lite";
import { useModalStores } from "./store/hooks";
import { formatUnixTime, TimeFormatToMinite } from "utils/time";

interface FormItemProps {
	label: string;
	children: any;
	index?: number;
}

const styles = require("./index.scss");

/**
 * 单个包裹添加延误通知
 */

const getFormItem = (props: FormItemProps) => {
	return (
		<Col span={24} key={props.index}>
			<div className={styles.formItem}>
				<div className={styles.label}>{props.label}</div>
				<div className={styles.children}>{props.children}</div>
			</div>
		</Col>
	);
};

const getChangedValue = (list, title) => {
	const find = list && list.length > 0 && list.find(elem => elem.title === title);
	return find;
};

const CommonNotification = () => {
	const store = useModalStores();
	const { changeStoreData, noticeTemplateList, currentNoticeTemplate, noticeLogList } = store;

	const getFields = () => {
		const fieldArr: FormItemProps[] = [
			{
				label: "通知模板",
				children: (
					<Select
						value={currentNoticeTemplate.title}
						style={{ width: "100%" }}
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
			}
		];
		const children = fieldArr.map((elem, index) =>
			getFormItem({
				label: elem.label,
				children: elem.children,
				index
			})
		);
		return children;
	};

	const getNoticeLog = () => {
		return (
			<div className={styles.noticeLog}>
				<p className={styles.title}>备注记录</p>
				<div className={styles.content}>
					{noticeLogList.map((elem, idx) => (
						<div key={idx} className={styles.module}>
							<p>
								{elem.createBy}{" "}
								<span>{formatUnixTime(elem.createDate, TimeFormatToMinite)}</span>
							</p>
							<div>{elem.notice}</div>
						</div>
					))}
				</div>
			</div>
		);
	};

	const addNotice = () => {
		return (
			<div className={styles.noticeAdd}>
				<p className={styles.title}>添加通知</p>
				<div className={styles.content}>
					<Row gutter={24}>{getFields()}</Row>
				</div>
			</div>
		);
	};

	return (
		<div>
			{noticeLogList && noticeLogList.length > 0 && getNoticeLog()}
			{addNotice()}
		</div>
	);
};

export default observer(CommonNotification);

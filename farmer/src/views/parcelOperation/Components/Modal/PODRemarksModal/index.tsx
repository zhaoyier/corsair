import * as React from "react";
import { Modal, Input } from "antd";
import { RemarksTypes } from "views/parcelOperation/store/constant";
import { OMSRemark } from "genServices/ezShipOMS/oms";
import { formatUnixTime, TimeFormatToMinite } from "utils/time";

export interface PODRemarksModalProps {
	visible: boolean;
	type: string;
	remarkLogs: any[];
	parcelList: any[];
	selectedRows: any[];
	selectedRowKeys: string[];
	onCancel?: () => void;
	onOk?: (remark) => void;
}
// 包裹备注  、 配送备注    弹窗

const styles = require("./index.scss");

const getTitle = (props: PODRemarksModalProps) => {
	const { type, selectedRowKeys, selectedRows } = props;
	// 配送备注，且选择的是单个包裹 --- 显示配送单号
	const top =
		type === "2" && selectedRowKeys && selectedRowKeys.length === 1
			? `配送单号：${selectedRows[0].deliveryId || selectedRows[0].shipmentId}`
			: RemarksTypes[type];
	return (
		<div>
			{top} &emsp;&emsp;
			<span className={styles.tip}>仓库仅展示最新的一条{RemarksTypes[type]}&emsp;</span>
		</div>
	);
};

const getParcelNum = parcelList => {
	return (
		<section className={styles.pkgNo}>
			包裹号：{" "}
			{parcelList &&
				parcelList.length > 0 &&
				parcelList.map((item, idx) => (
					<a
						key={idx}
						rel="opener"
						href={`/parcelOperation.html#/detail?packageCode=${item}`}
						target="_blank">
						{item}&emsp;
					</a>
				))}
		</section>
	);
};

const getRemarkLog = (remarkLogs: OMSRemark[], type) => {
	const deliveryMaxLogs = remarkLogs.slice(0, 5); // 配送备注记录 最多 5条
	const list = type === "1" ? remarkLogs : deliveryMaxLogs;
	return (
		<div className={styles.remarkLog}>
			<p className={styles.title}>备注记录</p>
			<div className={styles.content}>
				{list.map((elem, idx) => (
					<div className={styles.module} key={idx}>
						<p>
							{elem.updateBy}
							<span>{formatUnixTime(elem.updateDate, TimeFormatToMinite)}</span>
						</p>
						<div>{elem.remark}</div>
					</div>
				))}
			</div>
		</div>
	);
};

const addRemark = (remark, setRemark) => {
	return (
		<div className={styles.remarkLog}>
			<p className={styles.title}>添加备注</p>
			<div className={styles.content}>
				<Input.TextArea
					autoSize={{ minRows: 3, maxRows: 5 }}
					className={styles.textArea}
					value={remark}
					onChange={e => setRemark(e.target.value)}
				/>
			</div>
		</div>
	);
};

const PODRemarksModal = (props: PODRemarksModalProps) => {
	const { visible, type, selectedRowKeys, onCancel, onOk, remarkLogs, parcelList } = props;
	const title = getTitle(props);
	const [remark, setRemark] = React.useState("");
	const showLog =
		selectedRowKeys && selectedRowKeys.length === 1 && remarkLogs && remarkLogs.length > 0;
	const showParcelNum = type === "2" && selectedRowKeys && selectedRowKeys.length === 1;
	return (
		<Modal
			visible={visible}
			title={title}
			width={900}
			onCancel={onCancel}
			onOk={() => onOk(remark)}
			okText="确定"
			cancelText="关闭"
			afterClose={() => setRemark("")}>
			{/** 包裹备注：1.单条包裹： 有备注记录就显示 2.多条包裹： 不显示备注记录 */}
			{/** 配送备注：1.单条包裹： 显示配送单号，和配送单号下的包裹号，备注记录（有最多显示5条，没有不显示） 2.多条包裹：不显示备注记录，包裹号，配送单号 */}
			{showParcelNum && getParcelNum(parcelList)}
			{showLog && getRemarkLog(remarkLogs, type)}
			{addRemark(remark, setRemark)}
		</Modal>
	);
};

export default PODRemarksModal;

import * as React from "react";
import { Modal, message } from "antd";
import { CancelData } from "views/parcelOperation/store/types";
import { DeliveryPackage, DeliveryOrder } from "genServices/ezShipOMS/oms";
import ParcelTable from "../../Common/ParcelTable";
import NotOutboundOrderTable from "../../Common/NotOutboundOrderTable";

export interface SplitDeliveryModalProps {
	visible: boolean;
	record: any; // 当前包裹信息
	parcelList: DeliveryPackage[]; // 配送单下包裹信息
	orders: DeliveryOrder[];
	refresh?: () => void;
	onCancel?: () => void;
	onOk?: (data: CancelData, cb?: (result, refresh?: () => void) => void) => void;
}

const styles = require("./index.scss");

// 取消配送（强制取消包裹）  弹窗

const getTitle = (deliveryId: string, orders) => {
	return (
		<div className={styles.title}>
			<span className={styles.span}>拆分配送：单号{deliveryId}</span>
			{orders && orders.length > 0 && (
				<span className={styles.tips}>该配送单还有未出库的订单</span>
			)}
		</div>
	);
};

const SplitDeliveryModal = (props: SplitDeliveryModalProps) => {
	const { visible, record, parcelList, onCancel, onOk, orders, refresh } = props;
	const { deliveryId, shipmentId } = record;

	const [selectedRowKeys, setSelectRowKeys] = React.useState([]);
	// const [resVisible, setResVisible] = React.useState(false);
	const [result, setResult] = React.useState(null);

	const getSplitData = () => {
		return { deliveryId: deliveryId || shipmentId, parcelNumList: selectedRowKeys };
	};

	const closeAll = () => {
		setResult(null);
		onCancel();
		if (refresh) {
			refresh();
		}
	};

	const viewMore = (parcelCodes: string[]) => {
		if (parcelCodes && parcelCodes.length > 0) {
			window.open(`/parcelOperation.html#/detail?packageCode=${parcelCodes[0]}`, "_blank");
		} else {
			message.error("包裹号不存在");
		}
	};

	const SuccessModal = ({ result }) => {
		return (
			<Modal
				visible={result}
				title="操作成功"
				okText="查看"
				cancelText="取消"
				onCancel={() => closeAll()}
				onOk={() => viewMore(result.parcelCodes)}>
				<section className={styles.resultModal}>
					<p>
						生成了一笔新的配送单，配送单号为：<span>{result.shipmentId}</span>&emsp;
					</p>
					<p>当前配送单包含的包裹号有：</p>
					<p>
						<span>
							{result.parcelCodes &&
								result.parcelCodes.length > 0 &&
								result.parcelCodes.map((elem, idx) => (
									<a
										key={idx}
										rel="opener"
										href={`/parcelOperation.html#/detail?packageCode=${elem}`}
										target="_blank">
										{elem}&nbsp;
									</a>
								))}
						</span>
					</p>
				</section>
			</Modal>
		);
	};

	const title = getTitle(deliveryId || shipmentId, orders);
	// 拆分配送仅允许勾选包裹状态为【派送仓已上架 14600】【站点已上架 20700】【封箱出库 9000】的包裹，除此之外状态的包裹无勾选框
	// 当拆分配送的配送单并不完整时，可以拆分，但需要显示提示语“该配送单下还有订单未打包出库”。
	// 送单仅有多个包裹的情况，才可以拆
	//  https://www.tapd.cn/22347521/prong/stories/view/1122347521001079249  -- 修改允许勾选包裹状态
	return (
		<div>
			<Modal
				visible={visible}
				title={title}
				onCancel={onCancel}
				width={900}
				okButtonProps={{
					disabled: selectedRowKeys.length === 0 // 未选择包裹时不允许  确定拆分
				}}
				onOk={() =>
					onOk(getSplitData(), result => {
						setResult(result);
					})
				}
				afterClose={() => setSelectRowKeys([])}>
				<ParcelTable
					dataSource={parcelList}
					showCheckBox={true}
					showHeaderCheckBox={true}
					selectedRowKeys={selectedRowKeys}
					getCheckboxProps={record => ({
						disabled:
							String(record.packageStatus) !== "17100" &&
							String(record.packageStatus) !== "14600" &&
							String(record.packageStatus) !== "15000" &&
							String(record.packageStatus) !== "9000"
					})}
					setSelectRowKeys={selectedRowKeys => setSelectRowKeys(selectedRowKeys)}
				/>
				{orders && orders.length > 0 && <NotOutboundOrderTable dataSource={orders} />}
			</Modal>
			{result && <SuccessModal result={result} />}
		</div>
	);
};

export default SplitDeliveryModal;

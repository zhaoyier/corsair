import * as React from "react";
import { Modal, Checkbox } from "antd";
import { CancelData } from "views/parcelOperation/store/types";
import { DeliveryPackage, DeliveryOrder } from "genServices/ezShipOMS/oms";
import ParcelTable from "../../Common/ParcelTable";

export interface CancelParcelModalProps {
	visible: boolean;
	record: any; // 当前包裹信息
	parcelList: DeliveryPackage[]; // 配送单下包裹信息
	orders: DeliveryOrder[]; // 配送单下未出库订单
	refresh?: () => void;
	onCancel?: () => void;
	onOk?: (data: CancelData, cb?: (result, refresh?: () => void) => void) => void;
}

const styles = require("./index.scss");

// 取消配送（强制取消包裹）  弹窗

const getTitle = (deliveryId: string) => {
	return (
		<div className={styles.title}>
			<span className={styles.span}>取消配送({deliveryId})：</span>
			<span className={styles.tips}>注意！取消配送单仅修改包裹、订单状态，不涉及退款</span>
		</div>
	);
};

const CancelParcelModal = (props: CancelParcelModalProps) => {
	const { visible, record, parcelList, onCancel, onOk, refresh, orders } = props;
	const { deliveryId } = record;

	const [selectedRowKeys, setSelectRowKeys] = React.useState([]);

	const getFilterList = () => {
		const arr = [];
		parcelList.map(item => {
			if (String(item.packageStatus) === "9000" || String(item.packageStatus) === "14600") {
				arr.push(item);
			}
		});
		return arr;
	};
	const filterList = getFilterList();
	const title = getTitle(deliveryId);
	const hasOrder = orders && orders.length > 0;
	const cnacelDeliveryDisabled = hasOrder || filterList.length !== parcelList.length;
	const isCancelAll =
		!cnacelDeliveryDisabled &&
		filterList.length &&
		selectedRowKeys.length === filterList.length;

	const getCancelData = () => {
		return { deliveryId, parcelNumList: selectedRowKeys };
	};

	const successRes = result => {
		Modal.success({
			title: "操作成功",
			content: (
				<section className={styles.resultModal}>
					{isCancelAll ? (
						<p>
							配送单号：<span>{result.deliveryNumber}</span>&emsp;取消成功！
						</p>
					) : (
						<div>
							以下包裹取消成功：
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
						</div>
					)}
				</section>
			),
			okText: "关闭",
			onOk: () => {
				onCancel();
				refresh();
			}
		});
	};

	// 点击取消配送单号
	const cancelDeliveryNum = (checked: boolean, filterList) => {
		let selectedRowKeys = [];
		if (checked) {
			selectedRowKeys = filterList.length > 0 ? filterList.map(elem => elem.packageNum) : [];
		} else {
			selectedRowKeys = [];
		}
		setSelectRowKeys(selectedRowKeys);
	};

	// 1 --- 取消配送单
	// 1.	当配送单下还有订单未打包出库时，即配送单不完整，不可以取消配送单，禁用【取消整个配送单按钮】，并显示提示语“该配送单下还有订单未打包出库”。
	// 2. 只支持配送单下所有包裹状态都在【称重出库 9000】、【派送仓已上架 14600】的配送单才可支持取消配送单。
	// 3. 当选中【取消整个配送单】时，勾选配送单下所有可取消包裹。反选【取消整个配送单】按钮时，移除所有可取消包裹的选中。
	// 2 ---- 取消部分包裹
	// 	只支持在【称重出库】、【派送仓已上架】的包裹才可支持取消包裹，且配送单未锁定。
	// 当勾选所有包裹时，默认勾选【取消整个配送单】按钮。即按取消整个配送单处理，此时反选其中某个包裹时，取消【取消整个配送单】按钮的选中。
	// 对于选中的包裹进行包裹取消操作通知ALS 进行包裹取消。
	// 取消之后重新进行完整逻辑判断，通知ALS。
	// 取消包裹之后取消相应的订单
	return (
		<Modal
			visible={visible}
			title={title}
			onCancel={onCancel}
			width={900}
			onOk={() => onOk(getCancelData(), result => successRes(result))}
			afterClose={() => setSelectRowKeys([])}>
			<div className={styles.chooseAll}>
				<Checkbox
					disabled={cnacelDeliveryDisabled}
					onClick={(e: any) => cancelDeliveryNum(e.target.checked, filterList)}
					checked={isCancelAll}>
					{" "}
					取消整个配送单
				</Checkbox>
				&emsp;
				{hasOrder && <span className={styles.tips}>该配送单下还有订单未打包出库！</span>}
			</div>
			<ParcelTable
				dataSource={parcelList}
				hiddeHeader={true}
				showCheckBox={true}
				selectedRowKeys={selectedRowKeys}
				getCheckboxProps={record => ({
					disabled:
						String(record.packageStatus) !== "9000" &&
						String(record.packageStatus) !== "14600"
				})}
				setSelectRowKeys={selectedRowKeys => setSelectRowKeys(selectedRowKeys)}
			/>
		</Modal>
	);
};

export default CancelParcelModal;

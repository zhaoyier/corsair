import * as React from "react";
import { Modal, Row, Col, Table, InputNumber, notification } from "antd";
import { useObserver } from "mobx-react-lite";
import { store } from "../../store/helper/useStore";
import { toJS } from "mobx";
import { formatPrice, limitDecimals, priceToInt } from "utils/util";

interface ChangeMoneyModalProps {
	handleClose: () => void;
	showChangeMoneyModal: boolean;
}
export function ChangeMoneyModal(props: ChangeMoneyModalProps) {
	const { handleClose, showChangeMoneyModal } = props;
	const orderStore = store().orderStore;
	const [isEdit, setIsEdit] = React.useState(false);
	const [list, setList] = React.useState(toJS(orderStore.declaredAmount) || []);

	const changePrice = (value: number, index: number) => {
		setIsEdit(true);
		const newList = [...list];
		newList[index].declaredValue = priceToInt(value);
		setList(newList);
	};

	const columns = [
		{
			title: "序号",
			dataIndex: "orderItemId"
		},
		{
			title: "一级报关品类",
			dataIndex: "fstCategoryName"
		},
		{
			title: "二级报关品类",
			dataIndex: "secCategoryName"
		},
		{
			title: "单价",
			width: 200,
			dataIndex: "declaredValue",
			render: (text, r, i) => {
				console.log(r);
				const value = Number(formatPrice(text));
				return (
					<InputNumber
						min={0}
						value={value}
						onChange={value => changePrice(Number(value), i)}
						style={{ width: 150 }}
						precision={2}
						max={99999.99}
						step={0.01}
						parser={value => limitDecimals(value) as any}
					/>
				);
			}
		},
		{
			title: "币种",
			dataIndex: "currencyCode"
		},
		{
			title: "数量",
			dataIndex: "qty"
		}
		// {
		// 	title: "修改时间",
		// 	dataIndex: "updateDate",
		// 	render: text => formatUnixTime(text)
		// },
		// {
		// 	title: "更新作者",
		// 	dataIndex: "updateBy"
		// }
	];

	async function submitDeclaredAmount() {
		const findError =
			list &&
			list.length > 0 &&
			list.find(elem => Number(formatPrice(elem.declaredValue)) < 0.01);
		if (findError) {
			notification.error({
				message: "商品单价的有效数据范围：0.01 ~ 99999.99"
			});
			return;
		}
		const declaredAmountArr = list.map(elem => Number(elem.declaredValue));
		const reducer = (accumulator, currentValue) => accumulator + currentValue;
		const declaredAmount: any = declaredAmountArr.reduce(reducer);
		const itemInfos = list.map(elem => {
			return {
				itemID: elem.orderItemId,
				declaredValue: Number(elem.declaredValue),
				qty: elem.qty
			} as any;
		});
		console.log(declaredAmountArr, declaredAmount);
		await orderStore.modifyDeclaredAmount({
			orderId: orderStore.getRouterInfo().orderId,
			declaredAmount,
			itemInfos
		});
		await orderStore.initOrderDetail({ orderId: orderStore.orderId });
		handleClose();
	}

	return useObserver(() => (
		<Modal
			title={
				<React.Fragment>
					<span>修改申报金额</span>
					<span style={{ color: "#df2c42", fontSize: "12px", marginLeft: 4 }}>
						注意！ 申报金额为目的地国家币种!
					</span>
				</React.Fragment>
			}
			visible={showChangeMoneyModal}
			okText="确定"
			cancelText="取消"
			width={700}
			okButtonProps={{ disabled: !isEdit }}
			onOk={submitDeclaredAmount}
			onCancel={handleClose}>
			{/* <Row>
				<Col span={8} style={{ lineHeight: "32px" }}>
					申报总金额(当地币):
				</Col>
				<Col span={16}>
					<InputNumber
						min={0}
						onChange={e => setDeclaredAmount(e + "")}
						style={{ width: 315 }}
						precision={2}
						max={99999.99}
						step={0.01}
						parser={value => limitDecimals(value)}
					/>
				</Col>
			</Row> */}
			<Row style={{ marginTop: 20 }}>
				<Col span={24}>
					<Table
						columns={columns}
						dataSource={list || []}
						// dataSource={toJS(orderStore.declaredAmount)}
						pagination={false}
						rowKey="orderItemId"
						// rowKey={record => record.orderId}
						style={{ width: "100%" }}
					/>
				</Col>
			</Row>
		</Modal>
	));
}

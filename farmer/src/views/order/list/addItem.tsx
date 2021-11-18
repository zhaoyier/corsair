import * as React from "react";
import { Col, Row, Divider, Select, message, InputNumber } from "antd";
import { store } from "../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
const { Option } = Select;
import { ItemType } from "./item";
import { checkEmpty, limitInteger } from "utils/util";
export interface AddItemProps {
	item?: any;
	addItem: (item: ItemType) => void;
	handleClose: () => void;
}
export function AddItem(props: AddItemProps) {
	const [categoryId, setCategoryId] = React.useState<string>("");
	const [category, setCategory] = React.useState<string>("");
	const [qty, setQty] = React.useState<string>("");
	const { handleClose, addItem } = props;
	const commonStore = store().commonStore;
	function saveItem() {
		if (!categoryId || !category || checkEmpty(qty)) {
			message.error("请填入数据");
			return;
		}
		console.log(qty, typeof qty);
		if (checkEmpty(qty)) {
			message.error("请输入数字");
			return;
		}
		// console.log("保存item");
		addItem({ categoryId, category, qty });
		handleClose();
	}
	function cancelSave() {
		// console.log("取消");
		handleClose();
	}
	function handleSelect(value) {
		setCategoryId(value);
		setCategory(
			commonStore.cnHscodeList &&
				commonStore.cnHscodeList.CnHscodeList.find(i => i.id === value).name
		);
	}
	function handleInput(e) {
		setQty(e + "");
	}
	return useObserver(() => (
		<Row justify="center" style={{ margin: "-8px 0px" }}>
			<Divider />
			<Col span={12} style={{ textAlign: "center" }}>
				<Select
					allowClear
					showSearch
					placeholder="报关品类"
					style={{ width: 300 }}
					onChange={handleSelect}
					filterOption={(inputValue, option) =>
						option.props.children.toLowerCase().indexOf(inputValue.toLowerCase()) >= 0
					}>
					{(
						(commonStore.cnHscodeList && commonStore.cnHscodeList.CnHscodeList) ||
						[]
					).map(({ id, name }, index) => (
						<Option key={id + index} value={id}>
							{name}
						</Option>
					))}
				</Select>
			</Col>
			<Col span={8} style={{ textAlign: "center" }}>
				<InputNumber
					onChange={handleInput}
					placeholder="数量"
					max={9999}
					min={0}
					step={1}
					precision={0}
					parser={value => limitInteger(value)}
				/>
			</Col>
			<Col span={4}>
				<a onClick={saveItem}>保存</a>
				<Divider type="vertical" />
				<a onClick={cancelSave}>取消</a>
			</Col>
		</Row>
	));
}

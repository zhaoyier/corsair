import * as React from "react";
import { Col, Row, Select, InputNumber, Button } from "antd";
import { store } from "../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { checkEmpty, limitInteger } from "utils/util";

const { Option } = Select;
export interface ItemType {
	categoryId: string;
	category: string;
	qty: string;
}
export interface ItemProps {
	item: ItemType;
	index: number;
	delItem: (id: any) => void;
	changeItem: (item: ItemType, index: number) => void;
}
export function Item(props: ItemProps) {
	const [editItemId, setEditItemId] = React.useState<string>("");
	const [editItemName, setEditItemName] = React.useState<string>("");
	const [editItemNumber, setEditItemNumber] = React.useState<string>("");
	const commonStore = store().commonStore;
	const {
		// item: { category, categoryId, qty },
		delItem,
		changeItem,
		index
	} = props;
	function changeItemNameEdit(value) {
		setEditItemId(value);
		const editItemName =
			commonStore.cnHscodeList &&
			commonStore.cnHscodeList.CnHscodeList.find(i => i.id === value).name;
		setEditItemName(editItemName);
		changeItem({ categoryId: value, category: editItemName, qty: editItemNumber }, index);
	}

	function changeItemNumberEdit(e) {
		setEditItemNumber(e + "");
		changeItem({ categoryId: editItemId, category: editItemName, qty: e + "" }, index);
	}

	return useObserver(() => (
		<React.Fragment>
			<Row justify="center" style={{ padding: "10px 0px" }}>
				<Col span={17} style={{ textAlign: "center" }}>
					<Select
						allowClear
						showSearch
						placeholder="报关品类"
						value={editItemName}
						style={{ width: "90%" }}
						filterOption={(inputValue, option) =>
							option.props.children.toLowerCase().indexOf(inputValue.toLowerCase()) >=
							0
						}
						onChange={changeItemNameEdit}>
						{commonStore.cnHscodeList &&
							commonStore.cnHscodeList.CnHscodeList &&
							commonStore.cnHscodeList.CnHscodeList.length > 0 &&
							(commonStore.cnHscodeList.CnHscodeList || []).map(
								({ id, name }, index) => (
									<Option key={id + index} value={id}>
										{name}
									</Option>
								)
							)}
					</Select>
				</Col>
				<Col span={3} style={{ textAlign: "center" }}>
					<InputNumber
						value={!checkEmpty(editItemNumber) ? +editItemNumber : null}
						onChange={changeItemNumberEdit}
						placeholder="数量"
						max={9999}
						min={0}
						step={1}
						precision={0}
						parser={value => limitInteger(value)}
					/>
				</Col>
				<Col span={4} style={{ textAlign: "right" }}>
					<Button style={{ color: "#f00" }} onClick={() => delItem(index)}>
						删除
					</Button>
				</Col>
			</Row>
			{/* )} */}
		</React.Fragment>
	));
}

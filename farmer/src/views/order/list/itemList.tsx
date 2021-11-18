import * as React from "react";
import { Col, Form, Row } from "antd";
const styles = require("./index.scss");
import { Item, ItemType } from "./item";
import { store } from "../store/helper/useStore";
import { useObserver } from "mobx-react-lite";

export interface ItemListProps {
	save: (item: ItemType[]) => void;
}

export function ItemList(props: ItemListProps) {
	const initItem: ItemType = {
		categoryId: null,
		category: null,
		qty: null
	};
	const commonStore = store().commonStore;
	const [itemList, setItemList] = React.useState<ItemType[]>([]);

	function delItem(index) {
		itemList.splice(index, 1);
		setItemList([...itemList]);
		props.save(itemList);
	}

	function changedItem(item: ItemType, index: number) {
		itemList[index] = item;
		setItemList([...itemList]);
		props.save(itemList);
	}

	function addItem() {
		itemList.push({ ...initItem });
		setItemList([...itemList]);
		props.save(itemList);
	}

	React.useEffect(() => {
		commonStore.fetchCnHscodeList();
		addItem();
	}, []);

	return useObserver(() => (
		<Row justify="center" style={{ marginTop: 10 }}>
			<Col span={22}>
				<Form.Item noStyle name="transferNumber">
					<div>
						<div className={styles.itemList} onClick={() => addItem()}>
							+ 新增item
						</div>
						{itemList &&
							itemList.length > 0 &&
							itemList.map((it, index) => (
								<Item
									key={index}
									index={index}
									item={it}
									delItem={delItem}
									changeItem={changedItem}
								/>
							))}
					</div>
				</Form.Item>
			</Col>
		</Row>
	));
}

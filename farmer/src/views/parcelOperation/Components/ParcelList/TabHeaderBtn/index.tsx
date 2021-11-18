import * as React from "react";
import { useStores } from "views/parcelOperation/hooks";
import { observer } from "mobx-react-lite";
import { Button, Pagination } from "antd";

const styles = require("./index.scss");

const TabHeaderBtn = () => {
	const { modalStore, parcelListStore } = useStores();
	const { clickPODRemark, clickDelayNotification } = modalStore;
	const { selectedRows, selectedRowKeys, pageParam, changePage } = parcelListStore;
	const { total, limit, current } = pageParam;

	return (
		<section className={styles.tabHeaderBtns}>
			<section className={styles.leftBtns}>
				<Button
					type="primary"
					className={styles.btn}
					disabled={selectedRowKeys.length === 0}
					onClick={() => clickPODRemark("1", selectedRowKeys, selectedRows)}>
					包裹备注
				</Button>
				<Button
					type="primary"
					className={styles.btn}
					disabled={selectedRowKeys.length === 0}
					onClick={() => clickPODRemark("2", selectedRowKeys, selectedRows)}>
					配送备注
				</Button>
				<Button
					type="primary"
					className={styles.btn}
					disabled={selectedRowKeys.length === 0}
					onClick={() => clickDelayNotification(selectedRowKeys, selectedRows)}>
					延误通知
				</Button>
			</section>

			<section className={styles.paginate}>
				<Pagination
					current={current}
					total={Number(total)}
					pageSize={limit}
					showQuickJumper={true}
					showSizeChanger={true}
					showTotal={total => `共${total}条`}
					onChange={(current, pageSize) => changePage(current, pageSize)}
					onShowSizeChange={(current, pageSize) => changePage(current, pageSize)}
					pageSizeOptions={["50", "100", "150", "200"]}
				/>
			</section>
		</section>
	);
};

export default observer(TabHeaderBtn);

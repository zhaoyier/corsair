import * as React from "react";
import { Table, Button, Spin, Popover } from "antd";
import { useStores } from "views/parcelOperation/hooks";
import { observer } from "mobx-react-lite";
import { toJS } from "mobx";
import Loadable from "react-loadable";
import { ColumnsType, ColumnProps } from "antd/lib/table";
import ModaltStore from "views/parcelOperation/store/modal";
import { ChildTabStatus } from "views/parcelOperation/store/constant";
import { MyLoadingComponent } from "../../Common/Loading";
import { formatWeight } from "utils/util";
import { getTimeUnixDateStr } from "utils/time";
import { PackageDetail } from "genServices/ezShipOMS/oms";
import { OrderOriginEnum } from "genServices/ezShipOMS/public";
import { DownOutlined } from "@ant-design/icons";

const styles = require("./index.scss");

const ChildTab = Loadable({
	loader: () => import("../ChildTab"),
	loading: MyLoadingComponent
});

const tableComponents = {
	body: {
		row: props => {
			const classArr = [
				"ant-table-row ant-table-row-level-0",
				"ant-table-row ant-table-row-level-0 ant-table-row-selected"
			];
			if (classArr.find(elem => elem === props.className)) {
				const newProps = { ...props };
				const botRow = { ...props };
				const find = props.children.find(elme => elme.key === "deliveryAddress");
				const filter = props.children.filter(elme => elme.key !== "deliveryAddress");
				newProps.children = [...filter];
				botRow.children = [find];
				return (
					<React.Fragment>
						<tr {...newProps} />
						<tr className={styles.stationAddress}>{find}</tr>
					</React.Fragment>
				);
			}
			return <tr {...props} />;
		},
		cell: props => {
			if (props.className === "ant-table-cell subTd") {
				const newProps = { ...props };
				newProps.colSpan = 13;
				return <td {...newProps} />;
			} else {
				return <td {...props} />;
			}
		}
	}
};

const renderAction = (t, r: PackageDetail, store: ModaltStore) => {
	const {
		clickComplementSend,
		clickCancelParcel,
		clickSplitDelivery,
		clickEditDeliveryInfo,
		clickPrintLabel
	} = store;
	const isGlobal =
		r.packageOrigin &&
		r.packageOrigin.code === OrderOriginEnum.OrderOriginEnumEzshipGlobal;
	const disabled = isGlobal || r.regionCode === "MY";
	return (
		<div>
			<span style={{ display: "none" }}>{t}</span>
			<div style={{ display: "flex" }}>
				<Button
					type="link"
					disabled={disabled}
					onClick={() => clickPrintLabel(r)}>
					打印标签
				</Button>
				<Button
					disabled={disabled}
					type="link"
					onClick={() => clickComplementSend(r)}>
					补送
				</Button>
				<Button
					disabled={disabled}
					type="link"
					onClick={() => clickSplitDelivery(r)}>
					拆分
				</Button>
			</div>
			<div style={{ display: "flex" }}>
				<Button
					type="link"
					disabled={disabled}
					onClick={() => clickEditDeliveryInfo(r)}>
					修改配送信息
				</Button>
				<Button type="link" onClick={() => clickCancelParcel(r)}>
					强制取消包裹
				</Button>
			</div>
		</div>
	);
}

const getColumns = (store: ModaltStore): ColumnProps<PackageDetail>[] => {
	const columns = [

		{
			title: "操作",
			dataIndex: "action",
			align: "center",
			width: 80,
			render: (t, r: PackageDetail) => 
			<Popover placement="bottom"  zIndex={90} trigger="click" content={renderAction(t, r, store)}>
				<a>操作&nbsp;<DownOutlined /></a>
			</Popover>
		},
		{
			title: "包裹号",
			dataIndex: "packageCode",
			width: 160,
			render: (t, r) => (
				<div style={{display: "flex"}}>
					<a
						onClick={() =>
							window.open(`/parcelOperation.html#/detail?packageCode=${t}`, "_blank")
						}>
						{t}
					</a>
					<div className={styles.iconRow}>
						{r.isInsurance && <img src={require("../../../images/ic_baoxian.png")} />}
						{r.isAftersale && <img src={require("../../../images/ic_fix.png")} />}
						{r.isLocked && <img src={require("../../../images/ic_lock.png")} />}
					</div>
				</div>
			)
		},
		{
			title: "配送单号",
			dataIndex: "deliveryId",
			width: 150,
			render: (t, r) => <a onClick={() => store.clickDeliveryNum(r)}>{t}</a>
		},
		// {
		// 	title: "包裹来源",
		// 	dataIndex: "packageOrigin",
		// 	align: "center",
		// 	width: 100,
		// 	render: t => <span>{t ? t.name : null}</span>
		// },
		{
			title: "发货仓库",
			dataIndex: "warehouse",
			align: "center",
			width: 100
		},
		{
			title: "国家",
			dataIndex: "regionCode",
			align: "center",
			width: 100
		},
		{
			title: "运输方式",
			dataIndex: "transport",
			align: "center",
			width: 100
		},
		{
			title: "取货方式",
			dataIndex: "deliveryType",
			align: "center",
			width: 150
		},
		{
			title: "会员",
			dataIndex: "nickname",
			align: "center",
			width: 100
		},
		// {
		// 	title: "实际重量(kg)",
		// 	dataIndex: "actualWeight",
		// 	align: "center",
		// 	width: 100,
		// 	render: val => formatWeight(val)
		// },
		{
			title: "计费重(kg)",
			dataIndex: "chargWeight",
			align: "center",
			width: 100,
			render: val => formatWeight(val)
		},
		{
			title: "发货ETA",
			dataIndex: "etaDate",
			align: "center",
			width: 150,
			defaultSortOrder: "descend",
			sorter: (a, b) =>
				Number(getTimeUnixDateStr(a.etaDate)) - Number(getTimeUnixDateStr(b.etaDate))
		},
		{
			title: "包裹状态",
			dataIndex: "status",
			align: "center",
			width: 100
		},
		{
			title: "取货地址",
			dataIndex: "deliveryAddress",
			width: 20,
			className: "subTd",
			render: value => <span>取货地址:&nbsp;{value}</span>
		}
	];

	return columns as any;
};

const TabTable = () => {
	const { parcelListStore, modalStore } = useStores();
	const {
		parcelList,
		selectedRowKeys,
		changeStoreData,
		listLoading,
		getPackageLog,
		getTotalRemarks,
		changeParcelTable
	} = parcelListStore;
	const [expandedRowKeys, setExpandedRowKeys] = React.useState([]);

	React.useEffect(() => {
		setExpandedRowKeys([]);
	}, [listLoading]);

	const expandedRowRender = record => {
		const index = parcelList.findIndex(elem => elem.packageCode === record.packageCode);
		return (
			<Spin spinning={record.isLoading}>
				<ChildTab index={index} />
			</Spin>
		);
	};

	const onExpand = (expanded, record) => {
		const newExpandedRowKeys = [...expandedRowKeys];
		const index = parcelList.findIndex(elem => elem.packageCode === record.packageCode);
		if (expanded) {
			getPackageLog(record.packageCode, index);
			getTotalRemarks(record.packageCode, record.deliveryId, index);
			newExpandedRowKeys.push(record.packageCode);
		} else {
			const findIndex = newExpandedRowKeys.findIndex(elem => elem === record.packageCode);
			if (findIndex >= 0) {
				newExpandedRowKeys.splice(findIndex, 1);
			}
		}
		changeParcelTable(index, "childTabActiveKey", Object.keys(ChildTabStatus)[0]);
		setExpandedRowKeys(newExpandedRowKeys);
	};

	const expandable = {
		expandedRowRender,
		expandedRowKeys,
		onExpand
	};

	const getRowSelection = (selectedRowKeys: string[]) => {
		const rowSelection = {
			selectedRowKeys,
			onChange: (selectedRowKeys, selectedRows) => {
				changeStoreData("selectedRowKeys", selectedRowKeys);
				changeStoreData("selectedRows", selectedRows);
				console.log(`selectedRowKeys: ${selectedRowKeys}`, "selectedRows: ", selectedRows);
			}
		};
		return rowSelection;
	};

	return (
		<Spin tip="Loading..." spinning={listLoading}>
			<Table
				size="small"
				rowKey="packageCode"
				className={styles.columns}
				scroll={{ x: 1500, y: 500 }}
				columns={getColumns(modalStore) as ColumnsType}
				rowSelection={getRowSelection(selectedRowKeys)}
				dataSource={toJS(parcelList)}
				expandable={expandable}
				components={tableComponents}
				pagination={false}
			/>
		</Spin>
	);
};

export default observer(TabTable);

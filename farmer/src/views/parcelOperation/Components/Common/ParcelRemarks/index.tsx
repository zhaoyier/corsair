import * as React from "react";
import { Table } from "antd";
import { OMSRemark } from "genServices/ezShipOMS/oms";
import { IsRemarkList } from "views/parcelOperation/store/constant";
import { formatUnixTime } from "utils/time";
interface ParcelRemarksProps {
	dataSource: OMSRemark[];
}

const getColumns = () => {
	const columns = [
		{ title: "创建时间", dataIndex: "updateDate", render: t => formatUnixTime(t) },
		{
			title: "内容说明",
			dataIndex: "remark",
			// width: 400,
			render: t => (
				<div
					style={{
						width: "92%",
						whiteSpace: "pre-wrap",
						wordWrap: "break-word",
						wordBreak: "break-all"
					}}>
					{t}
				</div>
			)
		},
		{ title: "备注类型", dataIndex: "type", render: t => IsRemarkList[t] },
		{ title: "修改人", dataIndex: "updateBy" }
	];
	return columns;
};

const ParcelRemarks = (props: ParcelRemarksProps) => {
	const { dataSource } = props;
	return <Table style={{ marginLeft: 0 }} columns={getColumns()} dataSource={dataSource || []} />;
};

export default ParcelRemarks;

import * as React from "react";
import { Button, Modal, Table } from "antd";
import { store } from "../../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { AddedService, AddedServiceStatusEnum } from "genServices/ezShipOMS/oms";
import { formatUnixTime } from "utils/time";
import { addServiceTypeArr, addedServiceStatusEnum } from "views/order/constant";

interface AddServerModalProps {
	handleClose: () => void;
}

export function CatAddServerModal(props: AddServerModalProps) {
	const columns = [
		{
			title: "附加服务",
			dataIndex: "serviceType",
			render: text =>
				addServiceTypeArr.find(i => i.value === text) &&
				addServiceTypeArr.find(i => i.value === text).label
		},
		{
			title: "创建时间",
			dataIndex: "createDate",
			render: text => formatUnixTime(text)
		},
		{
			title: "创建人",
			dataIndex: "createBy"
		},
		{
			title: "备注信息",
			dataIndex: "remark"
		},
		{
			title: "状态",
			dataIndex: "status",
			render: text =>
				addedServiceStatusEnum.find(i => i.value === text) &&
				addedServiceStatusEnum.find(i => i.value === text).label
		},
		{
			title: "操作",
			dataIndex: "action",
			render: (...rest) => renderActionButton(rest[1])
		}
	];
	const orderStore = store().orderStore;
	// orderStore.delivery.packages
	const { handleClose } = props;
	const [addServerDate, setAddServerDate] = React.useState<AddedService[]>([]);
	// orderStore.cancelAddedService 删除增值服务
	async function initData() {
		const result = await orderStore.orderAddedService({ orderId: orderStore.orderId });
		setAddServerDate(result.addedServicList);
	}
	function downloadImageFromUrl(url: string) {
		fetch(url, {
			method: "GET",
			headers: {}
		})
			.then(res => {
				res.arrayBuffer().then(function(buffer) {
					const url = window.URL.createObjectURL(new Blob([buffer]));
					const link = document.createElement("a");
					link.href = url;
					link.setAttribute("download", "image.png");
					document.body.appendChild(link);
					link.click();
				});
			})
			.catch(err => {
				console.log(err);
			});
	}
	function renderActionButton(record) {
		if (record.enclosureLink) {
			return (
				<a onClick={() => record.enclosureLink.map(i => downloadImageFromUrl(i))}>
					下载附件
				</a>
			);
		} else if (record.status === AddedServiceStatusEnum.Pending_Status) {
			return (
				<a
					onClick={async () => {
						await orderStore.cancelAddedService({
							orderId: orderStore.orderId,
							serviceId: record.id
						});
						initData();
					}}>
					删除
				</a>
			);
		} else {
			return null;
		}
	}
	React.useEffect(() => {
		initData();
	}, []);
	return useObserver(() => (
		<Modal
			title="查看附加事件"
			visible={true}
			onOk={handleClose}
			onCancel={handleClose}
			okText="确定"
			cancelText="取消"
			width={900}
			footer={[
				<Button key="back" onClick={handleClose}>
					关闭
				</Button>
			]}>
			<div>
				<Table
					columns={columns}
					dataSource={addServerDate}
					pagination={false}
					rowKey={record => record.id}
				/>
			</div>
		</Modal>
	));
}

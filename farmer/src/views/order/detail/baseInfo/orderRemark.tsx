import * as React from "react";
import { Tabs, Table } from "antd";
import { getUrlParams } from "utils/url";
import { store } from "../../store/helper/useStore";
import { useObserver } from "mobx-react-lite";
import { RemarkTypeArr } from "views/order/constant";
import { formatUnixTime } from "utils/time";
import { ImgPreviewModal } from "components/ImgPreview";
const { TabPane } = Tabs;

export function OrderRemark() {
	const [visible, setVisible] = React.useState(false);
	const [img, setIMG] = React.useState("");
	const frontColumns = [
		{
			title: "时间",
			dataIndex: "createDate",
			render: text => formatUnixTime(text)
		},
		{
			title: "创建人",
			dataIndex: "createBy"
		},
		{
			title: "内容",
			dataIndex: "remarkContent"
		},
		{
			title: "附件",
			dataIndex: "enclosureLink",
			render: text =>
				text && (
					<a
						key={"enclosureLink"}
						onClick={() => {
							setIMG(text);
							setVisible(true);
						}}>
						查看
					</a>
				)
		}
	];
	const backColumns = [
		{
			title: "时间",
			dataIndex: "createDate",
			render: text => formatUnixTime(text)
		},
		{
			title: "创建人",
			dataIndex: "createBy"
		},
		{
			title: "备注类型",
			dataIndex: "remarkType",
			render: text =>
				RemarkTypeArr.find(item => item.value === text) &&
				RemarkTypeArr.find(item => item.value === text).label
		},
		{
			title: "内容",
			dataIndex: "remarkContent"
		},
		{
			title: "附件",
			dataIndex: "enclosureLink",
			render: text =>
				text && (
					<a
						onClick={() => {
							setIMG(text);
							setVisible(true);
						}}>
						查看
					</a>
				)
		}
	];
	const orderStore = store().orderStore;
	// const orderListUi = store().orderListUi;
	const remarkEle = React.useRef(null);
	React.useEffect(() => {
		initOrderRemark();
	}, []);
	function initOrderRemark() {
		const params = getUrlParams(window.location.href) as any;
		const anchor = params.anchor || "";
		jumpToAnchor(anchor);
	}
	function jumpToAnchor(anchor) {
		if (anchor === "remark") {
			remarkEle.current.scrollIntoView({
				top: remarkEle.current.getBoundingClientRect().y,
				behavior: "smooth"
			});
		}
	}
	function callback(key) {
		// orderListUi.changeIsFrontRemark();
		console.log(key);
	}
	return useObserver(() => (
		<div>
			<a ref={remarkEle}>{}</a>
			<Tabs onChange={callback} type="card">
				<TabPane tab="前台备注" key="front">
					<Table
						size="small"
						columns={frontColumns}
						dataSource={orderStore.remarks && orderStore.remarks.frontRemarks}
						rowKey={(record, index) => record.createDate + index}
						pagination={false}
					/>
				</TabPane>
				<TabPane tab="后台备注" key="back">
					<Table
						size="small"
						columns={backColumns}
						dataSource={orderStore.remarks && orderStore.remarks.backgroundRemarks}
						rowKey={record => record.createDate}
						pagination={false}
					/>
				</TabPane>
			</Tabs>
			<ImgPreviewModal
				key={"ImgPreviewModal"}
				visible={visible}
				urls={img.split(";").filter(i => i)}
				onCancel={() => setVisible(false)}
			/>
		</div>
	));
}

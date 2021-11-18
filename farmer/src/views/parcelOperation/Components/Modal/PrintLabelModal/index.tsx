import * as React from "react";
import { Modal, Button } from "antd";
import * as ReactDOMServer from "react-dom/server";
import SGLabelPrintForm, { labelPrint } from "./PrintForms/labelPrintForm";
import { injectPrintScript } from "ez-lodop";

export interface PrintLabelModalProps {
	visible: boolean;
	record: any;
	packageFaceInfo: any;
	onCancel?: () => void;
}

// 打印标签  弹窗

const getPrintContent = (markup, setMarkup, props) => {
	const param = {
		markup,
		packageFaceInfo: props.packageFaceInfo,
		onClose: () => setMarkup(true),
		packageNo: props.record.packageCode,
		isAutoUpload: false
	};
	return <SGLabelPrintForm {...param} />;
};

const PrintLabelModal = (props: PrintLabelModalProps) => {
	const { visible, record, onCancel } = props;
	const [markup, setMarkup] = React.useState(true);
	const title = "打印标签";

	React.useEffect(() => {
		injectPrintScript();
	}, [visible]);

	const onPreview = () => {
		setMarkup(false);
	};

	const labelPrints = () => {
		const html = ReactDOMServer.renderToStaticMarkup(getPrintContent(true, setMarkup, props));
		labelPrint(html, props.record.packageCode);
	};

	const getFooter = () => {
		return (
			<div>
				<Button onClick={onPreview} type="primary">
					预览
				</Button>
				<Button onClick={labelPrints} type="primary">
					打印
				</Button>
				<Button onClick={props.onCancel}>取消</Button>
			</div>
		);
	};

	return (
		<React.Fragment>
			<Modal
				visible={visible}
				title={title}
				width={600}
				onCancel={onCancel}
				footer={getFooter()}
				maskClosable={false}
				zIndex={100}>
				包裹号：{record.packageCode}
				{!markup && getPrintContent(false, setMarkup, props)}
			</Modal>
		</React.Fragment>
	);
};

export default PrintLabelModal;

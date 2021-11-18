import * as React from "react";

import * as Barcode from "react-barcode";
import * as QrCode from "qrcode.react";

import { getLodop, licences, LodopPrint } from "ez-lodop";
import { Base } from "./Base";
import { PRINT_INIT_STYLE } from "./printInitStyle";

const barcodeOption = {
	width: 1,
	height: 40,
	format: "CODE128",
	displayValue: true,
	fontOptions: "",
	font: "monospace",
	textAlign: "center",
	textPosition: "bottom",
	textMargin: 2,
	fontSize: 16,
	background: "#ffffff",
	lineColor: "#000",
	margin: 0
};

interface SGLabelPrintFormProps {
	packageFaceInfo: any;
	pages?: number;
	markup?: boolean;
}

export const labelPrint = (html, parcelCode) => {
	let LODOP = getLodop();
	LODOP.SET_LICENSES(...licences);
	LODOP.PRINT_INIT("INVOICE");
	LODOP.SET_PRINT_PAGESIZE(2, "40mm", "60mm", "allPrint");
	LODOP.ADD_PRINT_HTM("1mm", "1mm", "100%", "100%", html);
	LODOP.ADD_PRINT_BARCODE("3mm", "4mm", "56mm", "12mm", "128Auto", parcelCode);
	LODOP.ADD_PRINT_BARCODE("18mm", "36mm", "20mm", "20mm", "QRCode", parcelCode);
	LodopPrint(LODOP, parcelCode);
};

/**
 * 新加坡 包裹打印标签
 */
export default class SGLabelPrintForm extends Base<SGLabelPrintFormProps, {}> {
	constructor(props) {
		super(props);
	}
	printcode = packageNo => {
		const html = PRINT_INIT_STYLE + document.getElementById("invoice").outerHTML;
		labelPrint(html, packageNo);
	};

	render() {
		const { packageNo, packageFaceInfo, markup } = this.props;

		const printComponent = (
			<div
				style={{
					position: "relative",
					height: "36mm",
					overflow: "hidden",
					width: "58mm",
					backgroundColor: "#fff",
					margin: "0 auto"
				}}>
				{/* 包裹号条形码 */}
				<div style={{ height: "14mm", margin: "1mm auto", padding: 0 }}>
					<div style={{ textAlign: "center" }}>
						<div className="codePlaceholder">
							{packageNo && <Barcode value={packageNo} {...barcodeOption} />}
						</div>
					</div>
				</div>
				{/* 包裹信息 */}
				<div
					style={{
						position: "absolute",
						top: "17mm",
						left: "1mm",
						width: "36mm",
						height: "20mm"
					}}>
					{/* 客户昵称 */}
					<div style={{ lineHeight: "18px" }}>
						<span
							style={{
								fontSize: 13,
								fontWeight: "bold",
								wordWrap: "break-word",
								wordBreak: "break-all"
							}}>
							{packageFaceInfo.nickname}
						</span>
					</div>
					{/* 包裹号 */}
					<div
						style={{
							lineHeight: "18px",
							width: "36mm",
							whiteSpace: "nowrap",
							overflow: "hidden",
							textOverflow: "ellipsis"
						}}>
						<label style={{ fontSize: 13, marginRight: 10 }}>{packageNo}</label>
					</div>
					{/* 派送商或者派送方式 */}
					<div style={{ lineHeight: "18px" }}>
						<label style={{ paddingLeft: 0, fontSize: "13px" }}>
							{packageFaceInfo.deliveryMethod}
						</label>
					</div>

					{/* 国家 */}
					<div style={{ lineHeight: "18px" }}>
						<label style={{ paddingLeft: 0, fontSize: "13px" }}>SG</label>
					</div>
					{/* 包裹号二维码 */}
					<div style={{ position: "absolute", top: "2mm", right: "-20mm" }}>
						{packageNo && <QrCode size={60} value={packageNo} />}
					</div>
				</div>
			</div>
		);

		if (markup) {
			return printComponent;
		}
		return this.renderInBody(printComponent, () => this.printcode(packageNo), "40mm*60mm");
	}
}

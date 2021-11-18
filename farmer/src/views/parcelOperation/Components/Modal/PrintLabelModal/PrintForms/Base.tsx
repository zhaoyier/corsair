/**
 * 所有打印的基类，提供下载图片、预览功能
 */
import * as React from "react";
import { injectPrintScript } from "ez-lodop";
import * as ClassNames from "classnames";
import { message, Button } from "antd";
const styles = require("./print.scss");
import RenderInBody from "./renderInBody";
import { downloadImg, getQiniuUrl } from "utils/img";
import { SaveInvoiceURL } from "genServices/chaplin/invoice";
export function downloadFrom(selector: string, filename: string, onClose: () => void) {
	message.info(`开始${filename}图片`, 2);
	// 因为渲染qr需要时间，所以需要延迟一段时间
	setTimeout(function() {
		downloadImg(selector, filename + ".png")
			.then(() => {
				message.success(`${filename}生成成功`, 2);
				onClose();
			})
			.catch(err => {
				message.error(`${filename}生成失败` + err, 5);
			});
	}, 1500);
}
export function uploadb64From(selector: string, packageCode: string, directHtmlJson?: string[]) {
	// 因为渲染qr需要时间，所以需要延迟一段时间
	return new Promise((resolve, reject) => {
		setTimeout(function() {
			getQiniuUrl(selector, directHtmlJson)
				.then(async (ImgArr: string[]) => {
					if (ImgArr.length > 0) {
						try {
							await SaveInvoiceURL({
								packageCode,
								invoiceURL: ImgArr[1],
								wayBillURL: ImgArr[0]
							});
							resolve();
						} catch (err) {
							reject(err.message);
						}
					}
				})
				.catch(err => {
					message.error(`图片上传失败` + err, 5);
					reject(err);
				});
		}, 1500);
	});
}
export interface BasePrintFormProps {
	onClose: () => void;
	packageNo: string;
	isAutoDownFrom?: boolean;
	downQuery?: string;
	isAutoUpload: boolean;
}
export class Base<P = {}, S = {}> extends React.Component<P & BasePrintFormProps, S> {
	constructor(props) {
		super(props);
	}

	componentDidMount() {
		injectPrintScript();
		this.asyncDownloadImg(this.props.isAutoDownFrom); // 自动下载到本地
		this.asyncUploadImg(this.props.isAutoUpload); // 上传到七牛
	}
	asyncDownloadImg(downFromFlg: boolean) {
		const { packageNo, onClose, downQuery } = this.props;
		if (downFromFlg) {
			downloadFrom(downQuery || "#invoice", packageNo, onClose);
		}
	}
	asyncUploadImg(uploadFlag: boolean) {
		const { packageNo, downQuery } = this.props;
		if (uploadFlag) {
			uploadb64From(downQuery || "#invoice", packageNo)
				.then(resp => {
					console.log(resp);
					message.success("上传面单url成功");
				})
				.catch(e => message.error(e));
		}
	}
	renderInBody(child: React.ReactNode, internalPrint, size: string) {
		const { isAutoDownFrom } = this.props;
		return (
			<RenderInBody
				className={ClassNames(
					[isAutoDownFrom ? styles.autoDownFrom : styles.mask],
					[styles.sgAirbill],
					[styles.airbill]
				)}>
				<div className={styles.close} onClick={this.props.onClose} />
				<div className={styles.print}>
					<Button onClick={internalPrint}>打印</Button>
					<Button
						style={{ marginLeft: 20 }}
						onClick={() => {
							this.asyncDownloadImg(true);
						}}>
						下载
					</Button>
				</div>
				<div className={styles.tip}>{size}</div>
				<div id="invoice" className={styles.printContainer}>
					{child}
				</div>
			</RenderInBody>
		);
	}
}

/**
 * 参考代码
 * https://juejin.im/post/5a17c5e26fb9a04527254689
 */
import html2canvas from "html2canvas";
import { message } from "antd";
import { downloadURL } from "./url";
import { uploadBase64 } from "components/UploadImg";
/**
 *  将传入值转为整数
 */
function parseValue(value) {
	return parseInt(value, 10);
}
/**
 * 根据window.devicePixelRatio获取像素比
 */
function DPR() {
	if (window.devicePixelRatio && window.devicePixelRatio > 1) {
		return window.devicePixelRatio;
	}
	return 1;
}
/**
 * 绘制canvas
 */
export async function drawCanvas(selector) {
	// 获取想要转换的 DOM 节点
	const dom = document.querySelector(selector);
	const box = window.getComputedStyle(dom);
	// DOM 节点计算后宽高
	const width = parseValue(box.width);
	const height = parseValue(box.height);
	// 获取像素比
	const scaleBy = DPR();
	// 创建自定义 canvas 元素
	const canvas = document.createElement("canvas");

	// 设定 canvas 元素属性宽高为 DOM 节点宽高 * 像素比
	canvas.width = width * scaleBy;
	canvas.height = height * scaleBy;
	// 设定 canvas css宽高为 DOM 节点宽高
	canvas.style.width = `${width}px`;
	canvas.style.height = `${height}px`;
	// 获取画笔
	const context = canvas.getContext("2d");

	// 将所有绘制内容放大像素比倍
	context.scale(scaleBy, scaleBy);

	// 将自定义 canvas 作为配置项传入，开始绘制
	return await html2canvas(dom, { canvas });
}
/**
 * 图片转base64格式
 */
export function img2base64(url, crossOrigin) {
	return new Promise(resolve => {
		const img = new Image();

		img.onload = () => {
			const c = document.createElement("canvas");

			c.width = img.naturalWidth;
			c.height = img.naturalHeight;

			const cxt = c.getContext("2d");

			cxt.drawImage(img, 0, 0);
			// 得到图片的base64编码数据
			resolve(c.toDataURL("image/png"));
		};

		if (crossOrigin) {
			img.setAttribute("crossOrigin", crossOrigin);
		}
		img.src = url;
	});
}

export function createNode(json) {
	const template = `<div id='directHtmlDiv'>${json}</div>`;
	let tempNode = document.createElement("div");
	tempNode.innerHTML = template;
	tempNode.style.visibility = "hidden";
	return tempNode.firstChild;
}

// 直接点击保存并打印,ReactDOMServer不会调用Base的componentDidMount，所以需要directHtmlJson，生成真实Dom来生成图片
export function htmlToCanvas(selector, directHtmlJson?: string[]): Promise<HTMLCanvasElement[]> {
	// html2canvas配置 http://html2canvas.hertzen.com/configuration/
	const opts = {
		logging: false, // 日志开关，html2canvas的内部执行流程
		useCORS: true // 开启图片跨域配置
	};
	return new Promise(async (resolve, reject) => {
		if (directHtmlJson) {
			document.body.appendChild(createNode(directHtmlJson));
		}
		const querys = document.querySelectorAll(selector);
		const promise = [];
		for (let i = 0; i < querys.length; i++) {
			promise.push(html2canvas(querys[i], opts));
		}
		try {
			const canvasList = await Promise.all(promise);
			resolve(canvasList);
		} catch (err) {
			reject(err);
		}
	});
}

export function downloadImg(selector, filename): Promise<boolean> {
	return new Promise(async (resolve, reject) => {
		try {
			const canvas = await htmlToCanvas(selector);
			canvas.forEach(k => {
				downloadURL(k.toDataURL(), filename);
			});
			resolve(true);
		} catch (err) {
			message.error("下载图片失败:" + err.error || err);
			reject(err);
		}
	});
}
/**
 *  把html转换成canvas，然后再并发上传到七牛服务器，之后返回七牛的图片存储地址。
 */
export function getQiniuUrl(selector, directHtmlJson?: string[]): Promise<string[]> {
	return new Promise(async (resolve, reject) => {
		try {
			const canvas = await htmlToCanvas(selector, directHtmlJson);
			if (directHtmlJson) {
				const perDiv = document.getElementById("directHtmlDiv");
				document.body.removeChild(perDiv);
			}
			const promise = [];
			canvas.forEach(k => {
				promise.push(uploadBase64(k.toDataURL().split(",")[1]));
			});
			const ImgArr = [];
			const resp = await Promise.all(promise);
			resp.forEach(k => {
				ImgArr.push(`${k.baseUrl}${k.hash}`);
			});
			resolve(ImgArr);
		} catch (err) {
			message.error("上传图片失败:" + err.error || err);
			reject(err);
		}
	});
}

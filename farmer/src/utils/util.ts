// 接口返回的价格单位为 分，转为 元
export function formatPrice(price: string): string {
	if (String(price) && String(price) !== "undefined") {
		return (parseInt(price, 10) / 100).toFixed(2).toString();
	}
	return "";
}

// 接口返回的重量单位为 g，转为 kg
export function formatWeight(weight: string): string {
	if (String(weight) && String(weight) !== "undefined") {
		return (parseInt(weight, 10) / 1000).toFixed(2).toString();
	}
	return "";
}
// 接口返回的重量单位为 g，转为 kg, 后面接单位
export function formatWeightUnit(weight: string): string {
	return (parseInt(weight, 10) / 1000).toFixed(2).toString() + "（kg）";
}
// 元转为分
export function priceToInt(num: string | number): string {
	return String((Number(num) * 100).toFixed(0));
}
// kg转为g
export function weightToInt(weight: string): string {
	return "" + (parseFloat(weight) * 1000).toFixed(0);
}

export function isNumber(str: string): boolean {
	return !!/^(-)?\d+(\.\d+)?$/.exec(str);
}

/** 整数转化为 浮点数 */
export function integerToFloat(integer: number | string, bit: number = 2): number {
	const priceNum = Number(integer);
	if (!Number.isNaN(priceNum)) {
		return Number((priceNum / Math.pow(10, bit)).toFixed(bit));
	}
	return 0;
}

/** 浮点数转化为 整数 */
export function floatToInteger(float: number, bit: number = 2): number {
	const priceNum = Number(float);
	if (!Number.isNaN(priceNum)) {
		return Number((float * Math.pow(10, bit)).toFixed(bit));
	} else {
		return 0;
	}
}
// 去除 空参数
export const deleteEmpty = (obj, selectArr = []) => {
	const newObj = { ...obj };
	if (selectArr.length > 0) {
		selectArr.forEach(elem => {
			if (newObj[elem] === "全部") {
				delete newObj[elem];
			}
		});
	}
	Object.keys(newObj).forEach(key => {
		if (!newObj[key]) {
			delete newObj[key];
		}
		if (typeof newObj[key] === "string") {
			newObj[key] = newObj[key].trim();
		}
	});
	return newObj;
};

// 数组查询对象
export const findObj = (list, param, value) => {
	if (list && list.length > 0) {
		const find = list.find(elem => elem[param] === value);
		return find;
	}
	return null;
};

const getNumber = value => {
	return value.replace(/[^0-9]/gi, "");
};

// Input OnChange 小数
/**
 * @param canMinus 是否能输入负数
 * @param limit 最大保留几位小数
 */
export function limitDecimals(val, canMinus = false, limit = 2) {
	const value = val.replace(/[^0-9.-]/gi, "");
	const list = value.split(".");
	if (value.charAt(0) === ".") {
		return `${value.slice(1)}`;
	}
	const prefix = list[0].charAt(0) === "-" ? "-" : "";
	let before = prefix ? list[0].slice(1) : list[0];
	if (value.charAt(list[0].length) === ".") {
		return `${canMinus && prefix ? "-" : ""}${getNumber(before)}.${getNumber(
			list[1].slice(0, limit)
		)}`;
	} else {
		return `${canMinus && prefix ? "-" : ""}${getNumber(before)}`;
	}
}

// Input OnChange 整数
/**
 * @param canMinus 是否能输入负数
 */
export function limitInteger(val, canMinus = false) {
	const vals = String(val);
	const value = vals.replace(/[^0-9-]/gi, "");
	const prefix = vals.charAt(0) === "-" ? "-" : "";
	let num = prefix ? value.slice(1) : value;
	const before = canMinus ? `${prefix}${num}` : `${num}`;
	const res = `${canMinus && prefix ? "-" : ""}${getNumber(before)}`;
	return !checkEmpty(res) ? Number(`${canMinus && prefix ? "-" : ""}${getNumber(before)}`) : "";
}

// Input OnBlur
export function inputBlur(value) {
	if (value.charAt(value.length - 1) === "." || value === "-") {
		return value.slice(0, -1);
	}
	return value;
}

// base64转为为到blob二进制
function dataURItoBlob(dataURI, mimeString) {
	if (!mimeString) {
		mimeString = dataURI
			.split(",")[0]
			.split(":")[1]
			.split(";")[0]; // mime类型
		dataURI = dataURI.split(",")[1];
	}
	let byteString = atob(dataURI); // base64 解码
	let arrayBuffer = new ArrayBuffer(byteString.length); // 创建缓冲数组
	let intArray = new Uint8Array(arrayBuffer); // 创建视图
	for (let i = 0; i < byteString.length; i++) {
		intArray[i] = byteString.charCodeAt(i);
	}
	return new Blob([intArray], { type: mimeString });
}

function downloadURL(url: string, name: string = "") {
	const link = document.createElement("a");
	link.download = name;
	link.href = url;
	document.body.appendChild(link);
	link.click();
	document.body.removeChild(link);
}

export function downloadExcel(data: Blob, filename): void {
	const url = URL.createObjectURL(dataURItoBlob(data, "application/octet-stream"));
	downloadURL(url, filename);
	// 手动标记用于内存回收，防止内存泄漏
	URL.revokeObjectURL(url);
}

export const checkEmpty = num => {
	if (
		Number.isNaN(Number(num)) ||
		num === null ||
		num === undefined ||
		String(num).length === 0
	) {
		return true;
	}
	return false;
};

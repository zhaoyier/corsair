import * as moment from "moment";

export const TimeFormatTotalPositive = "YYYY-MM-DD HH:mm:ss";
export const TimeFormatTotalReverse = "HH:mm:ss DD-MM-YYYY";
export const TimeFormatToDay = "YYYY-MM-DD";
export const TimeFormatToDayReverse = "DD-MM-YYYY";
export const TimeFormatToHour = "YYYY-MM-DD HH";
export const TimeFormatToMinite = "YYYY-MM-DD HH:mm";
export const TimeFormatTotalWeekPositive = "yyyy-MM-dd EE hh:mm:ss"; // 2009-03-10 周二 08:09:04
export const TimeFormatTotalWeekPositive2 = "yyyy-MM-dd EE hh:mm:ss"; // 2009-03-10 星期二 08:09:04

// 判断登录是否超时
export const checkExp = expAt => {
	const current = moment().unix();
	if (current - expAt > 0) {
		return true;
	}
	return false;
};

/**
 * 时间戳格式化
 * t: 时间戳
 * format： 格式化类型   默认 YYYY-MM-DD HH:mm:ss
 * 注意： moment.unix(timestamp).format(format)---- timestamp为 10位 时间戳  单位是秒
 * 			 moment(+timestamp).format(fmt) ---- timestamp为 13位 时间戳  单位是毫秒
 */
export function formatUnixTime(
	t: number | string,
	format: string = TimeFormatTotalPositive
): string {
	let timestamp: number;
	// 传入时间戳为 null，空字符串，undefined，0
	if (!t) {
		return "-";
	}
	t = String(t);
	if (typeof t === "string") {
		timestamp = parseInt(t, 10);
	} else if (typeof t === "number") {
		timestamp = t;
	}

	if (timestamp) {
		return moment.unix(timestamp).format(format);
	}

	return "";
}

/**
 * 时间戳格式化   UTC
 * t: 时间戳
 * format： 格式化类型  默认 YYYY-MM-DD HH:mm:ss
 * https://www.cnblogs.com/x_wukong/p/5670638.html
 * 1.UTC时间 与 GMT时间：
 * 	 我们可以认为  格林威治时间就是时间协调时间（GMT=UTC），格林威治时间和UTC时间均用秒数来计算的。
 * 2.UTC时间 与 本地时
 * 	 UTC + 时区差 ＝ 本地时间
 *   时区差东为正，西为负。在此，把东八区时区差记为 +0800，
 *   UTC + (＋0800) = 本地（北京）时间
 *   那么，UTC = 本地时间（北京时间)）- 0800
 * 3.UTC 与 Unix时间戳
 *   UTC时间都是从（1970年01月01日 0:00:00)开始计算秒数的。
 *   所看到的UTC时间那就是从1970年这个时间点起到具体时间共有多少秒。 这个秒数就是Unix时间戳
 */
export function formatUtcTime(
	t: number | string,
	format: string = TimeFormatTotalPositive
): string {
	let timestamp: number;
	if (typeof t === "string") {
		timestamp = parseInt(t, 10);
	} else if (typeof t === "number") {
		timestamp = t;
	}

	if (timestamp) {
		return moment.utc(timestamp).format(format);
	}

	return "";
}

/**
 * 获取一天开始的时间戳
 */
export function getUnixTimeOfDayStart(m: moment.Moment): string {
	if (!m) {
		return null;
	}
	return m
		.startOf("day")
		.unix()
		.toString();
}

/**
 * 获取一天结束的时间戳
 */
export function getUnixTimeOfDayEnd(m: moment.Moment): string {
	if (!m) {
		return null;
	}
	return m
		.endOf("day")
		.unix()
		.toString();
}

/**
 * 获取时间戳 (不限制时间格式)
 */
export function getTimeUnixDateStr(m: string): string {
	if (!m) {
		return null;
	}
	let time: any = moment(m)
		.unix()
		.valueOf();
	return time;
}

/**
 * 获取UTC 时间戳 （不限制时间格式）  UTC
 */
export function getTimeUtcDateStr(m: string): string {
	if (!m) {
		return null;
	}
	let time: any = moment(m)
		.utc()
		.valueOf();
	return time;
}

/**
 * 获取UTC 时间戳 （限制时间格式）
 * 默认  YYYY-MM-DD
 */
export function dateToTimestamp(time, format: string = TimeFormatToDay) {
	let t = moment(time).format(format);
	return getTimeUtcDateStr(t);
}

// 国际化周
const weekCollection = {
	zh: {
		"0": "周日",
		"1": "周一",
		"2": "周二",
		"3": "周三",
		"4": "周四",
		"5": "周五",
		"6": "周六"
	},
	en: {
		"0": "Sunday",
		"1": "Monday",
		"2": "Tuesday",
		"3": "Wednesday",
		"4": "Thursday",
		"5": "Friday",
		"6": "Saturday"
	}
};
/**
 * 获取星期几 （区分中英文）
 */
export function getWeekDate(m: number | string, type: string = "zh"): string {
	const date = new Date(m);
	const weeks = date.getDay();

	return weekCollection[type][weeks] || "";
}

/**
 * 日期或时间戳  转化为  毫秒
 * @param date string|number "2019-12-01" or 10或13数字
 */
export const formateDateToMilliseconds = (date: string | number) => {
	let milliSeconds: number;
	if (/^\d+$/.test(String(date))) {
		if (String(date).length !== 10 && String(date).length !== 13) {
			return null;
		}
		milliSeconds = +date;
	} else {
		milliSeconds = +new Date(date);
	}
	if (String(milliSeconds).length === 10) {
		milliSeconds = milliSeconds * 1000;
	}
	return milliSeconds;
};
/**
 * 将 目标地时间  转化为 本地时间  单位： 秒
 * 本地： 东八区
 */
export const transformDateToLocalMMs = (date: string | number): number => {
	const cMilliSecond = formateDateToMilliseconds(date);
	const cDate = new Date(cMilliSecond);
	const cDateTime = cMilliSecond + (cDate.getTimezoneOffset() - 8 * 60) * 60 * 1000;
	return Math.round(cDateTime / 1000);
};

/**
 * 将 本地时间 转化为  目的地时间  单位： 秒
 * 本地： 东八区
 */
export const transformDateToLocalMs = (date: string | number): number => {
	const cMilliSecond = formateDateToMilliseconds(date);
	const cDate = new Date(cMilliSecond);
	const cDateTime = cMilliSecond + (-cDate.getTimezoneOffset() + 8 * 60) * 60 * 1000;
	return Math.round(cDateTime / 1000);
};

// antd 组件中 value  日期 显示
export function dataString(date) {
	return moment(date * 1000);
}

// antd DatePicker.RangePicker 日期 取值
export const changeDataPicker = date => {
	const newDate = {
		from: null,
		to: null
	};
	if (date.length !== 0) {
		newDate.from = moment(date[0])
			.startOf("day")
			.unix();
		newDate.to = moment(date[1])
			.endOf("day")
			.unix();
		return [newDate.from, newDate.to];
	} else {
		return [];
	}
};

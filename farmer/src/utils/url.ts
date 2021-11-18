import { countryCodes, fetchSystem } from "./country";

// export const isLocalEnv = /127\.0|192\.168|localhost|8990/i.test(location.host);
const testEnvHostReg = /65emall\.net|127\.0|192\.168|localhost/i;

export function isUAT(): boolean {
	return testEnvHostReg.test(location.hostname);
}

export function qiniuUrl(): string {
	return isUAT() ? "http://7xiata.com1.z0.glb.clouddn.com/" : "http://i.ezbuy.sg/";
}

export function uploadQiniuUrl(): string {
	return isUAT() ? "//up.qiniup.com" : "//up-as0.qiniup.com";
}

export const urlIsLITB = () => {
	return /lightinthebox/gi.test(location.href);
};
export function downloadURL(url: string, name: string = "") {
	const link = document.createElement("a");
	link.download = name;
	link.href = url;
	document.body.appendChild(link);
	link.click();
	document.body.removeChild(link);
}
export let WebDomains = {
	SG: "https://ezbuy.sg",
	MY: "https://ezbuy.my",
	TH: "https://ezbuy.co.th",
	ID: "https://ezbuy.co.id",
	PK: "https://pk.ezbuy.com",
	TWC: "https://tw.ezbuy.com",
	MC: "https://ezbuy.com",
	LITB: "https://www.lightinthebox.com",
	LITBMINI: "https://www.miniinthebox.com",
	LITBADOR: "https://www.ador.com"
};
// https://m.zewant.com/Promo?title=testing-123&id=5d3ff852222e9b7dea21ea27?
let MobileWebDomains = {
	SG: "https://m.ezbuy.sg",
	MY: "https://m.ezbuy.my",
	TH: "https://m.ezbuy.co.th",
	ID: "https://m.ezbuy.co.id",
	PK: "https://m-pk.ezbuy.com",
	TWC: "https://m-tw.ezbuy.com",
	THE: "https://m.zewant.com",
	MC: "https://m.ezbuy.com",
	LITB: "https://m.lightinthebox.com",
	LITBMINI: "https://m.miniinthebox.com",
	LITBADOR: "https://m.Ador.com"
};
if (isUAT()) {
	WebDomains = {
		SG: "sg.65emall.net",
		MY: "my.65emall.net",
		TH: "th.65emall.net",
		ID: "id.65emall.net",
		PK: "pk.65emall.net",
		TWC: "tw.65emall.net",
		MC: "https://web.ezbuy.dev",
		LITB: "https://www.lightinthebox.com",
		LITBMINI: "https://www.miniinthebox.com",
		LITBADOR: "https://www.ador.com"
	};
	MobileWebDomains = {
		SG: "m.sg.65emall.net",
		MY: "m.my.65emall.net",
		TH: "m.th.65emall.net",
		ID: "m.id.65emall.net",
		PK: "m.pk.65emall.net",
		TWC: "m.tw.65emall.net",
		THE: "https://m.zewant.net",
		MC: "https://m.ezbuy.dev",
		LITB: "https://mqa.lightinthebox.com",
		LITBMINI: "https://m.miniinthebox.com",
		LITBADOR: "https://m.Ador.com"
	};
}
// https://mqa.lightinthebox.com/promotions/{title}.html?id=xxx
export function getPageUrlByPath(
	path: string,
	countryCode: string = countryCodes[0],
	isPreview: boolean = false,
	platform?: string
) {
	// 支持 WEB
	let domain = WebDomains[countryCode.toUpperCase()] || WebDomains.SG;
	if (platform && platform.match(/mobile/gi)) {
		domain = MobileWebDomains[countryCode.toUpperCase()] || MobileWebDomains.SG;
	}
	const sys = fetchSystem();
	let TrimPath = path.trim();
	const hashIndex = TrimPath.lastIndexOf("#");
	const hash = TrimPath.substring(hashIndex);

	if (isPreview) {
		if (TrimPath.indexOf("?") > -1) {
			TrimPath += "&isPreview=true";
		} else {
			TrimPath += "?isPreview=true";
		}
	}
	if (/litb/gi.test(sys) && isUAT()) {
		if (TrimPath.indexOf("?") > -1) {
			TrimPath += "&litbuat=true";
		} else {
			TrimPath += "?litbuat=true";
		}
	}
	if (hashIndex > -1) {
		TrimPath = TrimPath.replace(hash, "") + hash;
	}

	if (/^(https?|\/\/)/i.test(TrimPath)) {
		return TrimPath;
	} else {
		if (/^(https)/i.test(domain)) {
			return `${domain}${TrimPath}`;
		}
		return `${location.protocol}//${domain}${TrimPath}`;
	}
}

function hasEncode(str: string): boolean {
	const preStr = str;
	try {
		str = decodeURIComponent(str);
	} catch (err) {
		str = preStr;
	}
	return preStr !== str;
}

function newEncodeURIComponent(str: string, isEncodeURIComponent = true): string {
	return hasEncode(str) ? str : isEncodeURIComponent ? encodeURIComponent(str) : encodeURI(str);
}

function replaceSearchKeyValue(keyValueString: string): string {
	return keyValueString.replace(/([^=]+)(=)(.+)/, (p1, p2, p3) => {
		return `${newEncodeURIComponent(p1, false)}${p2}${newEncodeURIComponent(p3)}`;
	});
}

export function transferToEncodedURL(link: string): string {
	let linkUrl = null;
	try {
		linkUrl = new URL(link);
	} catch (e) {
		console.error(e);
		return null;
	}
	let { search, pathname } = linkUrl;

	if (pathname) {
		linkUrl.pathname = newEncodeURIComponent(pathname, false);
	}
	if (search) {
		const splitStr: string = "&";
		const searchString: string = search.substring(1);
		const searchStringItems: string[] = searchString.split(splitStr);
		linkUrl.search = searchStringItems.reduce((preValue, cValue, index) => {
			const connection = index === searchStringItems.length - 1 ? "" : splitStr;
			return preValue + replaceSearchKeyValue(cValue) + connection;
		}, "?");
	}
	return linkUrl.href;
}

export function getSiteDomain(): string {
	const domain = window.location.hostname.split(".");
	if (domain.includes("65daigou")) {
		return "http://ezbuy.sg";
	} else if (domain.includes("pk") || domain.includes("pkadmin")) {
		return "http://pk.ezbuy.com";
	} else if (domain.includes("pkadmin4")) {
		return "http://pk4.65emall.net";
	} else {
		return "http://sg2.65emall.net";
	}
}

export function getApiEzSellerPrefix() {
	const location = window.location.origin;
	if (location.includes("65daigou")) {
		return "https://backend.65daigou.com/api/EzSeller";
	} else if (location.includes("65emall")) {
		return "http://webapi.sg.65emall.net/api/EzSeller";
	} else {
		return "/api/EzSeller";
	}
}

export function getSpikeOraclePrefix() {
	const location = window.location.origin;
	if (location.includes("65daigou")) {
		return "https://backend.65daigou.com/api/spike.Oracle";
	} else if (location.includes("65emall")) {
		return "http://webapi.sg.65emall.net/api/spike.Oracle";
	} else {
		return "/api/Oracle";
	}
}

export function getApiPantosPrefix() {
	const location = window.location.origin;
	if (location.includes("65daigou")) {
		return "https://backend.65daigou.com/api/pantos.Dispatch";
	} else if (location.includes("65emall")) {
		return "http://webapi.sg.65emall.net/api/pantos.Dispatch";
	} else {
		return "/api/EzSeller";
	}
}
// 截取地址栏参数
export function hrefObj() {
	let localhref = window.location.href;
	let localarr = localhref.split("?")[1].split("&");
	let tempObj = {};
	for (let i = 0; i < localarr.length; i++) {
		tempObj[localarr[i].split("=")[0]] = localarr[i].split("=")[1];
	}
	return tempObj;
}

export const isAnchor = (link: string) => {
	return /^#/.test(link);
};
export default qiniuUrl;

// 获取地址栏 参数
export const getUrlParams = function(url) {
	let params = {};
	(url + "?")
		.split("?")[1]
		.split("&")
		.forEach(function(pair: any) {
			pair = (pair + "=").split("=").map(decodeURIComponent);
			if (pair[0].length) {
				params[pair[0]] = pair[1];
			}
		});
	return params;
};

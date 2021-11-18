import { message } from "antd";
import { stringify } from "querystring";
import { downloadURL } from "utils/url";
import "whatwg-fetch";
import { loginTimeout } from "./auth";
import { getUserInfo } from "./user";

const testEnvHostReg = /65emall\.net|127\.0|192\.168|localhost/i;
const prefixNames: {
	[key: string]: string;
} = {
	Cms: "cms"
};

function isUAT(): boolean {
	return testEnvHostReg.test(location.hostname);
}
function getKflowDomain(): string {
	try {
		if (isUAT()) {
			return "localhost:8080";
		}
		const [, kenv = ""] = /((ktest|kdev)[0-9])\./.exec(location.hostname) || [];
		if (kenv) {
			return `sg-en-web-api.${kenv}-api.65emall.net`;
		}
		return "sg-en-web-api.kdev0-api.65emall.net";
	} catch (e) {
		return "";
	}
}

export function getApiDomain(serviceName) {
	if (isUAT()) {
		return getKflowDomain() || window.location.host.replace(/ezship\d+/, "dgadmin");
	}

	if (serviceName === "AdminHomepage" || serviceName === "AdminProduct") {
		return "//backend.65daigou.com";
	}
	return window.location.host;
}

export const apiDomain = `//${getKflowDomain() ||
	(isUAT() ? "webapi.65emall.net" : "webapi2.65daigou.com")}`;

export function checkStatus(response: Response): Response {
	if (!response.ok) {
		const error = new Error(response.statusText);
		switch (response.status) {
			case 401:
				window.location.pathname = "/login.html";
				break;
			default:
				message.warn(`${response.status}: ${response.statusText}`);
				return response;
		}
		throw error;
	}
	return response;
}

function getApiUrl(method: string) {
	const serviceName = method.split(".")[0];
	let apiPrefix = "/api/";
	const prefixName = prefixNames[serviceName];
	if (prefixName !== undefined) {
		apiPrefix += `${prefixName}/`;
	}

	if (serviceName === "oktights") {
		if (method.indexOf("/") !== -1) {
			return `//${apiDomain}/api/${method}`;
		}
		return `//${apiDomain}${apiPrefix}${method.replace(/\./g, "/")}`;
	} else {
		if (method.indexOf("/") !== -1) {
			return `//${getApiDomain(serviceName)}/api/${method}`;
		}
		return `//${getApiDomain(serviceName)}${apiPrefix}${method.replace(/\./g, "/")}`;
	}
}

export function doFetch(url: string, params?: any, fetchMethod?: string, headers?: any) {
	const method = fetchMethod ? fetchMethod : "post";
	const userInfo = getUserInfo();
	const options = {
		method,
		credentials: "include",
		headers
	} as any;

	if (userInfo && userInfo.token) {
		const arr = url.split("/");
		const fnName = arr[arr.length - 1];
		if (fnName !== "GetUploadInfo") {
			options.headers = {
				"pp-token": String(userInfo.token),
				"Content-Type": "application/json;charset=UTF-8"
			};
		}
	}

	if (params) {
		if (method === "post") {
			options.body = JSON.stringify(params);
		} else {
			url = url + "?" + stringify(params);
		}
	}

	return fetch(url, options).then(checkStatus);
}

export function webapi<T>(
	method: string,
	params: any,
	fetchMethod?: string,
	headers?: any
): Promise<T> {
	let url = fetchMethod ? fetchMethod : getApiUrl(method);
	if (url.indexOf("/api/AdminHomepage/") > 0) {
		// AdminHomepage 是个url特殊的服务
		url = url.replace("/api/AdminHomepage/", "/api/homepage/AdminHomepage/");
	}
	if (url.indexOf("/api/AdminProduct/") > 0) {
		// Fix AdminProduct URL Prefix
		url = url.replace("/api/AdminProduct/", "/api/AdminProduct/AdminProduct/");
	}
	return doFetch(url, params, fetchMethod, headers)
		.then(
			response =>
				response.json().then(json => ({ data: json, status: response.status })) as Promise<{
					data: any;
					status: number;
				}>
		)
		.then<T>(response => {
			if (response.status === 403) {
				return loginTimeout();
			}
			if (response.status >= 400) {
				return Promise.reject({
					status: response.status,
					code: response.data.code,
					message: response.data.message
				});
			}
			return response.data;
		});
}

// 服务器端提供流的形式下载
export function downloadFile(method: string, params: any, fileName: string = ""): Promise<void> {
	let url = getApiUrl(method);

	return doFetch(url, params)
		.then(resp => {
			if (resp.status === 400) {
				return resp.json().then(data => {
					throw new Error(data.message || "Download fail.");
				});
			} else {
				return resp.blob();
			}
		})
		.then(blob => URL.createObjectURL(blob))
		.then(url => downloadURL(url, fileName));
}

export function ajax(method: string, url: string, data: any): Promise<string> {
	return new Promise((resolve, reject) => {
		const xhttp = new XMLHttpRequest();
		xhttp.onreadystatechange = function() {
			if (this.readyState === 4 && this.status === 200) {
				resolve(this.responseText);
			}
		};
		xhttp.onerror = function() {
			reject("ajax error");
		};
		xhttp.open(method, url, true);
		xhttp.send(data);
	});
}

export function ajaxGet(url: string, data: any) {
	return ajax("GET", url, data);
}

export default webapi;

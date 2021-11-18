import { Modal } from "antd";
import { getUserInfo } from "./user";

const urlKey = "returnURL";

export function redirectToLogin(returnURL: string) {
	sessionStorage.setItem(urlKey, returnURL);
	window.location.href = "/login.html";
}

export function redirectToLastPage() {
	const returnURL = sessionStorage.getItem(urlKey);
	if (returnURL) {
		window.location.href = returnURL;
	} else {
		sessionStorage.removeItem(urlKey);
		window.location.href = "/index.html";
	}
}

export const loginTimeout = () => {
	let secondsToGo = 3;
	const modal = Modal.warning({
		title: "登录超时",
		content: `${secondsToGo} 秒后即将跳转登录页，请稍后...`
	});
	const timer = setInterval(() => {
		secondsToGo -= 1;
		modal.update({
			content: `${secondsToGo} 秒后即将跳转登录页，请稍后...`
		});
	}, 1000);
	setTimeout(() => {
		clearInterval(timer);
		modal.destroy();
		redirectToLogin(window.location.href);
	}, secondsToGo * 1000);
};

export const checkGroupId = (id: string) => {
	const userInfo = getUserInfo();
	if (userInfo && userInfo.groupsData) {
		const groupInfo = userInfo.groupsData.groupInfo || [];
		if (groupInfo && groupInfo.length > 0) {
			const find = groupInfo.find(elem => elem.groupId === id);
			if (find) {
				return true;
			}
		}
	}
	return false;
};

import { redirectToLogin } from "utils/auth";
import { LoginSuccessData } from "genServices/ezShipOMS/auth";

const loginKey = "loginInfo";
interface UserInfoType extends LoginSuccessData {
	username: string;
}

export type UserInfo = UserInfoType | null;

let userInfo: UserInfo = null;

export function getUserInfo(): UserInfo {
	if (userInfo) {
		return userInfo;
	}

	const info = sessionStorage.getItem(loginKey);
	if (!info) {
		return null;
	}

	try {
		userInfo = JSON.parse(info);
		return userInfo;
	} catch (e) {
		console.error(e);
	}

	if (!userInfo && !window.location.pathname.match(/login/)) {
		redirectToLogin(window.location.href);
	}
	return null;
}

export function setUserInfo(info: UserInfoType) {
	sessionStorage.setItem(loginKey, JSON.stringify(info));
	userInfo = info;
}

export function removeUserInfo() {
	sessionStorage.removeItem(loginKey);
	userInfo = null;
}

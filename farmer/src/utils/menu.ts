import { countryCodes } from "./country";
// import { doFetch } from "./webapi";
const menuData = require("../assets/menu.json");
export interface MenuItem {
	name: string;
	path: string;
	desc: string;
	icon: string;
	eName: string; // 模块entry name
	zhName: string;
	enName: string;
	href: string;
	submenu: MenuItem[];
	system: string;
	platForms?: string[]; // platform list
	excludes: string[]; // country list
	includes: string[]; // country list
	isHistory?: boolean;
}

function getMenuLimitLater(menu: MenuItem[], empty: any = []): MenuItem[] {
	if (!menu || menu.length <= 0) {
		return empty;
	}
	if (countryCodes.length > 1) {
		return menu;
	}

	let resultData: MenuItem[] = [];

	try {
		resultData = JSON.parse(JSON.stringify(menu));
	} catch (err) {
		console.log(err);
	}

	return resultData
		.map(item => {
			const { includes, excludes, platForms } = item;
			const system = window.sessionStorage.getItem("system");
			const findPlatForm =
				platForms && Array.isArray(platForms) && platForms.find(elem => elem === system);
			if (platForms && Array.isArray(platForms) && !findPlatForm) {
				return null;
			}

			if (
				(Array.isArray(includes) || Array.isArray(excludes)) &&
				((Array.isArray(includes) && !includes.includes(countryCodes[0])) ||
					(Array.isArray(excludes) && excludes.includes(countryCodes[0])))
			) {
				return null;
			}

			return {
				...item,
				submenu: getMenuLimitLater(item.submenu || [], null)
			};
		})
		.filter(ele => ele !== null);
}

export function fetchMenu(): Promise<MenuItem[]> {
	return new Promise(resolve => {
		resolve(getMenuLimitLater(menuData));
	});
}

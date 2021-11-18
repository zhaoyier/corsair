import { MenuItem } from "utils/menu";
interface TUseItem {
	link: string;
	count: number;
	name: string;
	lasttime: number;
}

const HABIT_KEY = "habit_key";

export function getKeys(): TUseItem[] | null {
	const str = localStorage.getItem(HABIT_KEY);

	try {
		const data: TUseItem[] = JSON.parse(str);

		if (Array.isArray(data)) {
			if (data.length === 1) {
				return data;
			}
			const lastTime = Math.max(...data.map(ele => ele.lasttime || 0));
			const lastItemIndex = data.findIndex(ele => ele.lasttime === lastTime);
			if (lastItemIndex < 0) {
				return data.sort((a, b) => b.count - a.count);
			}
			const lastItem = data[lastItemIndex];
			data.splice(lastItemIndex, 1);

			return [lastItem, ...data.sort((a, b) => b.count - a.count)].filter(ele => !!ele);
		}
		return [];
	} catch (err) {
		console.log(err);
		return [];
	}
}

export function saveClick(name: string, link: string): TUseItem {
	const data = getKeys();

	const lasttime = new Date().getTime();
	const i = data.findIndex(ele => ele.name === name);
	if (i > -1) {
		data[i] = {
			name,
			link,
			lasttime,
			count: data[i].count + 1
		};
	} else {
		data.push({
			name,
			link,
			lasttime,
			count: 1
		});
	}

	localStorage.setItem(HABIT_KEY, JSON.stringify(data));

	if (i > -1) {
		return data[i];
	} else {
		return data.slice(-1)[0];
	}
}

export function findAllMenu(data: MenuItem[], keyword: string = ""): MenuItem[] {
	return data
		.reduce((prev, next) => {
			const { submenu, name, href, system } = next;
			let d: MenuItem[] = [];

			if (submenu && submenu.length > 0) {
				d = d.concat(findAllMenu(submenu, keyword));
			} else {
				d.push({
					name,
					href,
					system
				} as MenuItem);
			}

			return [...prev, ...d];
		}, [])
		.filter(ele => {
			const { name, href, system } = ele;
			if (keyword) {
				const kreg = new RegExp(keyword, "gi");
				return [name, href, system].some(t => kreg.test(t));
			}

			return true;
		});
}

const litbCountries = require("utils/litbCountry.json");

export interface LitbCountryG {
	country: string;
	curreny: string;
	c3: string;
	[key: string]: string;
}

export interface SYSTEMTYPE {
	type: string;
	name: string;
}

export interface Country {
	code: string;
	name: string;
}
export const EzCountrys: Country[] = [
	{
		code: "SG",
		name: "Singapore"
	},
	{
		code: "MY",
		name: "Malaysia"
	},
	{
		code: "TH",
		name: "Thailand"
	},
	{
		code: "PK",
		name: "Pakistan"
	},
	{
		code: "TWC",
		name: "TaiWan-China"
	}
];

export const MultCountry: Country[] = [
	{
		code: "HK",
		name: "Hong Kong-China"
	},
	{
		code: "AUE",
		name: "Australia"
	}
];

export const CountryEnum = {
	CN: "CN",
	PK: "PK",
	SG: "SG",
	MY: "MY",
	US: "US",
	KR: "KR",
	JP: "JP",
	PKLocal: "PKLocal",
	AllNum: 0,
	PKNum: 600000000
};

export const countryCodes = (window.ezConfig && window.ezConfig.country.split(",")) || []; // 默认国家是第一个

export const controlCountry =
	countryCodes && countryCodes[0] === "PK"
		? ["PK"]
		: countryCodes[0] === "TWC"
		? ["TWC"]
		: ["SG", "MY", "TH"];

export const controlAllCountry: string[] =
	countryCodes[0] === "PK"
		? ["PK"]
		: countryCodes[0] === "TWC"
		? ["TWC"]
		: ["SG", "MY", "TH", "ID", "PK", "TWC"];

export const countrys = countryCodes.map(ele => EzCountrys.find(e => e.code === ele));

export const sellerCountryCodes = ["SGLocal", "MYLocal", "CN", "US", "KR", "PKLocal"];
export const userToSellerCountryCodes = {
	PK: [{ key: "PKLocal", num: 600000000, value: "PKLocal" }],
	SG: [
		{ key: "All", ad: "All", num: 0, value: "All" },
		{ key: "CN", ad: "CN", num: 1, value: "CN" },
		{ key: "US", ad: "US", num: 3, value: "US" },
		{ key: "KR", ad: "KR", num: 4, value: "KR" },
		{ key: "JP", ad: "JP", num: 5, value: "JP" },
		{ key: "SGLocal", ad: "SG", num: 100000000, value: "SGLocal" },
		{ key: "MYLocal", ad: "MY", num: 200000000, value: "MYLocal" },
		{ key: "PKLocal", ad: "PK", num: 600000000, value: "PKLocal" }
	]
};

export const ezCountryCodes = EzCountrys.map(({ code }) => code);
export const multCountryCodes = MultCountry.map(c => c.code);

export const hasPrimeCodes = ["SG", "MY", "TH"];

export function findCodeName(code: string): string {
	const item = countrys.find(ele => ele.code === code);

	if (item) {
		return item.name;
	} else {
		return "";
	}
}
/**
 * 兰亭地址栏title特殊处理
 * @param piece
 */
export function transformParamsForSeo(piece: string = "") {
	return piece.toLocaleLowerCase().replace(/\s*(&|\/)\s*|'|\s+/g, "-");
}
/**
 * 兰亭三位国家码映射为国家全称
 */
export const fetchLitbCountry = (code: Array<string>) => {
	if (!code) {
		return [];
	}
	let result = litbCountries.filter((c: LitbCountryG) => code.includes(c.c3));
	result = litbGroupFortmat(result);
	return result || [];
};
export const litbGroupFortmat = (group: LitbCountryG[]) => {
	if (!group || group.length === 0) {
		return [];
	}
	return group.map((nc: LitbCountryG) => {
		const { c3, country, currency } = nc;
		return {
			code: c3,
			name: country,
			currency
		};
	});
};
// 获取多国家站 国家
export const fetchMultCountry = (code: Array<string>): Country[] => {
	if (!code) {
		return [];
	}
	const result = MultCountry.filter(c => code.includes(c.code));
	return result || [];
};

export const fetchSystem = () => {
	return window.sessionStorage.getItem("system");
};

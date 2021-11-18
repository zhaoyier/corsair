export const PRICE_SYMBOL_SG = "S$";
export const PRICE_SYMBOL_MY = "RM";
export const PRICE_SYMBOL_AU = "AU$";
export const PRICE_SYMBOL_ID = "Rp";
export const PRICE_SYMBOL_TH = "฿";
export const PRICE_SYMBOL_PK = "Rs.";
export const PRICE_SYMBOL_TWC = "NT$";
export const PRICE_SYMBOL_KR = "₩";
export const PRICE_SYMBOL_CN = "￥";
export const PRICE_SYMBOL_US = "$";
export const PRICE_SYMBOL_JP = "円";
export const PRICE_SYMBOL_AUE = "A$";

// 根据国家 获取 价格单位
export function getPriceSymbol(code: string = "SG") {
	switch (code.toUpperCase()) {
		case "SGD":
		case "SG":
		case "SGLocal":
			return PRICE_SYMBOL_SG;
		case "MYR":
		case "MY":
		case "MYLocal":
			return PRICE_SYMBOL_MY;
		case "THB":
		case "TH":
			return PRICE_SYMBOL_TH;
		case "IDR":
		case "ID":
			return PRICE_SYMBOL_ID;
		case "PKR":
		case "PK":
		case "PKLocal":
			return PRICE_SYMBOL_PK;
		case "NTD":
		case "TW":
		case "TWC":
			return PRICE_SYMBOL_TWC;
		case "KR":
			return PRICE_SYMBOL_KR;
		case "CN":
			return PRICE_SYMBOL_CN;
		case "USD":
		case "US":
			return PRICE_SYMBOL_US;
		case "JP":
			return PRICE_SYMBOL_JP;
		case "AUD":
			return PRICE_SYMBOL_AUE;
		default:
			return PRICE_SYMBOL_SG;
	}
}

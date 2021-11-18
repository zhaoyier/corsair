export const LANGS = {
	EN: "en",
	ZH: "zh",
	MS: "ms",
	TH: "th",
	ZHTW: "zhtw",
	FR: "fr",
	ES: "es",
	DE: "de",
	IT: "it",
	PT: "pt",
	JA: "ja",
	RU: "ru",
	NL: "nl",
	AR: "ar",
	NO: "no",
	DA: "da",
	SV: "sv",
	KO: "ko",
	FI: "fi",
	HE: "he",
	TR: "tr",
	PL: "pl",
	CZ: "cz",
	GR: "gr",
	HR: "hr",
	RO: "ro",
	HU: "hu"
};
export const langNames = {
	[LANGS.ZH]: "简体中文",
	[LANGS.EN]: "English",
	[LANGS.MS]: "Bahasa Melayu",
	[LANGS.TH]: "ภาษาไทย",
	[LANGS.ZHTW]: "繁体中文"
};
export const litbLangNames = {
	[LANGS.EN]: "English",
	fr: "France",
	es: "Spanish",
	de: "Germany",
	it: "Italian",
	pt: "Portugal",
	ja: "Japanese",
	ru: "Russian",
	nl: "Nederlands",
	ar: "Arabic",
	no: "Norwegian",
	da: "Danish",
	sv: "Svenska",
	ko: "Korean",
	fi: "Finnish",
	he: "Hebrew",
	tr: "Turkish",
	pl: "Polish",
	cz: "Czech",
	gr: "Greek",
	hr: "Croatian",
	ro: "Romanian",
	hu: "Hungarian",
	[LANGS.TH]: "Thai"
};

export const langsBook = {
	SG: [LANGS.EN, LANGS.ZH],
	MY: [LANGS.EN, LANGS.MS, LANGS.ZH],
	TH: [LANGS.EN, LANGS.TH],
	PK: [LANGS.EN],
	TWC: [LANGS.EN, LANGS.ZHTW]
};

export function getLangsByCode(code: string) {
	return langsBook[code] || [LANGS.EN];
}

export function getAllLangCodes() {
	return Object.keys(langsBook).reduce((pValue, cValue) => {
		return [...pValue, ...langsBook[cValue].filter(code => !pValue.includes(code))];
	}, []);
}

export function hasLan(country?: string) {
	const ctry = country || window.sessionStorage.getItem("catalog") || "SG";
	return langsBook[ctry];
}

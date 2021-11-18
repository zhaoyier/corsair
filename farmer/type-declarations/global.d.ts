declare module "*.scss" {
	const content: any;
	export default content;
}

interface Window {
	require: any;
	AdminEntryInfo: string;
	monaco: any;
	ezConfig: {
		country: string;
	};
}

declare module "classnames";

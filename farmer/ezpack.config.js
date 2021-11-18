const path = require("path");
const fs = require("fs");
const { validateConfig } = require("@ezbuy/ezpack");
const antdEntry = require("./config/antd.entry");
// const antdCSSEntry = require("antd/dist/antd.min.css");

const REPOSITORY = "ezship";
const UAT_REPOSITORY = REPOSITORY + ".65emall.net";
const ONLINE_REPOSITORY = REPOSITORY + ".ezbuy.com";

const devPort = "8990";
const resolve = p => path.resolve(__dirname, p);
const srcPath = resolve("./src");
const entryPath = path.join(srcPath, "entries");
const apiPath = resolve("./src/genServices");
const globalCSSPath = [];

const entry = {};
fs.readdirSync(entryPath)
	.filter(file => file[0] !== ".")
	.forEach(file => {
		entry[file.replace(".tsx", "")] = path.join(entryPath, file);
	});

const regionInProcess = process.env.COUNTRY || process.env.REGION;
const region =
	typeof regionInProcess === "string" ? regionInProcess.toLowerCase() : regionInProcess; // PK
const isSpecificCountry = typeof region === "string";
// const repoPrefix = isSpecificCountry ? region.slice(0, 2) : "";

module.exports = validateConfig({
	name: REPOSITORY,
	webpack: {
		rootPath: srcPath,
		tsInclude: [srcPath],
		jsInclude: [],
		entry,
		dllEntry: {
			vendors: [
				"babel-polyfill",
				"react",
				"react-dom",
				"react-router-dom",
				"react-loadable",
				"classnames",
				"whatwg-fetch",
				"mobx",
				"mobx-react",
				"moment"
			],
			antd: antdEntry,
			antdStyles: ["antd/dist/antd.min.css"]
		},
		outputPath: path.resolve(__dirname, "./dist"),
		devPort,
		cdnPath: "/",
		cssModulePath: {
			include: srcPath,
			exclude: globalCSSPath
		},
		cssGlobalPath: {
			include: globalCSSPath
		},
		html: true,
		htmlOption: {
			title: "Ezship system",
			template: resolve("./src/index.html"),
			filename: "[name].html",
			customTplData: {
				region: isSpecificCountry ? region.toUpperCase() : "SG,MY,TH,ID,TWC,PK,AUE"
			}
		},
		copyPath: [
			{
				from: resolve("./src/assets"),
				to: "assets"
			}
		]
	},
	devServer: {
		port: devPort,
		initEntry: ["index", "login"],
		contentBase: [{ path: resolve("./src/assets"), prefix: "/assets" }],
		defaultProxyOrigin: {
			webapi: "dgadmin.65emall.net"
		},
		apiDefPath: apiPath
	},
	publish: {
		onlinePath: path.resolve("../", ONLINE_REPOSITORY),
		uatPath: path.resolve("../", UAT_REPOSITORY),
		uatEnv: [`${UAT_REPOSITORY}`, ...[2].map(num => `${REPOSITORY}${num}.65emall.net`)],
		jsDir: "js"
	}
});

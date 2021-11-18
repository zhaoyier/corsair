import { TAdminUploadInfo, GetUploadInfo } from "genServices/AdminHomepageService";

export function getToken(cb?: (info: TAdminUploadInfo) => void) {
	const errInfo: TAdminUploadInfo = {
		baseUrl: "",
		token: "",
		uploadUrl: ""
	};
	const errCb = (err: any) => {
		console.warn(`获取七牛的upload token错误, 错误信息是\n`);
		console.warn(err);

		if (cb) {
			cb(errInfo);
		}

		return errInfo;
	};

	return GetUploadInfo("")
		.then(data => {
			if (cb) {
				cb(data);
			}
			return data;
		}, errCb)
		.catch(errCb);
}

export default getToken;

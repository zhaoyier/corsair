import { checkStatus } from "utils/webapi";
import { GetUploadInfo } from "genServices/AdminHomepageService";

// 上传base64到七牛
export async function uploadBase64(pic): Promise<{ baseUrl: string; hash: string; key: string }> {
	const info = await GetUploadInfo("");
	return new Promise((resolve, reject) => {
		const xhr = new XMLHttpRequest();
		xhr.open("POST", info.uploadUrl + "/putb64/-1", true);
		xhr.setRequestHeader("Content-Type", "application/octet-stream");
		xhr.setRequestHeader("Authorization", `UpToken ${info.token}`);
		xhr.onreadystatechange = function() {
			if (this.readyState === 4) {
				if (this.status === 200) {
					resolve({ ...JSON.parse(this.responseText), status: this.status });
				} else {
					reject({ ...JSON.parse(this.responseText), status: this.status });
				}
			}
		};
		xhr.send(pic);
	})
		.then(checkStatus)
		.then(
			response =>
				new Promise(resolve => {
					resolve({ data: response, status: response.status });
				}) as Promise<{
					data: any;
					status: number;
				}>
		)
		.then(resp => {
			return { baseUrl: info.baseUrl, hash: resp.data.hash, key: resp.data.key };
		});
}

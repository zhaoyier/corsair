import { notification } from "antd";

/**
 * @param fn 需要执行的函数,必须是带有参数的Peomise函数
 * @param title 接口执行的提示事件名
 * @param isSuccess 成功是否需要提示
 * @returns promise函数执行后的结果 成功执行返回函数结果（ Promise )，失败则返回null并给出提示
 */
export async function tryCatchResult<T>(fn: Promise<T>, title?: string, isSuccess = false) {
	try {
		const result = await fn;
		if (result) {
			if (isSuccess) {
				notification.success({
					message: `${title} is Successful`,
					description: `${result}`
				});
			}
			return result;
		}
		return null;
	} catch (e) {
		notification.error({
			message: `${title} is Worry`,
			description: e.message
		});
		return null;
	}
}
// return function wrapped(...Args: any[]): Promise<[Error | null, T]> {
// 	return (fn.apply(this, arguments) as Promise <T>).then(
// 		(res: T) => {
// 			if (needSuccNotif) {
// 				notification.error({
// 					message: `${title} is Successful`,
// 					description: `${res}`
// 				});
// 			}
// 			return [null, res] as [null, T];
// 		},
// 		(err: Error) => {
// 			console.error(err);
// 			if (needSuccNotif) {
// 				notification.error({
// 					message: `${title} is Worry`,
// 					description: err.message
// 				});
// 			}
// 			return [err, {}] as [Error, T];
// 		}
// 	);
// };
/**
 * 用来捕获promise函数返回的一次性结果，只有成功与失败的接口适用，用于通用的返回成功或失败的接口
 * @param fn 需要执行的函数,必须是带有参数的Promise函数
 */
export async function tryCatchMessage(fn: Promise<any>) {
	try {
		const result = await fn;
		if (result.result || result) {
			notification.success({
				message: `${result.msg || "success"}`
			});
			return result.result;
		}
	} catch (e) {
		notification.error({
			message: `${e.message}`,
			description: e.message
		});
		return e.result;
	}
}

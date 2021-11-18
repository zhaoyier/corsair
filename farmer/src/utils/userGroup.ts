// import { GetGroupList, Bool, RetCode, GroupState } from "genServices/ugm/ugm";
// import { message } from "antd";

// export const getGroupByCatalog = (catalog: string) => {
// 	// const prevGroup = window.sessionStorage.getItem("eUserGroup");

// 	// if (prevGroup) {
// 	// 	return new Promise(resolve => {
// 	// 		resolve(JSON.parse(prevGroup));
// 	// 	});
// 	// }
// 	return GetGroupList({
// 		catalog,
// 		isExpired: Bool.BoolFalse,
// 		states: [GroupState.GroupStateReady]
// 	})
// 		.then(res => {
// 			if (res && res.ret && res.ret.code === RetCode.Ok) {
// 				window.sessionStorage.setItem("eUserGroup", JSON.stringify(res.groups));
// 				return res.groups;
// 			}
// 			message.error(res.ret.errMsg);
// 			return null;
// 		})
// 		.catch(error => {
// 			message.error(error.toString());
// 			return null;
// 		});
// };

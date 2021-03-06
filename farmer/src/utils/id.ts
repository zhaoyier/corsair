// 参考： http://stackoverflow.com/questions/105034/create-guid-uuid-in-javascript
// 生成唯一的ID
export function genUUID() {
	return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, function(c) {
		let r = (Math.random() * 16) | 0,
			v = c === "x" ? r : (r & 0x3) | 0x8;
		return v.toString(16);
	});
}

export default genUUID;

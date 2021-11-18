import { observable, action } from "mobx";
import { message } from "antd";
import {
	GetDelayNoticeTemplate,
	GetDelayNotice,
	DelayNoticeInfo,
	CreateDelayNotice
} from "genServices/ezShipOMS/oms";
import { DelayNoticeTemplate } from "genServices/ezShipOMS/public";

export default class DelayNotificationModalStore {
	@observable loading: boolean = false;
	@observable mailChecked: boolean = true;
	@observable noticeTemplateList: DelayNoticeTemplate[] = []; // 通知模板列表
	@observable mailTemplateList: DelayNoticeTemplate[] = []; // 邮件模板列表
	@observable currentNoticeTemplate: DelayNoticeTemplate = {}; // 当前通知模板
	@observable currentMailTemplate: DelayNoticeTemplate = {}; // 当前邮件模板
	@observable noticeLogList: DelayNoticeInfo[] = []; // 通知记录列表
	// @observable packingNo: string = ""; // 该包裹号对应的封箱号
	// @observable packingNos: string[] = []; // 封箱号下所有的包裹

	@action
	changeStoreData = (param: string, value, subParam?: string, secParam?: string) => {
		if (subParam) {
			if (secParam) {
				this[param][subParam][secParam] = value;
			} else {
				this[param][subParam] = value;
			}
		} else {
			this[param] = value;
		}
	};

	// 获取模板
	@action
	getDelayNoticeTemplate = () => {
		GetDelayNoticeTemplate({})
			.then(resp => {
				if (resp.noticeTemplate && resp.noticeTemplate.length > 0) {
					this.changeStoreData("currentNoticeTemplate", { ...resp.noticeTemplate[0] });
				}
				if (resp.mailNoticeTemplate && resp.mailNoticeTemplate.length > 0) {
					this.changeStoreData("currentMailTemplate", { ...resp.mailNoticeTemplate[0] });
				}
				this.changeStoreData("noticeTemplateList", resp.noticeTemplate || []);
				this.changeStoreData("mailTemplateList", resp.mailNoticeTemplate || []);
			})
			.catch(err => message.error(err.message));
	};

	// 获取通知记录信息
	@action
	getDelayNotice = parcelCode => {
		GetDelayNotice({ parcelCode })
			.then(resp => {
				this.changeStoreData("noticeLogList", resp.info || []);
				// this.changeStoreData("packingNo", resp.packingNo);
				// this.searchPkgNo(resp.packingNo);
			})
			.catch(err => message.error(err.message));
	};

	// 获取封箱号下所有包裹
	// @action
	// searchPkgNo = packingNo => {
	// 	if (!packingNo) {
	// 		return;
	// 	}
	// 	// 延误通知下的封箱号搜索
	// 	DelayPackingNoSearch({ packingNo })
	// 		.then(resp => {
	//       this.changeStoreData("packingNos", resp.packingNos);
	// 		})
	// 		.catch(err => message.error(err.message));
	// };

	// 发起延误通知
	@action
	createDelayNotice = (parcelCodes: string[], cb: () => void) => {
		if (!this.currentNoticeTemplate.content) {
			message.warning("请填写通知内容！");
			return;
		}
		this.changeStoreData("loading", true);
		CreateDelayNotice({
			parcelCodes,
			// packingNos: new Array(this.packingNo),
			notice: this.currentNoticeTemplate.content,
			mailNotice: this.mailChecked ? this.currentMailTemplate.content : ""
		})
			.then(resp => {
				this.changeStoreData("loading", false);
				if (resp.result.result) {
					message.success("添加延误通知成功！");
					cb();
				} else {
					message.error(resp.result.msg);
				}
			})
			.catch(err => {
				message.error(err.message);
				this.changeStoreData("loading", false);
			});
	};
}

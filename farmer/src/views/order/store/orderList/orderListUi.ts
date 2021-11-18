import { action, observable } from "mobx";
import { SearchOrderReq } from "genServices/ezShipOMS/oms";
// import { CommonStore } from "../common";
export class OrderListUi {
	@observable offset: number = 0;
	@observable page: number = 1;
	@observable pageSize: number = 0;
	@observable total: number = 0;
	@observable limit: number = 50;
	@observable searchParams: SearchOrderReq;
	// @observable isFrontRemark: boolean = true;

	// @action
	// setIsFrontRemark = () => {
	// 	this.isFrontRemark = !this.isFrontRemark;
	// }
	@action
	setOrderListPage(pageSize, total) {
		this.pageSize = pageSize;
		this.total = total;
	}
	@action
	setOffset = (offset: number) => {
		this.offset = offset;
	};
	@action
	setPageSize = (pageSize: number) => {
		this.pageSize = pageSize;
	};
	@action
	setTotal = (total: number) => {
		this.total = total;
	};
	@action
	setPage = (page: number) => {
		this.page = page;
	};
}

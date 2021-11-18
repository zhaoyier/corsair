import { observable, action } from "mobx";
import { DeliveryTypes, getDeliveryTypeById, defaultAddresses } from "./constant";
import { message } from "antd";
import {
	DeliveryInfoHome,
	DeliveryInfoEzCollection,
	DeliveryInfoNeighbourhoodStation,
	DeliveryInfoMRT,
	DeliveryInfoSelfCollection,
	GetDefaultDeliveryInfoResp,
	TmsDeliveryHomeAddress,
	TmsDeliveryStation,
	TmsPickupPeriod,
	GetDefaultDeliveryInfo,
	GetDeliveryInfo,
	GetDeliveryNeighbourhoodStation,
	GetDeliveryPickupDate,
	GetDeliveryPickupPeriod,
	UpdateDeliveryInfo
} from "genServices/ezShipOMS/external";
import { deleteEmpty } from "utils/util";

export default class EditDeliveryModalStore {
	@observable disabled: boolean = false; // 不能修改配送信息
	@observable shipmentId: string = null;
	@observable deliveryTypeId: string = DeliveryTypes[Object.keys(DeliveryTypes)[0]].value; // id
	@observable deliveryTypeInfo: Partial<
		DeliveryInfoHome &
			DeliveryInfoEzCollection &
			DeliveryInfoNeighbourhoodStation &
			DeliveryInfoMRT &
			DeliveryInfoSelfCollection
	> = {};
	@observable info: GetDefaultDeliveryInfoResp = null; // 全部默认信息
	@observable defaultInfo:
		| DeliveryInfoHome
		| DeliveryInfoEzCollection
		| DeliveryInfoNeighbourhoodStation
		| DeliveryInfoMRT
		| DeliveryInfoSelfCollection = null; // 默认取货方式信息
	@observable defaultDeliveryTypeId: string = "";
	@observable defaultStationId: string = "";
	@observable defaultAddress: string = "";
	@observable loading: boolean = false;
	@observable addressList: TmsDeliveryHomeAddress[] = [];
	@observable neiStationList: TmsDeliveryStation[] = []; // 邻里站点名称列表
	@observable stationList: TmsDeliveryStation[] = []; // 站点名称列表
	@observable pickUpDatesList: string[] = []; // 选择日期列表
	@observable periodsList: TmsPickupPeriod[] = []; // 选择时间列表
	@observable isForcedChange: boolean = false;

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

	// 根据配送单号 -- 获取初始配送信息
	@action
	getDefaultDeliveryInfo = record => {
		this.changeStoreData("disabled", false);
		this.changeStoreData("loading", true);
		GetDefaultDeliveryInfo({ shipmentId: record.deliveryId })
			.then(resp => {
				if (!resp.editable) {
					message.warning("当前配送单不支持修改配送信息！");
					this.changeStoreData("disabled", true);
					this.changeStoreData("loading", false);
				}
				const deliveryTypeId = resp.deliveryTypeId || DeliveryTypes.Home.value;
				const deliveryResult = resp[getDeliveryTypeById(deliveryTypeId).result];
				const defaultInfo = {
					...deliveryResult,
					addresses:
						deliveryResult.address !== undefined
							? new Array(deliveryResult.address)
							: null,
					pickUpDates:
						deliveryResult.pickUpDate !== undefined
							? new Array(deliveryResult.pickUpDate)
							: null,
					PickupPeriods:
						deliveryResult.PickupPeriod !== undefined
							? new Array(deliveryResult.PickupPeriod)
							: null,
					stations:
						deliveryResult.station !== undefined
							? new Array(deliveryResult.station)
							: null,
					neighbourhoodStations:
						deliveryResult.neighbourhoodStation !== undefined
							? new Array(deliveryResult.neighbourhoodStation)
							: null
				};
				[
					"address",
					"pickUpDate",
					"PickupPeriod",
					"station",
					"neighbourhoodStation"
				].forEach(x => delete defaultInfo[x]);
				const totalData = {
					deliveryTypeId,
					deliveryTypeInfo: defaultInfo,
					defaultInfo,
					info: resp,
					defaultDeliveryTypeId: deliveryTypeId,
					defaultStationId: deliveryResult.station ? deliveryResult.station.stationId : 0,
					defaultAddress: deliveryResult.addresses
						? deliveryResult.addresses.addressId
						: "",
					loading: false
				};
				this.getdefaultData(totalData, record);
			})
			.catch(err => {
				message.error(err.message);
				this.changeStoreData("loading", false);
			});
	};

	@action
	getdefaultData = (totalData, record) => {
		const {
			deliveryTypeId,
			deliveryTypeInfo,
			defaultInfo,
			info,
			defaultDeliveryTypeId,
			defaultStationId,
			defaultAddress,
			loading
		} = totalData;
		this.shipmentId = record.deliveryId;
		this.deliveryTypeId = deliveryTypeId;
		this.deliveryTypeInfo = deliveryTypeInfo;
		this.defaultInfo = defaultInfo;
		this.info = info;
		this.defaultDeliveryTypeId = defaultDeliveryTypeId;
		this.defaultStationId = defaultStationId;
		this.defaultAddress = defaultAddress;
		this.loading = loading;
		this.getDeliveryInfo(deliveryTypeId, record.deliveryId);
	};

	// 获取配送信息|切换运输方式联动配送信息
	@action
	getDeliveryInfo = (deliveryTypeId: string, shipmentId?: string) => {
		this.changeStoreData("loading", true);
		GetDeliveryInfo({ shipmentId: shipmentId ? shipmentId : this.shipmentId, deliveryTypeId })
			.then(resp => {
				const deliveryResult = resp[getDeliveryTypeById(deliveryTypeId).result];
				let neiStationList = [],
					pickUpDatesList = [],
					addressList = [],
					periodsList = [],
					deliveryTypeInfo = deliveryResult;
				switch (deliveryTypeId) {
					case DeliveryTypes.Home.value:
						pickUpDatesList = deliveryResult.pickUpDates;
						periodsList = deliveryResult.PickupPeriods.map(x => ({
							periodName: x,
							periodId: ""
						}));
						addressList = deliveryResult.addresses;

						deliveryTypeInfo = Object.assign(deliveryResult, {
							addresses: new Array(defaultAddresses),
							pickUpDates: [],
							PickupPeriods: []
						});
						break;
					case DeliveryTypes.neighbourhoodStation.value:
						neiStationList = deliveryResult.neighbourhoodStations;
						deliveryTypeInfo = {
							...deliveryResult,
							neighbourhoodStations: [],
							stations: []
						};
						break;
					default:
						deliveryTypeInfo = { ...deliveryResult, stations: [] };
						break;
				}
				const newDeliveryTypeInfo =
					deliveryTypeId === this.defaultDeliveryTypeId
						? this.defaultInfo || deliveryTypeInfo
						: deliveryTypeInfo;
				this.changeStoreData("deliveryTypeId", deliveryTypeId);
				this.changeStoreData("deliveryTypeInfo", newDeliveryTypeInfo);
				this.changeStoreData("neiStationList", neiStationList);
				this.changeStoreData("stationList", deliveryResult.stations || []);
				this.changeStoreData("pickUpDatesList", pickUpDatesList);
				this.changeStoreData("periodsList", periodsList);
				this.changeStoreData("addressList", addressList);
				this.changeStoreData("loading", false);
				if (
					deliveryTypeId === this.defaultDeliveryTypeId &&
					deliveryTypeId === DeliveryTypes.neighbourhoodStation.value
				) {
					const defaultNeibour = this.info.infoNeighbourhoodStation;
					if (defaultNeibour && defaultNeibour.neighbourhoodStation) {
						this.onRegionChange(defaultNeibour.neighbourhoodStation.stationId, true);
					}
				}
			})
			.catch(err => {
				message.error(err.message);
				this.changeStoreData("loading", false);
			});
	};

	// 邻里:区域改变，站点名称列表改变
	@action
	onRegionChange = (val, isDefault?: boolean) => {
		if (!val) {
			return;
		}
		// isDefault=true邻里站点地址有默认值时展示，切换时清空
		GetDeliveryNeighbourhoodStation({
			deliveryTypeId: DeliveryTypes.neighbourhoodStation.value,
			regionId: val
		})
			.then(resp => {
				const deliveryTypeInfo = isDefault
					? this.deliveryTypeInfo
					: {
							...this.deliveryTypeInfo,
							stations: [],
							pickUpDates: [],
							PickupPeriods: []
							// tslint:disable-next-line:indent
					  };
				this.changeStoreData("stationList", resp.stations);
				this.changeStoreData("periodsList", []);
				this.changeStoreData("deliveryTypeInfo", deliveryTypeInfo);
			})
			.catch(err => message.error(err.message));
	};

	// 取货方式改变
	@action
	onDeliveryChange = e => {
		this.isForcedChange = false;
		this.deliveryTypeId = e.target.value;
		this.getDeliveryInfo(e.target.value);
	};

	// 站点名称改变，送货日期范围改变
	@action
	onStationChange = () => {
		if (this.deliveryTypeInfo && this.deliveryTypeId !== DeliveryTypes.ezcollection.value) {
			// ezCollection,没有日期选择
			const hasStationId =
				this.deliveryTypeInfo.stations && this.deliveryTypeInfo.stations.length > 0;
			if (hasStationId) {
				GetDeliveryPickupDate({
					deliveryTypeId: this.deliveryTypeId,
					stationsId: hasStationId ? this.deliveryTypeInfo.stations[0].stationId : ""
				})
					.then(resp => {
						this.changeStoreData("pickUpDatesList", resp.pickUpDates);
					})
					.catch(err => message.error(err.message));
			}
		}
	};

	// 送货日期改变，送货时间段联动，除了home
	@action
	onPickUpDateChange = () => {
		if (this.deliveryTypeId === DeliveryTypes.Home.value) {
			return;
		}
		const hasStationId =
			this.deliveryTypeInfo.stations && this.deliveryTypeInfo.stations.length > 0;
		if (hasStationId) {
			GetDeliveryPickupPeriod({
				deliveryTypeId: this.deliveryTypeId,
				stationsId: hasStationId ? this.deliveryTypeInfo.stations[0].stationId : "",
				pickupDate: this.deliveryTypeInfo.pickUpDates[0]
			})
				.then(resp => {
					this.changeStoreData("periodsList", resp.periods);
				})
				.catch(err => {
					message.error(err.message);
					this.changeStoreData("periodsList", []);
				});
		}
	};

	getInfo = () => {
		const type = getDeliveryTypeById(this.deliveryTypeId).result;
		console.log(type);
		const {
			pickUpDates,
			PickupPeriods,
			addresses,
			addressToName,
			addressToPhone,
			zipCode,
			street,
			block,
			unitStart,
			unitEnd,
			companyName,
			buildingName,
			stations,
			neighbourhoodStations
		} = this.deliveryTypeInfo;
		const newPickUpDates =
			pickUpDates && pickUpDates.length > 0 && pickUpDates.filter(elem => Number(elem) > 0);
		const newPickupPeriods =
			PickupPeriods &&
			PickupPeriods.length > 0 &&
			PickupPeriods.filter(elem => elem.length > 0);
		switch (type) {
			case "infoHome":
				return {
					pickUpDates: newPickUpDates,
					PickupPeriods: newPickupPeriods,
					addresses:
						addresses && addresses.length === 0
							? [{ addressId: "0", addressName: "" }]
							: addresses,
					addressToName,
					addressToPhone,
					zipCode,
					street,
					block,
					unitStart,
					unitEnd,
					companyName,
					buildingName
				};
			case "infoEzCollection":
				return {
					stations,
					addressToName,
					addressToPhone
				};
			case "infoNeighbourhoodStation":
				return {
					pickUpDates: newPickUpDates,
					PickupPeriods: newPickupPeriods,
					neighbourhoodStations,
					stations,
					addressToName,
					addressToPhone
				};
			case "infoMRT":
				return {
					pickUpDates: newPickUpDates,
					PickupPeriods: newPickupPeriods,
					stations,
					addressToName,
					addressToPhone
				};
			case "infoSelfCollection":
				return {
					pickUpDates: newPickUpDates,
					PickupPeriods: newPickupPeriods,
					stations,
					addressToName,
					addressToPhone
				};
			default:
				return null;
		}
	};

	// 更新配送信息
	@action
	updateDelivery = (cb: () => void) => {
		this.changeStoreData("loading", true);
		if (
			this.isForcedChange &&
			this.deliveryTypeId === DeliveryTypes.Home.value &&
			!this.deliveryTypeInfo.PickupPeriods[0]
		) {
			message.error("强制大小货时请选择送货时间！");
			this.changeStoreData("loading", false);
			return;
		}
		const updateDeliveryInfo = {
			shipmentId: this.shipmentId, // 配送单号
			deliveryTypeId: this.deliveryTypeId, // 配送方式
			[getDeliveryTypeById(this.deliveryTypeId).result]: deleteEmpty(this.getInfo()),
			defaultDeliveryTypeId: this.defaultDeliveryTypeId, // 默认配送方式
			defaultStationId: this.defaultStationId, // 默认地址站点Id
			defaultAddress: this.defaultAddress, // 默认地址名
			isForcedChange: this.isForcedChange // 是否无视重量，强制指定大小货送货时间
		};
		console.log(deleteEmpty(updateDeliveryInfo));
		UpdateDeliveryInfo(deleteEmpty(updateDeliveryInfo))
			.then(resq => {
				if (resq.result.code === "0") {
					message.success("update Delivery successs!");
					cb();
				} else {
					message.error(
						resq.result.message ? resq.result.message : "update Delivery failed!"
					);
				}
				this.changeStoreData("loading", false);
			})
			.catch(err => {
				message.error(err.message ? err.message : "update Delivery failed!");
				this.changeStoreData("loading", false);
			});
	};
}

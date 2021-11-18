import { TmsPickupPeriod, TmsDeliveryHomeAddress } from "genServices/mithril/tms";

const homeArr = [
	"PickupPeriods",
	"pickUpDates",
	"isForcedChange",
	"addresses",
	"addressToName",
	"addressToPhone",
	"zipCode",
	"street",
	"block",
	"unitStart",
	"companyName",
	"buildName"
];
const ezcollectionArr = ["stations", "addressToName", "addressToPhone", "stationAddress"];
const neighbourhoodStationArr = [
	"neighbourhoodStations",
	"stations",
	"PickupPeriods",
	"pickUpDates",
	"addressToName",
	"addressToPhone",
	"stationAddress"
];
const mRTArr = [
	"stations",
	"PickupPeriods",
	"pickUpDates",
	"addressToName",
	"addressToPhone",
	"stationAddress"
];
const selfCollectionArr = [
	"stations",
	"PickupPeriods",
	"pickUpDates",
	"addressToName",
	"addressToPhone",
	"stationAddress"
];
export const totalArr = [
	"PickupPeriods",
	"pickUpDates",
	"isForcedChange",
	"addresses",
	"addressToName",
	"addressToPhone",
	"zipCode",
	"street",
	"block",
	"unitStart",
	"companyName",
	"buildName",
	"stations",
	"stationAddress",
	"neighbourhoodStations",
	"PickupPeriods",
	"pickUpDates"
];
export const DeliveryTypes = {
	Home: {
		value: "4",
		code: "Home",
		label: "Home",
		fields: homeArr,
		result: "infoHome"
	},
	ezcollection: {
		value: "5",
		code: "ezCollection",
		label: "ezcollection",
		fields: ezcollectionArr,
		result: "infoEzCollection"
	},
	neighbourhoodStation: {
		value: "2",
		code: "NeighbourhoodStation",
		label: "邻里",
		fields: neighbourhoodStationArr,
		result: "infoNeighbourhoodStation"
	},
	MRT: {
		value: "3",
		code: "MRT",
		label: "MRT",
		fields: mRTArr,
		result: "infoMRT"
	},
	SelfCollection: {
		value: "1",
		code: "SelfCollection",
		label: "自取仓",
		fields: selfCollectionArr,
		result: "infoSelfCollection"
	}
};

export function getDeliveryTypeById(id: string) {
	if (!id) {
		return "";
	}
	const item = Object.keys(DeliveryTypes).find(item => DeliveryTypes[item].value === id);
	return item ? DeliveryTypes[item] : "";
}
export function getDeliveryTypeByCode(code: string) {
	if (!code) {
		return "";
	}
	const item = Object.keys(DeliveryTypes).find(item => DeliveryTypes[item].code === code);
	return item ? DeliveryTypes[item].id : "";
}
// 全部时间段 ，强制修改大小货的时候展示
export const allPickTime: TmsPickupPeriod[] = [
	{
		periodId: "",
		periodName: "10:00 AM – 6:00 PM"
	},
	// {
	// 	periodId: "",
	// 	periodName: "6:30 PM – 10:30 PM"
	// },
	{
		periodId: "",
		periodName: "9:00 AM – 4:00 PM"
	},
	{
		periodId: "",
		periodName: "3:00 PM – 10:00 PM"
	}
];

export const defaultAddresses: TmsDeliveryHomeAddress = {
	addressId: "0",
	addressName: ""
};

export const emptyPickupPeriods: TmsPickupPeriod = {
	periodId: "",
	periodName: ""
};

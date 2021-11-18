/**
 * This file is auto-generated by protobufgen
 * Don't change manually
 */

import { Result as EzShipOMSResult, RemarkTypeEnum as EzShipOMSRemarkTypeEnum } from "../ezShipOMS/public";
import { Empty as CommonEmpty } from "../common/empty";
import webapi from "../webapi";

export enum ServiceType {
	PHOTOGRAPH = "PHOTOGRAPH",
	REPACKAGE = "REPACKAGE",
}


export enum ArrivalType {
	ARRIVAL_TYPE_TRANSPORT = "ARRIVAL_TYPE_TRANSPORT",
	ARRIVAL_TYPE_PURCHASE = "ARRIVAL_TYPE_PURCHASE",
	ARRIVAL_TYPE_ALLOT = "ARRIVAL_TYPE_ALLOT",
	ARRIVAL_TYPE_ALLOT_WITHDRAW = "ARRIVAL_TYPE_ALLOT_WITHDRAW",
}


export enum OrderArrivalStatus {
	ARRIVAL_STATUS_CREATED = "ARRIVAL_STATUS_CREATED",
	ARRIVAL_STATUS_SIGN = "ARRIVAL_STATUS_SIGN",
	ARRIVAL_STATUS_INSTORAGE = "ARRIVAL_STATUS_INSTORAGE",
	ARRIVAL_STATUS_UPSHELF = "ARRIVAL_STATUS_UPSHELF",
	ARRIVAL_STATUS_CANCELED = "ARRIVAL_STATUS_CANCELED",
}


export enum StockOutType {
	STOCK_OUT_ORDER = "STOCK_OUT_ORDER",
	STOCK_OUT_ALLOT = "STOCK_OUT_ALLOT",
	STOCK_OUT_WITHDRAW = "STOCK_OUT_WITHDRAW",
}


export enum ServiceStatus {
	SERVICE_STATUS_CANCEL = "SERVICE_STATUS_CANCEL",
	SERVICE_STATUS_CREATED = "SERVICE_STATUS_CREATED",
	SERVICE_STATUS_DONE = "SERVICE_STATUS_DONE",
}


export enum StockOutStatus {
	STOCK_OUT_CREATED = "STOCK_OUT_CREATED",
	STOCK_OUT_WAVE = "STOCK_OUT_WAVE",
	STOCK_OUT_PICKED = "STOCK_OUT_PICKED",
	STOCK_OUT_PACKED = "STOCK_OUT_PACKED",
	STOCK_OUT_WEIGH = "STOCK_OUT_WEIGH",
	STOCK_OUT_OUTSTOCK = "STOCK_OUT_OUTSTOCK",
	STOCK_OUT_CANCELDED = "STOCK_OUT_CANCELDED",
}


export enum SendMethod {
	DOOR_SMALL = "DOOR_SMALL",
	EZCOLLECTION = "EZCOLLECTION",
	MRT = "MRT",
	NEIGHBOR = "NEIGHBOR",
	WAREHOUSE = "WAREHOUSE",
	EZHOME = "EZHOME",
	DOOR_BIG = "DOOR_BIG",
}


export enum GoodsOrigin {
	GOODS_ORIGIN_LITB = "GOODS_ORIGIN_LITB",
	GOODS_ORIGIN_EZBUY = "GOODS_ORIGIN_EZBUY",
	GOODS_ORIGIN_EZSHIP = "GOODS_ORIGIN_EZSHIP",
}


export enum ALSMsgType {
	PARCEL = "PARCEL",
	ORDER = "ORDER",
	PICKUP = "PICKUP",
}


export enum ALSEventType {
	STATUS = "STATUS",
	PROPERTY = "PROPERTY",
	ARRIVED = "ARRIVED",
	EDITABLE = "EDITABLE",
	PLANDATE = "PLANDATE",
}


export enum OmsNotifyAlsType {
	OmsNotifyAlsType_Invalid = "OmsNotifyAlsType_Invalid",
	OmsNotifyAlsType_EditDelivery = "OmsNotifyAlsType_EditDelivery",
	OmsNotifyAlsType_CancelDelivery = "OmsNotifyAlsType_CancelDelivery",
}



export interface ModifyArrivalNoticeReq {
	notifyNumber?: string;
	status?: OrderArrivalStatus;
	warehouseCode?: string;
	trackNumber?: string;
}

export interface ModifyArrivalNoticeResp {
	result?: EzShipOMSResult;
}

export interface OMSUpdateArrivalNoticeReq {
	/**
	 * @pattern ^\d+$
	 */
	warehouseId?: string;
	warehouseCode?: string;
	trackingNo?: string;
	orderArrivalNo?: string;
	logisticsName?: string;
	arrivalType?: ArrivalType;
	goodsOrigin?: GoodsOrigin;
	order?: OrderParam;
}

export interface OMSUpdateArrivalNoticeResp {
	result?: EzShipOMSResult;
}

export interface OMSCancelArrivalNoticeReq {
	result?: EzShipOMSResult;
}

export interface OMSCancelArrivalNoticeResp {
	result?: EzShipOMSResult;
}

export interface WmsAddedServiceCreateReq {
	orderNo?: string;
	serviceType?: ServiceType;
	/**
	 * @pattern ^\d+$
	 */
	outValAddServiceId?: string;
}

export interface WmsAddServiceCreateResp {
	result?: EzShipOMSResult;
	data?: string;
}

export interface WMSCancelAddedServiceReq {
	str?: string;
}

export interface WMSCancelAddedServiceResp {
	result?: EzShipOMSResult;
}

export interface ModifyAddedServiceReq {
	/**
	 * @pattern ^\d+$
	 */
	addedServiceId?: string;
	status?: ServiceStatus;
	remark?: string;
	enclosureLink?: string[];
	/**
	 * @pattern ^\d+$
	 */
	succeedDate?: string;
	operator?: string;
}

export interface ModifyAddedServiceResp {
	result?: EzShipOMSResult;
}

export interface ModifyOutboundNotice {
	notifyNumber?: string;
	status?: StockOutStatus;
}

export interface ModifyOutboundNoticeReq {
	outboundNotices?: ModifyOutboundNotice[];
}

export interface ModifyOutboundNoticeResp {
	result?: EzShipOMSResult;
}

export interface Sku {
	skuNumber?: string;
	/**
	 * @pattern ^\d+$
	 */
	categoryId?: string;
}

export interface ModifyOrderReq {
	orderNumber?: string;
	transportMethod?: string[];
	reason?: string;
	/**
	 * @minimum 0
	 */
	weight?: number;
	/**
	 * @pattern ^\d+$
	 */
	length?: string;
	/**
	 * @pattern ^\d+$
	 */
	width?: string;
	/**
	 * @pattern ^\d+$
	 */
	height?: string;
	/**
	 * @minimum 0
	 */
	volumeWeight?: number;
	skus?: Sku[];
}

export interface ModifyOrderResp {
	result?: EzShipOMSResult;
}

export interface FrontRemarkReplyCountReq {
	orderNumber?: string[];
}

export interface FrontRemarkReplyCountResp {
	result?: EzShipOMSResult;
	data?: string[];
}

export interface AddFrontRemarkReq {
	remarkType?: EzShipOMSRemarkTypeEnum;
	remarkContent?: string;
	enclosureLink?: string[];
	/**
	 * @pattern ^\d+$
	 */
	createDate?: string;
	createBy?: string;
	orderNumber?: string;
}

export interface AddFrontRemarkResp {
	result?: EzShipOMSResult;
}

export interface OrderRemarksReq {
	orderNumber?: string[];
}

export interface RemarkHistory {
	remarkType?: EzShipOMSRemarkTypeEnum;
	remarkContent?: string;
	enclosureLink?: string;
	isReply?: boolean;
	/**
	 * @pattern ^\d+$
	 */
	createDate?: string;
	createBy?: string;
	orderNumber?: string;
}

export interface OrderRemarksResp {
	result?: EzShipOMSResult;
	data?: RemarkHistory[];
}

export interface OrderArrivalCreateParam {
	trackingNo?: string;
	outOrderArrivalNo?: string;
	logisticsName?: string;
	arrivalType?: ArrivalType;
	goodsOrigin?: GoodsOrigin;
	warehouseCode?: string;
	order?: OrderParam;
}

export interface OrderParam {
	nickname?: string;
	countryName?: string;
	countryCode?: string;
	orderNo?: string;
	identifier?: string;
	orderItems?: OrderItemParam[];
}

export interface OrderItemParam {
	itemImages?: string;
	itemName?: string;
	sku?: string;
	/**
	 * @minimum 0
	 */
	quantity?: number;
	/**
	 * @pattern ^\d+$
	 */
	declareCategoryId?: string;
	declareCategoryName?: string;
	/**
	 * @minimum 0
	 */
	unitPrice?: number;
	orderLineNo?: string;
	currencyCode?: string;
}

export interface OrderArrivalCreateResp {
	result?: EzShipOMSResult;
	data?: string;
}

export interface AddedServiceCreateParam {
	orderNo?: string;
	serviceType?: ServiceType;
	/**
	 * @pattern ^\d+$
	 */
	outValAddServiceId?: string;
}

export interface AddServiceCreateResp {
	result?: EzShipOMSResult;
	data?: string;
}

export interface StockOutCreateParam {
	param?: StockOutParam[];
}

export interface StockOutParam {
	invoiceNo?: string;
	goodsOrigin?: GoodsOrigin;
	stockOutType?: StockOutType;
	warehouseCode?: string;
	transportMethodCode?: string;
	sendMethod?: SendMethod;
	/**
	 * @pattern ^\d+$
	 */
	stationId?: string;
	stationCode?: string;
	countryCode?: string;
	routeCode?: string;
	routeName?: string;
	nickname?: string;
	receiverName?: string;
	stateName?: string;
	cityName?: string;
	address?: string;
	mobile?: string;
	itemList?: StockOutItemParam[];
	address2?: string;
	postcode?: string;
	/**
	 * @pattern ^\d+$
	 */
	expressId?: string;
}

export interface StockOutItemParam {
	orderLineNo?: string;
	sku?: string;
	orderNo?: string;
	/**
	 * @minimum 0
	 */
	quantity?: number;
	itemName?: string;
	itemAttribute?: string;
	/**
	 * @minimum 0
	 */
	unitPrice?: number;
}

export interface StockOutCreateResp {
	result?: EzShipOMSResult;
	data?: StockOutCreateModel;
}

export interface StockOutCreateModel {
	stockOutNo?: StockOutCreateResultModel[];
}

export interface StockOutCreateResultModel {
	needSplit?: boolean;
	invoiceNo?: string;
	stockOutNo?: string;
	splitOrderNos?: StockOutSplitOrderNosModel[];
}

export interface StockOutSplitOrderNosModel {
	orderNos?: string[];
}

export interface WMSParcel {
	parcelCode?: string;
	/**
	 * @minimum 0
	 */
	weight?: number;
	/**
	 * @minimum 0
	 */
	volumeWeight?: number;
	/**
	 * @pattern ^\d+$
	 */
	length?: string;
	/**
	 * @pattern ^\d+$
	 */
	width?: string;
	/**
	 * @pattern ^\d+$
	 */
	height?: string;
	orderNumber?: string[];
	trackingNumber?: string;
	parcelColor?: string;
}

export interface SyncParcelReq {
	parcels?: WMSParcel[];
	updateBy?: string;
}

export interface SyncParcelResp {
	result?: EzShipOMSResult;
}

export interface SealedOrShelfNumber {
	parcelCode?: string;
	sealedNumber?: string;
	shelfNumber?: string;
}

export interface SyncSealedOrShelfNumberReq {
	numbers?: SealedOrShelfNumber[];
	updateBy?: string;
}

export interface SyncSealedOrShelfNumberResp {
	result?: EzShipOMSResult;
}

export interface StringParam {
	str?: string;
}

export interface OMSCancelAddedServiceReq {
	wmsServiceId?: string;
}

export interface OMSCancelAddedServiceResp {
	result?: EzShipOMSResult;
}

export interface OMSCancelArrivalNoticeDocumentReq {
	/**
	 * @pattern ^\d+$
	 */
	orderId?: string;
}

export interface OMSCancelArrivalNoticeDocumentResp {
	result?: EzShipOMSResult;
}

export interface NoticeOMSLoseOrderReq {
	sku?: string;
}

export interface NoticeOMSLoseOrderResp {
	result?: EzShipOMSResult;
}

export interface Parcel {
	parcelCode?: string;
	trackNumber?: string;
	/**
	 * @pattern ^\d+$
	 */
	shippedDate?: string;
	/**
	 * @pattern ^\d+$
	 */
	logisticsCo?: string;
	desCode?: string;
	/**
	 * @pattern ^\d+$
	 */
	warehouseId?: string;
	/**
	 * @pattern ^\d+$
	 */
	status?: string;
	eventDesc?: string;
	/**
	 * @pattern ^\d+$
	 */
	eventTime?: string;
}

export interface ParcelsList {
	result?: EzShipOMSResult;
	data?: Parcel[];
}

export interface TrackByNumberAndTimeParam {
	tracknumber?: string;
	/**
	 * @pattern ^\d+$
	 */
	evenTime?: string;
}

export interface TrackByNumberModel {
	tracknumber?: string;
	status?: string;
	evenTime?: string;
	evenDesc?: string;
}

export interface TrackResult {
	/**
	 * @pattern ^\d+$
	 */
	Code?: string;
}

export interface ParcelTracks {
	result?: TrackResult;
	data?: TrackByNumberModel[];
}

export interface DeliveryLogisticsParam {
	/**
	 * @minimum 0
	 */
	length?: number;
	/**
	 * @minimum 0
	 */
	width?: number;
	/**
	 * @minimum 0
	 */
	high?: number;
	/**
	 * @minimum 0
	 */
	weight?: number;
	countryCode?: string;
	/**
	 * @pattern ^\d+$
	 */
	warehouse?: string;
	logisticsType?: string;
	deliveryMethod?: string;
	zipcode?: string;
	/**
	 * @pattern ^\d+$
	 */
	orderId?: string;
}

export interface DeliveryLogisticsResp {
	result?: EzShipOMSResult;
	data?: DeliveryLogisticsModel;
}

export interface DeliveryLogisticsModel {
	/**
	 * @pattern ^\d+$
	 */
	logisticsId?: string;
	logisticsName?: string;
	kpCode?: string;
}

export interface HscodeResp {
	result?: EzShipOMSResult;
	data?: HscodeModel[];
}

export interface HscodeModel {
	/**
	 * @pattern ^\d+$
	 */
	id?: string;
	codeName?: string;
	hscode?: string;
	codeNameCn?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	status?: number;
}

export interface LogisticsTypeModel {
	code?: string;
	name?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	status?: number;
	originId?: string;
	originName?: string;
}

export interface LogisticsTypeResp {
	data?: LogisticsTypeModel[];
}

export interface UpdateDeliveryInfoReq {
	/**
	 * @pattern ^\d+$
	 */
	shipmentId?: string;
	/**
	 * @pattern ^\d+$
	 */
	deliveryTypeId?: string;
	infoHome?: DeliveryInfoHome;
	infoEzCollection?: DeliveryInfoEzCollection;
	infoNeighbourhoodStation?: DeliveryInfoNeighbourhoodStation;
	infoMRT?: DeliveryInfoMRT;
	infoSelfCollection?: DeliveryInfoSelfCollection;
	/**
	 * @pattern ^\d+$
	 */
	defaultDeliveryTypeId?: string;
	/**
	 * @pattern ^\d+$
	 */
	defaultStationId?: string;
	defaultAddress?: string;
	isForcedChange?: boolean;
}

export interface UpdateDeliveryInfoResp {
	result?: EzShipOMSResult;
}

export interface DeliveryInfoReq {
	/**
	 * @pattern ^\d+$
	 */
	shipmentId?: string;
}

export interface GetDefaultDeliveryInfoResp {
	infoHome?: DefaultDeliveryInfoHome;
	infoEzCollection?: DefaultDeliveryInfoEzCollection;
	infoNeighbourhoodStation?: DefaultDeliveryInfoNeighbourhoodStation;
	infoMRT?: DefaultDeliveryInfoMRT;
	infoSelfCollection?: DefaultDeliveryInfoSelfCollection;
	/**
	 * @pattern ^\d+$
	 */
	weight?: string;
	items?: ParcelDeliveryItem[];
	editable?: boolean;
	/**
	 * @pattern ^\d+$
	 */
	deliveryTypeId?: string;
}

export interface DefaultDeliveryInfoHome {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDate?: string;
	PickupPeriod?: string;
	address?: TmsDeliveryHomeAddress;
	addressToName?: string;
	addressToPhone?: string;
	zipCode?: string;
	street?: string;
	block?: string;
	unitStart?: string;
	unitEnd?: string;
	companyName?: string;
	buildingName?: string;
}

export interface TmsDeliveryHomeAddress {
	/**
	 * @pattern ^\d+$
	 */
	addressId?: string;
	addressName?: string;
}

export interface DefaultDeliveryInfoEzCollection {
	station?: TmsDeliveryStation;
	addressToName?: string;
	addressToPhone?: string;
}

export interface TmsDeliveryStation {
	/**
	 * @pattern ^\d+$
	 */
	stationId?: string;
	stationName?: string;
	stationAddress?: string;
}

export interface DefaultDeliveryInfoNeighbourhoodStation {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDate?: string;
	PickupPeriod?: string;
	neighbourhoodStation?: TmsDeliveryStation;
	station?: TmsDeliveryStation;
	addressToName?: string;
	addressToPhone?: string;
}

export interface DeliveryInfoHome {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDates?: string[];
	PickupPeriods?: string[];
	addresses?: TmsDeliveryHomeAddress[];
	addressToName?: string;
	addressToPhone?: string;
	zipCode?: string;
	street?: string;
	block?: string;
	unitStart?: string;
	unitEnd?: string;
	companyName?: string;
	buildingName?: string;
}

export interface DeliveryInfoEzCollection {
	stations?: TmsDeliveryStation[];
	addressToName?: string;
	addressToPhone?: string;
}

export interface DeliveryInfoNeighbourhoodStation {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDates?: string[];
	PickupPeriods?: string[];
	neighbourhoodStations?: TmsDeliveryStation[];
	stations?: TmsDeliveryStation[];
	addressToName?: string;
	addressToPhone?: string;
}

export interface DefaultDeliveryInfoMRT {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDate?: string;
	PickupPeriod?: string;
	station?: TmsDeliveryStation;
	addressToName?: string;
	addressToPhone?: string;
}

export interface DefaultDeliveryInfoSelfCollection {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDate?: string;
	PickupPeriod?: string;
	station?: TmsDeliveryStation;
	addressToName?: string;
	addressToPhone?: string;
}

export interface ParcelDeliveryItem {
	packageNum?: string;
	deliverEta?: string;
	transportCode?: string;
	transport?: string;
	/**
	 * @pattern ^\d+$
	 */
	packageStatus?: string;
	packageStatusName?: string;
	/**
	 * @pattern ^\d+$
	 */
	warehouseId?: string;
	warehouse?: string;
}

export interface DeliveryInfoMRT {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDates?: string[];
	PickupPeriods?: string[];
	stations?: TmsDeliveryStation[];
	addressToName?: string;
	addressToPhone?: string;
}

export interface DeliveryInfoSelfCollection {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDates?: string[];
	PickupPeriods?: string[];
	stations?: TmsDeliveryStation[];
	addressToName?: string;
	addressToPhone?: string;
}

export interface GetDeliveryInfoReq {
	/**
	 * @pattern ^\d+$
	 */
	shipmentId?: string;
	/**
	 * @pattern ^\d+$
	 */
	deliveryTypeId?: string;
}

export interface GetDeliveryInfoResp {
	infoHome?: DeliveryInfoHome;
	infoEzCollection?: DeliveryInfoEzCollection;
	infoNeighbourhoodStation?: DeliveryInfoNeighbourhoodStation;
	infoMRT?: DeliveryInfoMRT;
	infoSelfCollection?: DeliveryInfoSelfCollection;
}

export interface GetDeliveryNeighbourhoodStationReq {
	/**
	 * @pattern ^\d+$
	 */
	deliveryTypeId?: string;
	/**
	 * @pattern ^\d+$
	 */
	regionId?: string;
}

export interface GetDeliveryNeighbourhoodStationResp {
	stations?: TmsDeliveryStation[];
}

export interface GetDeliveryPickupDateReq {
	/**
	 * @pattern ^\d+$
	 */
	deliveryTypeId?: string;
	/**
	 * @pattern ^\d+$
	 */
	stationsId?: string;
}

export interface GetDeliveryPickupDateResp {
	/**
	 * @pattern ^\d+$
	 */
	pickUpDates?: string[];
}

export interface GetDeliveryPickupPeriodReq {
	/**
	 * @pattern ^\d+$
	 */
	deliveryTypeId?: string;
	/**
	 * @pattern ^\d+$
	 */
	stationsId?: string;
	/**
	 * @pattern ^\d+$
	 */
	pickupDate?: string;
}

export interface GetDeliveryPickupPeriodResp {
	periods?: TmsPickupPeriod[];
}

export interface TmsPickupPeriod {
	/**
	 * @pattern ^\d+$
	 */
	periodId?: string;
	periodName?: string;
}

export interface AlsNotifyOmsAlterationReq {
	msgType?: ALSMsgType;
	msgId?: string;
	eventType?: ALSEventType;
	sign?: string;
	packNo?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	packStatus?: number;
	/**
	 * @pattern ^\d+$
	 */
	packStatusTime?: string;
	packStatusDesc?: string;
	signature?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	countryId?: number;
	shelfCode?: string;
	returned?: boolean;
	boxNo?: string;
	driverName?: string;
	driverMobile?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	sendCount?: number;
	remark?: string;
	stationName?: string;
	stationAddress?: string;
	images?: string;
	stationShelfCode?: string;
	orderNo?: string;
	editable?: boolean;
	/**
	 * @pattern ^\d+$
	 */
	orderStatusTime?: string;
	taskNo?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	taskStatus?: number;
	/**
	 * @pattern ^\d+$
	 */
	taskStatusTime?: string;
	/**
	 * @pattern ^\d+$
	 */
	arrived?: string;
	/**
	 * @pattern ^\d+$
	 */
	planDate?: string;
}

export interface NotifyResult {
	/**
	 * @pattern ^\d+$
	 */
	code?: string;
}

export interface AlsNotifyOmsAlterationResp {
	result?: NotifyResult;
}

export interface OmsNotifyAlsReq {
	alsOrderNumber?: string;
	address?: string;
	mobile?: string;
	receiverName?: string;
	postcode?: string;
	/**
	 * @pattern ^\d+$
	 */
	periodId?: string;
	/**
	 * @pattern ^\d+$
	 */
	stationId?: string;
	sendMethod?: string;
	/**
	 * @pattern ^\d+$
	 */
	planDate?: string;
	type?: OmsNotifyAlsType;
}

export interface OrderArrivalUpdateParam {
	arrivalType?: ArrivalType;
	goodsOrigin?: GoodsOrigin;
	logisticsName?: string;
	order?: OrderData;
	orderArrivalNo?: string;
	trackingNo?: string;
	warehouseCode?: string;
}

export interface OrderData {
	countryCode?: string;
	identifier?: string;
	nickname?: string;
	orderItems?: OrderItemData[];
	orderNo?: string;
}

export interface OrderItemData {
	currencyCode?: string;
	/**
	 * @pattern ^\d+$
	 */
	declareCategoryId?: string;
	declareCategoryName?: string;
	itemName?: string;
	orderLineNo?: string;
	/**
	 * @minimum 0
	 */
	quantity?: number;
	sku?: string;
	/**
	 * @minimum 0
	 */
	unitPrice?: number;
}

export interface OrderArrivalUpdateResp {
	result?: EzShipOMSResult;
	data?: string;
}

export interface NotifyALSSupplementOrderReq {
	address?: string;
	mobile?: string;
	note?: string;
	/**
	 * @pattern ^\d+$
	 */
	periodId?: string;
	/**
	 * @pattern ^\d+$
	 */
	planDate?: string;
	postcode?: string;
	receiverName?: string;
	nickName?: string;
	sendMethod?: string;
	sendPeriod?: string;
	/**
	 * @pattern ^\d+$
	 */
	stationId?: string;
	newOutOrderNo?: string;
	outOrderNo?: string;
	parcelParams?: NotifyALSSupplementParcel[];
}

export interface NotifyALSSupplementParcel {
	boxNo?: string;
	packageColor?: string;
	containerNo?: string;
	packageEta?: string;
	orderNo?: string;
	packageRemark?: string;
	trackingNumber?: string;
	/**
	 * @pattern ^\d+$
	 */
	packageWeight?: string;
	newParcelNo?: string;
	oldParcelNo?: string;
}

export interface NotifyALSSupplementOrderResp {
	alsOrderNo?: string;
	alsPackageNo?: string;
	/**
	 * @pattern ^\d+$
	 */
	alsPackageStatus?: string;
}



export function GetNewPackageCodeList(payload: Partial<CommonEmpty>) {
	return webapi<ParcelsList>("ezShipOMS.External/GetNewPackageCodeList", payload);
}

export function ModifyArrivalNoticeDocument(payload: Partial<ModifyArrivalNoticeReq>) {
	return webapi<ModifyArrivalNoticeResp>("ezShipOMS.External/ModifyArrivalNoticeDocument", payload);
}

export function ModifyAddedService(payload: Partial<ModifyAddedServiceReq>) {
	return webapi<ModifyAddedServiceResp>("ezShipOMS.External/ModifyAddedService", payload);
}

export function ModifyOutboundNoticeDocument(payload: Partial<ModifyOutboundNoticeReq>) {
	return webapi<ModifyOutboundNoticeResp>("ezShipOMS.External/ModifyOutboundNoticeDocument", payload);
}

export function ModifyOrder(payload: Partial<ModifyOrderReq>) {
	return webapi<ModifyOrderResp>("ezShipOMS.External/ModifyOrder", payload);
}

export function GetOrderRemarksReplyCount(payload: Partial<FrontRemarkReplyCountReq>) {
	return webapi<FrontRemarkReplyCountResp>("ezShipOMS.External/GetOrderRemarksReplyCount", payload);
}

export function AddFrontRemark(payload: Partial<AddFrontRemarkReq>) {
	return webapi<AddFrontRemarkResp>("ezShipOMS.External/AddFrontRemark", payload);
}

export function GetOrderRemarks(payload: Partial<OrderRemarksReq>) {
	return webapi<OrderRemarksResp>("ezShipOMS.External/GetOrderRemarks", payload);
}

export function SyncParcelInfo(payload: Partial<SyncParcelReq>) {
	return webapi<SyncParcelResp>("ezShipOMS.External/SyncParcelInfo", payload);
}

export function SyncSealedOrShelfNumber(payload: Partial<SyncSealedOrShelfNumberReq>) {
	return webapi<SyncSealedOrShelfNumberResp>("ezShipOMS.External/SyncSealedOrShelfNumber", payload);
}

export function CancelAddedService(payload: Partial<OMSCancelAddedServiceReq>) {
	return webapi<OMSCancelAddedServiceResp>("ezShipOMS.External/CancelAddedService", payload);
}

export function CancelArrivalNoticeDocument(payload: Partial<OMSCancelArrivalNoticeDocumentReq>) {
	return webapi<OMSCancelArrivalNoticeDocumentResp>("ezShipOMS.External/CancelArrivalNoticeDocument", payload);
}

export function UpdateArrivalNoticeDocument(payload: Partial<OrderArrivalUpdateParam>) {
	return webapi<OrderArrivalUpdateResp>("ezShipOMS.External/UpdateArrivalNoticeDocument", payload);
}

export function UpdateDeliveryInfo(payload: Partial<UpdateDeliveryInfoReq>) {
	return webapi<UpdateDeliveryInfoResp>("ezShipOMS.External/UpdateDeliveryInfo", payload);
}

export function GetDefaultDeliveryInfo(payload: Partial<DeliveryInfoReq>) {
	return webapi<GetDefaultDeliveryInfoResp>("ezShipOMS.External/GetDefaultDeliveryInfo", payload);
}

export function GetDeliveryInfo(payload: Partial<GetDeliveryInfoReq>) {
	return webapi<GetDeliveryInfoResp>("ezShipOMS.External/GetDeliveryInfo", payload);
}

export function GetDeliveryNeighbourhoodStation(payload: Partial<GetDeliveryNeighbourhoodStationReq>) {
	return webapi<GetDeliveryNeighbourhoodStationResp>("ezShipOMS.External/GetDeliveryNeighbourhoodStation", payload);
}

export function GetDeliveryPickupDate(payload: Partial<GetDeliveryPickupDateReq>) {
	return webapi<GetDeliveryPickupDateResp>("ezShipOMS.External/GetDeliveryPickupDate", payload);
}

export function GetDeliveryPickupPeriod(payload: Partial<GetDeliveryPickupPeriodReq>) {
	return webapi<GetDeliveryPickupPeriodResp>("ezShipOMS.External/GetDeliveryPickupPeriod", payload);
}

export function AlsNotifyOmsAlteration(payload: Partial<AlsNotifyOmsAlterationReq>) {
	return webapi<AlsNotifyOmsAlterationResp>("ezShipOMS.External/AlsNotifyOmsAlteration", payload);
}

export function OmsNotifyAlsUpdateOrder(payload: Partial<OmsNotifyAlsReq>) {
	return webapi<CommonEmpty>("ezShipOMS.External/OmsNotifyAlsUpdateOrder", payload);
}

export function WmsNotifyOmsLoseOrder(payload: Partial<NoticeOMSLoseOrderReq>) {
	return webapi<NoticeOMSLoseOrderResp>("ezShipOMS.External/WmsNotifyOmsLoseOrder", payload);
}

export function NotifyALSSupplementOrder(payload: Partial<NotifyALSSupplementOrderReq>) {
	return webapi<NotifyALSSupplementOrderResp>("ezShipOMS.External/NotifyALSSupplementOrder", payload);
}


export default {
	GetNewPackageCodeList,
	ModifyArrivalNoticeDocument,
	ModifyAddedService,
	ModifyOutboundNoticeDocument,
	ModifyOrder,
	GetOrderRemarksReplyCount,
	AddFrontRemark,
	GetOrderRemarks,
	SyncParcelInfo,
	SyncSealedOrShelfNumber,
	CancelAddedService,
	CancelArrivalNoticeDocument,
	UpdateArrivalNoticeDocument,
	UpdateDeliveryInfo,
	GetDefaultDeliveryInfo,
	GetDeliveryInfo,
	GetDeliveryNeighbourhoodStation,
	GetDeliveryPickupDate,
	GetDeliveryPickupPeriod,
	AlsNotifyOmsAlteration,
	OmsNotifyAlsUpdateOrder,
	WmsNotifyOmsLoseOrder,
	NotifyALSSupplementOrder,
};
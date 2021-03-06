/**
 * This file is auto-generated by tgen
 * Don't change manually
 */
import webapi from "./webapi";
import * as Common from "./CommonService";
import * as Product from "./ProductService";


export enum EzProductAPIExceptionErrCode {
	MaxLimit = 3,
	ProductBlocked = 1,
	ProductOutOfStock = 2,
	Unknown = 0,
}
export enum ProductDetailEntrance {
	MaxLimit = 2,
	Normal = 1,
	Unknown = 0,
}


export interface Store {
	belongToEzbuy: boolean;
	name: string;
	productsCount: number;
	satisfaction: string;
	link: string;
}
export interface TCat {
	dcid: number;
	dcname: string;
}
export interface TJoinPrimeInfo {
	url: string;
	title: string;
	params: string[];
}
export interface TPrice {
	unitPrice: number;
	originalPrice?: number;
	discount?: number;
	symbol: string;
}
export interface TProduct {
	refId: string;
	productName: string;
	productNameTrans: TTranslation;
	productUrl: string;
	primaryImage: string;
	images: string[];
	vendorName: string;
	brandName?: string;
	description?: string;
	properties: TProperty[];
	skus: TProductSku[];
	originCode: string;
	price: TPrice;
	isPrime: boolean;
	isEzbuy: boolean;
	extra: TProductExtra;
	localPrice: TPrice;
	descAttrs?: string[];
	shopName: string;
	domesticShippingFee?: TShippingFee;
	originCountry: string;
	primeEta?: TProductEta;
	ezbuyEta?: TProductEta;
	containGSTFee: boolean;
	isBlockDetail: boolean;
	isShowSeller: boolean;
	gpid: number;
	cid: number;
	isBuyforme: boolean;
	tCats: TCat[];
	notice: string;
	pcid: number;
	defaultShippingName: string;
	defaultShippingFee: string;
	gstFee: string;
	primeShippingName: string;
	primeShipmentType: number;
	ezbuyShipmentType: number;
	specialHandlingFeeMessage: string;
	specialHandlingFeePercent: number;
	unavailableReason?: string;
	sellerProductCount: number;
	exchangeRate: number;
	selectedSkuId: string;
	joinPrime?: TJoinPrimeInfo;
	titleIcons: Common.TTitleIcon[];
	sellerBanner: string;
	ruleTable: Common.TTable;
	banner: Product.TProductBanner;
	useSellerShop: boolean;
	disableRemark: boolean;
	limiterStage?: string;
	minUnitPrice: number;
	maxUnitPrice: number;
	skuShipments: XSkuShipments[];
	store: Store;
	videos?: Product.TVideo[];
	layout: number;
	promotionPrice: TPrice;
	originProductUrl: string;
}
export interface TProductEta {
	minimumDays: number;
	maximumDays: number;
}
export interface TProductExtra {
	flashSale?: TProductExtraFlashSale;
	cashOff?: TProductExtraCashOff;
	freeShipping?: TProductExtraFreeShipping;
	premium?: TProductExtraPremium;
	mncashoff: XDetailMNCashoff;
	fastDelivery?: TProductExtraFastDelivery;
	tag?: TProductExtraTag[];
}
export interface TProductExtraCashOff {
	available: boolean;
	cashOffZoneKey?: string;
	cashOffZoneName?: string;
	cashOffTagColor?: string;
	link: string;
}
export interface TProductExtraFastDelivery {
	available: boolean;
	bannerUrl: string;
}
export interface TProductExtraFlashSale {
	available: boolean;
	price: TPrice;
	beginAt: number;
	endAt: number;
	quantity: number;
	orderLimit: number;
	localPrice: TPrice;
	couponAvailable: boolean;
	beginTS: number;
	endTS: number;
	settingId: string;
}
export interface TProductExtraFreeShipping {
	available: boolean;
	url: string;
	name: string;
}
export interface TProductExtraPremium {
	available: boolean;
	bannerUrl: string;
}
export interface TProductExtraTag {
	name: string;
	desc: string;
	img: string;
	link: string;
}
export interface TProductSku {
	skuId: string;
	price: TPrice;
	quantity: number;
	propValues: string[];
	estWeight?: number;
	estVolumeWeight?: number;
	localPrice: TPrice;
	propIds: string[];
	skuUrl: string;
	skuTitle: {[key: string]: string};
	imgs: string[];
	gpid: string;
	promotionPrice: TPrice;
	shipment: TProductSkuShipment;
}
export interface TProductSkuShipment {
	ezbuyShipmentIds: number[];
	primeShipmentIds: number[];
}
export interface TProperty {
	prop: string;
	propTrans: TTranslation;
	propItems: TProperyItem[];
	propId: string;
}
export interface TProperyItem {
	value: string;
	valueTrans: TTranslation;
	propValue: string;
	propValueTrans: TTranslation;
	imageUrl?: string;
	valueId: string;
}
export interface TShippingFee {
	fee: TPrice;
	localFee: TPrice;
}
export interface TTranslation {
	EN?: string;
	CN?: string;
	MY?: string;
	TH?: string;
	ZHTW?: string;
}
export interface XDetailMNCashoff {
	tagImg: string;
	url: string;
	name: string;
	discountInfo?: XDetailMNCashoffDiscountInfo;
	tag?: XDetailMNCashoffTag;
}
export interface XDetailMNCashoffDiscountInfo {
	desc: string;
	priceDesc: string;
	price: string;
	imgUrl: string;
	unitMinPrice: number;
	unitMaxPrice: number;
}
export interface XDetailMNCashoffTag {
	text: string;
	color: string;
	img: string;
}
export interface XProductPrice {
	price: string;
}
export interface XSkuShipments {
	id: number;
	name: string;
	eta: string;
	desc: string;
	tips: string;
	onlyForPrime: boolean;
	fee: {[key: string]: XProductPrice};
	icon: string;
}


export function GetProduct(catalogCode: string, identifier: string, entrance: ProductDetailEntrance, src: string, userInfo: Product.TProductUserInfo, loadLocal: boolean): Promise<TProduct> {
	return webapi<TProduct>("EzProduct.GetProduct", { catalogCode, identifier, entrance, src, userInfo, loadLocal });
}



export default {
	GetProduct,
};


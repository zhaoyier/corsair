/**
 * This file is auto-generated by protobufgen
 * Don't change manually
 */

import { XExtraProductInfo as AppcommonXExtraProductInfo, XListViewOption as AppcommonXListViewOption } from "../appcommon/list_public";

export enum BadgeStyle {
	BadgeStyleNormal = "BadgeStyleNormal",
	BadgeStyleDiscount = "BadgeStyleDiscount",
	BadgeStyleFreeShipping = "BadgeStyleFreeShipping",
	BadgeStyleSku = "BadgeStyleSku",
	BadgeStyleImg = "BadgeStyleImg",
}



export interface ResultResp {
	result?: boolean;
	msg?: string;
}

export interface SimpleProduct {
	/**
	 * @pattern ^\d+$
	 */
	gpid?: string;
	url?: string;
	name?: string;
	pictureUrl?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	localUnitPrice?: number;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	originalLocalUnitPrice?: number;
	originCode?: string;
	discountRate?: string;
	isCashOff?: boolean;
	cashOffColor?: string;
	vendorName?: string;
	isBuyforme?: boolean;
	originName?: string;
	tag?: ProductTag;
	titleIcons?: TitleIcon[];
	discountInfo?: XDiscountInfo;
	/**
	 * @minimum 0
	 */
	commentGradeAvg?: number;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	commentCount?: number;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	orderQuantity?: number;
	price?: string;
	originPrice?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	pcid?: number;
	icon?: ExtraIcon;
	longBadges?: Badge[];
	leftView?: ProductLeftView;
	rightView?: ProductRightView;
	extraProductInfo?: AppcommonXExtraProductInfo;
	viewOption?: AppcommonXListViewOption;
	vendorId?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	manufacturerId?: number;
	manufacturer?: string;
}

export interface ExtraIcon {
	text?: string;
	link?: string;
}

export interface XDiscountInfo {
	desc?: string;
	priceDesc?: string;
	price?: string;
	imgUrl?: string;
}

export interface TitleIcon {
	icon?: string;
	text?: string;
	link?: string;
	linkText?: string;
}

export interface ProductTag {
	text?: string;
	color?: string;
	img?: string;
}

export interface TCommonOptionItemWithId {
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	id?: number;
	name?: string;
	itemDescription?: string;
}

export interface SearchFilterField {
	name?: string;
	/**
	 * @minimum 0
	 * @TJS-type integer
	 */
	productCount?: number;
}

export interface SearchFilter {
	name?: string;
	fields?: SearchFilterField[];
	emphasized?: boolean;
}

export interface SearchFilterCond {
	filterName?: string;
	fieldName?: string;
}

export interface SearchSortCond {
	sort?: string;
	isDesc?: boolean;
}

export interface SortOption {
	code?: string;
	name?: string;
	descTitle?: string;
	ascTitle?: string;
}

export interface Badge {
	style?: BadgeStyle;
	text?: string;
	img?: string;
	color?: string;
}

export interface ProductLeftView {
	/**
	 * @minimum 0
	 */
	rateScore?: number;
	text?: string;
}

export interface ProductRightView {
	text?: string;
}
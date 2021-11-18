/**
 * This file is auto-generated by protobufgen
 * Don't change manually
 */



export enum RespParty {
	RespPartyNone = "RespPartyNone",
	RespPartySeller = "RespPartySeller",
	RespPartyEzbuy = "RespPartyEzbuy",
}


export enum AftersaleOrderType {
	AftersaleOrderTypeNone = "AftersaleOrderTypeNone",
	AftersaleOrderTypeReissue = "AftersaleOrderTypeReissue",
	AftersaleOrderTypeRefund = "AftersaleOrderTypeRefund",
	AftersaleOredrTypeAll = "AftersaleOredrTypeAll",
	AftersaleOredrTypeExchang = "AftersaleOredrTypeExchang",
}



export interface OrderItemPair {
	/**
	 * @pattern ^\d+$
	 */
	orderId?: string;
	/**
	 * @pattern ^\d+$
	 */
	orderItemId?: string;
}

export interface OrderItemPairWithQty {
	order?: OrderItemPair;
	/**
	 * @pattern ^\d+$
	 */
	qty?: string;
}

export interface Result {
	/**
	 * @pattern ^\d+$
	 */
	code?: string;
	message?: string;
}
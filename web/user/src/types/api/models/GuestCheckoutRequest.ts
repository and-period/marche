/* tslint:disable */
/* eslint-disable */
/**
 * Marche Online
 * マルシェ購入者用API
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { GuestCheckoutAddress } from './GuestCheckoutAddress';
import {
    GuestCheckoutAddressFromJSON,
    GuestCheckoutAddressFromJSONTyped,
    GuestCheckoutAddressToJSON,
} from './GuestCheckoutAddress';
import type { GuestCheckoutCreditCard } from './GuestCheckoutCreditCard';
import {
    GuestCheckoutCreditCardFromJSON,
    GuestCheckoutCreditCardFromJSONTyped,
    GuestCheckoutCreditCardToJSON,
} from './GuestCheckoutCreditCard';
import type { PaymentMethodType } from './PaymentMethodType';
import {
    PaymentMethodTypeFromJSON,
    PaymentMethodTypeFromJSONTyped,
    PaymentMethodTypeToJSON,
} from './PaymentMethodType';

/**
 * 
 * @export
 * @interface GuestCheckoutRequest
 */
export interface GuestCheckoutRequest {
    /**
     * 支払いキー(重複判定用)
     * @type {string}
     * @memberof GuestCheckoutRequest
     */
    requestId: string;
    /**
     * コーディネータID
     * @type {string}
     * @memberof GuestCheckoutRequest
     */
    coordinatorId: string;
    /**
     * 箱の通番（箱単位で購入する場合のみ）
     * @type {number}
     * @memberof GuestCheckoutRequest
     */
    boxNumber: number;
    /**
     * プロモーションコード（割引適用時のみ）
     * @type {string}
     * @memberof GuestCheckoutRequest
     */
    promotionCode: string;
    /**
     * 
     * @type {PaymentMethodType}
     * @memberof GuestCheckoutRequest
     */
    paymentMethod: PaymentMethodType;
    /**
     * 決済ページからの遷移先URL
     * @type {string}
     * @memberof GuestCheckoutRequest
     */
    callbackUrl: string;
    /**
     * 支払い合計金額（誤り検出用）
     * @type {number}
     * @memberof GuestCheckoutRequest
     */
    total: number;
    /**
     * 購入者メールアドレス
     * @type {string}
     * @memberof GuestCheckoutRequest
     */
    email: string;
    /**
     * 配送先住所を請求先住所と同一にする
     * @type {boolean}
     * @memberof GuestCheckoutRequest
     */
    isSameAddress: boolean;
    /**
     * 
     * @type {GuestCheckoutAddress}
     * @memberof GuestCheckoutRequest
     */
    billingAddress?: GuestCheckoutAddress;
    /**
     * 
     * @type {GuestCheckoutAddress}
     * @memberof GuestCheckoutRequest
     */
    shippingAddress: GuestCheckoutAddress;
    /**
     * 
     * @type {GuestCheckoutCreditCard}
     * @memberof GuestCheckoutRequest
     */
    creditCard: GuestCheckoutCreditCard;
}

/**
 * Check if a given object implements the GuestCheckoutRequest interface.
 */
export function instanceOfGuestCheckoutRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "requestId" in value;
    isInstance = isInstance && "coordinatorId" in value;
    isInstance = isInstance && "boxNumber" in value;
    isInstance = isInstance && "promotionCode" in value;
    isInstance = isInstance && "paymentMethod" in value;
    isInstance = isInstance && "callbackUrl" in value;
    isInstance = isInstance && "total" in value;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "isSameAddress" in value;
    isInstance = isInstance && "shippingAddress" in value;
    isInstance = isInstance && "creditCard" in value;

    return isInstance;
}

export function GuestCheckoutRequestFromJSON(json: any): GuestCheckoutRequest {
    return GuestCheckoutRequestFromJSONTyped(json, false);
}

export function GuestCheckoutRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): GuestCheckoutRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'requestId': json['requestId'],
        'coordinatorId': json['coordinatorId'],
        'boxNumber': json['boxNumber'],
        'promotionCode': json['promotionCode'],
        'paymentMethod': PaymentMethodTypeFromJSON(json['paymentMethod']),
        'callbackUrl': json['callbackUrl'],
        'total': json['total'],
        'email': json['email'],
        'isSameAddress': json['isSameAddress'],
        'billingAddress': !exists(json, 'billingAddress') ? undefined : GuestCheckoutAddressFromJSON(json['billingAddress']),
        'shippingAddress': GuestCheckoutAddressFromJSON(json['shippingAddress']),
        'creditCard': GuestCheckoutCreditCardFromJSON(json['creditCard']),
    };
}

export function GuestCheckoutRequestToJSON(value?: GuestCheckoutRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'requestId': value.requestId,
        'coordinatorId': value.coordinatorId,
        'boxNumber': value.boxNumber,
        'promotionCode': value.promotionCode,
        'paymentMethod': PaymentMethodTypeToJSON(value.paymentMethod),
        'callbackUrl': value.callbackUrl,
        'total': value.total,
        'email': value.email,
        'isSameAddress': value.isSameAddress,
        'billingAddress': GuestCheckoutAddressToJSON(value.billingAddress),
        'shippingAddress': GuestCheckoutAddressToJSON(value.shippingAddress),
        'creditCard': GuestCheckoutCreditCardToJSON(value.creditCard),
    };
}


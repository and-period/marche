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
import type { DiscountType } from './DiscountType';
import {
    DiscountTypeFromJSON,
    DiscountTypeFromJSONTyped,
    DiscountTypeToJSON,
} from './DiscountType';
import type { PromotionStatus } from './PromotionStatus';
import {
    PromotionStatusFromJSON,
    PromotionStatusFromJSONTyped,
    PromotionStatusToJSON,
} from './PromotionStatus';

/**
 * プロモーション情報
 * @export
 * @interface Promotion
 */
export interface Promotion {
    /**
     * プロモーションID
     * @type {string}
     * @memberof Promotion
     */
    id: string;
    /**
     * タイトル
     * @type {string}
     * @memberof Promotion
     */
    title: string;
    /**
     * 詳細説明
     * @type {string}
     * @memberof Promotion
     */
    description: string;
    /**
     * 
     * @type {PromotionStatus}
     * @memberof Promotion
     */
    status: PromotionStatus;
    /**
     * 
     * @type {DiscountType}
     * @memberof Promotion
     */
    discountType: DiscountType;
    /**
     * 割引額(単位:円/%)
     * @type {number}
     * @memberof Promotion
     */
    discountRate: number;
    /**
     * クーポンコード
     * @type {string}
     * @memberof Promotion
     */
    code: string;
    /**
     * クーポン利用可能開始日時(unixtime)
     * @type {number}
     * @memberof Promotion
     */
    startAt: number;
    /**
     * クーポン利用可能終了日時(unixtime)
     * @type {number}
     * @memberof Promotion
     */
    endAt: number;
}

/**
 * Check if a given object implements the Promotion interface.
 */
export function instanceOfPromotion(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "title" in value;
    isInstance = isInstance && "description" in value;
    isInstance = isInstance && "status" in value;
    isInstance = isInstance && "discountType" in value;
    isInstance = isInstance && "discountRate" in value;
    isInstance = isInstance && "code" in value;
    isInstance = isInstance && "startAt" in value;
    isInstance = isInstance && "endAt" in value;

    return isInstance;
}

export function PromotionFromJSON(json: any): Promotion {
    return PromotionFromJSONTyped(json, false);
}

export function PromotionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Promotion {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'title': json['title'],
        'description': json['description'],
        'status': PromotionStatusFromJSON(json['status']),
        'discountType': DiscountTypeFromJSON(json['discountType']),
        'discountRate': json['discountRate'],
        'code': json['code'],
        'startAt': json['startAt'],
        'endAt': json['endAt'],
    };
}

export function PromotionToJSON(value?: Promotion | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'title': value.title,
        'description': value.description,
        'status': PromotionStatusToJSON(value.status),
        'discountType': DiscountTypeToJSON(value.discountType),
        'discountRate': value.discountRate,
        'code': value.code,
        'startAt': value.startAt,
        'endAt': value.endAt,
    };
}


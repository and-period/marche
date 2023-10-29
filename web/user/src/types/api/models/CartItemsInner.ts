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
/**
 * 
 * @export
 * @interface CartItemsInner
 */
export interface CartItemsInner {
    /**
     * 商品ID
     * @type {string}
     * @memberof CartItemsInner
     */
    productId?: string;
    /**
     * 数量
     * @type {number}
     * @memberof CartItemsInner
     */
    quantity: number;
}

/**
 * Check if a given object implements the CartItemsInner interface.
 */
export function instanceOfCartItemsInner(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "quantity" in value;

    return isInstance;
}

export function CartItemsInnerFromJSON(json: any): CartItemsInner {
    return CartItemsInnerFromJSONTyped(json, false);
}

export function CartItemsInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): CartItemsInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'productId': !exists(json, 'productId') ? undefined : json['productId'],
        'quantity': json['quantity'],
    };
}

export function CartItemsInnerToJSON(value?: CartItemsInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'productId': value.productId,
        'quantity': value.quantity,
    };
}


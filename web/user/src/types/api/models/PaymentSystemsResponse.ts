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

import { mapValues } from '../runtime';
import type { PaymentSystem } from './PaymentSystem';
import {
    PaymentSystemFromJSON,
    PaymentSystemFromJSONTyped,
    PaymentSystemToJSON,
} from './PaymentSystem';

/**
 * 
 * @export
 * @interface PaymentSystemsResponse
 */
export interface PaymentSystemsResponse {
    /**
     * 決済システム状態一覧
     * @type {Array<PaymentSystem>}
     * @memberof PaymentSystemsResponse
     */
    systems: Array<PaymentSystem>;
}

/**
 * Check if a given object implements the PaymentSystemsResponse interface.
 */
export function instanceOfPaymentSystemsResponse(value: object): value is PaymentSystemsResponse {
    if (!('systems' in value) || value['systems'] === undefined) return false;
    return true;
}

export function PaymentSystemsResponseFromJSON(json: any): PaymentSystemsResponse {
    return PaymentSystemsResponseFromJSONTyped(json, false);
}

export function PaymentSystemsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PaymentSystemsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'systems': ((json['systems'] as Array<any>).map(PaymentSystemFromJSON)),
    };
}

export function PaymentSystemsResponseToJSON(value?: PaymentSystemsResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'systems': ((value['systems'] as Array<any>).map(PaymentSystemToJSON)),
    };
}


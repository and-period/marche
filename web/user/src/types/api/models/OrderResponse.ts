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
import type { Order } from './Order';
import {
    OrderFromJSON,
    OrderFromJSONTyped,
    OrderToJSON,
} from './Order';
import type { Product } from './Product';
import {
    ProductFromJSON,
    ProductFromJSONTyped,
    ProductToJSON,
} from './Product';
import type { Coordinator } from './Coordinator';
import {
    CoordinatorFromJSON,
    CoordinatorFromJSONTyped,
    CoordinatorToJSON,
} from './Coordinator';
import type { Promotion } from './Promotion';
import {
    PromotionFromJSON,
    PromotionFromJSONTyped,
    PromotionToJSON,
} from './Promotion';

/**
 * 
 * @export
 * @interface OrderResponse
 */
export interface OrderResponse {
    /**
     * 
     * @type {Order}
     * @memberof OrderResponse
     */
    order: Order;
    /**
     * 
     * @type {Coordinator}
     * @memberof OrderResponse
     */
    coordinator: Coordinator;
    /**
     * 
     * @type {Promotion}
     * @memberof OrderResponse
     */
    promotion: Promotion;
    /**
     * 
     * @type {Array<Product>}
     * @memberof OrderResponse
     */
    products: Array<Product>;
}

/**
 * Check if a given object implements the OrderResponse interface.
 */
export function instanceOfOrderResponse(value: object): value is OrderResponse {
    if (!('order' in value) || value['order'] === undefined) return false;
    if (!('coordinator' in value) || value['coordinator'] === undefined) return false;
    if (!('promotion' in value) || value['promotion'] === undefined) return false;
    if (!('products' in value) || value['products'] === undefined) return false;
    return true;
}

export function OrderResponseFromJSON(json: any): OrderResponse {
    return OrderResponseFromJSONTyped(json, false);
}

export function OrderResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): OrderResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'order': OrderFromJSON(json['order']),
        'coordinator': CoordinatorFromJSON(json['coordinator']),
        'promotion': PromotionFromJSON(json['promotion']),
        'products': ((json['products'] as Array<any>).map(ProductFromJSON)),
    };
}

export function OrderResponseToJSON(value?: OrderResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'order': OrderToJSON(value['order']),
        'coordinator': CoordinatorToJSON(value['coordinator']),
        'promotion': PromotionToJSON(value['promotion']),
        'products': ((value['products'] as Array<any>).map(ProductToJSON)),
    };
}


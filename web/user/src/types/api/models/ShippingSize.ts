/* tslint:disable */
 
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


/**
 * 配送時の箱の大きさ
 * @export
 */
export const ShippingSize = {
    UNKNOWN: 0,
    SIZE60: 1,
    SIZE80: 2,
    SIZE100: 3
} as const;
export type ShippingSize = typeof ShippingSize[keyof typeof ShippingSize];


export function ShippingSizeFromJSON(json: any): ShippingSize {
    return ShippingSizeFromJSONTyped(json, false);
}

export function ShippingSizeFromJSONTyped(json: any, ignoreDiscriminator: boolean): ShippingSize {
    return json as ShippingSize;
}

export function ShippingSizeToJSON(value?: ShippingSize | null): any {
    return value as any;
}


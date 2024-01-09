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


/**
 * 商品販売ステータス
 * @export
 */
export const ProductStatus = {
    UNKNOWN: 0,
    PRESALE: 1,
    FOR_SALE: 2,
    OUT_OF_SALES: 3
} as const;
export type ProductStatus = typeof ProductStatus[keyof typeof ProductStatus];


export function ProductStatusFromJSON(json: any): ProductStatus {
    return ProductStatusFromJSONTyped(json, false);
}

export function ProductStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProductStatus {
    return json as ProductStatus;
}

export function ProductStatusToJSON(value?: ProductStatus | null): any {
    return value as any;
}


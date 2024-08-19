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
/**
 * 
 * @export
 * @interface UploadUrlResponse
 */
export interface UploadUrlResponse {
    /**
     * アップロード後の状態参照用キー
     * @type {string}
     * @memberof UploadUrlResponse
     */
    key: string;
    /**
     * 署名付きアップロードURL
     * @type {string}
     * @memberof UploadUrlResponse
     */
    url: string;
}

/**
 * Check if a given object implements the UploadUrlResponse interface.
 */
export function instanceOfUploadUrlResponse(value: object): value is UploadUrlResponse {
    if (!('key' in value) || value['key'] === undefined) return false;
    if (!('url' in value) || value['url'] === undefined) return false;
    return true;
}

export function UploadUrlResponseFromJSON(json: any): UploadUrlResponse {
    return UploadUrlResponseFromJSONTyped(json, false);
}

export function UploadUrlResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): UploadUrlResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'key': json['key'],
        'url': json['url'],
    };
}

export function UploadUrlResponseToJSON(value?: UploadUrlResponse | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'key': value['key'],
        'url': value['url'],
    };
}


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
 * @interface UpdateAuthUserAccountIdRequest
 */
export interface UpdateAuthUserAccountIdRequest {
    /**
     * ユーザーID(検索用)(32文字まで)
     * @type {string}
     * @memberof UpdateAuthUserAccountIdRequest
     */
    accountId: string;
}

/**
 * Check if a given object implements the UpdateAuthUserAccountIdRequest interface.
 */
export function instanceOfUpdateAuthUserAccountIdRequest(value: object): value is UpdateAuthUserAccountIdRequest {
    if (!('accountId' in value) || value['accountId'] === undefined) return false;
    return true;
}

export function UpdateAuthUserAccountIdRequestFromJSON(json: any): UpdateAuthUserAccountIdRequest {
    return UpdateAuthUserAccountIdRequestFromJSONTyped(json, false);
}

export function UpdateAuthUserAccountIdRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateAuthUserAccountIdRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'accountId': json['accountId'],
    };
}

export function UpdateAuthUserAccountIdRequestToJSON(value?: UpdateAuthUserAccountIdRequest | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'accountId': value['accountId'],
    };
}


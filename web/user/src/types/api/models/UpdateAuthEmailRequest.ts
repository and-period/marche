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
 * @interface UpdateAuthEmailRequest
 */
export interface UpdateAuthEmailRequest {
    /**
     * メールアドレス
     * @type {string}
     * @memberof UpdateAuthEmailRequest
     */
    email: string;
}

/**
 * Check if a given object implements the UpdateAuthEmailRequest interface.
 */
export function instanceOfUpdateAuthEmailRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "email" in value;

    return isInstance;
}

export function UpdateAuthEmailRequestFromJSON(json: any): UpdateAuthEmailRequest {
    return UpdateAuthEmailRequestFromJSONTyped(json, false);
}

export function UpdateAuthEmailRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateAuthEmailRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': json['email'],
    };
}

export function UpdateAuthEmailRequestToJSON(value?: UpdateAuthEmailRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'email': value.email,
    };
}


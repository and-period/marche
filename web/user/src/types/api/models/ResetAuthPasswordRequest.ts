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
 * @interface ResetAuthPasswordRequest
 */
export interface ResetAuthPasswordRequest {
    /**
     * メールアドレス
     * @type {string}
     * @memberof ResetAuthPasswordRequest
     */
    email: string;
    /**
     * 検証コード
     * @type {string}
     * @memberof ResetAuthPasswordRequest
     */
    verifyCode: string;
    /**
     * パスワード(8~32文字, 英小文字,数字を少なくとも1文字ずつは含む)
     * @type {string}
     * @memberof ResetAuthPasswordRequest
     */
    password: string;
    /**
     * パスワード(確認用)
     * @type {string}
     * @memberof ResetAuthPasswordRequest
     */
    passwordConfirmation: string;
}

/**
 * Check if a given object implements the ResetAuthPasswordRequest interface.
 */
export function instanceOfResetAuthPasswordRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "email" in value;
    isInstance = isInstance && "verifyCode" in value;
    isInstance = isInstance && "password" in value;
    isInstance = isInstance && "passwordConfirmation" in value;

    return isInstance;
}

export function ResetAuthPasswordRequestFromJSON(json: any): ResetAuthPasswordRequest {
    return ResetAuthPasswordRequestFromJSONTyped(json, false);
}

export function ResetAuthPasswordRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ResetAuthPasswordRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'email': json['email'],
        'verifyCode': json['verifyCode'],
        'password': json['password'],
        'passwordConfirmation': json['passwordConfirmation'],
    };
}

export function ResetAuthPasswordRequestToJSON(value?: ResetAuthPasswordRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'email': value.email,
        'verifyCode': value.verifyCode,
        'password': value.password,
        'passwordConfirmation': value.passwordConfirmation,
    };
}


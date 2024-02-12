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
import type { Prefecture } from './Prefecture';
import {
    PrefectureFromJSON,
    PrefectureFromJSONTyped,
    PrefectureToJSON,
} from './Prefecture';

/**
 * 
 * @export
 * @interface UpdateAddressRequest
 */
export interface UpdateAddressRequest {
    /**
     * 姓(16文字まで)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    lastname: string;
    /**
     * 名(16文字まで)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    firstname: string;
    /**
     * 姓(かな)(ひらがな,32文字まで)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    lastnameKana: string;
    /**
     * 名(かな)(ひらがな,32文字まで)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    firstnameKana: string;
    /**
     * 郵便番号(ハイフンなし)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    postalCode: string;
    /**
     * 
     * @type {Prefecture}
     * @memberof UpdateAddressRequest
     */
    prefectureCode: Prefecture;
    /**
     * 市区町村(32文字まで)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    city: string;
    /**
     * 町名・番地(64文字まで)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    addressLine1: string;
    /**
     * ビル名・号室など(64文字まで)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    addressLine2: string;
    /**
     * 電話番号 (ハイフンあり)
     * @type {string}
     * @memberof UpdateAddressRequest
     */
    phoneNumber: string;
    /**
     * デフォルト設定
     * @type {boolean}
     * @memberof UpdateAddressRequest
     */
    isDefault: boolean;
}

/**
 * Check if a given object implements the UpdateAddressRequest interface.
 */
export function instanceOfUpdateAddressRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "lastname" in value;
    isInstance = isInstance && "firstname" in value;
    isInstance = isInstance && "lastnameKana" in value;
    isInstance = isInstance && "firstnameKana" in value;
    isInstance = isInstance && "postalCode" in value;
    isInstance = isInstance && "prefectureCode" in value;
    isInstance = isInstance && "city" in value;
    isInstance = isInstance && "addressLine1" in value;
    isInstance = isInstance && "addressLine2" in value;
    isInstance = isInstance && "phoneNumber" in value;
    isInstance = isInstance && "isDefault" in value;

    return isInstance;
}

export function UpdateAddressRequestFromJSON(json: any): UpdateAddressRequest {
    return UpdateAddressRequestFromJSONTyped(json, false);
}

export function UpdateAddressRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UpdateAddressRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'lastname': json['lastname'],
        'firstname': json['firstname'],
        'lastnameKana': json['lastnameKana'],
        'firstnameKana': json['firstnameKana'],
        'postalCode': json['postalCode'],
        'prefectureCode': PrefectureFromJSON(json['prefectureCode']),
        'city': json['city'],
        'addressLine1': json['addressLine1'],
        'addressLine2': json['addressLine2'],
        'phoneNumber': json['phoneNumber'],
        'isDefault': json['isDefault'],
    };
}

export function UpdateAddressRequestToJSON(value?: UpdateAddressRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'lastname': value.lastname,
        'firstname': value.firstname,
        'lastnameKana': value.lastnameKana,
        'firstnameKana': value.firstnameKana,
        'postalCode': value.postalCode,
        'prefectureCode': PrefectureToJSON(value.prefectureCode),
        'city': value.city,
        'addressLine1': value.addressLine1,
        'addressLine2': value.addressLine2,
        'phoneNumber': value.phoneNumber,
        'isDefault': value.isDefault,
    };
}


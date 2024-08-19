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
import type { Prefecture } from './Prefecture';
import {
    PrefectureFromJSON,
    PrefectureFromJSONTyped,
    PrefectureToJSON,
} from './Prefecture';

/**
 * アドレス情報
 * @export
 * @interface Address
 */
export interface Address {
    /**
     * アドレス帳ID
     * @type {string}
     * @memberof Address
     */
    id: string;
    /**
     * デフォルト設定フラグ
     * @type {boolean}
     * @memberof Address
     */
    isDefault: boolean;
    /**
     * 姓
     * @type {string}
     * @memberof Address
     */
    lastname: string;
    /**
     * 名
     * @type {string}
     * @memberof Address
     */
    firstname: string;
    /**
     * 姓（かな）
     * @type {string}
     * @memberof Address
     */
    lastnameKana: string;
    /**
     * 名（かな）
     * @type {string}
     * @memberof Address
     */
    firstnameKana: string;
    /**
     * 郵便番号
     * @type {string}
     * @memberof Address
     */
    postalCode: string;
    /**
     * 都道府県
     * @type {string}
     * @memberof Address
     */
    prefecture: string;
    /**
     * 
     * @type {Prefecture}
     * @memberof Address
     */
    prefectureCode: Prefecture;
    /**
     * 市区町村
     * @type {string}
     * @memberof Address
     */
    city: string;
    /**
     * 町名・番地
     * @type {string}
     * @memberof Address
     */
    addressLine1: string;
    /**
     * ビル名・号室など
     * @type {string}
     * @memberof Address
     */
    addressLine2: string;
    /**
     * 電話番号
     * @type {string}
     * @memberof Address
     */
    phoneNumber: string;
}



/**
 * Check if a given object implements the Address interface.
 */
export function instanceOfAddress(value: object): value is Address {
    if (!('id' in value) || value['id'] === undefined) return false;
    if (!('isDefault' in value) || value['isDefault'] === undefined) return false;
    if (!('lastname' in value) || value['lastname'] === undefined) return false;
    if (!('firstname' in value) || value['firstname'] === undefined) return false;
    if (!('lastnameKana' in value) || value['lastnameKana'] === undefined) return false;
    if (!('firstnameKana' in value) || value['firstnameKana'] === undefined) return false;
    if (!('postalCode' in value) || value['postalCode'] === undefined) return false;
    if (!('prefecture' in value) || value['prefecture'] === undefined) return false;
    if (!('prefectureCode' in value) || value['prefectureCode'] === undefined) return false;
    if (!('city' in value) || value['city'] === undefined) return false;
    if (!('addressLine1' in value) || value['addressLine1'] === undefined) return false;
    if (!('addressLine2' in value) || value['addressLine2'] === undefined) return false;
    if (!('phoneNumber' in value) || value['phoneNumber'] === undefined) return false;
    return true;
}

export function AddressFromJSON(json: any): Address {
    return AddressFromJSONTyped(json, false);
}

export function AddressFromJSONTyped(json: any, ignoreDiscriminator: boolean): Address {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'],
        'isDefault': json['isDefault'],
        'lastname': json['lastname'],
        'firstname': json['firstname'],
        'lastnameKana': json['lastnameKana'],
        'firstnameKana': json['firstnameKana'],
        'postalCode': json['postalCode'],
        'prefecture': json['prefecture'],
        'prefectureCode': PrefectureFromJSON(json['prefectureCode']),
        'city': json['city'],
        'addressLine1': json['addressLine1'],
        'addressLine2': json['addressLine2'],
        'phoneNumber': json['phoneNumber'],
    };
}

export function AddressToJSON(value?: Address | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'id': value['id'],
        'isDefault': value['isDefault'],
        'lastname': value['lastname'],
        'firstname': value['firstname'],
        'lastnameKana': value['lastnameKana'],
        'firstnameKana': value['firstnameKana'],
        'postalCode': value['postalCode'],
        'prefecture': value['prefecture'],
        'prefectureCode': PrefectureToJSON(value['prefectureCode']),
        'city': value['city'],
        'addressLine1': value['addressLine1'],
        'addressLine2': value['addressLine2'],
        'phoneNumber': value['phoneNumber'],
    };
}


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


import * as runtime from '../runtime';
import type {
  AddressResponse,
  AddressesResponse,
  CreateAddressRequest,
  ErrorResponse,
  PostalCodeResponse,
  UpdateAddressRequest,
} from '../models/index';
import {
    AddressResponseFromJSON,
    AddressResponseToJSON,
    AddressesResponseFromJSON,
    AddressesResponseToJSON,
    CreateAddressRequestFromJSON,
    CreateAddressRequestToJSON,
    ErrorResponseFromJSON,
    ErrorResponseToJSON,
    PostalCodeResponseFromJSON,
    PostalCodeResponseToJSON,
    UpdateAddressRequestFromJSON,
    UpdateAddressRequestToJSON,
} from '../models/index';

export interface V1CreateAddressRequest {
    body: CreateAddressRequest;
}

export interface V1DeleteAddressRequest {
    addressId: string;
}

export interface V1GetAddressRequest {
    addressId: string;
}

export interface V1ListAddressesRequest {
    limit?: number;
    offset?: number;
}

export interface V1SearchPostalCodeRequest {
    postalCode: string;
}

export interface V1UpdateAddressRequest {
    addressId: string;
    body: UpdateAddressRequest;
}

/**
 * 
 */
export class AddressApi extends runtime.BaseAPI {

    /**
     * アドレス登録
     */
    async v1CreateAddressRaw(requestParameters: V1CreateAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AddressResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling v1CreateAddress().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/v1/addresses`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body'] as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AddressResponseFromJSON(jsonValue));
    }

    /**
     * アドレス登録
     */
    async v1CreateAddress(requestParameters: V1CreateAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AddressResponse> {
        const response = await this.v1CreateAddressRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * アドレス削除
     */
    async v1DeleteAddressRaw(requestParameters: V1DeleteAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['addressId'] == null) {
            throw new runtime.RequiredError(
                'addressId',
                'Required parameter "addressId" was null or undefined when calling v1DeleteAddress().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/v1/addresses/{addressId}`.replace(`{${"addressId"}}`, encodeURIComponent(String(requestParameters['addressId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * アドレス削除
     */
    async v1DeleteAddress(requestParameters: V1DeleteAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1DeleteAddressRaw(requestParameters, initOverrides);
    }

    /**
     * アドレス取得
     */
    async v1GetAddressRaw(requestParameters: V1GetAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AddressResponse>> {
        if (requestParameters['addressId'] == null) {
            throw new runtime.RequiredError(
                'addressId',
                'Required parameter "addressId" was null or undefined when calling v1GetAddress().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/v1/addresses/{addressId}`.replace(`{${"addressId"}}`, encodeURIComponent(String(requestParameters['addressId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AddressResponseFromJSON(jsonValue));
    }

    /**
     * アドレス取得
     */
    async v1GetAddress(requestParameters: V1GetAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AddressResponse> {
        const response = await this.v1GetAddressRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * アドレス一覧取得
     */
    async v1ListAddressesRaw(requestParameters: V1ListAddressesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AddressesResponse>> {
        const queryParameters: any = {};

        if (requestParameters['limit'] != null) {
            queryParameters['limit'] = requestParameters['limit'];
        }

        if (requestParameters['offset'] != null) {
            queryParameters['offset'] = requestParameters['offset'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/v1/addresses`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AddressesResponseFromJSON(jsonValue));
    }

    /**
     * アドレス一覧取得
     */
    async v1ListAddresses(requestParameters: V1ListAddressesRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AddressesResponse> {
        const response = await this.v1ListAddressesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * 郵便番号情報検索
     */
    async v1SearchPostalCodeRaw(requestParameters: V1SearchPostalCodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<PostalCodeResponse>> {
        if (requestParameters['postalCode'] == null) {
            throw new runtime.RequiredError(
                'postalCode',
                'Required parameter "postalCode" was null or undefined when calling v1SearchPostalCode().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/postal-codes/{postalCode}`.replace(`{${"postalCode"}}`, encodeURIComponent(String(requestParameters['postalCode']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => PostalCodeResponseFromJSON(jsonValue));
    }

    /**
     * 郵便番号情報検索
     */
    async v1SearchPostalCode(requestParameters: V1SearchPostalCodeRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<PostalCodeResponse> {
        const response = await this.v1SearchPostalCodeRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * アドレス更新
     */
    async v1UpdateAddressRaw(requestParameters: V1UpdateAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['addressId'] == null) {
            throw new runtime.RequiredError(
                'addressId',
                'Required parameter "addressId" was null or undefined when calling v1UpdateAddress().'
            );
        }

        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling v1UpdateAddress().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/v1/addresses/{addressId}`.replace(`{${"addressId"}}`, encodeURIComponent(String(requestParameters['addressId']))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body'] as any,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * アドレス更新
     */
    async v1UpdateAddress(requestParameters: V1UpdateAddressRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1UpdateAddressRaw(requestParameters, initOverrides);
    }

}

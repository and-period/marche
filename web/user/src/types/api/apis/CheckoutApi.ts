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
  CheckoutProductRequest,
  CheckoutResponse,
  CheckoutStateResponse,
  ErrorResponse,
  GuestCheckoutProductRequest,
  GuestCheckoutResponse,
  GuestCheckoutStateResponse,
} from '../models/index';
import {
    CheckoutProductRequestFromJSON,
    CheckoutProductRequestToJSON,
    CheckoutResponseFromJSON,
    CheckoutResponseToJSON,
    CheckoutStateResponseFromJSON,
    CheckoutStateResponseToJSON,
    ErrorResponseFromJSON,
    ErrorResponseToJSON,
    GuestCheckoutProductRequestFromJSON,
    GuestCheckoutProductRequestToJSON,
    GuestCheckoutResponseFromJSON,
    GuestCheckoutResponseToJSON,
    GuestCheckoutStateResponseFromJSON,
    GuestCheckoutStateResponseToJSON,
} from '../models/index';

export interface V1CheckoutProductRequest {
    body: CheckoutProductRequest;
}

export interface V1GetCheckoutStateRequest {
    transactionId: string;
}

export interface V1GetGuestCheckoutStateRequest {
    transactionId: string;
}

export interface V1GuestCheckoutProductRequest {
    body: GuestCheckoutProductRequest;
}

/**
 * 
 */
export class CheckoutApi extends runtime.BaseAPI {

    /**
     * 商品購入
     */
    async v1CheckoutProductRaw(requestParameters: V1CheckoutProductRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CheckoutResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling v1CheckoutProduct().'
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
            path: `/v1/checkouts/products`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body'] as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CheckoutResponseFromJSON(jsonValue));
    }

    /**
     * 商品購入
     */
    async v1CheckoutProduct(requestParameters: V1CheckoutProductRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CheckoutResponse> {
        const response = await this.v1CheckoutProductRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * 注文情報の取得
     */
    async v1GetCheckoutStateRaw(requestParameters: V1GetCheckoutStateRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CheckoutStateResponse>> {
        if (requestParameters['transactionId'] == null) {
            throw new runtime.RequiredError(
                'transactionId',
                'Required parameter "transactionId" was null or undefined when calling v1GetCheckoutState().'
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
            path: `/v1/checkouts/{transactionId}`.replace(`{${"transactionId"}}`, encodeURIComponent(String(requestParameters['transactionId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CheckoutStateResponseFromJSON(jsonValue));
    }

    /**
     * 注文情報の取得
     */
    async v1GetCheckoutState(requestParameters: V1GetCheckoutStateRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CheckoutStateResponse> {
        const response = await this.v1GetCheckoutStateRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * ゲスト注文情報の取得
     */
    async v1GetGuestCheckoutStateRaw(requestParameters: V1GetGuestCheckoutStateRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GuestCheckoutStateResponse>> {
        if (requestParameters['transactionId'] == null) {
            throw new runtime.RequiredError(
                'transactionId',
                'Required parameter "transactionId" was null or undefined when calling v1GetGuestCheckoutState().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/guests/checkouts/{transactionId}`.replace(`{${"transactionId"}}`, encodeURIComponent(String(requestParameters['transactionId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GuestCheckoutStateResponseFromJSON(jsonValue));
    }

    /**
     * ゲスト注文情報の取得
     */
    async v1GetGuestCheckoutState(requestParameters: V1GetGuestCheckoutStateRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GuestCheckoutStateResponse> {
        const response = await this.v1GetGuestCheckoutStateRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * ゲスト商品購入
     */
    async v1GuestCheckoutProductRaw(requestParameters: V1GuestCheckoutProductRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GuestCheckoutResponse>> {
        if (requestParameters['body'] == null) {
            throw new runtime.RequiredError(
                'body',
                'Required parameter "body" was null or undefined when calling v1GuestCheckoutProduct().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/guests/checkouts/products`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['body'] as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GuestCheckoutResponseFromJSON(jsonValue));
    }

    /**
     * ゲスト商品購入
     */
    async v1GuestCheckoutProduct(requestParameters: V1GuestCheckoutProductRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GuestCheckoutResponse> {
        const response = await this.v1GuestCheckoutProductRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

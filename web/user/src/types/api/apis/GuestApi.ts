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
  ErrorResponse,
  GuestCheckoutRequest,
  GuestCheckoutResponse,
  GuestCheckoutStateResponse,
} from '../models/index';
import {
    ErrorResponseFromJSON,
    ErrorResponseToJSON,
    GuestCheckoutRequestFromJSON,
    GuestCheckoutRequestToJSON,
    GuestCheckoutResponseFromJSON,
    GuestCheckoutResponseToJSON,
    GuestCheckoutStateResponseFromJSON,
    GuestCheckoutStateResponseToJSON,
} from '../models/index';

export interface V1GetGuestCheckoutStateRequest {
    transactionId: string;
}

export interface V1GuestCheckoutRequest {
    body: GuestCheckoutRequest;
}

/**
 * 
 */
export class GuestApi extends runtime.BaseAPI {

    /**
     * ゲスト注文情報の取得
     */
    async v1GetGuestCheckoutStateRaw(requestParameters: V1GetGuestCheckoutStateRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GuestCheckoutStateResponse>> {
        if (requestParameters.transactionId === null || requestParameters.transactionId === undefined) {
            throw new runtime.RequiredError('transactionId','Required parameter requestParameters.transactionId was null or undefined when calling v1GetGuestCheckoutState.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/guests/checkouts/{transactionId}`.replace(`{${"transactionId"}}`, encodeURIComponent(String(requestParameters.transactionId))),
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
    async v1GuestCheckoutRaw(requestParameters: V1GuestCheckoutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GuestCheckoutResponse>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling v1GuestCheckout.');
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
            path: `/v1/guests/checkouts`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.body as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GuestCheckoutResponseFromJSON(jsonValue));
    }

    /**
     * ゲスト商品購入
     */
    async v1GuestCheckout(requestParameters: V1GuestCheckoutRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GuestCheckoutResponse> {
        const response = await this.v1GuestCheckoutRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

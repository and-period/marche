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
  AuthResponse,
  ErrorResponse,
  ForgotAuthPasswordRequest,
  GetUploadUrlRequest,
  RefreshAuthTokenRequest,
  ResetAuthPasswordRequest,
  SignInRequest,
  UpdateAuthPasswordRequest,
  UploadUrlResponse,
} from '../models/index';
import {
    AuthResponseFromJSON,
    AuthResponseToJSON,
    ErrorResponseFromJSON,
    ErrorResponseToJSON,
    ForgotAuthPasswordRequestFromJSON,
    ForgotAuthPasswordRequestToJSON,
    GetUploadUrlRequestFromJSON,
    GetUploadUrlRequestToJSON,
    RefreshAuthTokenRequestFromJSON,
    RefreshAuthTokenRequestToJSON,
    ResetAuthPasswordRequestFromJSON,
    ResetAuthPasswordRequestToJSON,
    SignInRequestFromJSON,
    SignInRequestToJSON,
    UpdateAuthPasswordRequestFromJSON,
    UpdateAuthPasswordRequestToJSON,
    UploadUrlResponseFromJSON,
    UploadUrlResponseToJSON,
} from '../models/index';

export interface V1ForgotAuthPasswordRequest {
    body: ForgotAuthPasswordRequest;
}

export interface V1GetUserThumbnailUploadUrlRequest {
    body: GetUploadUrlRequest;
}

export interface V1RefreshAuthTokenRequest {
    body: RefreshAuthTokenRequest;
}

export interface V1ResetAuthPasswordRequest {
    body: ResetAuthPasswordRequest;
}

export interface V1SignInRequest {
    body: SignInRequest;
}

export interface V1UpdateUserPasswordRequest {
    body: UpdateAuthPasswordRequest;
}

/**
 * 
 */
export class AuthApi extends runtime.BaseAPI {

    /**
     * パスワードリセット
     */
    async v1ForgotAuthPasswordRaw(requestParameters: V1ForgotAuthPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling v1ForgotAuthPassword.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/auth/forgot-password`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.body as any,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * パスワードリセット
     */
    async v1ForgotAuthPassword(requestParameters: V1ForgotAuthPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1ForgotAuthPasswordRaw(requestParameters, initOverrides);
    }

    /**
     * トークン検証
     */
    async v1GetAuthRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AuthResponse>> {
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
            path: `/v1/auth`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AuthResponseFromJSON(jsonValue));
    }

    /**
     * トークン検証
     */
    async v1GetAuth(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AuthResponse> {
        const response = await this.v1GetAuthRaw(initOverrides);
        return await response.value();
    }

    /**
     * 購入者サムネイルアップロード用URL取得
     */
    async v1GetUserThumbnailUploadUrlRaw(requestParameters: V1GetUserThumbnailUploadUrlRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UploadUrlResponse>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling v1GetUserThumbnailUploadUrl.');
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
            path: `/v1/upload/users/thumbnail`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.body as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UploadUrlResponseFromJSON(jsonValue));
    }

    /**
     * 購入者サムネイルアップロード用URL取得
     */
    async v1GetUserThumbnailUploadUrl(requestParameters: V1GetUserThumbnailUploadUrlRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UploadUrlResponse> {
        const response = await this.v1GetUserThumbnailUploadUrlRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * トークン更新
     */
    async v1RefreshAuthTokenRaw(requestParameters: V1RefreshAuthTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AuthResponse>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling v1RefreshAuthToken.');
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
            path: `/v1/auth/refresh-token`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.body as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AuthResponseFromJSON(jsonValue));
    }

    /**
     * トークン更新
     */
    async v1RefreshAuthToken(requestParameters: V1RefreshAuthTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AuthResponse> {
        const response = await this.v1RefreshAuthTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * パスワードリセット - コード検証
     */
    async v1ResetAuthPasswordRaw(requestParameters: V1ResetAuthPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling v1ResetAuthPassword.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/auth/forgot-password/verified`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.body as any,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * パスワードリセット - コード検証
     */
    async v1ResetAuthPassword(requestParameters: V1ResetAuthPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1ResetAuthPasswordRaw(requestParameters, initOverrides);
    }

    /**
     * サインイン
     */
    async v1SignInRaw(requestParameters: V1SignInRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AuthResponse>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling v1SignIn.');
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
            path: `/v1/auth`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.body as any,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AuthResponseFromJSON(jsonValue));
    }

    /**
     * サインイン
     */
    async v1SignIn(requestParameters: V1SignInRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AuthResponse> {
        const response = await this.v1SignInRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * サインアウト
     */
    async v1SignOutRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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
            path: `/v1/auth`,
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * サインアウト
     */
    async v1SignOut(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1SignOutRaw(initOverrides);
    }

    /**
     * パスワード更新
     */
    async v1UpdateUserPasswordRaw(requestParameters: V1UpdateUserPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling v1UpdateUserPassword.');
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
            path: `/v1/auth/password`,
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters.body as any,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * パスワード更新
     */
    async v1UpdateUserPassword(requestParameters: V1UpdateUserPasswordRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1UpdateUserPasswordRaw(requestParameters, initOverrides);
    }

}

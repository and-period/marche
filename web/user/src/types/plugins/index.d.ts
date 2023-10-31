import 'pinia'
import { AuthApi, ProductApi } from '../api'
import { CustomErrorMessage } from '~/plugins/api-error-handler'
import VueI18n from 'vue-i18n'

declare module 'pinia' {
  export interface PiniaCustomProperties {
    /**
     * apiクライアントのエラーをハンドリングする関数
     * @param error 発生したエラーオブジェクト
     * @param customObject エラーメッセージをカスタマイズするオブジェクト
     * @returns Promise.rejectを返す。呼び出す側で再度returnすることを想定している。
     */
    errorHandler: (
      error: unknown,
      customObject?: CustomErrorMessage,
    ) => Promise<never>
    authApiClient: (token?: string | undefined) => AuthApi
    productApiClient: (token?: string | undefined) => ProductApi
    i18n: VueI18
  }
}

declare module '#app' {
  interface NuxtApp {
    $md: MarkdownIt
  }
}

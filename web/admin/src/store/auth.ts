import axios from 'axios'
import { getToken, isSupported } from 'firebase/messaging'
import { defineStore } from 'pinia'
import Cookies from 'universal-cookie'

import dayjs, { Dayjs } from 'dayjs'
import { useCommonStore } from './common'
import { messaging } from '~/plugins/firebase'
import {
  AuthResponse,
  SignInRequest,
  UpdateAuthEmailRequest,
  UpdateAuthPasswordRequest,
  VerifyAuthEmailRequest
} from '~/types/api'
import {
  AuthError,
  ConflictError,
  ConnectionError,
  InternalServerError,
  PreconditionError,
  ValidationError
} from '~/types/exception'
import { apiClient } from '~/plugins/api-client'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    redirectPath: '/',
    isAuthenticated: false,
    user: undefined as AuthResponse | undefined,
    expiredAt: undefined as Dayjs | undefined
  }),

  getters: {
    accessToken (state): string | undefined {
      return state.user?.accessToken
    }
  },

  actions: {
    async signIn (payload: SignInRequest): Promise<string> {
      try {
        const res = await apiClient.authApi().v1SignIn(payload)
        this.setExpiredAt(res.data)
        this.isAuthenticated = true
        this.user = res.data

        const cookies = new Cookies()
        cookies.set('refreshToken', this.user.refreshToken, { secure: true })

        // Push通知の許可設定
        this.getDeviceToken()
          .then((deviceToken) => {
            if (deviceToken === '') {
              return // Push通知が未有効化状態
            }
            const currentToken: string = cookies.get('deviceToken')
            if (deviceToken === currentToken) {
              return // API側へ登録済み
            }
            return this.registerDeviceToken(deviceToken)
          })
          .catch((err) => {
            console.log('push notifications are disabled.', err)
          })

        return this.redirectPath
      } catch (err) {
        if (axios.isAxiosError(err)) {
          if (!err.response) {
            return Promise.reject(new ConnectionError(err))
          }
          switch (err.response.status) {
            case 400:
            case 401:
              return Promise.reject(
                new ValidationError(
                  'ユーザー名またはパスワードが違います。',
                  err
                )
              )
            default:
              return Promise.reject(new InternalServerError(err))
          }
        }
        throw new InternalServerError(err)
      }
    },

    async passwordUpdate (payload: UpdateAuthPasswordRequest): Promise<void> {
      try {
        await apiClient.authApi().v1UpdateAuthPassword(payload)
        const commonStore = useCommonStore()
        commonStore.addSnackbar({
          message: 'パスワードを更新しました。',
          color: 'info'
        })
      } catch (error) {
        if (axios.isAxiosError(error)) {
          if (!error.response) {
            return Promise.reject(new ConnectionError(error))
          }
          const statusCode = error.response.status
          switch (statusCode) {
            case 401:
              return Promise.reject(
                new AuthError('認証エラー。再度ログインをしてください。', error)
              )
            case 400:
              return Promise.reject(
                new ValidationError('入力値に誤りがあります。', error)
              )
            default:
              return Promise.reject(new InternalServerError(error))
          }
        }
        throw new InternalServerError(error)
      }
    },

    async emailUpdate (payload: UpdateAuthEmailRequest): Promise<void> {
      try {
        await apiClient.authApi().v1UpdateAuthEmail(payload)
        const commonStore = useCommonStore()
        commonStore.addSnackbar({
          message: '認証コードを送信しました。',
          color: 'info'
        })
      } catch (error) {
        if (axios.isAxiosError(error)) {
          if (!error.response) {
            return Promise.reject(new ConnectionError(error))
          }
          const statusCode = error.response.status
          switch (statusCode) {
            case 400:
              return Promise.reject(
                new ValidationError('入力内容に誤りがあります。', error)
              )
            case 401:
              return Promise.reject(
                new AuthError('認証エラー。再度ログインをしてください。', error)
              )
            case 409:
              return Promise.reject(
                new ConflictError(
                  'このメールアドレスはすでに登録されているため、変更できません。',
                  error
                )
              )
            case 412:
              return Promise.reject(
                new PreconditionError(
                  '変更前のメールアドレスと同じため、変更できません。',
                  error
                )
              )
            case 500:
            default:
              return Promise.reject(new InternalServerError(error))
          }
        }
        throw new InternalServerError(error)
      }
    },

    async codeVerify (payload: VerifyAuthEmailRequest): Promise<void> {
      try {
        await apiClient.authApi().v1VerifyAuthEmail(payload)
        const commonStore = useCommonStore()
        commonStore.addSnackbar({
          message: 'メールアドレスが変更されました。',
          color: 'info'
        })
      } catch (error) {
        return this.errorHandler(error, {
          409: 'このメールアドレスはすでに利用されているため使用できません。'
        })
      }
    },

    async getAuthByRefreshToken (refreshToken: string): Promise<void> {
      try {
        const res = await apiClient.authApi().v1RefreshAuthToken({
          refreshToken
        })
        this.setExpiredAt(res.data)
        this.isAuthenticated = true
        this.user = res.data
        this.user.refreshToken = refreshToken
      } catch (error) {
        const cookies = new Cookies()
        cookies.remove('refreshToken')
        if (axios.isAxiosError(error)) {
          if (!error.response) {
            return Promise.reject(new ConnectionError(error))
          }
          if (error.response.status === 401) {
            return Promise.reject(
              new AuthError('認証エラー。再度ログインをしてください。', error)
            )
          } else {
            return Promise.reject(new InternalServerError(error))
          }
        }
        throw new InternalServerError(error)
      }
    },

    async registerDeviceToken (deviceToken: string): Promise<void> {
      try {
        await apiClient.authApi().v1RegisterAuthDevice({ device: deviceToken })

        const cookies = new Cookies()
        cookies.set('deviceToken', deviceToken, { secure: true })
      } catch (error) {
        if (axios.isAxiosError(error)) {
          if (!error.response) {
            return Promise.reject(new ConnectionError(error))
          }
          switch (error.response.status) {
            case 401:
              return Promise.reject(
                new AuthError('認証エラー。再度ログインをしてください。', error)
              )
            default:
              return Promise.reject(new InternalServerError(error))
          }
        }
        throw new InternalServerError(error)
      }
    },

    async getDeviceToken (): Promise<string> {
      const runtimeConfig = useRuntimeConfig()

      const supported = await isSupported()
      if (!supported) {
        console.log('this browser does not support push notificatins.')
        return '' // Push通知未対応ブラウザ
      }

      return await getToken(messaging, {
        vapidKey: runtimeConfig.public.FIREBASE_VAPID_KEY
      })
        .then((currentToken) => {
          return currentToken
        })
        .catch((err) => {
          console.log('failed to get device token', err)
          return ''
        })
    },

    setRedirectPath (payload: string) {
      this.redirectPath = payload
    },

    setExpiredAt (user: AuthResponse) {
      this.expiredAt = dayjs().add(user.expiresIn, 'second')
    },

    logout () {
      try {
        apiClient.authApi().v1SignOut()
        const cookies = new Cookies()
        cookies.remove('refreshToken')
        this.$reset()
      } catch (error) {
        console.log('APIでエラーが発生しました。', error)
      }
    }
  }
})

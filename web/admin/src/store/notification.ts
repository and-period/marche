import axios from 'axios'
import { defineStore } from 'pinia'

import { useAuthStore } from './auth'
import { useCommonStore } from './common'

import ApiClientFactory from '~/plugins/factory'
import {
  CreateNotificationRequest,
  NotificationApi,
  NotificationResponse,
  NotificationsResponse,
  UpdateNotificationRequest
} from '~/types/api'
import {
  AuthError,
  ConnectionError,
  InternalServerError,
  NotFoundError,
  ValidationError
} from '~/types/exception'

export const useNotificationStore = defineStore('Notification', {
  state: () => {
    const apiClient = (token: string) => {
      const factory = new ApiClientFactory()
      return factory.create(NotificationApi, token)
    }
    return {
      notifications: [] as NotificationsResponse['notifications'],
      totalItems: 0,
      apiClient
    }
  },
  actions: {
    /**
     * 登録済みのお知らせ一覧を取得する非同期関数
     * @param limit 取得上限数
     * @param offset 取得開始位置
     * @returns
     */
    async fetchNotifications (limit = 20, offset = 0): Promise<void> {
      try {
        const accessToken = this.getAccessToken()
        const res = await this.apiClient(accessToken).v1ListNotifications(
          limit,
          offset
        )
        const { notifications, total }: NotificationsResponse = res.data

        this.notifications = notifications
        this.totalItems = total
      } catch (err) {
        return this.errorHandler(err)
      }
    },

    /**
     * お知らせを登録する非同期関数
     * @param payload
     */
    async createNotification (
      payload: CreateNotificationRequest
    ): Promise<void> {
      try {
        const authStore = useAuthStore()
        const accessToken = authStore.accessToken
        if (!accessToken) {
          return Promise.reject(
            new AuthError('認証エラー。再度ログインをしてください。')
          )
        }

        await this.apiClient(accessToken).v1CreateNotification(payload)
        const commonStore = useCommonStore()
        commonStore.addSnackbar({
          message: `${payload.title}を作成しました。`,
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
            case 500:
            default:
              return Promise.reject(new InternalServerError(error))
          }
        }

        throw new InternalServerError(error)
      }
    },

    /**
     * お知らせを削除する非同期関数
     * @param id お知らせID
     */
    async deleteNotification (id: string): Promise<void> {
      const commonStore = useCommonStore()
      try {
        const authStore = useAuthStore()
        const accessToken = authStore.accessToken
        if (!accessToken) {
          return Promise.reject(new Error('認証エラー'))
        }

        await this.apiClient(accessToken).v1DeleteNotification(id)
        commonStore.addSnackbar({
          message: '品物削除が完了しました',
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
                new ValidationError(
                  '削除できませんでした。管理者にお問い合わせしてください。',
                  error
                )
              )
            case 401:
              return Promise.reject(
                new AuthError('認証エラー。再度ログインをしてください。', error)
              )
            case 404:
              return Promise.reject(
                new NotFoundError(
                  '削除するお知らせが見つかりませんでした。',
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
      this.fetchNotifications()
    },

    /**
     * お知らせIDからお知らせ情報情報を取得する非同期関数
     * @param id お知らせID
     * @returns お知らせ情報
     */
    async getNotification (id: string): Promise<NotificationResponse> {
      try {
        const authStore = useAuthStore()
        const accessToken = authStore.accessToken
        if (!accessToken) {
          return Promise.reject(new Error('認証エラー'))
        }
        const res = await this.apiClient(accessToken).v1GetNotification(id)
        return res.data
      } catch (error) {
        if (axios.isAxiosError(error)) {
          if (!error.response) {
            return Promise.reject(new ConnectionError(error))
          }
          const statusCode = error.response.status
          switch (statusCode) {
            case 401:
              return Promise.reject(
                new AuthError('認証エラー。再度ログインをしてください', error)
              )
            case 404:
              return Promise.reject(
                new NotFoundError(
                  '一致するお知らせ情報が見つかりませんでした。',
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

    /**
     * お知らせ情報を編集する非同期関数
     * @param id セールID
     * @param payload
     */
    async editNotification (
      id: string,
      payload: UpdateNotificationRequest
    ): Promise<void> {
      const commonStore = useCommonStore()
      try {
        const authStore = useAuthStore()
        const accessToken = authStore.accessToken
        if (!accessToken) {
          return Promise.reject(new Error('認証エラー'))
        }
        await this.apiClient(accessToken).v1UpdateNotification(id, payload)
        commonStore.addSnackbar({
          message: 'お知らせ情報の編集が完了しました',
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
                new AuthError('認証エラー。再度ログインをしてください', error)
              )
            case 404:
              return Promise.reject(
                new NotFoundError(
                  '一致するお知らせ情報が見つかりませんでした。',
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

    getAccessToken (): string {
      const authStore = useAuthStore()
      const accessToken = authStore.accessToken
      if (!accessToken) {
        throw new AuthError('認証エラー。再度ログインをしてください。')
      }
      return accessToken
    }
  }
})

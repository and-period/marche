import { defineStore } from 'pinia'

import { useCommonStore } from './common'
import {
  CategoriesResponse,
  CategoriesResponseCategoriesInner,
  CreateCategoryRequest,
  UpdateCategoryRequest
} from '~/types/api'
import { apiClient } from '~/plugins/api-client'

export const useCategoryStore = defineStore('category', {
  state: () => ({
    categories: [] as CategoriesResponse['categories'],
    total: 0
  }),

  actions: {
    /**
     * カテゴリ一覧を取得する非同期関数
     * @param limit 取得上限数
     * @param offset 取得開始位置
     * @param orders ソートキー
     */
    async fetchCategories (limit = 20, offset = 0, orders = []): Promise<void> {
      try {
        const res = await listCategories(limit, offset, '', orders)
        this.categories = res.categories
        this.total = res.total
      } catch (err) {
        return this.errorHandler(err)
      }
    },

    /**
     * カテゴリを検索をする非同期関数
     * @param name カテゴリ名(あいまい検索)
     * @param categoryIds stateの更新時に残しておく必要があるカテゴリ情報
     */
    async searchCategories (name = '', categoryIds: string[] = []): Promise<void> {
      try {
        const res = await listCategories(undefined, undefined, name, [])
        const categories: CategoriesResponseCategoriesInner[] = []
        this.categories.forEach((category: CategoriesResponseCategoriesInner): void => {
          if (!categoryIds.includes(category.id)) {
            return
          }
          categories.push(category)
        })
        res.categories.forEach((category: CategoriesResponseCategoriesInner): void => {
          if (categories.find((v): boolean => v.id === category.id)) {
            return
          }
          categories.push(category)
        })
        this.categories = categories
        this.total = res.total
      } catch (err) {
        return this.errorHandler(err)
      }
    },

    /**
     * カテゴリを追加取得する非同期関数
     * @param limit 取得上限数
     * @param offset 取得開始位置
     * @param orders ソートキー
     */
    async moreCategories (limit = 20, offset = 0, orders = []): Promise<void> {
      try {
        const res = await listCategories(limit, offset, '', orders)
        this.categories.push(...res.categories)
        this.total = res.total
      } catch (err) {
        return this.errorHandler(err)
      }
    },

    /**
     * カテゴリを新規登録する非同期関数
     * @param payload
     */
    async createCategory (payload: CreateCategoryRequest): Promise<void> {
      const commonStore = useCommonStore()
      try {
        const res = await apiClient.categoryApi().v1CreateCategory(payload)
        this.categories.unshift(res.data)
        commonStore.addSnackbar({
          message: 'カテゴリーを追加しました。',
          color: 'info'
        })
      } catch (err) {
        return this.errorHandler(err, { 409: 'このカテゴリー名はすでに登録されています。' })
      }
    },

    /**
     * カテゴリを編集する非同期関数
     * @param categoryId カテゴリID
     * @param payload
     */
    async updateCategory (categoryId: string, payload: UpdateCategoryRequest) {
      const commonStore = useCommonStore()
      try {
        await apiClient.categoryApi().v1UpdateCategory(categoryId, payload)
        commonStore.addSnackbar({
          message: '変更しました。',
          color: 'info'
        })
      } catch (err) {
        return this.errorHandler(err, { 409: 'このカテゴリー名はすでに登録されています。' })
      }
      this.fetchCategories()
    },

    /**
     * カテゴリを削除する非同期関数
     * @param categoryId カテゴリID
     */
    async deleteCategory (categoryId: string): Promise<void> {
      const commonStore = useCommonStore()
      try {
        await apiClient.categoryApi().v1DeleteCategory(categoryId)
        commonStore.addSnackbar({
          message: 'カテゴリー削除が完了しました',
          color: 'info'
        })
      } catch (err) {
        return this.errorHandler(err)
      }
      this.fetchCategories()
    }
  }
})

async function listCategories (limit = 20, offset = 0, name = '', orders: string[] = []): Promise<CategoriesResponse> {
  const res = await apiClient.categoryApi().v1ListCategories(limit, offset, name, orders.join(','))
  return { ...res.data }
}

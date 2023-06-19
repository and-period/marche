import { CreateProductTagRequest, ProductTagResponse, ProductTagsResponse, UpdateProductTagRequest } from "~/types/api";
import { apiClient } from '~/plugins/api-client'
import { useCommonStore } from "./common";

export const useProductTagStore = defineStore('productTag', {
  state: () => ({
    productTag: {} as ProductTagResponse,
    productTags: [] as ProductTagsResponse['productTags'],
    total: 0
  }),

  actions: {
    /**
     * 商品タグ一覧を取得する非同期関数
     * @param limit 取得上限数
     * @param offset 取得開始位置
     * @param orders ソートキー
     */
    async fetchProductTags(limit = 20, offset = 0, orders: string[] = []): Promise<void> {
      try {
        const res = await apiClient.productTagApi().v1ListProductTags(limit, offset, '', orders.join(','))
        this.productTags = res.data.productTags
        this.total = res.data.total
      } catch (err) {
        return this.errorHandler(err)
      }
    },

    /**
     * 商品タグを新規登録する非同期関数
     * @param payload
     */
    async createProductTag (payload: CreateProductTagRequest): Promise<void> {
      const commonStore = useCommonStore()
      try {
        const res = await apiClient.productTagApi().v1CreateProductTag(payload)
        this.productTags.unshift(res.data)
        commonStore.addSnackbar({
          message: '商品タグを追加しました。',
          color: 'info'
        })
      } catch (err) {
        return this.errorHandler(err, { 409: 'この商品タグ名はすでに登録されています。' })
      }
    },

    /**
     * 商品タグを更新する非同期関数
     * @param productTagId 商品タグID
     * @param payload
     */
    async updateProductTag (productTagId: string, payload: UpdateProductTagRequest): Promise<void> {
      const commonStore = useCommonStore()
      try {
        await apiClient.productTagApi().v1UpdateProductTag(productTagId, payload)
        commonStore.addSnackbar({
          message: '商品タグを更新しました。',
          color: 'info'
        })
      } catch (err) {
        return this.errorHandler(err, { 409: 'この商品タグ名はすでに登録されています。' })
      }
    },

    /**
     * 商品タグを削除する非同期関数
     * @param productTagId 商品タグID
     */
    async deleteProductTag (productTagId: string): Promise<void> {
      const commonStore = useCommonStore()
      try {
        await apiClient.productTagApi().v1DeleteProductTag(productTagId)
        commonStore.addSnackbar({
          message: '商品タグを削除しました。',
          color: 'info'
        })
      } catch (err) {
        this.errorHandler(err)
      }
    },
  }
})

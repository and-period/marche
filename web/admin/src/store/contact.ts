import { defineStore } from 'pinia'

import { useAuthStore } from './auth'

import ApiClientFactory from '~/plugins/factory'
import { ContactApi, ContactResponse, ContactsResponse } from '~/types/api'

export const useContactStore = defineStore('Contact', {
  state: () => ({
    contacts: [] as ContactsResponse['contacts'],
  }),
  actions: {
    async fetchContacts(limit: number = 20, offset: number = 0): Promise<void> {
      try {
        const authStore = useAuthStore()
        const accessToken = authStore.accessToken
        if (!accessToken) {
          return Promise.reject(new Error('認証エラー'))
        }
        const factory = new ApiClientFactory()
        const contactsApiClient = factory.create(ContactApi, accessToken)
        const res = await contactsApiClient.v1ListContacts(limit, offset)
        console.log(res)
        this.contacts = res.data.contacts
      } catch (error) {
        console.log(error)
        throw new Error('Internal Server Error')
      }
    },
    async getContact(id: string): Promise<ContactResponse> {
      try{
        const authStore = useAuthStore()
        const accessToken = authStore.accessToken
        if (!accessToken) throw new Error('認証エラー')
        const factory = new ApiClientFactory()
        const contactsApiClient = factory.create(ContactApi, accessToken)
        const res = await contactsApiClient.v1GetContact(id)
        return res.data
      } catch (error) {
        console.log(error)
        throw new Error('Internal Server Error')
      }
    }
  },
})

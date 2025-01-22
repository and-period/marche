<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { MOCK_RECOMMEND_ITEMS } from '~/constants/mock'
import { useTopPageStore } from '~/store/home'
import { useInstagramStore } from '~/store/instagram'
import type { BannerItem } from '~/types/props'
import type { I18n } from '~/types/locales'

const i18n = useI18n()

const router = useRouter()

const topPageStore = useTopPageStore()
const { archives, lives } = storeToRefs(topPageStore)
const { getHomeContent } = topPageStore

const instagramStore = useInstagramStore()
const { OEmbedHTML } = storeToRefs(instagramStore)
const { getInstagramOEmbed } = instagramStore

const tt = (str: keyof I18n['base']['top']) => {
  return i18n.t(`base.top.${str}`)
}

const isInItLoading = ref<boolean>(false)

const archiveRef = ref<HTMLDivElement | null>(null)
const archiveRefScrollLeft = ref<number>(0)

const updateScrollLeft = () => {
  if (archiveRef.value) {
    archiveRefScrollLeft.value = archiveRef.value.scrollLeft
  }
}

const isEnglish = computed<boolean>(() => i18n.locale.value === 'en')

useAsyncData('home-content', async () => {
  isInItLoading.value = true
  await getHomeContent()
  isInItLoading.value = false
})

useAsyncData('instagram-oembed', async () => {
  await getInstagramOEmbed()
})

onMounted(() => {
  if (!document.querySelector('script[src="//www.instagram.com/embed.js"]')) {
    const instagramScript = document.createElement('script')
    instagramScript.src = '//www.instagram.com/embed.js'
    instagramScript.async = true
    document.body.appendChild(instagramScript)
  }
})

onMounted(() => {
  if (archiveRef.value) {
    archiveRef.value.addEventListener('scroll', updateScrollLeft)
  }
})

onUnmounted(() => {
  if (archiveRef.value) {
    archiveRef.value.removeEventListener('scroll', updateScrollLeft)
  }
})

const handleClickArchiveLeftButton = () => {
  if (archiveRef.value) {
    archiveRef.value.scrollTo({
      left: archiveRef.value.scrollLeft - 368,
      behavior: 'smooth',
    })
    updateScrollLeft()
  }
}

const handleClickArchiveRightButton = () => {
  if (archiveRef.value) {
    archiveRef.value.scrollTo({
      left: archiveRef.value.scrollLeft + 368,
      behavior: 'smooth',
    })
    updateScrollLeft()
  }
}

const handleClickLiveItem = (id: string) => {
  router.push(`/live/${id}`)
}

const handleClickLiveMore = () => {
  router.push(`/marches`)
}

const banners: BannerItem[] = [
  {
    imgSrc: '/img/banner2.jpg',
    link: '/about',
    isInternalLink: true,
  },
  { imgSrc: '/img/banner.png', link: '/about', isInternalLink: true },
  {
    imgSrc: '/img/banner3.png',
    link: '/live/p6XURyWhSk2EerwWiYYzDU',
    isInternalLink: true,
  },
]

const isOpen = ref<boolean>(false)
const handleClickMoreViewButton = () => {
  isOpen.value = !isOpen.value
}

const handleClickAllArchive = () => {
  router.push(`/marches`)
}

const handleClickAllItem = () => {
  router.push(`/items`)
}

useSeoMeta({
  title: 'トップページ',
})
</script>

<template>
  <div>
    <!-- 動画部分 -->
    <div class="relative md:h-[calc(100vh-180px)] h-[calc(100vh-140px)] w-full">
      <div class="absolute bg-black/50 w-full h-full" />
      <div
        class="absolute w-full h-full z-10 flex flex-col md:gap-40 justify-center"
      >
        <div
          class="text-white md:text-[48px] text-[28px] font-bold w-full text-center tracking-wider md:grow-0 grow flex flex-col justify-center"
        >
          <p>&#035;{{ tt('deepJapanText') }}</p>
          <p>{{ tt('localVideoMediaText') }}</p>
        </div>

        <div
          class="md:py-0 pb-12 px-4 xl:w-[40%] md:w-[80%] grid grid-cols-2 items-center justify-center md:gap-16 gap-4 mx-auto md:text-[18px] text-[14px] text-white font-bold"
        >
          <nuxt-link
            to="/items"
            class="flex flex-col sm:flex-row justify-center sm:items-baseline bg-base/50 rounded-xl px-8 py-2 hover:bg-base/60 tracking-wide text-center"
          >
            <div class="">
              <span class="md:text-[24px] text-[18px]">{{ tt('discoverProductLinkText') }}</span>
              <span
                v-if="isEnglish"
                class="sm:mr-1"
              >{{ tt('discoverConjunctionText') }}</span>
              <span v-else>{{ tt('discoverConjunctionText') }}</span>
            </div>
            {{ tt('discoverText') }}
          </nuxt-link>
          <nuxt-link
            to="/experiences"
            class="flex flex-col sm:flex-row justify-center sm:items-baseline bg-base/50 rounded-xl px-8 py-2 hover:bg-base/60 tracking-wide text-center"
          >
            <div class="flex flex-row items-baseline">
              <span class="md:text-[24px] text-[18px]">{{ tt('discoverExperienceLinkText') }}</span>
              <span
                v-if="isEnglish"
                class="sm:mr-1"
              >{{ tt('discoverConjunctionText') }}</span>
              <span v-else>{{ tt('discoverConjunctionText') }}</span>
            </div>
            {{ tt('discoverText') }}
          </nuxt-link>
        </div>
      </div>
      <video
        webkit-playsinline
        playsinline
        muted
        autoplay
        loop
        class="h-full object-cover w-full"
      >
        <source
          src="/video/furumaru.webm"
          type="video/webm"
        >
      </video>
    </div>

    <the-carousel
      v-if="false"
      :items="banners"
      :line-add-friend-image-url="tt('lineAddFriendImageUrl')"
      :line-add-friend-image-alt="tt('lineAddFriendImageAlt')"
      :line-coupon-text="tt('lineCouponText')"
    />

    <div class="mb-[72px] mt-8 flex flex-col gap-y-16 md:mt-[76px]">
      <the-content-box
        v-if="true"
        title="live"
        :sub-title="tt('marcheListSubTitle')"
      >
        <template v-if="isInItLoading" />
        <template v-if="lives.length === 0">
          <div class="flex justify-center">
            <img
              src="~/assets/img/furuneko-sleep.png"
              alt="furuneko sleep"
              width="120"
              height="136"
              class="block"
            >
          </div>
          <div class="mt-8 text-center text-[14px] text-main md:text-[16px]">
            <p>{{ tt("noMarcheItemFirstText") }}</p>
            <p class="md:mt-4">
              {{ tt("noMarcheItemSecondText") }}
            </p>
          </div>
          <div
            class="my-4 grid w-full justify-center md:mt-10 md:flex md:gap-x-16"
          >
            <button
              class="w-60 bg-main py-2 text-white"
              @click="handleClickAllArchive"
            >
              {{ tt("pastMarcheLinkText") }}
            </button>
            <button
              class="mt-4 w-60 bg-main py-2 text-white md:mt-0"
              @click="handleClickAllItem"
            >
              {{ tt("productsLinkText") }}
            </button>
          </div>
        </template>
        <template v-if="lives.length > 0">
          <div
            class="mx-auto grid max-w-7xl gap-x-10 gap-y-8 px-2 md:grid-cols-2 lg:grid-cols-3"
          >
            <transition-group
              enter-active-class="duration-300 ease-in-out"
              enter-from-class="opacity-0 h-0"
              enter-to-class="opacity-100 h-full"
              leave-active-class="duration-300 ease-in-out"
              leave-from-class="opacity-100 h-full"
              leave-to-class="opacity-0 h-0"
            >
              <the-live-item
                v-for="liveItem in lives"
                :id="liveItem.scheduleId"
                :key="liveItem.scheduleId"
                :title="liveItem.title"
                :img-src="liveItem.thumbnailUrl"
                :start-at="liveItem.startAt"
                :is-live-status="liveItem.status"
                :marche-name="liveItem.coordinator.marcheName"
                :address="liveItem.coordinator.city"
                :cn-name="liveItem.coordinator.username"
                :cn-img-src="liveItem.coordinator.thumbnailUrl"
                :live-streaming-text="tt('liveStreamingText')"
                :live-upcoming-text="tt('liveUpcomingText')"
                @click="handleClickLiveItem(liveItem.scheduleId)"
              />
            </transition-group>
          </div>
          <div
            v-if="false"
            class="mb-4 mt-10 flex w-full justify-center"
          >
            <button
              class="relative w-60 bg-main py-2 text-white"
              @click="handleClickMoreViewButton"
            >
              {{ tt("viewMoreText") }}
              <div class="absolute bottom-3.5 right-4">
                <the-up-arrow-icon
                  v-show="isOpen"
                  fill="white"
                />
                <the-down-arrow-icon
                  v-show="!isOpen"
                  fill="white"
                />
              </div>
            </button>
          </div>
        </template>
      </the-content-box>

      <the-content-box
        title="video"
        :sub-title="tt('archiveListSubTitle')"
      >
        <div class="relative mx-auto flex max-w-[1440px]">
          <div class="absolute left-4 flex h-[208px] items-center">
            <the-icon-button
              class="hidden bg-white/50 hover:bg-white md:block"
              @click="handleClickArchiveLeftButton"
            >
              <the-left-arrow-icon />
            </the-icon-button>
          </div>
          <div
            ref="archiveRef"
            class="hidden-scrollbar flex w-full flex-col gap-8 md:flex-row md:overflow-x-scroll"
          >
            <the-archive-item
              v-for="archive in archives"
              :id="archive.scheduleId"
              :key="archive.scheduleId"
              :title="archive.title"
              :img-src="archive.thumbnailUrl"
              :start-at="archive.startAt"
              :end-at="archive.endAt"
              :width="368"
              :archived-stream-text="tt('archivedStreamText')"
              class="cursor-pointer md:min-w-[368px] md:max-w-[368px]"
              @click="handleClickLiveItem(archive.scheduleId)"
            />
          </div>
          <div class="absolute right-4 flex h-[208px] items-center">
            <the-icon-button
              class="hidden bg-white/50 hover:bg-white md:block"
              @click="handleClickArchiveRightButton"
            >
              <the-right-arrow-icon />
            </the-icon-button>
          </div>
        </div>

        <div class="mb-4 mt-10 flex w-full justify-center">
          <button
            class="w-60 bg-main py-2 text-white"
            @click="handleClickLiveMore"
          >
            {{ tt("archivesLinkText") }}
          </button>
        </div>
      </the-content-box>

      <the-content-box
        title="instagram"
        sub-title="人気の動画"
      >
        <blockquote
          class="instagram-media"
          data-instgrm-captioned
          data-instgrm-permalink="https://www.instagram.com/p/DBwRdmPvSy7/?utm_source=ig_embed&amp;utm_campaign=loading"
          data-instgrm-version="14"
          style=" background:#FFF; border:0; border-radius:3px; box-shadow:0 0 1px 0 rgba(0,0,0,0.5),0 1px 10px 0 rgba(0,0,0,0.15); margin: 1px; max-width:540px; min-width:326px; padding:0; width:99.375%; width:-webkit-calc(100% - 2px); width:calc(100% - 2px);"
        >
          <div style="padding:16px;">
            <a
              href="https://www.instagram.com/p/DBwRdmPvSy7/?utm_source=ig_embed&amp;utm_campaign=loading"
              style=" background:#FFFFFF; line-height:0; padding:0 0; text-align:center; text-decoration:none; width:100%;"
              target="_blank"
            > <div style=" display: flex; flex-direction: row; align-items: center;"> <div style="background-color: #F4F4F4; border-radius: 50%; flex-grow: 0; height: 40px; margin-right: 14px; width: 40px;" /> <div style="display: flex; flex-direction: column; flex-grow: 1; justify-content: center;"> <div style=" background-color: #F4F4F4; border-radius: 4px; flex-grow: 0; height: 14px; margin-bottom: 6px; width: 100px;" /> <div style=" background-color: #F4F4F4; border-radius: 4px; flex-grow: 0; height: 14px; width: 60px;" /></div></div><div style="padding: 19% 0;" /> <div style="display:block; height:50px; margin:0 auto 12px; width:50px;"><svg
              width="50px"
              height="50px"
              viewBox="0 0 60 60"
              version="1.1"
              xmlns="http://www.w3.org/2000/svg"
              xmlns:xlink="http://www.w3.org/1999/xlink"
            ><g
              stroke="none"
              stroke-width="1"
              fill="none"
              fill-rule="evenodd"
            ><g
              transform="translate(-511.000000, -20.000000)"
              fill="#000000"
            ><g><path d="M556.869,30.41 C554.814,30.41 553.148,32.076 553.148,34.131 C553.148,36.186 554.814,37.852 556.869,37.852 C558.924,37.852 560.59,36.186 560.59,34.131 C560.59,32.076 558.924,30.41 556.869,30.41 M541,60.657 C535.114,60.657 530.342,55.887 530.342,50 C530.342,44.114 535.114,39.342 541,39.342 C546.887,39.342 551.658,44.114 551.658,50 C551.658,55.887 546.887,60.657 541,60.657 M541,33.886 C532.1,33.886 524.886,41.1 524.886,50 C524.886,58.899 532.1,66.113 541,66.113 C549.9,66.113 557.115,58.899 557.115,50 C557.115,41.1 549.9,33.886 541,33.886 M565.378,62.101 C565.244,65.022 564.756,66.606 564.346,67.663 C563.803,69.06 563.154,70.057 562.106,71.106 C561.058,72.155 560.06,72.803 558.662,73.347 C557.607,73.757 556.021,74.244 553.102,74.378 C549.944,74.521 548.997,74.552 541,74.552 C533.003,74.552 532.056,74.521 528.898,74.378 C525.979,74.244 524.393,73.757 523.338,73.347 C521.94,72.803 520.942,72.155 519.894,71.106 C518.846,70.057 518.197,69.06 517.654,67.663 C517.244,66.606 516.755,65.022 516.623,62.101 C516.479,58.943 516.448,57.996 516.448,50 C516.448,42.003 516.479,41.056 516.623,37.899 C516.755,34.978 517.244,33.391 517.654,32.338 C518.197,30.938 518.846,29.942 519.894,28.894 C520.942,27.846 521.94,27.196 523.338,26.654 C524.393,26.244 525.979,25.756 528.898,25.623 C532.057,25.479 533.004,25.448 541,25.448 C548.997,25.448 549.943,25.479 553.102,25.623 C556.021,25.756 557.607,26.244 558.662,26.654 C560.06,27.196 561.058,27.846 562.106,28.894 C563.154,29.942 563.803,30.938 564.346,32.338 C564.756,33.391 565.244,34.978 565.378,37.899 C565.522,41.056 565.552,42.003 565.552,50 C565.552,57.996 565.522,58.943 565.378,62.101 M570.82,37.631 C570.674,34.438 570.167,32.258 569.425,30.349 C568.659,28.377 567.633,26.702 565.965,25.035 C564.297,23.368 562.623,22.342 560.652,21.575 C558.743,20.834 556.562,20.326 553.369,20.18 C550.169,20.033 549.148,20 541,20 C532.853,20 531.831,20.033 528.631,20.18 C525.438,20.326 523.257,20.834 521.349,21.575 C519.376,22.342 517.703,23.368 516.035,25.035 C514.368,26.702 513.342,28.377 512.574,30.349 C511.834,32.258 511.326,34.438 511.181,37.631 C511.035,40.831 511,41.851 511,50 C511,58.147 511.035,59.17 511.181,62.369 C511.326,65.562 511.834,67.743 512.574,69.651 C513.342,71.625 514.368,73.296 516.035,74.965 C517.703,76.634 519.376,77.658 521.349,78.425 C523.257,79.167 525.438,79.673 528.631,79.82 C531.831,79.965 532.853,80.001 541,80.001 C549.148,80.001 550.169,79.965 553.369,79.82 C556.562,79.673 558.743,79.167 560.652,78.425 C562.623,77.658 564.297,76.634 565.965,74.965 C567.633,73.296 568.659,71.625 569.425,69.651 C570.167,67.743 570.674,65.562 570.82,62.369 C570.966,59.17 571,58.147 571,50 C571,41.851 570.966,40.831 570.82,37.631" /></g></g></g></svg></div><div style="padding-top: 8px;"> <div style=" color:#3897f0; font-family:Arial,sans-serif; font-size:14px; font-style:normal; font-weight:550; line-height:18px;">この投稿をInstagramで見る</div></div><div style="padding: 12.5% 0;" /> <div style="display: flex; flex-direction: row; margin-bottom: 14px; align-items: center;"><div> <div style="background-color: #F4F4F4; border-radius: 50%; height: 12.5px; width: 12.5px; transform: translateX(0px) translateY(7px);" /> <div style="background-color: #F4F4F4; height: 12.5px; transform: rotate(-45deg) translateX(3px) translateY(1px); width: 12.5px; flex-grow: 0; margin-right: 14px; margin-left: 2px;" /> <div style="background-color: #F4F4F4; border-radius: 50%; height: 12.5px; width: 12.5px; transform: translateX(9px) translateY(-18px);" /></div><div style="margin-left: 8px;"> <div style=" background-color: #F4F4F4; border-radius: 50%; flex-grow: 0; height: 20px; width: 20px;" /> <div style=" width: 0; height: 0; border-top: 2px solid transparent; border-left: 6px solid #f4f4f4; border-bottom: 2px solid transparent; transform: translateX(16px) translateY(-4px) rotate(30deg)" /></div><div style="margin-left: auto;"> <div style=" width: 0px; border-top: 8px solid #F4F4F4; border-right: 8px solid transparent; transform: translateY(16px);" /> <div style=" background-color: #F4F4F4; flex-grow: 0; height: 12px; width: 16px; transform: translateY(-4px);" /> <div style=" width: 0; height: 0; border-top: 8px solid #F4F4F4; border-left: 8px solid transparent; transform: translateY(-4px) translateX(8px);" /></div></div> <div style="display: flex; flex-direction: column; flex-grow: 1; justify-content: center; margin-bottom: 24px;"> <div style=" background-color: #F4F4F4; border-radius: 4px; flex-grow: 0; height: 14px; margin-bottom: 6px; width: 224px;" /> <div style=" background-color: #F4F4F4; border-radius: 4px; flex-grow: 0; height: 14px; width: 144px;" /></div></a><p style=" color:#c9c8cd; font-family:Arial,sans-serif; font-size:14px; line-height:17px; margin-bottom:0; margin-top:8px; overflow:hidden; padding:8px 0 7px; text-align:center; text-overflow:ellipsis; white-space:nowrap;">
              <a
                href="https://www.instagram.com/p/DBwRdmPvSy7/?utm_source=ig_embed&amp;utm_campaign=loading"
                style=" color:#c9c8cd; font-family:Arial,sans-serif; font-size:14px; font-style:normal; font-weight:normal; line-height:17px; text-decoration:none;"
                target="_blank"
              >Instagram(@instagram)がシェアした投稿</a>
            </p>
          </div>
        </blockquote>
        <div
          class="mt-4"
          :v-html="OEmbedHTML"
        />
      </the-content-box>>

      <the-content-box
        v-if="false"
        title="recommend"
        sub-title="おすすめの商品"
      >
        <div
          class="mx-auto grid max-w-[1440px] grid-cols-2 gap-x-8 gap-y-6 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5"
        >
          <the-product-list-item
            v-for="productItem in MOCK_RECOMMEND_ITEMS"
            :id="productItem.id"
            :key="productItem.id"
            :name="productItem.name"
            :price="productItem.price"
            :img-src="productItem.imgSrc"
            :inventory="productItem.inventory"
            :address="productItem.address"
            :cn-name="productItem.cnName"
            :cn-img-src="productItem.cnImgSrc"
          />
        </div>
      </the-content-box>
    </div>
  </div>
</template>

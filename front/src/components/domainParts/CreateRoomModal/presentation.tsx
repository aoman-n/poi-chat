import React, { useState, useRef, useCallback } from 'react'
import { CSSTransition } from 'react-transition-group'
import Swiper, { ReactIdSwiperProps, SwiperRefNode } from 'react-id-swiper'
import ClientOnlyPortal from '@/components/parts/ClientOnlyPortal'
import Icon from '@/components/parts/Icon'
import CircleArrowButton from '@/components/parts/CircleArrowButton'
import { isExistsRef } from '@/utils/elements'
import styles from './index.module.scss'
import 'swiper/swiper.min.css'
import 'swiper/components/pagination/pagination.min.css'

export type CreateRoomModalProps = {
  open: boolean
  handleClose: () => void
}

type Image = {
  id: string
  url: string
}

const images: Image[] = [
  {
    id: '1',
    url: 'roomBg1.jpg',
  },
  {
    id: '2',
    url: 'roomBg2.png',
  },
  {
    id: '3',
    url: 'roomBg3.jpg',
  },
]

const CreateRoomModal: React.FC<CreateRoomModalProps> = ({
  open,
  handleClose,
}) => {
  const [selectedImageId, setSelectedImageId] = useState(images[0].id)
  const swiperRef = useRef<SwiperRefNode>(null)

  const params: ReactIdSwiperProps = {
    pagination: {
      el: '.swiper-pagination',
      type: 'custom',
      clickable: true,
      dynamicBullets: true,
    },
    spaceBetween: 30,
    noSwiping: true,
  }

  const goPrev = useCallback(() => {
    if (isExistsRef(swiperRef)) {
      swiperRef.current.swiper?.slideNext()
    }
  }, [])

  const goNext = useCallback(() => {
    if (isExistsRef(swiperRef)) {
      swiperRef.current.swiper?.slidePrev()
    }
  }, [])

  return (
    <ClientOnlyPortal selector="#modal-root">
      <CSSTransition
        in={open}
        unmountOnExit
        timeout={200}
        classNames={{
          enter: styles.backdropEnter,
          enterActive: styles.backdropEnterActive,
          exit: styles.backdropExit,
          exitActive: styles.backdropExitActive,
        }}
      >
        <div className={styles.backdrop} onClick={handleClose}>
          <div className={styles.modal} onClick={(e) => e.stopPropagation()}>
            <div className="bg-white" style={{ width: '680px' }}>
              <h3 className="py-6 text-center border-b-2 border-gray-200 text-2xl text-gray-800">
                チャットルーム作成
              </h3>
              <div className="border-b-2 border-gray-200 py-12 px-12">
                <div className="mb-8 text-lg">
                  <label
                    htmlFor="username"
                    className="block mb-3 text-gray-700"
                  >
                    ルーム名
                  </label>
                  <input
                    id="username"
                    type="text"
                    className="rounded-sm px-4 py-3 bg-gray-100 w-full focus:outline-none text-gray-700"
                    placeholder="ルーム名を入力"
                  />
                </div>
                <div className="mb-5 text-lg">
                  <label
                    htmlFor="username"
                    className="block mb-3 text-gray-700"
                  >
                    壁紙選択
                  </label>
                  <div className="relative">
                    <Swiper {...params} ref={swiperRef}>
                      {images.map((image) => (
                        <div key={image.id} className="relative">
                          <img src={image.url} alt={image.id} />
                          {image.id === selectedImageId ? (
                            <div className="absolute top-0 left-0 w-full h-full z-50 flex justify-center items-center duration-100">
                              <div className="h-24 w-24 rounded-full border-2 border-white flex justify-center items-center bg-green-500 bg-opacity-75 duration-100">
                                <Icon
                                  type="check"
                                  color="white"
                                  className="h-14 w-14"
                                />
                              </div>
                            </div>
                          ) : (
                            <div className="group cursor-pointer absolute top-0 left-0 w-full h-full z-50 hover:bg-gray-600 hover:opacity-90 flex justify-center items-center duration-100">
                              <div className="invisible group-hover:visible h-24 w-24 rounded-full  flex justify-center items-center bg-opacity-75 duration-100">
                                <Icon
                                  type="check"
                                  color="white"
                                  className="h-16 w-16"
                                />
                              </div>
                            </div>
                          )}
                        </div>
                      ))}
                    </Swiper>
                    <CircleArrowButton
                      arrowType="right"
                      onClick={goPrev}
                      classNames="absolute top-1/2 -right-8 transform -translate-y-2/4 z-50"
                    />
                    <CircleArrowButton
                      arrowType="left"
                      onClick={goNext}
                      classNames="absolute top-1/2 -left-8 transform -translate-y-2/4 z-50"
                    />
                  </div>
                </div>
              </div>
              <div className="flex justify-end py-4 px-8 bg-gray-50 space-x-4">
                <button
                  type="submit"
                  className="focus:outline-none py-3 px-7 text-gray-700 font-bold rounded-sm bg-gray-50 duration-200 hover:opacity-90 text-lg border-2 border-gray-300 hover:border-gray-600"
                >
                  キャンセル
                </button>
                <button
                  type="submit"
                  className="focus:outline-none py-3 px-7 text-white font-semibold rounded-sm bg-gray-800 duration-200 hover:opacity-90 text-lg"
                >
                  作成
                </button>
              </div>
            </div>
          </div>
        </div>
      </CSSTransition>
    </ClientOnlyPortal>
  )
}

export default CreateRoomModal

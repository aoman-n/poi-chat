import React from 'react'
import { CSSTransition } from 'react-transition-group'
import ClientOnlyPortal from '@/components/parts/ClientOnlyPortal'
import styles from './index.module.scss'

export type CreateRoomModalProps = {
  open: boolean
  handleClose: () => void
  children: React.ReactNode
}

const Modal: React.FC<CreateRoomModalProps> = ({
  open,
  handleClose,
  children,
}) => {
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
            {children}
          </div>
        </div>
      </CSSTransition>
    </ClientOnlyPortal>
  )
}

export default Modal

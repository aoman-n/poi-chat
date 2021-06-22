import React, { useRef, useEffect, useState } from 'react'
import { createPortal } from 'react-dom'

type ClientOnlyPortalProps = {
  selector: string
}

const ClientOnlyPortal: React.FC<ClientOnlyPortalProps> = ({
  children,
  selector,
}) => {
  const ref = useRef<HTMLDivElement | null>(null)
  const [mounted, setMounted] = useState(false)

  useEffect(() => {
    ref.current = document.querySelector(selector)
    setMounted(true)
  }, [selector])

  return mounted && ref.current ? createPortal(children, ref.current) : null
}

export default ClientOnlyPortal

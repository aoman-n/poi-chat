import '../src/styles/globals.css'
import 'swiper/swiper.min.css'
import 'swiper/components/pagination/pagination.min.css'

export const parameters = {
  actions: { argTypesRegex: "^on[A-Z].*" },
}

const modalRoot = document.createElement('div');
modalRoot.setAttribute('id', 'modal-root');
document.body.append(modalRoot);

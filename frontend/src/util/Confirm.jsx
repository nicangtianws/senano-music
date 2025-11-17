// utils/confirm.jsx
import { createRoot } from 'react-dom/client'
import ConfirmBox from '../components/ConfirmBox'

const MESSAGE_CONTAINER_ID = 'confirm-container'

const ensureContainer = () => {
  let container = document.getElementById(MESSAGE_CONTAINER_ID)
  if (!container) {
    container = document.createElement('div')
    container.id = MESSAGE_CONTAINER_ID
    document.body.appendChild(container)
  }
  return container
}

// Promise 版渲染方法
const renderConfirm = (type, message, title) => {
  return new Promise((resolve) => {
    const container = ensureContainer()
    const wrapper = document.createElement('div')
    container.appendChild(wrapper)

    const root = createRoot(wrapper)

    const handleResolve = () => {
      cleanup()
      resolve(true)
    }

    const handleCancel = () => {
      cleanup()
      resolve(false)
    }

    const cleanup = () => {
      root.unmount()
      container.removeChild(wrapper)
    }

    root.render(
      <ConfirmBox
        type={type}
        title={title}
        message={message}
        onResolve={handleResolve}
        onCancel={handleCancel}
        onClose={cleanup}
      />
    )
  })
}

const Confirm = {
  confirm: (msg, title = '确认') => renderConfirm('confirm', msg, title),
  info: (msg, title = '提示') => renderConfirm('info', msg, title),
  success: (msg, title = '成功') => renderConfirm('success', msg, title),
  error: (msg, title = '错误') => renderConfirm('error', msg, title),
  warning: (msg, title = '警告') => renderConfirm('warning', msg, title)
}

export default Confirm

// utils/message.jsx
import { createRoot } from 'react-dom/client'
import MessageItem from '../components/MessageItem'

let messageIdCounter = 0

const MESSAGE_CONTAINER_ID = 'message-container'

const createContainer = () => {
  let container = document.getElementById(MESSAGE_CONTAINER_ID)
  if (!container) {
    container = document.createElement('div')
    container.id = MESSAGE_CONTAINER_ID
    document.body.appendChild(container)
  }
  return container
}

const renderMessage = (type, message, title) => {
  const container = createContainer()
  const messageWrapper = document.createElement('div')
  const messageId = `message-${messageIdCounter++}`
  messageWrapper.id = messageId
  
  const root = createRoot(messageWrapper)
  
  const onRemove = () => {
    root.unmount()
    messageWrapper.remove()
  }
  
  root.render(
    <MessageItem 
      type={type} 
      message={message} 
      title={title} 
      onRemove={onRemove}
    />
  )
  
  container.appendChild(messageWrapper)
}

const Message = {
  info: (msg, title = '提示') => renderMessage('info', msg, title),
  success: (msg, title = '成功') => renderMessage('success', msg, title),
  error: (msg, title = '错误') => renderMessage('error', msg, title),
  warning: (msg, title = '警告') => renderMessage('warning', msg, title)
}

export default Message
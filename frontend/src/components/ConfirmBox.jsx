// components/ConfirmBox.jsx
import styled from 'styled-components'

const ConfirmBoxDiv = styled.div`
  position: fixed;
  top: 30%;
  left: 50%;
  transform: translate(-50%, -50%);
  z-index: 9999;
  width: 300px;
  color: #333;
  background: #fff;
  border: 1px solid #ccc;
  box-shadow: 0 2px 10px rgba(0,0,0,0.1);
  border-radius: 8px;
  padding: 16px;
  font-family: sans-serif;
`

export default function ConfirmBox({ title, message, onResolve, onCancel, onClose }) {
  const handleConfirm = () => {
    onResolve?.()
    onClose?.()
  }

  const handleCancel = () => {
    onCancel?.()
    onClose?.()
  }

  return (
    <ConfirmBoxDiv>
      <h3>{title}</h3>
      <p style={{ margin: '12px 0' }}>{message}</p>
      <div style={{ textAlign: 'right' }}>
        <button onClick={handleConfirm}>确认</button>
        <button onClick={handleCancel} style={{ marginLeft: 10 }}>取消</button>
      </div>
    </ConfirmBoxDiv>
  )
}

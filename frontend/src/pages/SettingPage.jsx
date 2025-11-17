import { Link } from 'react-router'
import { MusicScan } from '../../wailsjs/go/api/App'
import Message from '../util/Message'
import Confirm from '../util/Confirm'
import styled from 'styled-components'

const SettingItemDiv = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
`

export default function SettingPage() {
  return (
    <>
      <h1>Settings</h1>
      <Link to="/">
        <i className="icon bi bi-arrow-90deg-left"></i>
      </Link>
      <br />
      <div>
        <SettingItemDiv>
          <div>重新扫描</div>
          <div>
            <button
              onClick={(e) => {
                Confirm.warning('是否确认重新扫描？').then((res) => {
                  if(res) {
                    MusicScan().then((res) => {
                      if (!res) {
                        Message.error('未知错误')
                        return
                      }
                      const response = JSON.parse(res)
                      if (response.code !== 200) {
                        Message.error('扫描失败，' + response.message)
                        return
                      }
                      Message.info('扫描完成')
                    })
                  }
                })
              }}
            >
              扫描
            </button>
          </div>
        </SettingItemDiv>
      </div>
    </>
  )
}

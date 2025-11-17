import { useEffect, useState } from 'react'
import { Link } from 'react-router'
import styled from 'styled-components'
import MusicItem from '../components/MusicItem'
import { MusicList } from '../../wailsjs/go/api/App'
import FooterBox from '../components/FooterBox'

const BodyDiv = styled.div`
  display: grid;
  grid-template-columns: 3fr 1fr;
  grid-template-rows: 50px 1fr;
  grid-template-areas:
    'header header'
    'main sidebar';
`

const HeaderDiv = styled.div`
  grid-area: header;
  display: grid;
  grid-template-columns: 2fr 3fr 6fr;
  background-color: ${(props) => props.theme.bg3};
  height: 50px;
`
const OperateDiv = styled.div`
  padding: 5px 30px;
  display: flex;
  justify-content: flex-end;
  align-items: center;
`

const MainDiv = styled.div`
  grid-area: main;
  display: grid;
  background-color: gray;
  padding-bottom: 100px;
`

const SidebarDiv = styled.div`
  grid-area: sidebar;
  display: grid;
  background-color: lightgray;
`

const FooterDiv = styled.div`
  grid-area: footer;
  display: grid;
  background-color: lightblue;
  width: 100%;
  height: 100px;
  position: fixed;
  bottom: 0;
`

export default function HomePage() {
  const [musicInfos, setMusicInfos] = useState([])
  const [currentMusic, setCurrentMusic] = useState({})

  const play = (musicInfo) => {
    console.log('play music', musicInfo)
    setCurrentMusic(musicInfo)
  }

  useEffect(() => {
    MusicList().then((res) => {
      if (!res) {
        message('加载失败！')
      }
      let response = JSON.parse(res)
      console.log('MusicInfos', response)
      if (response.code != 200) {
        message(response.message, '加载失败！')
        return
      }
      let newMusicInfos = response.data.map((musicInfo, index) => {
        return (
          <MusicItem
            musicInfo={musicInfo}
            key={'music-info-' + index}
            play={play}
          ></MusicItem>
        )
      })
      setMusicInfos(newMusicInfos)
    })
  }, [])

  return (
    <>
      <BodyDiv>
        <HeaderDiv>
          <div className="logo"></div>
          <div className="page-name"></div>
          <OperateDiv>
            <Link to={'setting'}>
              <i className="bi bi-gear icon"></i>
            </Link>
          </OperateDiv>
        </HeaderDiv>
        <MainDiv>
          <div boxtype="full" id="search-box"></div>
          <div boxtype="full" id="list-box">
            <ul>{musicInfos}</ul>
          </div>
        </MainDiv>
        <SidebarDiv>
          <div boxtype="full"></div>
          <div boxtype="full"></div>
        </SidebarDiv>
        <FooterDiv>
          <FooterBox currentMusic={currentMusic}></FooterBox>
        </FooterDiv>
      </BodyDiv>
    </>
  )
}

export default function MusicItem({ musicInfo, play }) {
  return (
    <>
      <li onClick={() => {
        play(musicInfo)
      }}>{musicInfo.title}</li>
    </>
  )
}

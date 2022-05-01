const audioPlay = (): void => {
  const audio: HTMLAudioElement = new Audio('../public/bell.mp3')
  audio.play().then(() => {
    console.log("Audio started!")
  })
    .catch(error => console.warn(error))
}

export default audioPlay;
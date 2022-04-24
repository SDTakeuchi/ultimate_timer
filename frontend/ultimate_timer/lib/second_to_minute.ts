const secondToMinute = (second: number): object => {
  const hh: number = Number(("00"+ ~~(second / 3600)).slice(-2));
  const mm: number = Number(("00"+ ~~(~~(second / 60) % 60)).slice(-2));
  const ss: number = Number(("00"+ ~~(second % 60)).slice(-2));
  return {
    hour: hh,
    min: mm,
    sec: ss,
  };
}

export default secondToMinute;

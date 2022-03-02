const zeroPad = (num, places) => String(num).padStart(places, "0");

export const convertToDateString = (timestamp) => {
  var date = new Date(timestamp);
  var dd = date.getDate();
  var mm = date.getMonth() + 1;
  var yy = date.getFullYear();
  var hour = date.getHours();
  var min = date.getMinutes();
  return (
    dd + "/" + mm + "/" + yy + " " + zeroPad(hour, 2) + ":" + zeroPad(min, 2)
  );
};

export const isEmpty = (obj) => {
  for (var i in obj) {
    return false;
  }
  return true;
};

export function removeEmptyFields(data) {
  Object.keys(data).forEach((key) => {
    if (data[key] === "" || data[key] == null) {
      delete data[key];
    }
  });
}

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

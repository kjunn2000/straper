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

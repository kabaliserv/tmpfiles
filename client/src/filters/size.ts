import Vue from "vue";

let Size = {
  TB: 2 ** 40,
  GB: 2 ** 30,
  MB: 2 ** 20,
  KB: 2 ** 10,
};

Vue.filter("size", function(value: number): string {
  let result = "";
  switch (true) {
    case value > Size.TB:
      result = (value / Size.TB).toFixed(2) + " TB";
      break;
    case value > Size.GB:
      result = (value / Size.GB).toFixed(2) + " GB";
      break;
    case value > Size.MB:
      result = (value / Size.MB).toFixed(2) + " MB";
      break;
    case value > Size.KB:
      result = (value / Size.KB).toFixed(2) + " KB";
      break;
    default:
        result = value + ' Octet'
  }
  return result;
});

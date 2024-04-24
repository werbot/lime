import { apiGet } from "@/utils/api";

export function getCookie(cname: string): string {
  const name = cname + "=";
  const decodedCookie = decodeURIComponent(document.cookie);
  const ca = decodedCookie.split(";");

  for (const cookie of ca) {
    let c = cookie.trim();
    if (c.startsWith(name)) {
      return c.substring(name.length);
    }
  }
  return "";
}

export function setCookie(cName: string, cValue: string, days: number): void {
  const date = new Date();
  date.setDate(date.getDate() + days);
  const value = `${cValue}${days === null ? "" : "; expires=" + date.toUTCString()}`;
  document.cookie = `${cName}=${value}`;
}

export function delCookie(cName: string): void {
  document.cookie = `${cName}=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;`;
}

export function priceFormat(cost: string): string {
  return Number(cost) ? (Number(cost) / 100).toFixed(2) : "0.00";
}

export const termObj = [
  /*{ name: "", color: "" },*/
  { name: "hour", color: "gray" },
  { name: "day", color: "pink" },
  { name: "week", color: "indigo" },
  { name: "month", color: "purple" },
  { name: "year", color: "green" },
];

export const actionObj = [
  /*{ name: "", color: "" },*/
  { name: "onSendMail", color: "purple" },
  { name: "onSignIn", color: "gray" },
  { name: "onSignOut", color: "gray" },
  { name: "onAdd", color: "green" },
  { name: "onUpdate", color: "yellow" },
  { name: "onDelete", color: "red" },
  { name: "onClone", color: "green" },
];

export const sectionsObj = [
  /*{ name:"", ico: "" },*/
  { name: "System", ico: "tooth" },
  { name: "Setting", ico: "tooth" },
  { name: "Customer", ico: "users" },
  { name: "Pattern", ico: "pattern" },
  { name: "License", ico: "ticket" },
  { name: "Payment", ico: "banknotes" },
];

export const paymentStatusObj = [
  /*{ name: "", color: "" },*/
  { name: "paid", color: "green" },
  { name: "unpaid", color: "red" },
  { name: "processed", color: "yellow" },
  { name: "canceled", color: "red" },
  { name: "failed", color: "red" },
];

export const paymentProvidersObj = [
  /*{ name: "", color: "" },*/
  { name: "None", color: "gray" },
  { name: "Stripe", color: "purple" },
];

export const currencyObj = [
  /*{ name: "", symbol: ""},*/
  { name: "EUR", symbol: "€" },
  { name: "USD", symbol: "＄" },
  { name: "JPY", symbol: "JP¥" },
  { name: "GBP", symbol: "£" },
  { name: "AUD", symbol: "A＄" },
  { name: "CAD", symbol: "CA＄" },
  { name: "CHF", symbol: "₣" },
  { name: "CNY", symbol: "¥" },
  { name: "SEK", symbol: "kr" },
];

export const monthNamesObj = [
  /*{ short: "", large: ""},*/
  { short: "Jan", large: "January" },
  { short: "Feb", large: "February" },
  { short: "Mar", large: "March" },
  { short: "Apr", large: "April" },
  { short: "May", large: "May" },
  { short: "Jun", large: "June" },
  { short: "Jul", large: "July" },
  { short: "Aug", large: "August" },
  { short: "Sep", large: "September" },
  { short: "Oct", large: "October" },
  { short: "Nov", large: "November" },
  { short: "Dec", large: "December" },
];

export function arrayToObject(array, valueExtractor) {
  return array.reduce((accumulator, item, index) => {
    const key = index + 1; // Start keys from 1
    accumulator[key] = valueExtractor(item);
    return accumulator;
  }, {});
}

export function reduceToObject(items, parseFunc) {
  return items.reduce((obj, item) => {
    obj[item.key] = parseFunc(item.value);
    return obj;
  }, {});
}

export function formatDate(timestamp): string {
  if (!isNaN(timestamp)) {
    timestamp *= 1000;
  }

  const date = new Date(timestamp);
  return `${date.getDate()} ${
    monthNamesObj[date.getMonth()].short
  } ${date.getFullYear()}, ${date.toLocaleTimeString()}`;
}

export function randomString(length: number): string {
  const characters =
    "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
  let result = "";
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
}

export interface DataItem {
  key: string;
  value: number;
}

export function firstLetter(sentence: string) {
  if (sentence && typeof sentence === "string") {
    return sentence.charAt(0).toUpperCase() + sentence.slice(1);
  }
  return sentence;
}

export function costFormat(cost) {
  return Number(cost) ? (Number(cost) / 100).toFixed(2) : "0.00";
}

export function costStripe(cost) {
  return Math.round(Number(cost) * 100);
}

/*
export async function fetchListData(url, listType, keyPath, valuePath) {
  try {
    const res = await apiGet(url, {});
    if (res.code === 200 && Array.isArray(res.result[listType])) {
      const partsKeyPath = keyPath.split('.');
      const partsValuePath = valuePath.split('.');

      return res.result[listType].reduce((accumulator, item) => {
        let key = partsKeyPath.reduce((subItem, part) => subItem?.[part], item);
        let value = partsValuePath.reduce((subItem, part) => subItem?.[part], item);
        if (key !== undefined && value !== undefined) {
          accumulator[key] = value;
        }
        return accumulator;
      }, {});
    }
  } catch (error) {
    console.error(`Error fetching ${listType}:`, error);
  }

  return {};
}
*/
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

export const termFormat = [
  { name: "", color: "" },
  { name: "hour", color: "gray" },
  { name: "day", color: "pink" },
  { name: "week", color: "indigo" },
  { name: "month", color: "purple" },
  { name: "year", color: "green" },
];

export const actionFormat = [
  { name: "onSendMail", color: "purple" },
  { name: "onSignIn", color: "gray" },
  { name: "onSignOut", color: "gray" },
  { name: "onAdd", color: "green" },
  { name: "onUpdate", color: "yellow" },
  { name: "onDelete", color: "red" },
  { name: "onClone", color: "green" },
];

export const paymentStatusFormat = [
  { name: "", color: "" },
  { name: "paid", color: "green" },
  { name: "unpaid", color: "red" },
  { name: "processed", color: "yellow" },
  { name: "canceled", color: "red" },
  { name: "failed", color: "red" },
];

export const term = ["hour", "day", "week", "month", "year"];

export const sections = ["System", "Setting", "Customer", "Pattern", "License"];

export const currency = [
  "EUR",
  "USD",
  "JPY",
  "GBP",
  "AUD",
  "CAD",
  "CHF",
  "CNY",
  "SEK",
];

export const monthNames = [
  "Jan",
  "Feb",
  "Mar",
  "Apr",
  "May",
  "Jun",
  "Jul",
  "Aug",
  "Sep",
  "Oct",
  "Nov",
  "Dec",
];

export function formatDate(timestamp): string {
  if (!isNaN(timestamp)) {
    timestamp *= 1000;
  }

  const date = new Date(timestamp);
  return `${date.getDate()} ${
    monthNames[date.getMonth()]
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

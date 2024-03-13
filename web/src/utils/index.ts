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
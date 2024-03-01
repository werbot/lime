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

export function setCookie(cname: string, cvalue: string, days: number): void {
  const date = new Date();
  date.setDate(date.getDate() + days);
  const value = `${cvalue}${days === null ? "" : "; expires=" + date.toUTCString()}`;
  document.cookie = `${cname}=${value}`;
}
